package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os/exec"
	"regexp"
	"runtime"
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
	SwapRAM             string `json:"swapRam"`
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

var serverInfo = &ServerInfo{}

func GetServerInfo() ServerInfo {
	if runtime.GOOS != "windows" && runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		log.Fatalf("unsupported OS: %s", runtime.GOOS)
	}
	if serverInfo.Project.Version == "" {
		serverInfo.Project.Version = config.WardGoVserion
	}
	if serverInfo.Setup.ServerName == "" {
		serverInfo.Setup.ServerName = config.GlobalConfig.Setup.ServerName
	}

	//获取cpu信息
	if serverInfo.Processor.Name == "" ||
		serverInfo.Processor.ClockSpeed == "" ||
		serverInfo.Processor.CoreCount == "" ||
		serverInfo.Processor.BitDepth == "" {
		info, err := cpu.Info()
		if err != nil {
			log.Println("error getting cpu info: ", err)
		}
		if len(info) == 0 {
			log.Println("no cpu info available...")
		}
		//cpu型号
		cpuName := info[0].ModelName
		cpuNameSplit := strings.Split(cpuName, "@")
		cpuName = strings.TrimSpace(cpuNameSplit[0])
		serverInfo.Processor.Name = cpuName
		//cpu 频率
		mhz := info[0].Mhz
		ghz := mhz / 1000.0
		ghzStr := fmt.Sprintf("%.1f", ghz)
		serverInfo.Processor.ClockSpeed = ghzStr + " GHz"

		//cpu架构
		host, err := host.Info()
		if err != nil {
			log.Println("error getting cpu info: ", err)
		}
		serverInfo.Processor.BitDepth = host.KernelArch
		//cpu 核心
		counts, err := cpu.Counts(true)
		if err != nil {
			log.Println("error getting cpu counts: ", err)
		}
		serverInfo.Processor.CoreCount = strconv.Itoa(counts) + " Cores"

	}

	if serverInfo.Machine.TotalRAM == "" ||
		serverInfo.Machine.SwapRAM == "" ||
		serverInfo.Machine.RAMTypeOrOSBitDepth == "" ||
		serverInfo.Machine.OperatingSystem == "" {
		hostinfo, err := host.Info()
		if err != nil {
			log.Println("error getting cpu info: ", err)
		}
		opSystem := fmt.Sprintf("%s %s,%s", hostinfo.Platform, hostinfo.PlatformVersion, hostinfo.PlatformFamily)
		serverInfo.Machine.OperatingSystem = opSystem
		//内存
		memory, err := mem.VirtualMemory()
		if err != nil {
			log.Println("error getting memory info: ", err)
		}
		//1024.0 * 1024.0 * 1024.0 = 1073741824
		totalRam := float64(memory.Total) / 1073741824.0
		gRam := fmt.Sprintf("%.1f", totalRam)
		serverInfo.Machine.TotalRAM = gRam + "GiB Ram"
		// 内存swap大小
		swapInfo, err := mem.SwapMemory()
		if err != nil {
			log.Println("get swap memory failed:", err)
		}
		swapRam := float64(swapInfo.Total) / 1073741824.0
		swapGRam := fmt.Sprintf("%.1f", swapRam)
		serverInfo.Machine.SwapRAM = swapGRam + "GiB SwapRam"
		// 内存类型
		ramType := GetMachineRamType()
		serverInfo.Machine.RAMTypeOrOSBitDepth = ramType
	}
	//进程数量
	hostinfo, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
	}
	serverInfo.Machine.ProcCount = strconv.FormatUint(hostinfo.Procs, 10)

	//获取存储相关信息
	if serverInfo.Storage.MainStorage == "" ||
		serverInfo.Storage.DiskCount == "" ||
		serverInfo.Storage.Total == "" ||
		serverInfo.Storage.SwapAmount == "" {
		driveInfo := GetMainHardDriveInfo()
		serverInfo.Storage = driveInfo
	}
	//动态变化
	host, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
	}
	//获取启动时长信息
	uptime := host.Uptime
	uptime2Seperate := ConvertUptime2Seperate(uptime)
	serverInfo.Uptime.Days = strconv.FormatUint(uptime2Seperate[0], 10)
	serverInfo.Uptime.Hours = strconv.FormatUint(uptime2Seperate[1], 10)
	serverInfo.Uptime.Minutes = strconv.FormatUint(uptime2Seperate[2], 10)
	serverInfo.Uptime.Seconds = strconv.FormatUint(uptime2Seperate[3], 10)

	return *serverInfo
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
			SwapRAM:             "",
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
	cpuName = strings.TrimSpace(cpuNameSplit[0])
	serverInfo.Processor.Name = cpuName
	//cpu 频率
	mhz := info[0].Mhz
	ghz := mhz / 1000.0
	ghzStr := fmt.Sprintf("%.1f", ghz)
	serverInfo.Processor.ClockSpeed = ghzStr + " GHz"

	//cpu架构
	host, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
		return serverInfo
	}
	serverInfo.Processor.BitDepth = host.KernelArch
	//cpu 核心
	counts, err := cpu.Counts(true)
	if err != nil {
		log.Println("error getting cpu counts: ", err)
		return serverInfo
	}
	serverInfo.Processor.CoreCount = strconv.Itoa(counts) + " Cores"

	//获取machine相关信息
	opSystem := fmt.Sprintf("%s %s,%s", host.Platform, host.PlatformVersion, host.PlatformFamily)
	serverInfo.Machine.OperatingSystem = opSystem
	//内存
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Println("error getting memory info: ", err)
		return serverInfo
	}
	//1024.0 * 1024.0 * 1024.0 = 1073741824
	totalRam := float64(memory.Total) / 1073741824.0
	gRam := fmt.Sprintf("%.1f", totalRam)
	serverInfo.Machine.TotalRAM = gRam + "GiB Ram"
	//
	ramType := GetMachineRamType()
	serverInfo.Machine.RAMTypeOrOSBitDepth = ramType
	serverInfo.Machine.ProcCount = strconv.FormatUint(host.Procs, 10)

	//获取存储相关信息
	driveInfo := GetMainHardDriveInfo()
	serverInfo.Storage = driveInfo

	//获取启动时长信息
	uptime := host.Uptime
	uptime2Seperate := ConvertUptime2Seperate(uptime)
	serverInfo.Uptime.Days = strconv.FormatUint(uptime2Seperate[0], 10)
	serverInfo.Uptime.Hours = strconv.FormatUint(uptime2Seperate[1], 10)
	serverInfo.Uptime.Minutes = strconv.FormatUint(uptime2Seperate[2], 10)
	serverInfo.Uptime.Seconds = strconv.FormatUint(uptime2Seperate[3], 10)

	return serverInfo
}

