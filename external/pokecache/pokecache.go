package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	for url, cacheEntry := range c.cache {
		if time.Since(cacheEntry.createdAt) > interval {
			c.mu.Lock()
			delete(c.cache, url)
			c.mu.Unlock()
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) AddCache(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	if ok {
		return value.val, ok
	}
	return nil, ok
}
