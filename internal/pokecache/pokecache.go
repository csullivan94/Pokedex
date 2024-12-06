package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]cacheEntry
	mu   sync.RWMutex
}

type cacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Data: make(map[string]cacheEntry),
	}
	ticker := time.NewTicker(interval)
	go cache.reapLoop(*ticker, interval)
	return cache
}

func (c *Cache) Add(key string, v []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = cacheEntry{
		CreatedAt: time.Now(),
		val:       v,
	}

}

func (c *Cache) Get(key string) (v []byte, a bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.Data[key]
	if !exists {
		return nil, false
	}

	return value.val, true
}

func (c *Cache) reapLoop(ticker time.Ticker, interval time.Duration) {
	for {
		<-ticker.C
		for i := range c.Data {
			elapsed := time.Since(c.Data[i].CreatedAt)
			if elapsed >= interval {
				delete(c.Data, i)
			}
		}
	}
}
