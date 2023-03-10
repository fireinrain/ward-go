//go:build windows
// +build windows

package sys

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetTCPCount() (int, error) {
	return getActiveNetworkPorts("TCP")
}

func GetUDPCount() (int, error) {
	return getActiveNetworkPorts("UDP")
}

func getActiveNetworkPorts(portType string) (int error) {
	cmdStr := fmt.Sprintf("powershell networkstat -an | findstr %s | Measure-Object -Line | Select -ExpandProperty Lines", portType)
	out, err := exec.Command(cmdStr).Output()
	if err != nil {
		fmt.Printf("error running powershell: %v\n", err)
		return 0, errors.New("error running powershell: " + err)
	}
	s := string(out)
	space := strings.TrimSpace(s)
	atoi, err := strconv.Atoi(space)
	return atoi, err
}
