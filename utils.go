package bati

import (
	"crypto/md5"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/hashicorp/go-uuid"
)

func getlocalhostid() string {
	var ip net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(fmt.Sprintf("failed to check interfaces: %s", err.Error()))
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}

	if len(ip) > 0 {
		return ip.String()
	}

	name, err := os.Hostname()
	if err != nil {
		return ""
	}

	return name
}

func getNowMillisecs() uint64 {
	return uint64(time.Now().UnixMilli())
}

func GenMsgId() string {
	s, _ := uuid.GenerateUUID()
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))[:12]
}

func (x *ServiceMsg) Str() string {
	switch x.Type {
	case ServiceMsg_ConnJoin:
		return fmt.Sprintf("[service conn-join-msg] id: %s, cid: %s, uid: %s, join-service: %v, rids: %v", x.Id, x.JoinData.Cid, x.JoinData.Uid, x.JoinData.JoinService, x.JoinData.Rooms)
	case ServiceMsg_ConnQuit:
		return fmt.Sprintf("[service conn-quit-msg] id: %s, cid: %s, uid: %s, quit-service: %v, rids: %v", x.Id, x.QuitData.Cid, x.QuitData.Uid, x.QuitData.QuitService, x.QuitData.Rooms)
	case ServiceMsg_Biz:
		return fmt.Sprintf("[service biz-msg] id: %s, cid: %s, uid: %s, quit-service: %v, rids: %v", x.Id, x.QuitData.Cid, x.QuitData.Uid, x.QuitData.QuitService, x.QuitData.Rooms)
	}
	return fmt.Sprintf("id: %s, typ: %d", x.Id, x.Type)
}
