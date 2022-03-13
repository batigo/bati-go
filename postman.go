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

type PostmanConf struct {
	Servcie string   `json:"service" ini:"service"`
	Kafka   *KafkaSt `json:"kafka" ini:"kafka"`
}

type KafkaSt struct {
	Hostports string `json:"hostports" ini:"hostports"`
	GroupId   string `json:"groupid" ini:"groupid"`
	Readers   int    `json:"readers" ini:"readers"`
	Writers   int    `json:"writers" ini:"writers"`
}

type PostmanType int8

type BatiMsgHandler func(msg BatiMsg, service string) error

func NewPostman(conf PostmanConf, msghandler BatiMsgHandler) (*Postman, error) {
	if msghandler == nil {
		return nil, ErrPosterMsgHandlerNull
	}

	postman := &Postman{
		msghandler: msghandler,
		Service:    conf.Servcie,
		msgchan:    make(chan ServiceMsg, 1000),
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

	mq         mqbroker
	msghandler func(BatiMsg, string) error
	mutex      sync.RWMutex
	// objects protected by mutex
	stoped  bool
	msgchan chan ServiceMsg
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

type mqbroker interface {
	run()
	stop()
}

func newmqbroker(postman *Postman, conf PostmanConf) (mqbroker, error) {
	if conf.Kafka != nil && conf.Kafka.Hostports != "" {
		if conf.Kafka.Readers <= 0 {
			conf.Kafka.Readers = 32
		}
		if conf.Kafka.Writers <= 0 {
			conf.Kafka.Writers = 32
		}
		gid := conf.Kafka.GroupId
		if gid == "" {
			return nil, fmt.Errorf("group id missing")
		}

		brokers := strings.Split(conf.Kafka.Hostports, ",")
		return &kfkbroker{
			postman: postman,
			brokers: brokers,
			group:   gid,
			readers: conf.Kafka.Readers,
			writers: conf.Kafka.Writers,
			mutext:  sync.RWMutex{},
		}, nil
	}

	return nil, fmt.Errorf("mq broker conf not found")
}
