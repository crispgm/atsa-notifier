// Package global loads global data
package global

import (
	"sync"

	"github.com/crispgm/atsa-notifier/internal/conf"
)

var (
	globalData = make(map[string]any)
	mu         sync.RWMutex
)

// LoadGlobalData .
func LoadGlobalData(cfg *conf.Conf) {
	mu.Lock()
	defer mu.Unlock()

	dbpath := cfg.ATSADB.LocalPath
	players, err := conf.LoadPlayerFromLocalDB(dbpath)
	if err != nil {
		panic(err)
	}
	globalData["players"] = players
	globalData["templates"] = cfg.Templates
}

// GetGlobalData .
func GetGlobalData(key string) any {
	mu.RLock()
	defer mu.RUnlock()
	return globalData[key]
}
