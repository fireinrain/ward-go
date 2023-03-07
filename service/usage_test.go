package service

import (
	"fmt"
	"testing"
)

func TestGetPlatformDiskPaths(t *testing.T) {

	paths, err := GetPlatformDiskPaths(2)
	if err != nil {
		fmt.Println("error: ", err)
	}
	for _, valued := range paths {
		fmt.Printf("path: %s\n", valued)

	}
}

func TestCalculateDiskUsage(t *testing.T) {
	paths, err := GetPlatformDiskPaths(2)
	if err != nil {
		fmt.Println("error: ", err)
	}
	for _, valued := range paths {
		fmt.Printf("path: %s\n", valued)

		//readlink, err := os.Readlink(valued)
		//if err != nil {
		//	fmt.Println("Error:", err)
		//}
		//fmt.Println("link target:", readlink)

		//diskInfo, err := disk.Usage(valued)
		//if err != nil {
		//	log.Println("get disk usage error:", err)
		//}
		//fmt.Printf("total space: %v bytes\n", diskInfo.Total)
		//fmt.Printf("free space: %v bytes\n", diskInfo.Free)
		//fmt.Printf("used space: %v bytes\n", diskInfo.Used)
		//fmt.Printf("usage percent: %.2f%%\n", diskInfo.UsedPercent)
		path, err := GetDiskMountedPath(valued)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("mount path:", path)

		fmt.Println("--------------")
	}

}
