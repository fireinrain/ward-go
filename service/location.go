package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

type LocationInfo struct {
	//主机名称
	HostName string `json:"hostName"`
	//Ip地址
	IpAddress string `json:"ipAddress"`
	//国家
	CountryName string `json:"countryName"`
	//国家代码
	CountryCode string `json:"countryCode"`
	//国旗标志
	CountryFlag string `json:"countryFlag"`
	//地区
	RegionName string `json:"regionName"`
	//城市
	CityName string `json:"cityName"`
	//邮编
	PostCode string `json:"postCode"`
	//经度
	Latitude string `json:"latitude"`
	//纬度
	Longitude string `json:"longitude"`
}

// GetLocationInfoByIPv4
//
//	@Description: 根据ip查询地理信息
//	@param ipv4
//	@return LocationInfo
func GetLocationInfoByIPv4(ipv4 string) LocationInfo {
	locationInfo := LocationInfo{}
	resp, err := http.Get("https://www.geodatatool.com/en/?ip=" + ipv4)
	if err != nil {
		log.Println("get location by ip error: ", err)
		return locationInfo
	}
	defer resp.Body.Close()

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get location by ip error: ", err)
		return locationInfo
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Println("request getdata failed: ", err)
		return locationInfo
	}
	//抽取
	val, _ := doc.Find("//script").First().Attr("src")

	// Print the response body as a string
	fmt.Println(val)
	fmt.Println(string(body))
	return locationInfo
}
