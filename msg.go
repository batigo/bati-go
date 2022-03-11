package bati

import (
	"encoding/json"
	"fmt"
)

// 和上游service的消息
// QoS:
// 0: most once
// 1: at least once
// 2: exact once
// 现阶段只支持 QoS = 0

type ServiceMsgType int8

const (
	// service下发消息的四种类型, service -> bati
	// 单个session
	ServiceMsgTypeSession ServiceMsgType = 1

	// 这个支持给所有的service用户和service下一个room的所有用户（ServiceMsg.RoomId有值）发送消息
	ServiceMsgTypeService ServiceMsgType = 2

	// 给bati所有的长连接发消息
	ServiceMsgTypeBroadCast ServiceMsgType = 3

	// 一个room里面多个用户发送消息
	ServiceMsgTypeRoomUsers ServiceMsgType = 10

	// 下面三个类型是用来维护业务service & room的关系数据, service -> bati
	// 用户注册service和room
	ServiceMsgTypeRegService ServiceMsgType = 4
	// 用户注销room
	ServiceMsgTypeUnregRoom ServiceMsgType = 5
	// 用户注销service
	ServiceMsgTypeUnregService ServiceMsgType = 6

	// 客户端断开连接，bati -> service
	ServiceMsgTypeClientQuit ServiceMsgType = 7
	// 客户端连接带了注册的service，bati会想对应的service发送这个消息，bati -> service
	ServiceMsgTypeClientJoin ServiceMsgType = 8
)

// ServiceMsg是bati <-> service之间的消息协议
type ServiceMsg struct {
	Id string `json:"id"`

	// 消息类型，bati支持的类型见上
	// bati -> service的消息只有ServiceMsgTypeClientQuit会有值，空值表示业务消息
	// service -> bati消息类型见上
	Type ServiceMsgType `json:"t,omitempty"`

	// bati -> service, 标明消息是从哪个session发送的
	// service -> bati, 消息的目标session，以下消息类型需要:
	// ServiceMsgTypeSession: 必须
	// ServiceMsgTypeRegService: SessionId和Mid必须有一个有效值, SessionId有效值优先
	// ServiceMsgTypeUnregRoom: SessionId和Mid必须有一个有效值, SessionId有效值优先
	// ServiceMsgTypeUnregService: SessionId和Mid必须有一个有效值, SessionId有效值优先
	SessionId string `json:"sid,omitempty"`

	// bati -> service: 触发消息的用户id
	// service -> bati: 业务方操作的用户id，比如
	Mid int64 `json:"mid,omitempty"`

	// 消息的目标service
	// serviceid不用在队列消息数据里面体现，因为消息都是从固定的service(topic)读取消息, 这个字段只是为了往下游client发送消息带的字段
	ServiceId string `json:"cid,omitempty"`
	// 消息的目前roomid, 如果没有这个字段会给service下面所有的长连接发送消息
	RoomId string `json:"rid,omitempty"`

	// ServiceMsgTypeRoomUsers 消息类型的用户id列表
	Mids []int64 `json:"mids,omitempty"`

	// 上游service发送的消息需要bati保证的qos，现阶段不做这块，只做qos=0
	Qos int8 `json:"qos,omitempty"`
	// 上下游业务之间的定义的消息数据结构
	Data json.RawMessage `json:"d,omitempty"`

	// extension字段
	MqSuffix string `json:"ms,omitempty"`
	T        int64  `json:"ts,omitempty"`
	// 广播消息的发送比例，[0-100]， 0， 100表示全部发送，n(0<n<100)表示n%的发送比例
	BroadcastRate int8    `json:"broadcast_rate,omitempty"`
	ExcludeMids   []int64 `json:"exclude_mids,omitempty"`
	IncludeMids   []int64 `json:"include_mids,omitempty"`
	// bati -> service会有这个字段，长连接客户端ip
	IP string `json:"ip,omitempty"`
}

func (m ServiceMsg) String() string {
	return fmt.Sprintf("id: %s, sid: %s, typ: %d", m.Id, m.SessionId, m.Type)
}

// 维护service&room群组关系的数据
type ServiceDataSt struct {
	// service字段，业务方可以不再填写该字段，bati会自动填充好
	Service string `json:"cid,omitempty"`
	Room    string `json:"rid,omitempty"`
	// service支持多房间可以用这个字段来注册/注销多个房间
	Rooms []string `json:"rids,omitempty"`
}
