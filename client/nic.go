package client

import (
	"fmt"
	"net"
	"strings"
)

var excludeNicNamePrefix = []string{"docker", "veth", "lo", "br-"}

func exclude(name string) bool {
	for _, prefix := range excludeNicNamePrefix {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func GetNicInfo() (nicArray []map[string]string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, iface := range ifaces {
		if exclude(iface.Name) {
			continue
		}
		nic := make(map[string]string)
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			if ip = ip.To4(); ip == nil {
				continue
			}
			nic["address"] = ip.String()
			nic["mac"] = iface.HardwareAddr.String()
			nic["name"] = iface.Name
			nicArray = append(nicArray, nic)
		}
	}
	return
}