// ConvertUptime2Seperate
//
//	@Description: 将启动时间秒 转化为天小时分钟秒格式
//	@param uptime
//	@return []uint64
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

// GetMachineRamType
//
//	@Description: 获取ram类型
//	@return string
func GetMachineRamType() string {
	ramType := "Unknown"

	if runtime.GOOS == "windows" {
		out, err := exec.Command("powershell", "wmic", "memorychip", "get SMBIOSMemoryType").Output()
		if err != nil {
			fmt.Printf("error running powershell: %v\n", err)
			return ramType
		}

		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			//转换为数字
			ramTypeNum, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				continue
			}
			if ramTypeNum == 21 {
				ramType = "DDR2"
				break
			}
			if ramTypeNum == 22 {
				ramType = "DDR2 FB-DIMM"
				break
			}
			if ramTypeNum == 24 {
				ramType = "DDR3"
				break
			}
			if ramTypeNum == 26 {
				ramType = "DDR4"
				break
			}
		}
		return ramType
	}
	if runtime.GOOS == "linux" {
		out, err := exec.Command("dmidecode", "-t", "memory").Output()
		if err != nil {
			fmt.Printf("error running dmidecode: %v\n", err)
			return ramType
		}

		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Type:") && strings.Contains(line, "DDR") {
				ramType = strings.TrimSpace(strings.Split(line, ":")[1])
				break
			}
		}
		return ramType
	}
	if runtime.GOOS == "darwin" {
		out, err := exec.Command("system_profiler", "SPMemoryDataType").Output()
		if err != nil {
			fmt.Printf("error running system_profiler: %v\n", err)
			return ramType
		}
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Type:") && strings.Contains(line, "DDR") {
				ramType = strings.TrimSpace(strings.Split(line, ":")[1])
				break
			}
		}
		return ramType
	}
	return ramType
}

