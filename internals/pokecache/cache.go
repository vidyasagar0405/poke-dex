package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mutex:        &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.cacheEntries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, v := range c.cacheEntries {
		if v.createdAt.Before(now.Add(-last)){
			delete(c.cacheEntries, k)
		}
	}
}
