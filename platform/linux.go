package platform

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
	"ward-go/cache"
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

type ServerPlatform interface {
	GetStatusInfo() ServerInfo
}

type LinuxPlatform struct {
	PlatformName string         `json:"platformName"`
	Cache        *gocache.Cache `json:"cache"`
}

func NewLinuxPlatform() *LinuxPlatform {
	return &LinuxPlatform{
		PlatformName: "linux",
		Cache:        cache.GlobalCache,
	}
}

func (pf *LinuxPlatform) GetStatusInfo() ServerInfo {
	//TODO implement me
	panic("implement me")
}
