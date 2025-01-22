package utils

import (
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"net"
)

func GetIPList() (ipList []string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		logrus.Error(err)
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ipv4 := ipNet.IP.To4()
			if ipv4 == nil {
				continue
			}
			ipList = append(ipList, ipv4.String())
		}
	}
	return ipList
}

func PrintSystemInfo() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, i := range ipList {
			global.Logger.Infof("[gvb]  backend 运行在 http://%s:%d/api", i, port)
			global.Logger.Infof("[gvb]  swagger运行在 http://%s:%d/swagger/index.html#", i, port)
		}
	} else {
		global.Logger.Infof("[gvb]  backend 运行在 http://%s:%d/api", ip, port)
		global.Logger.Infof("[gvb]  swagger运行在 http://%s:%d/swagger/index.html#", ip, port)
	}

}
