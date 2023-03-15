package platform

import gocache "github.com/patrickmn/go-cache"

type WindowsPlatform struct {
	PlatformName string         `json:"platformName"`
	Cache        *gocache.Cache `json:"cache"`
}
