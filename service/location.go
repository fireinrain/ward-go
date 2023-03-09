package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//https://www.geodatatool.com/ data

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

// https://ip-api.com/ data

type IPApiInfo struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

func GetLocationInfoByIpApi(ipv4 string) (*IPApiInfo, error) {
	normalIpv4Address := CheckNormalIpv4Address(ipv4)
	if !normalIpv4Address {
		return nil, errors.New("args not a valid ipv4 address: " + ipv4)
	}
	//TODO warn I use the free pricing mode,so request does not have ssl or https protocol
	var requestUrl = "http://ip-api.com/json/" + ipv4
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println("get location by ip-api error: ", err)
		return nil, errors.New("get location by ip-api error: " + err.Error())
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh,en;q=0.9,zh-TW;q=0.8,zh-CN;q=0.7,ja;q=0.6")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Origin", "http://ip-api.com")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "http://ip-api.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var ipApiInfo IPApiInfo
	err = json.NewDecoder(resp.Body).Decode(&ipApiInfo)
	if err != nil {
		log.Println("get location by ip-api error: ", err)
		return nil, errors.New("get location by ip-api error: " + err.Error())
	}
	if ipApiInfo.Status != "success" {
		return nil, errors.New("failed to get location by ip-api")
	}
	return &ipApiInfo, nil
}

// GetFlagEmoji
//
//	@Description: 根据国家code获取国旗emoji
//	@param countryCode
//	@return string
func GetFlagEmoji(countryCode string) string {
	var codePoints []int
	for _, char := range strings.ToUpper(countryCode) {
		codePoints = append(codePoints, 127397+int(char))
	}
	runes := make([]rune, len(codePoints))
	for i, cp := range codePoints {
		runes[i] = rune(cp)
	}
	return string(runes)
}

// GetFlagEmojiSimple
//
//	@Description: 根据国家code获取国旗emoji
//	@param countryCode
//	@return string
func GetFlagEmojiSimple(countryCode string) string {
	var codePoints []rune
	for _, char := range strings.ToUpper(countryCode) {
		codePoints = append(codePoints, rune(127397+int(char)))
	}
	return string(codePoints)
}

// CheckNormalIpv4Address
//
//	@Description: 判断一个字符串是否是一个合法的ipv4字符串
//	@param someString
//	@return bool
func CheckNormalIpv4Address(someString string) bool {
	pattern := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	match, err := regexp.MatchString(pattern, someString)
	if err != nil {
		log.Println("error for check normal ipv4 address:", err)
		return false
	}
	if match {
		return true
	}
	return false
}
