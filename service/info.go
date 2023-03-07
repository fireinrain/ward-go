package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"log"
	"strconv"
	"strings"
	"ward-go/config"
)

type ServerInfo struct {
	Processor Processor `json:"processor"`
	Machine   Machine   `json:"machine"`
	Storage   Storage   `json:"storage"`
	Uptime    Uptime    `json:"uptime"`
	Setup     Setup     `json:"setup"`
	Project   Project   `json:"project"`
}
type Processor struct {
	Name       string `json:"name"`
	CoreCount  string `json:"coreCount"`
	ClockSpeed string `json:"clockSpeed"`
	BitDepth   string `json:"bitDepth"`
}
type Machine struct {
	OperatingSystem     string `json:"operatingSystem"`
	TotalRAM            string `json:"totalRam"`
	RAMTypeOrOSBitDepth string `json:"ramTypeOrOSBitDepth"`
	ProcCount           string `json:"procCount"`
}
type Storage struct {
	MainStorage string `json:"mainStorage"`
	Total       string `json:"total"`
	DiskCount   string `json:"diskCount"`
	SwapAmount  string `json:"swapAmount"`
}
type Uptime struct {
	Days    string `json:"days"`
	Hours   string `json:"hours"`
	Minutes string `json:"minutes"`
	Seconds string `json:"seconds"`
}
type Setup struct {
	ServerName string `json:"serverName"`
}
type Project struct {
	Version string `json:"version"`
}

// GetServerInfoService
//
//	@Description: 获取服务器状态指标
//	@return ServerInfo
func GetServerInfoService() ServerInfo {
	serverInfo := ServerInfo{
		Processor: Processor{
			Name:       "",
			CoreCount:  "",
			ClockSpeed: "",
			BitDepth:   "",
		},
		Machine: Machine{
			OperatingSystem:     "",
			TotalRAM:            "",
			RAMTypeOrOSBitDepth: "",
			ProcCount:           "",
		},
		Storage: Storage{
			MainStorage: "",
			Total:       "",
			DiskCount:   "",
			SwapAmount:  "",
		},
		Uptime: Uptime{
			Days:    "",
			Hours:   "",
			Minutes: "",
			Seconds: "",
		},
		Setup: Setup{
			ServerName: config.GlobalConfig.Setup.ServerName,
		},
		Project: Project{
			Version: "v1.0",
		},
	}
	//获取cpu信息
	info, err := cpu.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
		return serverInfo
	}
	if len(info) == 0 {
		log.Println("no cpu info available...")
		return serverInfo
	}
	//cpu型号
	cpuName := info[0].ModelName
	cpuNameSplit := strings.Split(cpuName, "@")
	cpuName = cpuNameSplit[0]
	serverInfo.Processor.Name = cpuName
	//cpu 频率
	mhz := info[0].Mhz
	ghz := mhz / 1000.0
	ghzStr := fmt.Sprintf("%.1f", ghz)
	serverInfo.Processor.ClockSpeed = ghzStr

	//cpu架构
	stat, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
		return serverInfo
	}
	serverInfo.Processor.BitDepth = stat.KernelArch
	//cpu 核心
	counts, err := cpu.Counts(true)
	if err != nil {
		log.Println("error getting cpu counts: ", err)
		return serverInfo
	}
	serverInfo.Processor.CoreCount = strconv.Itoa(counts)

	//获取machine相关信息

	//获取存储相关信息

	//获取启动时长信息

	return serverInfo

}

func ConvertUptime2Seperate(uptime uint64) []uint64 {
	if uptime == 0 {
		return []uint64{0, 0, 0, 0}
	}
	var result []uint64
	//days, hours, minutes seconds

	dayCount := uptime / (60 * 60 * 24)

	result = append(result, dayCount)

	dayLeftSecond := uptime % (60 * 60 * 24)

	hourCount := dayLeftSecond / (60 * 60)

	result = append(result, hourCount)

	hourLeftSecond := dayLeftSecond % (60 * 60)

	minuteCount := hourLeftSecond / (60)

	result = append(result, minuteCount)

	minuteLeftSecond := hourLeftSecond % (60)

	result = append(result, minuteLeftSecond)

	return result
}
