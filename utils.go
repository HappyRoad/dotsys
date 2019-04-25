package dotsys

import (
	"net"
	"strings"
)

func getInnerIp() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				if strings.HasPrefix(ip, "192.168") {
					return
				}
			}
		}
	}
	return
}
