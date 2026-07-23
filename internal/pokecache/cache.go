package pokecache

import (
	"sync"
	"time"
)

// Create a struct called Cache to hold a map[string]cacheEntry

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	// Add a map to hold the cached data
	data map[string]cacheEntry
	// Add a mutex to protect the map across goroutines
	mu sync.Mutex
}

// Expose a NewCache() function that creates a new cache
// It should take as input a configurable interval using the time.Duration method from
// the time package. This will be used to determine how long
// to keep items in the cache before they are considered stale.
// func NewCache() *Cache {
// use a time.Ticker inside a goroutine started by NewCache.
// In a loop like for range ticker.C { ... }, check the entries
// and remove any whose createdAt is older than the cache's interval.
//
//		go func() {
//			cache := &Cache{
//				data: make(map[string]cacheEntry),
//			}
//			cache.reapLoop(5 * time.Second)
//		}()
//		return &Cache{
//			data: make(map[string]cacheEntry),
//		}
//	}
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

// Create a cache.Add() method that adds a new entry to the cache.
// It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Create a cache.Get() method that retrieves an entry from the cache.
// It should take a key (a string) and return a val ([]byte) and a bool
// indicating if the entry was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

// Create a cache.reapLoop() method that is called when the cache is created
// by the NewCache function. Each time an interval which is the time.Duration passed
// to NewCache passes it should remove any entries that are older than the interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.data {
				if time.Since(entry.createdAt) > interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
