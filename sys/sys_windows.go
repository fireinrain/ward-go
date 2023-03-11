//go:build windows
// +build windows

package sys

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"os/exec"
	"strconv"
	"strings"
)

func GetTCPCount() (int, error) {
	stats, err := net.Connections("tcp")
	if err != nil {
		return 0, err
	}
	return len(stats), nil
	//return getActiveNetworkPorts("TCP")
}

func GetUDPCount() (int, error) {
	stats, err := net.Connections("udp")
	if err != nil {
		return 0, err
	}
	return len(stats), nil
	//return getActiveNetworkPorts("UDP")
}

// getActiveNetworkPorts
// issue: 在windows上耗时较长
//
//	@Description:
//	@param portType
//	@return int
//	@return error
func getActiveNetworkPorts(portType string) (int, error) {
	cmdStr := fmt.Sprintf("netstat -an | findstr '%s' | Measure-Object -Line | Select -ExpandProperty Lines", portType)
	out, err := exec.Command("powershell", cmdStr).Output()
	if err != nil {
		fmt.Printf("error running powershell: %v\n", err)
		return 0, errors.New("error running powershell: " + err.Error())
	}
	s := string(out)
	space := strings.TrimSpace(s)
	atoi, err := strconv.Atoi(space)
	return atoi, err
}
