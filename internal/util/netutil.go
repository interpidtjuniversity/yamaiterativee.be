package util

import "net"

func GetLocalIPv4Address() (ipv4Address string) {
	//获取所有网卡
	addrs, _ := net.InterfaceAddrs()
	//遍历
	for _, addr := range addrs {
		//取网络地址的网卡的信息
		ipNet, isIpNet := addr.(*net.IPNet)
		//是网卡并且不是本地环回网卡
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			//能正常转成ipv4
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}
	return
}
