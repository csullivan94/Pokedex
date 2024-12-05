package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
	}
	ticker := time.NewTicker(interval)
	cache.reapLoop(*ticker, interval)
	return cache
}

func (c *Cache) Add(key string, v []byte) {

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}

}

func (c *Cache) Get(key string) (v []byte, a bool) {

	value, exists := c.data[key]
	if !exists {
		return nil, false
	}

	return value.val, true
}

func (c *Cache) reapLoop(ticker time.Ticker, interval time.Duration) {

	for i := range c.data {
		difference := (<-ticker.C).Sub(c.data[i].createdAt)
		if difference > interval {
			c.mu.Lock()
			defer c.mu.Unlock()
			delete(c.data, i)
		}
	}
}
