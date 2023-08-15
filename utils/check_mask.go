package utils

import (
	"net"
)

// 判断是否在白名单子网范围内
func InWhiteSubnets(targetIp string, whiteSubnets []string) (result bool) {
	for _, subnet := range whiteSubnets {
		_, ipNet, _ := net.ParseCIDR(subnet)
		if ipNet.Contains(net.ParseIP(targetIp)) {
			// fmt.Println(targetIp + "is in white subnets")
			result = true
			break
		} else {
			// fmt.Println(targetIp + "is not in white subnets")
			result = false
		}
	}
	return result
}

// 判断是否是白名单地址内
func InWhiteAddrs(targetIp string, whiteAddrs []string) (result bool) {
	for _, addr := range whiteAddrs {
		if addr == targetIp {
			result = true
			break
		} else {
			result = false
		}
	}
	return result
}
