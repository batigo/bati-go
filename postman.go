package bati

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	ErrPosterStopped        = errors.New("postman stopped")
	ErrPosterMsgHandlerNull = errors.New("msghandler is null")
	ErrPosterBlocked        = errors.New("postman writer bloced")
)

const (
	// 这个指postman消息发送方向
	// postman将消息往上游发送，接收上游的响应详细
	PostmanTypeSenderUpper PostmanType = 1
	// postman接收下游的消息，将响应发往下游
	PostmanTypeSenderDowner PostmanType = 2
)

type PostmanConf = ServiceConf

type ServiceConf struct {
	Servcie    string   `json:"service" ini:"service"`
	Kafka      *KafkaSt `json:"kafka" ini:"kafka"`
	QuitNotify bool     `json:"quit_notify" ini:"quit_notify"`
	// MultiRooms = true: 该channel对应的业务允许一个长连接attach多个room
	MultiRooms bool `json:"multi_rooms" ini:"multi_rooms"`
	AutoReg    bool `json:"auto_reg" ini:"auto_reg"`
	// liverpool自动注册后会发送ServiceMsgTypeClientJoin类消息给channel
	AutoRegNotify bool `json:"auto_reg_notify" ini:"auto_reg_notify"`
}

type KafkaSt struct {
	Hostports string `json:"hostports" ini:"hostports"`
	GroupId   string `json:"groupid" ini:"groupid"`
	Readers   int    `json:"readers" ini:"readers"`
	Writers   int    `json:"writers" ini:"writers"`
}

type PostmanType int8

type ServiceMsgHandler func(msg ServiceMsg, channel string) error

func (t PostmanType) WriterChannelPrefix() string {
	switch t {
	case PostmanTypeSenderUpper:
		return "_up"
	case PostmanTypeSenderDowner:
		return "_down"
	}

	return ""
}

func (t PostmanType) ReadChannelPrefix() string {
	switch t {
	case PostmanTypeSenderUpper:
		return "_down"
	case PostmanTypeSenderDowner:
		return "_up"
	}

	return ""
}

func NewPostman(typ PostmanType, conf ServiceConf, msghandler ServiceMsgHandler) (*Postman, error) {
	if msghandler == nil {
		return nil, ErrPosterMsgHandlerNull
	}

	postman := &Postman{
		typ:           typ,
		msghandler:    msghandler,
		Service:       conf.Servcie,
		msgchan:       make(chan ServiceMsg, 1000),
		broadcastchan: make(chan ServiceMsg, 1000),
		QuitNotify:    conf.QuitNotify,
	}

	var err error
	postman.mq, err = newmqbroker(postman, conf)
	if err != nil {
		return nil, err
	}

	return postman, nil
}

type Postman struct {
	Service    string
	QuitNotify bool

	typ        PostmanType
	mq         mqbroker
	msghandler func(ServiceMsg, string) error
	mutex      sync.RWMutex
	// objects protected by mutex
	stoped        bool
	msgchan       chan ServiceMsg
	broadcastchan chan ServiceMsg
}

func (p *Postman) Run() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.stoped {
		return ErrPosterStopped
	}

	p.mq.run()
	return nil
}

func (p *Postman) Stop() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.stoped {
		return
	}

	p.stoped = true
	close(p.msgchan)
	close(p.broadcastchan)
	p.mq.stop()
}

func (p *Postman) SendMsg(msg ServiceMsg) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.stoped {
		return ErrPosterStopped
	}

	select {
	case p.msgchan <- msg:
	//
	default:
		return ErrPosterBlocked
	}

	return nil
}

// 这个只能是上游的app（PostmanTypeSenderDowner）调用
func (p *Postman) BroadcastMsg(msg ServiceMsg) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.stoped {
		return ErrPosterStopped
	}

	select {
	case p.broadcastchan <- msg:
	//
	default:
		return ErrPosterBlocked
	}

	return nil
}

type mqbroker interface {
	run()
	stop()
}

func newmqbroker(postman *Postman, conf PostmanConf) (mqbroker, error) {
	localid := getlocalhostid()
	if localid == "" {
		return nil, fmt.Errorf("failed to get localid")
	}

	if conf.Kafka != nil && conf.Kafka.Hostports != "" {
		if conf.Kafka.Readers <= 0 {
			conf.Kafka.Readers = 32
		}
		if conf.Kafka.Writers <= 0 {
			conf.Kafka.Writers = 32
		}
		gid := conf.Kafka.GroupId
		if gid == "" {
			gid = "liverpool-consumer"
		}
		if postman.typ == PostmanTypeSenderUpper {
			// 往上游发消息的postman在收消息用不同的consumer group
			gid += "_" + localid
		}

		brokers := strings.Split(conf.Kafka.Hostports, ",")
		return &kfkbroker{
			postman: postman,
			brokers: brokers,
			group:   gid,
			localid: localid,
			readers: conf.Kafka.Readers,
			writers: conf.Kafka.Writers,
			mutext:  sync.RWMutex{},
		}, nil
	}

	return nil, fmt.Errorf("mq broker conf not found")
}
