package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Data map[string]CacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	var cache Cache
	cache.Data = make(map[string]CacheEntry)
	ticker := time.NewTicker(interval * time.Millisecond)
	go func() {
		for range ticker.C {
			cache.reapLoop(interval)
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = CacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	entry, ok := c.Data[key]
	if !ok {
		return []byte{}, false
	}

	return entry.Val, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.Data {
		if time.Since(entry.CreatedAt) > interval {
			delete(c.Data, key)
		}
	}
}
