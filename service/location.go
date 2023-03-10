package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

const IpGeoLocationApiKey = "71a993e55ea64df29c3caa7c094f7099"

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

// GetLocationInfoByGeoDataTool
//
//	@Description: 根据ip查询地理信息
//	@param ipv4
//	@return LocationInfo
func GetLocationInfoByGeoDataTool(ipv4 string) (*LocationInfo, error) {
	normalIpv4Address := CheckNormalIpv4Address(ipv4)
	if !normalIpv4Address {
		return nil, errors.New("args not a valid ipv4 address: " + ipv4)
	}
	var locationInfo = &LocationInfo{}
	var requestUrl = "https://www.geodatatool.com/en/?ip=" + ipv4
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println("get location by geodatatool error: ", err)
		return nil, errors.New("get location by geodatatool error: " + err.Error())
	}
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh,en;q=0.9,zh-TW;q=0.8,zh-CN;q=0.7,ja;q=0.6")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.geodatatool.com")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", "https://www.geodatatool.com/en/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("parse geodatatool document error: ", err)
		return locationInfo, errors.New("parse geodatatool document error: " + err.Error())
	}
	var resultList []string

	dom.Find("body > div.container > div > div > div > div > div > div.col-md-4.column > div.sidebar-data.hidden-xs.hidden-sm > div > span:nth-child(2)").Each(func(i int, selection *goquery.Selection) {
		cleanStr := strings.TrimSpace(selection.Text())
		resultList = append(resultList, cleanStr)
		//fmt.Println(cleanStr)
	})
	country := resultList[3]
	countryStrs := strings.Split(country, " ")
	country = countryStrs[0]
	resultList[3] = country

	locationInfo.HostName = resultList[0]
	locationInfo.IpAddress = resultList[1]
	locationInfo.CountryName = resultList[2]
	locationInfo.CountryCode = resultList[3]
	locationInfo.CountryFlag = GetFlagEmojiSimple(resultList[3])
	locationInfo.RegionName = resultList[4]
	locationInfo.CityName = resultList[5]
	locationInfo.PostCode = resultList[6]
	locationInfo.Latitude = resultList[7]
	locationInfo.Longitude = resultList[8]
	return locationInfo, nil
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

