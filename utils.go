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

const milliunit = int64(time.Millisecond)

func getNowMillisecs() int64 {
	return time.Now().UnixNano() / milliunit
}

func GenMsgId() string {
	s, _ := uuid.GenerateUUID()
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))[:12]
}
