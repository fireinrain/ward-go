package service

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetReadableStr(t *testing.T) {
	status := NetStatus{
		UploadDataSize:   1843421802,
		DownloadDataSize: 3600899960,
		UploadSpeed:      6218,
		DownloadSpeed:    6131,
		TCPConnections:   12,
		UDPConnections:   35,
	}
	str := status.GetReadableStr()
	fmt.Println(str)
}

func TestGetFloat(t *testing.T) {
	a := float64(1758)
	b := float64(1024)
	f := a / b
	fmt.Println(f)
	float := strconv.FormatFloat(f, 'f', 1, 64)
	fmt.Println(float)
}

func TestFindPropertyNetUnitStr(t *testing.T) {
	str := FindPropertyNetUnitStr(123456778)
	fmt.Println(str)
}