// GetLocationInfoByIpApi
//
//	@Description: 通过ipapi.com获取ip的location信息
//	@param ipStr support ipv4 or ipv6
//	@return *IPApiInfo
//	@return error
func GetLocationInfoByIpApi(ipStr string) (*IPApiInfo, error) {
	normalIpv4Address := CheckStrIsIpAddress(ipStr)
	if !normalIpv4Address {
		return nil, errors.New("args not a valid ipStr address: " + ipStr)
	}
	//TODO warn I use the free pricing mode,so request does not have ssl or https protocol
	var requestUrl = "http://ip-api.com/json/" + ipStr
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

// https://ipgeolocation.io/ data
// 1k request for per day

type IpgeolocationInfo struct {
	IP             string   `json:"ip"`
	Hostname       string   `json:"hostname"`
	ContinentCode  string   `json:"continent_code"`
	ContinentName  string   `json:"continent_name"`
	CountryCode2   string   `json:"country_code2"`
	CountryCode3   string   `json:"country_code3"`
	CountryName    string   `json:"country_name"`
	CountryCapital string   `json:"country_capital"`
	StateProv      string   `json:"state_prov"`
	District       string   `json:"district"`
	City           string   `json:"city"`
	Zipcode        string   `json:"zipcode"`
	Latitude       string   `json:"latitude"`
	Longitude      string   `json:"longitude"`
	IsEu           bool     `json:"is_eu"`
	CallingCode    string   `json:"calling_code"`
	CountryTld     string   `json:"country_tld"`
	Languages      string   `json:"languages"`
	CountryFlag    string   `json:"country_flag"`
	GeonameID      string   `json:"geoname_id"`
	Isp            string   `json:"isp"`
	ConnectionType string   `json:"connection_type"`
	Organization   string   `json:"organization"`
	Asn            string   `json:"asn"`
	Currency       Currency `json:"currency"`
	TimeZone       TimeZone `json:"time_zone"`
}
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type TimeZone struct {
	Name            string  `json:"name"`
	Offset          int     `json:"offset"`
	CurrentTime     string  `json:"current_time"`
	CurrentTimeUnix float64 `json:"current_time_unix"`
	IsDst           bool    `json:"is_dst"`
	DstSavings      int     `json:"dst_savings"`
}

//"71a993e55ea64df29c3caa7c094f7099"
//"https://api.ipgeolocation.io/ipgeo?apiKey=API_KEY&ip=8.8.8.8"

// GetIpgeolocationInfo
//
//	@Description: 使用ipgeolocation 获取geoip信息
//	@param ipString
//	@return *IpgeolocationInfo
//	@return error
func GetIpgeolocationInfo(ipString string) (*IpgeolocationInfo, error) {
	normalIpv4Address := CheckStrIsIpAddress(ipString)
	if !normalIpv4Address {
		return nil, errors.New("args not a valid ipStr address: " + ipString)
	}
	var requestUrl = fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", IpGeoLocationApiKey, ipString)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println("get location by ipgeolocation error: ", err)
		return nil, errors.New("get location by ipgeolocation error: " + err.Error())
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh,en;q=0.9,zh-TW;q=0.8,zh-CN;q=0.7,ja;q=0.6")
	req.Header.Add("Cache-Control", "no-cache")
	//req.Header.Add("Origin", "http://ip-api.com")
	req.Header.Add("Pragma", "no-cache")
	//req.Header.Add("Referer", "http://ip-api.com/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var ipGeoInfo IpgeolocationInfo
	err = json.NewDecoder(resp.Body).Decode(&ipGeoInfo)
	if err != nil {
		log.Println("get location by ipgeolocation error: ", err)
		return nil, errors.New("get location by ipgeolocation error: " + err.Error())
	}
	return &ipGeoInfo, nil
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

// CheckNormalIpv6Address
//
//	@Description: 判断字符串是否是一个合法的ipv6地址
//	@param someString
//	@return bool
func CheckNormalIpv6Address(someString string) bool {
	ipv6Regex := regexp.MustCompile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)

	if ipv6Regex.MatchString(someString) {
		return true
	}
	return false
}

// CheckNormalIpAddress
//
//	@Description: 判断字符串是否是合法ip地址
//	@param someString
//	@return bool
func CheckNormalIpAddress(someString string) bool {
	return !(CheckNormalIpv4Address(someString) && CheckNormalIpv6Address(someString))
}

// CheckStrIsIpAddress
//
//	@Description: 判断str是否为合格的ip str
//	@param str
//	@return bool
func CheckStrIsIpAddress(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil
}

// GetMachineAllIps
//
//	@Description: 获取机器所有的ip
//	@return []string
func GetMachineAllIps() ([]string, error) {
	result := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("get machine all ips error: ", err)
		return result, errors.New("get machine all ips error: " + err.Error())
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//log.Println(ipnet.IP.String())
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result, nil
}

// GetMachineLocalIP
//
//	@Description: 获取局域网内机器分配的ip
//	@return string
//	@return error
func GetMachineLocalIP() (string, error) {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		log.Println("error getting ip address: ", err)
		return "", errors.New("error getting ip address: " + err.Error())
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	log.Println("your ip address is:", localAddr.IP)
	return localAddr.IP.String(), nil
}

// GetMachineISPIP
//
//	@Description: 获取公网ip
//	@return string
//	@return error
func GetMachineISPIP() (string, error) {
	//请求ifconfig.me 获得结果
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		log.Println("get isp ip address error: ", err.Error())
		return "", errors.New("get isp ip address error: " + err.Error())
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("get isp ip address error: ", err.Error())
		return "", errors.New("get isp ip address error: " + err.Error())
	}
	result := strings.TrimSpace(string(all))
	return result, nil
}
