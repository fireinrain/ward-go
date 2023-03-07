package service

import (
	"errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const PlatformWindows = 0
const PlatformLinux = 1
const PlatformOSX = 2

// ServerUsageInfo
// @Description: 服务器资源使用信息
type ServerUsageInfo struct {
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
}

// GetUsageInfoService
//
//	@Description: 获取服务器资源使用信息
//	@return ServerUsageInfo
func GetUsageInfoService() ServerUsageInfo {
	usageInfo := ServerUsageInfo{
		Processor: "0",
		Ram:       "0",
		Storage:   "0",
	}
	// cpu usage
	percent, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("get cpu percent error:", err)
		return usageInfo
	}
	cpuUsage := strconv.FormatFloat(percent[0], 'f', 0, 64)
	usageInfo.Processor = cpuUsage
	// ram usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Println("get ram usage error:", err)
		return usageInfo
	}
	ramUsage := strconv.FormatFloat(memInfo.UsedPercent, 'f', 0, 64)
	usageInfo.Ram = ramUsage
	//fmt.Printf("total ram: %v bytes\n", memInfo.Total)
	//fmt.Printf("available ram: %v bytes\n", memInfo.Available)
	//fmt.Printf("used ram: %v bytes\n", memInfo.Used)
	//fmt.Printf("ram used percent: %.2f%%\n", memInfo.UsedPercent)

	// disk usage
	if runtime.GOOS == "windows" {
		//log.Println("current platform is windows")
		//获取所有的磁盘 然后计算总usage
	}
	if runtime.GOOS == "linux" {

	}
	if runtime.GOOS == "darwin" {
		diskPaths, err := GetPlatformDiskPaths(2)
		if err != nil {
			log.Println("get disk path error:", err)
		}
		// diff to the real diskPaths
		diskPaths = ExtractRealDiskPath(diskPaths)

		var realDiskPath []string
		for _, diskPath := range diskPaths {
			path, err := GetDiskMountedPath(diskPath)
			if err != nil {
				//排除掉无法mount的路径
				continue
			}
			realDiskPath = append(realDiskPath, path)
		}

		var diskTotal uint64 = 0
		var diskUsed uint64 = 0
		for _, path := range realDiskPath {
			diskInfo, err := disk.Usage(path)

			if err != nil {
				log.Println("get disk usage error:", err)
				return usageInfo
			}
			diskTotal += diskInfo.Total
			diskUsed += diskInfo.Used
		}
		diskUsage := strconv.FormatFloat((float64(diskUsed)/float64(diskTotal))*100.0, 'f', 0, 64)
		//fmt.Println(diskUsage)
		//fmt.Printf("total space: %v bytes\n", diskInfo.Total)
		//fmt.Printf("free space: %v bytes\n", diskInfo.Free)
		//fmt.Printf("used space: %v bytes\n", diskInfo.Used)
		//fmt.Printf("usage percent: %.2f%%\n", diskInfo.UsedPercent)
		//log.Println("current platform is linux or osx")
		//diskUsage := strconv.FormatFloat(diskInfo.UsedPercent, 'f', 0, 64)
		usageInfo.Storage = diskUsage
	}

	return usageInfo
}

// GetPlatformDiskPaths
//
//	@Description: 获取平台所有挂载的disk路径
//	@param diskMountPath
//	@return []string
func GetPlatformDiskPaths(platform int) ([]string, error) {
	var disks []string

	//windows
	if platform == 0 {
		for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			drivePath := string(drive) + ":\\"
			_, err := os.Stat(drivePath)
			if err == nil {
				disks = append(disks, drivePath)
			}
		}
	}
	//linux
	if platform == 1 {
		err := filepath.Walk("/sys/block", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && strings.HasPrefix(info.Name(), "sd") {
				disks = append(disks, "/dev/"+info.Name())
			}
			return nil
		})

		if err != nil {
			if err != nil {
				log.Println("get platform all disk path error: ", err)
				return disks, err
			}
		}
	}
	//macos
	if platform == 2 {
		err := filepath.Walk("/dev", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//跳过disk0
			if strings.Contains(info.Name(), "disk0") {
				return nil
			}
			if info.Mode()&os.ModeDevice != 0 && strings.HasPrefix(info.Name(), "disk") {
				disks = append(disks, "/dev/"+info.Name())
			}
			return nil
		})

		if err != nil {
			if strings.HasSuffix(err.Error(), "bad file descriptor") {
				err = nil
			} else {
				log.Println("get platform all disk path error: ", err)
				return disks, err
			}
		}
		//fmt.Println(disks)
	}
	//log.Println(disks)
	return disks, nil
}

// CalculateDiskUsage
//
//	@Description: 计算磁盘使用
//	@param diskPath
//	@return float64
func CalculateDiskUsage(diskPath []string) float64 {

	return 2.34
}

// GetDiskMountedPath
//
//	@Description: 获取磁盘的挂载点(macos)
//	@param diskPath
//	@return string
//	@return error
func GetDiskMountedPath(diskPath string) (string, error) {
	if diskPath == "" {
		log.Fatalf("diskPath cant be empty")
	}
	cmd := exec.Command("mount")
	var out strings.Builder
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Println("get disk mounted path error: ", err)
		return "", err
	}

	output := out.String()
	mountPoint := ""

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == diskPath {
			mountPoint = fields[2]
			break
		}
	}

	if mountPoint == "" {
		log.Println("disk is not currently mounted")
		return "", errors.New("disk is not currently mounted")
	}
	//fmt.Printf("disk is mounted at %s\n", mountPoint)
	return mountPoint, nil
}

// ExtractRealDiskPath
//
//	@Description: 抽取真是的disk path
//
// macos 上 同一个disk会被挂载到好几个path上
//
//	@param paths
//	@return []string
func ExtractRealDiskPath(paths []string) []string {
	var result []string
	m := map[string]byte{}

	// 定义一个正则表达式对象，匹配/dev/diskXsY格式的字符串
	r, _ := regexp.Compile(`^(/dev/disk\d+)s\d+$`)
	// 在输入字符串中查找匹配正则表达式的子字符串
	for _, path := range paths {
		match := r.FindStringSubmatch(path)
		if len(match) > 1 {
			m[match[1]] = '0'
			//fmt.Printf("Matched: %s\n", match[1]) // 输出第一个捕获组中的结果"/dev/disk1"
		}
	}
	for key := range m {
		key += "s1"
		result = append(result, key)
	}
	return result
}
