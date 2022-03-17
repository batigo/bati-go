package bati

import (
	"encoding/json"
	"fmt"
)

// 将一个conn加入一组房间,
// joinService = true, 这个conn会加入到service整体群组里面
// len(rooms) > 0 && joinService = true这个conn会加入到对应的房间里，同时加入service集体群组
func (p *Postman) SendConnJoinMsg(conn string, rooms []string, joinService bool) error {
	if len(rooms) == 0 && !joinService {
		return fmt.Errorf("bad params: rooms empty & joinService not enable")
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeConnJoin,
		JoinData: &JoinData{
			Cid:         conn,
			JoinService: joinService,
			Rooms:       rooms,
		},
		Ts: getNowMillisecs(),
	})
}

// 将一个conn退出一组房间,
// quitService = true, 这个conn会退出service整体群组
// len(rooms) > 0 && quitService = true这个conn会退出对应的房间，同时退出service集体群组
func (p *Postman) SendConnQuitMsg(conn string, rooms []string, quitService bool) error {
	if len(rooms) == 0 && !quitService {
		return fmt.Errorf("bad params: rooms empty & quitService not enable")
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeConnQuit,
		QuitData: &QuitData{
			Cid:         conn,
			QuitService: quitService,
			Rooms:       rooms,
		},
		Ts: getNowMillisecs(),
	})
}

// 将一个user对应的所有conn加入一组房间,
// joinService = true, 这个user会加入到service整体群组里面
// len(rooms) > 0 && joinService = true这个user会加入到对应的房间里，同时加入service集体群组
func (p *Postman) SendUidJoinMsg(uid string, rooms []string, joinService bool) error {
	if len(rooms) == 0 && !joinService {
		return fmt.Errorf("bad params: rooms empty & joinService not enable")
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeConnJoin,
		JoinData: &JoinData{
			Uid:         uid,
			JoinService: joinService,
			Rooms:       rooms,
		},
		Ts: getNowMillisecs(),
	})
}

// 将一个user对应的所有conn退出一组房间,
// quitService = true, 这个user会退出service整体群组
// len(rooms) > 0 && quitService = true这个user会加入到对应的房间里，同时加入service集体群组
func (p *Postman) SendUidQuitMsg(uid string, rooms []string, quitService bool) error {
	if len(rooms) == 0 && !quitService {
		return fmt.Errorf("bad params: rooms empty & quitService not enable")
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeConnQuit,
		QuitData: &QuitData{
			Uid:         uid,
			QuitService: quitService,
			Rooms:       rooms,
		},
		Ts: getNowMillisecs(),
	})
}

// 向一组conn发送消息
// room != "" 表示只有conn在room里面才发送
func (p *Postman) SendConnsBizMsg(conns []string, room string, data interface{}) error {
	if len(conns) == 0 {
		return fmt.Errorf("bad params: conns empty")
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type: BizMsgTypeUsers,
			Cids: conns,
			Room: room,
			Data: bs,
		},
		Ts: getNowMillisecs(),
	})
}

// 向一组user对应的所有conn发送消息
// room != "" 表示只有conn在room里面才发送
func (p *Postman) SendUsersBizMsg(uids []string, room string, data interface{}) error {
	if len(uids) == 0 {
		return fmt.Errorf("bad params: uids empty")
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type: BizMsgTypeUsers,
			Uids: uids,
			Room: room,
			Data: bs,
		},
		Ts: getNowMillisecs(),
	})
}

// 向一个房间内所有conn发送消息
func (p *Postman) SendRoomBizMsg(room string, data interface{}) error {
	if len(room) == 0 {
		return fmt.Errorf("bad params: room empty")
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type: BizMsgTypeRoom,
			Room: room,
			Data: bs,
		},
		Ts: getNowMillisecs(),
	})
}

// 向一个room内所有conn广播消息, 额外附带一些控制条件
// ratio=x: 广播比率, 向room内x%的conn广播消息
// ratio < 100 的情况下，whiteUids对应用conn不受ratio参数影响
// blackUids对应的conn不会广播消息
func (p *Postman) SendRoomBizMsgCond(room string, data interface{}, ratio uint8, whiteUids, blackUids []string) error {
	if len(room) == 0 || ratio == 0 {
		return fmt.Errorf("bad params: room empty or ration is zero")
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type:           BizMsgTypeRoom,
			Room:           room,
			Data:           bs,
			BroadcastRatio: ratio,
			WhiteUids:      whiteUids,
			BlackUids:      blackUids,
		},
		Ts: getNowMillisecs(),
	})
}

func (p *Postman) SendServiceBizMsg(data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type: BizMsgTypeService,
			Data: bs,
		},
		Ts: getNowMillisecs(),
	})
}

// 向一个service内所有conn发送消息, 额外附带一些控制条件
// ratio=x: 广播比率, 向room内x%的conn广播消息
// ratio < 100 的情况下，whiteUids对应用conn不受ratio参数影响
// blackUids对应的conn不会广播消息
func (p *Postman) SendServiceBizMsgCond(data interface{}, ratio uint8, whiteUids, blackUids []string) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.SendMsg(ServiceMsg{
		Id:   GenMsgId(),
		Type: ServiceMsgTypeBiz,
		BizData: &BizData{
			Type:           BizMsgTypeService,
			Data:           bs,
			BroadcastRatio: ratio,
			WhiteUids:      whiteUids,
			BlackUids:      blackUids,
		},
		Ts: getNowMillisecs(),
	})
}
