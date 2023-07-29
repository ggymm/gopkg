package utils

import "net"

// GetLocalIp 获取本机的IP地址
func GetLocalIp() string {
	var (
		err        error
		addrs      []net.Addr
		interfaces []net.Interface
	)

	// 1. 获取所有网卡
	interfaces, err = net.Interfaces()
	if err != nil {
		return ""
	}
	// 2. 获取所有网卡的IP地址
	for _, inter := range interfaces {
		addrs, err = inter.Addrs()
		if err != nil {
			return ""
		}

		// 3. 过滤掉虚拟网卡
		if inter.Flags&net.FlagUp == 0 {
			continue
		}

		// 5. 过滤掉回环地址
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				// 6. 过滤掉IPV6地址
				if ipNet.IP.To4() != nil {
					// 7. 返回IP地址
					return ipNet.IP.String()
				}
			}
		}
	}

	return ""
}

func GetLocalIpList(local bool) []string {
	var (
		ipList []string

		err        error
		addrs      []net.Addr
		interfaces []net.Interface
	)

	// 1. 获取所有网卡
	interfaces, err = net.Interfaces()
	if err != nil {
		return ipList
	}

	// 2. 获取所有网卡的IP地址
	for _, inter := range interfaces {
		addrs, err = inter.Addrs()
		if err != nil {
			return ipList
		}

		// 3. 过滤掉虚拟网卡
		if inter.Flags&net.FlagUp == 0 {
			continue
		}

		// 5. 过滤掉回环地址
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				// 6. 过滤掉IPV6地址
				if ipNet.IP.To4() != nil {
					// 7. 返回IP地址
					ipList = append(ipList, ipNet.IP.String())
				}
			}
		}
	}

	if local {
		ipList = append(ipList, "127.0.0.1")
		ipList = append(ipList, "localhost")
	}

	return ipList
}
