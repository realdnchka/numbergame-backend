package utils

import(
	"time"

	"github.com/patrickmn/go-cache"
)

var leaderboardCache *cache.Cache

func init() {
	leaderboardCache = cache.New(30*time.Minute, 10*time.Minute)
}

func GetFromCache(key string) 	(interface{}, bool) {
	return leaderboardCache.Get(key)
}

func SetToCache(key string, value interface{}) {
	leaderboardCache.Set(key, value, cache.DefaultExpiration)
}