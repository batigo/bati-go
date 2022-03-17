package bati

import (
	"encoding/json"
	"fmt"
)

// uid for user id
// cid for conn id

type ServiceMsgType int8

const (
	ServiceMsgTypeConnJoin ServiceMsgType = 1
	ServiceMsgTypeConnQuit ServiceMsgType = 2
	ServiceMsgTypeBiz      ServiceMsgType = 3
)

// ServiceMsg是bati <-> service之间的消息协议
type ServiceMsg struct {
	Id       string         `json:"id"`
	Type     ServiceMsgType `json:"type"`
	JoinData *JoinData      `json:"join_data,omitempty"`
	QuitData *QuitData      `json:"quit_data,omitempty"`
	BizData  *BizData       `json:"biz_data,omitempty"`
	Ts       int64          `json:"ts"`
}

type JoinData struct {
	Cid         string   `json:"cid,omitempty"`
	Uid         string   `json:"uid,omitempty"`
	JoinService bool     `json:"join_service,omitempty"`
	Rooms       []string `json:"rids,omitempty"`
}

type QuitData struct {
	Cid         string   `json:"cid,omitempty"`
	Uid         string   `json:"uid,omitempty"`
	QuitService bool     `json:"quit_service,omitempty"`
	Rooms       []string `json:"rids,omitempty"`
}

type BizType = uint8

const (
	BizMsgTypeUsers   BizType = 1
	BizMsgTypeRoom    BizType = 2
	BizMsgTypeService BizType = 3
	BizMsgTypeAll     BizType = 4
)

type BizData struct {
	Type           BizType  `json:"type"`
	Cids           []string `json:"cids,omitempty"`
	Uids           []string `json:"uids,omitempty"`
	Room           string   `json:"rid,omitempty"`
	BroadcastRatio uint8    `json:"broadcast_ratio"`
	BlackUids      []string `json:"black_uids,omitempty"`
	WhiteUids      []string `json:"white_uids,omitempty"`
	Data           []byte   `json:"data,omitempty"`
}

func (m ServiceMsg) String() string {
	switch m.Type {
	case ServiceMsgTypeConnJoin:
		return fmt.Sprintf("[service conn-join-msg] id: %s, cid: %s, uid: %s, join-service: %v, rids: %v", m.Id, m.JoinData.Cid, m.JoinData.Uid, m.JoinData.JoinService, m.JoinData.Rooms)
	case ServiceMsgTypeConnQuit:
		return fmt.Sprintf("[service conn-quit-msg] id: %s, cid: %s, uid: %s, quit-service: %v, rids: %v", m.Id, m.QuitData.Cid, m.QuitData.Uid, m.QuitData.QuitService, m.QuitData.Rooms)
	case ServiceMsgTypeBiz:
		return fmt.Sprintf("[service biz-msg] id: %s, cid: %s, uid: %s, quit-service: %v, rids: %v", m.Id, m.QuitData.Cid, m.QuitData.Uid, m.QuitData.QuitService, m.QuitData.Rooms)
	}
	return fmt.Sprintf("id: %s, typ: %d", m.Id, m.Type)
}

type BatiMsgType = uint8

const (
	BatiMsgTypeBiz      BatiMsgType = 1
	BatiMsgTypeConnQuit BatiMsgType = 2
)

type BatiMsg struct {
	Id   string          `json:"id"`
	Type BatiMsgType     `json:"type"`
	Data json.RawMessage `json:"data"`
	Cid  string          `json:"cid"`
	Uid  string          `json:"uid"`
	Ip   string          `json:"ip"`
	Ts   int64           `json:"ts"`
}
