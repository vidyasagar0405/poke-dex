package pokecache_test

import (
	"testing"
	"time"

	"github.com/vidyasagar0405/pokedexcli/internals/pokecache"
)

func TestCacheExpiration(t *testing.T) {
	cache := pokecache.NewCache(100 * time.Millisecond)
	key := "testKey"
	val := []byte("testValue")

	cache.Add(key, val)

	// Immediate check
	if _, found := cache.Get(key); !found {
		t.Error("Entry not found immediately after addition")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)
	if _, found := cache.Get(key); found {
		t.Error("Entry not expired after TTL")
	}
}
