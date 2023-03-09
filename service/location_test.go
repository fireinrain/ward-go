package service

import (
	"fmt"
	"testing"
)

func TestGetLocationInfoByIPv4(t *testing.T) {
	pv4, err := GetLocationInfoByIPv4("216.127.164.234")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", *pv4)
}

func TestGetLocationInfoByIpApi(t *testing.T) {
	api, err := GetLocationInfoByIpApi("208.95.112.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("info: %v \n", api)
	api2, err2 := GetLocationInfoByIpApi("2403:71c0:2000:a0c1:afc1::")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("info: %v \n", api2)

}

func TestGetIpgeolocationInfo(t *testing.T) {
	info, err := GetIpgeolocationInfo("208.95.112.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("info: %v \n", info)
}

func TestGetFlagEmoji(t *testing.T) {
	emoji := GetFlagEmoji("US")
	fmt.Println(emoji)
}

func TestGetFlagEmojiSimple(t *testing.T) {
	simple := GetFlagEmojiSimple("US")
	fmt.Println(simple)
}

func TestCheckNormalIpv6Address(t *testing.T) {
	address := CheckNormalIpv6Address("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
}

func TestCheckNormalIpAddress(t *testing.T) {
	address := CheckNormalIpAddress("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
	address = CheckNormalIpAddress("127.2.3.1")
	fmt.Println(address)
}

func TestCheckStrIsIpAddress(t *testing.T) {
	address := CheckStrIsIpAddress("403:71c0:2000:a0c1:afc1::")
	fmt.Println(address)
	address = CheckStrIsIpAddress("127.2.3.")
	fmt.Println(address)

}
