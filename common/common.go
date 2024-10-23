package common

import (
	"fmt"
	"net"
	"strings"
)

// 收集日志项
type CollectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// GetOutboundIP 获取本机IP的函数
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "notgetip", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
