package platform

import gocache "github.com/patrickmn/go-cache"

type DarwinPlatform struct {
	PlatformName string         `json:"platformName"`
	Cache        *gocache.Cache `json:"cache"`
}
