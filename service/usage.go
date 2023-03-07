package service

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"path/filepath"
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
	} else {
		diskInfo, err := disk.Usage("/")
		if err != nil {
			log.Println("get disk usage error:", err)
			return usageInfo
		}
		//fmt.Printf("total space: %v bytes\n", diskInfo.Total)
		//fmt.Printf("free space: %v bytes\n", diskInfo.Free)
		//fmt.Printf("used space: %v bytes\n", diskInfo.Used)
		//fmt.Printf("usage percent: %.2f%%\n", diskInfo.UsedPercent)
		//log.Println("current platform is linux or osx")
		diskUsage := strconv.FormatFloat(diskInfo.UsedPercent, 'f', 0, 64)
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
