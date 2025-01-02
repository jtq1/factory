package cache

import (
	models "appTalleres"
	"fmt"
	"runtime"
	"time"
)

type cacheManager struct {
	cacheManagers []models.CacheManager
}

func NewCacheManager(cms ...models.CacheManager) *cacheManager {
	return &cacheManager{cms}
}

func (c *cacheManager) AddCachePair(cds ...models.CacheManager) {
	c.cacheManagers = append(c.cacheManagers, cds...)
}

func (c *cacheManager) SyncCache() {
	fmt.Printf("SyncCache Started")
	ticker := time.NewTicker(30 * time.Second)

	syncFunc := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic inside SyncCache: %v\n", r)

				// Create a buffer to hold the stack trace
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, true)             // Capture the stack trace
				fmt.Printf("Stack trace:\n%s\n", buf[:n]) // Print the stack trace
			}
		}()

		for i := range c.cacheManagers {
			err := c.cacheManagers[i].SyncCache()
			if err != nil {
				fmt.Printf("Error SyncCache: %v - %T\n", err, c.cacheManagers[i])
			}
		}
	}

	go func() {
		for range ticker.C {
			syncFunc()
		}
	}()
}
