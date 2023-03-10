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
	"time"
	"ward-go/config"
)

type ServerInfo struct {
	Processor Processor `json:"processor"`
	Machine   Machine   `json:"machine"`
	Storage   Storage   `json:"storage"`
	Network   Network   `json:"network"`
	Location  Location  `json:"location"`
	Uptime    Uptime    `json:"uptime"`
	Setup     Setup     `json:"setup"`
	Project   Project   `json:"project"`
	TimeStamp time.Time `json:"timestamp"`
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
type Network struct {
	UploadData     string `json:"uploadData"`
	DownloadData   string `json:"downloadData"`
	UploadSpeed    string `json:"uploadSpeed"`
	DownloadSpeed  string `json:"downloadSpeed"`
	TCPConnections string `json:"tcpConnections"`
	UDPConnections string `json:"udpConnections"`
}
type Location struct {
	Country         string `json:"country"`
	CountryCode     string `json:"countryCode"`
	CountryFlag     string `json:"countryFlag"`
	Timezone        string `json:"timezone"`
	TimezoneOffset  string `json:"timezoneOffset"`
	CurrentDateTime string `json:"currentDateTime"`
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
	serverInfo.TimeStamp = time.Now()
	if runtime.GOOS != "windows" && runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		log.Fatalf("unsupported OS: %s", runtime.GOOS)
	}
	if serverInfo.Project.Version == "" {
		serverInfo.Project.Version = config.WardGoVserion
	}
	if serverInfo.Setup.ServerName == "" {
		serverInfo.Setup.ServerName = config.GlobalConfig.Setup.ServerName
	}

	//??????cpu??????
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
		//cpu??????
		cpuName := info[0].ModelName
		cpuNameSplit := strings.Split(cpuName, "@")
		cpuName = strings.TrimSpace(cpuNameSplit[0])
		serverInfo.Processor.Name = cpuName
		//cpu ??????
		mhz := info[0].Mhz
		ghz := mhz / 1000.0
		ghzStr := fmt.Sprintf("%.1f", ghz)
		serverInfo.Processor.ClockSpeed = ghzStr + " GHz"

		//cpu??????
		host, err := host.Info()
		if err != nil {
			log.Println("error getting cpu info: ", err)
		}
		serverInfo.Processor.BitDepth = host.KernelArch
		//cpu ??????
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
		//??????
		memory, err := mem.VirtualMemory()
		if err != nil {
			log.Println("error getting memory info: ", err)
		}
		//1024.0 * 1024.0 * 1024.0 = 1073741824
		totalRam := float64(memory.Total) / 1073741824.0
		gRam := fmt.Sprintf("%.1f", totalRam)
		serverInfo.Machine.TotalRAM = gRam + "GiB Ram"
		// ??????swap??????
		swapInfo, err := mem.SwapMemory()
		if err != nil {
			log.Println("get swap memory failed:", err)
		}
		swapRam := float64(swapInfo.Total) / 1073741824.0
		swapGRam := fmt.Sprintf("%.1f", swapRam)
		serverInfo.Machine.SwapRAM = swapGRam + "GiB SwapRam"
		// ????????????
		ramType := GetMachineRamType()
		serverInfo.Machine.RAMTypeOrOSBitDepth = ramType
	}
	//????????????
	hostinfo, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
	}
	serverInfo.Machine.ProcCount = strconv.FormatUint(hostinfo.Procs, 10)

	//????????????????????????
	if serverInfo.Storage.MainStorage == "" ||
		serverInfo.Storage.DiskCount == "" ||
		serverInfo.Storage.Total == "" ||
		serverInfo.Storage.SwapAmount == "" {
		driveInfo := GetMainHardDriveInfo()
		serverInfo.Storage = driveInfo
	}
	//????????????
	serverInfo.Network = networkStatus.CurNetStatus.GetReadableStr()
	//????????????
	if serverInfo.Location.CountryFlag == "" ||
		serverInfo.Location.Country == "" ||
		serverInfo.Location.CountryCode == "" ||
		serverInfo.Location.Timezone == "" ||
		serverInfo.Location.TimezoneOffset == "" {
		//????????????isp ip
		ispip, err := GetMachineISPIP()
		if err != nil {
			log.Println("get machine isp ip error: " + err.Error())
		} else {
			geoLocation, err := GetIpgeolocationInfo(ispip)
			if err != nil {
				log.Println("get ipgeolocation info error: " + err.Error())
			} else {
				serverInfo.Location.Country = geoLocation.CountryName
				serverInfo.Location.CountryFlag = GetFlagEmojiSimple(geoLocation.CountryName)
				serverInfo.Location.CountryCode = geoLocation.CountryCode2
				serverInfo.Location.Timezone = geoLocation.TimeZone.Name
				serverInfo.Location.TimezoneOffset = strconv.Itoa(geoLocation.TimeZone.Offset)
			}
		}
	}
	//???????????????????????????????????????????????????
	if serverInfo.Location.Timezone == "" || serverInfo.Location.TimezoneOffset == "" {
		//????????????????????????????????????????????????
		currentTime := time.Now()
		formatDateTime := currentTime.Format("2006-01-02 15:04:05")
		serverInfo.Location.CurrentDateTime = formatDateTime
	} else {
		//?????????????????????????????????????????????????????????
		location, err := time.LoadLocation(serverInfo.Location.Timezone)
		if err != nil {
			log.Println("get location error: " + err.Error())
		} else {
			locationTime := time.Now().In(location)
			format := locationTime.Format("2006-01-02 15:04:05")
			serverInfo.Location.CurrentDateTime = format
		}
	}

	//????????????
	host, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
	}
	//????????????????????????
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
//	@Description: ???????????????????????????
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
	//??????cpu??????
	info, err := cpu.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
		return serverInfo
	}
	if len(info) == 0 {
		log.Println("no cpu info available...")
		return serverInfo
	}
	//cpu??????
	cpuName := info[0].ModelName
	cpuNameSplit := strings.Split(cpuName, "@")
	cpuName = strings.TrimSpace(cpuNameSplit[0])
	serverInfo.Processor.Name = cpuName
	//cpu ??????
	mhz := info[0].Mhz
	ghz := mhz / 1000.0
	ghzStr := fmt.Sprintf("%.1f", ghz)
	serverInfo.Processor.ClockSpeed = ghzStr + " GHz"

	//cpu??????
	host, err := host.Info()
	if err != nil {
		log.Println("error getting cpu info: ", err)
		return serverInfo
	}
	serverInfo.Processor.BitDepth = host.KernelArch
	//cpu ??????
	counts, err := cpu.Counts(true)
	if err != nil {
		log.Println("error getting cpu counts: ", err)
		return serverInfo
	}
	serverInfo.Processor.CoreCount = strconv.Itoa(counts) + " Cores"

	//??????machine????????????
	opSystem := fmt.Sprintf("%s %s,%s", host.Platform, host.PlatformVersion, host.PlatformFamily)
	serverInfo.Machine.OperatingSystem = opSystem
	//??????
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

	//????????????????????????
	driveInfo := GetMainHardDriveInfo()
	serverInfo.Storage = driveInfo

	//????????????????????????
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
//	@Description: ?????????????????? ?????????????????????????????????
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
//	@Description: ??????ram??????
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
			//???????????????
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
//	@Description: ??????????????????
//	@return string
func GetMainHardDriveInfo() Storage {
	storage := Storage{
		MainStorage: "Unknown",
		Total:       "Unknown",
		DiskCount:   "Unknown",
		SwapAmount:  "Unknown",
	}
	if runtime.GOOS == "windows" {
		//MediaType Model                               Size
		//--------- -----                               ----
		//HDD       HGST HTS721010A9E630       1000204886016
		//SSD       SAMSUNG MZNTY128HDHP-000H1  128035676160
		out, err := exec.Command("powershell", "Get-PhysicalDisk | Select-Object MediaType,Model,Size").Output()
		if err != nil {
			fmt.Printf("error running powershell: %v\n", err)
		} else {
			pattern := "[0-9]$"
			var diskPair [][]string
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				matched, _ := regexp.MatchString(pattern, line)
				if matched {
					diskInfo := strings.Split(line, " ")
					var tempDiskInfos []string
					for _, value := range diskInfo {
						if value == "" {
							continue
						}
						tempDiskInfos = append(tempDiskInfos, value)

					}
					disk := []string{strings.TrimSpace(tempDiskInfos[1]), strings.TrimSpace(tempDiskInfos[0]), strings.TrimSpace(tempDiskInfos[3])}
					diskPair = append(diskPair, disk)
				}
			}
			//??????????????????????????????
			var totalSize uint64
			for _, disk := range diskPair {
				uintVal, err := strconv.ParseUint(disk[2], 10, 64)
				if err != nil {
					log.Println("parse string to uint64 error: " + err.Error())
				} else {
					totalSize += uintVal
				}
			}
			//?????????
			uintVal, err := strconv.ParseUint(diskPair[0][2], 10, 64)
			if err != nil {
				log.Println("parse string to uint64 error: " + err.Error())
				//set as default if parse string error
				uintVal = 0
			}
			mainStorageSize := FindPropertyNetUnitStr(uintVal)
			storage.MainStorage = diskPair[0][0] + "-" + diskPair[0][1] + "-" + mainStorageSize
			storage.DiskCount = strconv.Itoa(len(diskPair)) + " Disks"
			storage.Total = FindPropertyNetUnitStr(totalSize)
			storage.SwapAmount = ""

			return storage
		}
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
//	@Description: ??????macos????????????
//	@return int
func GetMacosDiskCountAndTotalSize() (int, float64) {
	// ???????????????????????????????????????????????????????????????
	diskCount := 0
	//GB ?????? ??????????????????
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
			//????????????
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
//	@Description: ???????????????????????????GB ?????????????????????
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
