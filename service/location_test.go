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
}

func TestGetFlagEmoji(t *testing.T) {
	emoji := GetFlagEmoji("US")
	fmt.Println(emoji)
}

func TestGetFlagEmojiSimple(t *testing.T) {
	simple := GetFlagEmojiSimple("US")
	fmt.Println(simple)
}
