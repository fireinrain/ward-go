package service

import (
	"fmt"
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
