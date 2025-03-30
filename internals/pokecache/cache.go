package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		cacheEntries: map[string]cacheEntry{},
		mutex:        sync.Mutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if c.cacheEntries[key].val != nil {
		return c.cacheEntries[key].val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {

}
