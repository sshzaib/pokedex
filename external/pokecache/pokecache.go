package pokecache

import "time"

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]CacheEntry),
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