// GetMainHardDriveInfo
//
//	@Description: 获取主存储名
//	@return string
func GetMainHardDriveInfo() Storage {
	storage := Storage{
		MainStorage: "Unknown",
		Total:       "Unknown",
		DiskCount:   "Unknown",
		SwapAmount:  "Unknown",
	}
	if runtime.GOOS == "windows" {

	}
	if runtime.GOOS == "linux" {

	}
	if runtime.GOOS == "darwin" {
		out, err := exec.Command("diskutil", "info", "/").Output()
		var mainStorage string
		if err != nil {
			fmt.Printf("error running diskutil: %v\n", err)
		} else {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.Contains(line, "Volume Name:") {
					segments := strings.Split(line, ":")
					mainStorage = strings.TrimSpace(segments[len(segments)-1])
					continue
				}

				//ssd
				if strings.Contains(line, "Solid State:") {
					segments := strings.Split(line, ":")
					isSSD := strings.TrimSpace(segments[len(segments)-1])
					if isSSD == "Yes" {
						mainStorage = mainStorage + "-SSD"
					} else {
						mainStorage = mainStorage + "-HDD"
					}
					storage.MainStorage = mainStorage
					continue
				}
			}
		}
		//get mainstorage swap size
		output2, err2 := exec.Command("sysctl", "vm.swapusage").Output()
		if err2 != nil {
			fmt.Printf("error running sysctl: %v\n", err)
		} else {
			lines := strings.Split(string(output2), "\n")
			for _, line := range lines {
				if strings.Contains(line, "total = ") {
					segments := strings.Split(line, " ")
					swapSize := strings.TrimSpace(segments[3])
					swapSize = strings.Replace(swapSize, "M", "", 1)
					//convertGB
					swapSizeFloat, err := strconv.ParseFloat(swapSize, 2)
					if err != nil {
						fmt.Printf("convert swap size error: %s", err)
					} else {
						f := swapSizeFloat / 1024.0
						result := fmt.Sprintf("%.1f", f)
						storage.SwapAmount = result + " Gib Swap"
						break
					}
				}
			}
		}
		diskCount, totalSize := GetMacosDiskCountAndTotalSize()
		storage.DiskCount = strconv.Itoa(diskCount) + " Disks"
		storage.Total = strconv.FormatFloat(totalSize, 'f', 1, 64) + " Gib Total"
	}

	return storage
}

// GetMacosDiskCountAndTotalSize
//
//	@Description: 获取macos磁盘数量
//	@return int
func GetMacosDiskCountAndTotalSize() (int, float64) {
	// 在输入字符串中查找匹配正则表达式的子字符串
	diskCount := 0
	//GB 格式 保留一位小数
	totalSizeCount := 0.0

	re := regexp.MustCompile(`[\*\+]\s*(\d+\.\d+\s*[KMGTP]B)`)

	output, err := exec.Command("diskutil", "list").Output()
	if err != nil {
		fmt.Printf("error running sysctl: %v\n", err)
	} else {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "/dev/disk") {
				if strings.HasPrefix(line, "/dev/disk0") {
					continue
				}
				diskCount += 1
			}
			//磁盘容量
			if strings.HasPrefix(line, "   0:") {
				if strings.Contains(line, "GUID_partition_scheme") {
					continue
				}
				match := re.FindStringSubmatch(line)
				diskSizeStr := match[1]
				size, err := Convert2EqualGbSize(diskSizeStr)
				if err == nil {
					totalSizeCount += size
				}
			}
		}
	}
	return diskCount, totalSizeCount
}

// Convert2EqualGbSize
//
//	@Description: 将容量字符串转换为GB 保留小数点一位
//	@param sizeStr
//	@return float64
func Convert2EqualGbSize(sizeStr string) (float64, error) {
	sizeStr = strings.TrimSpace(sizeStr)
	suffixes := map[string]float64{
		"KB": 1 / (1024.0 * 1024),
		"MB": 1 / 1024,
		"GB": 1,
		"TB": 1024,
		"PB": 1024 * 1024,
	}
	for suffix, factor := range suffixes {
		if strings.HasSuffix(sizeStr, suffix) {
			trimSuffix := strings.TrimSuffix(sizeStr, suffix)
			trimSuffix = strings.TrimSpace(trimSuffix)
			value, err := strconv.ParseFloat(trimSuffix, 64)
			if err != nil {
				return 0, err
			}
			return value * factor, nil
		}
	}
	return 0, fmt.Errorf("invalid size suffix")
}
