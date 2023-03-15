package cache

import "time"
import "github.com/patrickmn/go-cache"

var GlobalCache = cache.New(5*time.Minute, 30*time.Minute)
