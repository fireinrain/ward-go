package service

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"net"
	"testing"
	"time"
)

func TestGetServerInfoService(t *testing.T) {
	service := GetServerInfoService()
	info, err := host.Info()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", info)

	fmt.Printf("service: %v\n", service)

}

func TestConvertUptime2Seperate(t *testing.T) {
	seperate := ConvertUptime2Seperate(1029499)
	for _, value := range seperate {
		fmt.Println(value)
	}
}

func TestGetMainHardDriveDiskName(t *testing.T) {
	name := GetMainHardDriveInfo()
	fmt.Println(name)
}

func TestGetMacosDiskCountAndTotalSize(t *testing.T) {
	size, i := GetMacosDiskCountAndTotalSize()
	fmt.Println(size, i)
}

func TestGetServerInfo(t *testing.T) {
	_ = &ServerInfo{}
	serverInfo := GetServerInfo()
	fmt.Println(serverInfo.Storage.MainStorage)

}

func TestMapKeyExist(t *testing.T) {
	dict := map[string]int{"key1": 1, "key2": 2}
	if value, ok := dict["key1"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("key1 不存在")
	}
}

func TestGetMachineAllIps(t *testing.T) {
	ips, err := GetMachineAllIps()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", ips)
}

func TestGetIPaddress(t *testing.T) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ipnet.IP.IsLoopback() {
				continue
			}
			if ipnet.IP.To4() == nil {
				continue
			}
			fmt.Println(ipnet.IP.String())
		}
	}

}

func TestGetIp(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

func TestGetMachineLocalIP(t *testing.T) {
	ispip, err := GetMachineLocalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ispip)
}

func TestGetMachineISPIP(t *testing.T) {
	ispip, _ := GetMachineISPIP()
	fmt.Println(ispip)
}

func TestGetCurrentDateTime(t *testing.T) {
	currentTime := time.Now()
	formatDateTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(formatDateTime)
}
