package service

import (
	"github.com/shirou/gopsutil/net"
	"log"
	"strconv"
	"time"
	"ward-go/sys"
)

// 单位换算
var unitConversion = []NetUnit{
	{
		Name:  "B",
		Value: 1.0,
	},
	{
		Name:  "KB",
		Value: 1024.0,
	},
	{
		Name:  "MB",
		Value: 1024 * 1024.0,
	},
	{
		Name:  "GB",
		Value: 1024 * 1024 * 1024.0,
	},
	{
		Name:  "TB",
		Value: 1024 * 1024 * 1024 * 1024.0,
	},
	{
		Name:  "PB",
		Value: 1024 * 1024 * 1024 * 1024 * 1024.0,
	},
}

type NetUnit struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type NetworkStatus struct {
	PreNetStatus NetStatus `json:"preNetStatus"`
	CurNetStatus NetStatus `json:"curNetStatus"`
}

type NetStatus struct {
	//上传数据大小 单位都是byte
	UploadDataSize uint64 `json:"uploadDataSize"`
	//下载数据大小
	DownloadDataSize uint64 `json:"downloadDataSize"`
	//上传速度 单位也是byte/s
	UploadSpeed uint64 `json:"uploadSpeed"`
	//下载速度 单位byte/s
	DownloadSpeed uint64 `json:"downloadSpeed"`
	//tcp链接数
	TCPConnections int `json:"tcpConnections"`
	//udp连接数
	UDPConnections int `json:"udpConnections"`
}

var networkStatus = &NetworkStatus{}

func init() {
	ioStats, err := net.IOCounters(false)
	if err != nil {
		log.Println("get io counters failed:", err)
	} else if len(ioStats) > 0 {
		ioStat := ioStats[0]
		networkStatus.PreNetStatus.UploadDataSize = ioStat.BytesSent
		networkStatus.PreNetStatus.DownloadDataSize = ioStat.BytesRecv

	} else {
		log.Println("can not find io counters")
	}
	go recordNetworkStatus(networkStatus)
}

// recordNetworkStatus
//
//	@Description: 记录网络状态数据
//	@param network
func recordNetworkStatus(network *NetworkStatus) {
	for {
		now := time.Now()
		time.Sleep(1 * time.Second)
		ioStats, err := net.IOCounters(false)
		if err != nil {
			log.Println("get io counters failed:", err)
		} else if len(ioStats) > 0 {
			ioStat := ioStats[0]
			network.CurNetStatus.UploadDataSize = ioStat.BytesSent
			network.CurNetStatus.DownloadDataSize = ioStat.BytesRecv

			duration := time.Now().Sub(now)
			seconds := float64(duration) / float64(time.Second)
			up := uint64(float64(network.CurNetStatus.UploadDataSize-network.PreNetStatus.UploadDataSize) / seconds)
			down := uint64(float64(network.CurNetStatus.DownloadDataSize-network.PreNetStatus.DownloadDataSize) / seconds)
			network.CurNetStatus.UploadSpeed = up
			network.CurNetStatus.DownloadSpeed = down

			status := network.CurNetStatus
			network.PreNetStatus = status
		} else {
			log.Println("can not find io counters")
		}

		tcpCount, err := sys.GetTCPCount()
		network.CurNetStatus.TCPConnections = tcpCount
		if err != nil {
			log.Println("get tcp connections failed:", err)
		}

		udpCount, err := sys.GetUDPCount()
		network.CurNetStatus.UDPConnections = udpCount
		if err != nil {
			log.Println("get udp connections failed:", err)
		}
	}
}

// GetReadableStr
//
//	@Description: 获取可读的网络数据表示
//	@receiver network
//	@return Network
func (network *NetStatus) GetReadableStr() Network {
	return Network{
		UploadData:     FindPropertyNetUnitStr(network.UploadDataSize),
		DownloadData:   FindPropertyNetUnitStr(network.DownloadDataSize),
		UploadSpeed:    FindPropertyNetUnitStr(network.UploadSpeed) + "/S",
		DownloadSpeed:  FindPropertyNetUnitStr(network.DownloadSpeed) + "/S",
		TCPConnections: strconv.Itoa(network.TCPConnections),
		UDPConnections: strconv.Itoa(network.UDPConnections),
	}
}

// FindPropertyNetUnitStr
//
//	@Description: 寻找合适的网络数据单位
//	@param netUnitUint
//	@return string
func FindPropertyNetUnitStr(netUnitUint uint64) string {
	f := float64(netUnitUint)
	for index, value := range unitConversion {
		bigValue := f / value.Value
		if index == len(unitConversion)-1 {
			return strconv.FormatFloat(bigValue, 'f', 2, 64) + " " + unitConversion[index].Name
		}
		smallValue := f / unitConversion[index+1].Value

		if bigValue > 0 && smallValue <= 1 {
			result := strconv.FormatFloat(bigValue, 'f', 2, 64)
			//fmt.Println(result)
			return result + " " + unitConversion[index].Name
		}
	}
	return strconv.FormatUint(netUnitUint, 10) + " " + unitConversion[0].Name
}
