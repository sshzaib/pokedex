package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    sync.Mutex
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func (c *Cache) ReapFunc() {
	interval := 5 * time.Hour
	for {
		for url, cacheEntry := range c.cache {
			if time.Since(cacheEntry.createdAt) > interval {
				c.mu.Lock()
				delete(c.cache, url)
				c.mu.Unlock()
			}
		}
	}
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]CacheEntry),
		mu:    sync.Mutex{},
	}
}

func (c *Cache) AddCache(key string, val []byte) {
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	value, ok := c.cache[key]
	if ok {
		return value.val, ok
	}
	return nil, ok
}
