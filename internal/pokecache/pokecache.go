package pokecache

import (
    "time"
    "sync"
)

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}

type Cache struct {
    cache map[string]cacheEntry
    mux   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
    c := Cache{
        cache: make(map[string]cacheEntry),
        mux:   &sync.Mutex{},
    }
    go c.reapLoop(interval)
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mux.Lock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val:       val,
    }
    defer c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.Lock()
    data, exists := c.cache[key]
    if exists {
        return data.val, true
    }
    defer c.mux.Unlock()
    return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(interval)
    }
}

func (c *Cache) reap(interval time.Duration) {
    c.mux.Lock()
    timeThreshold := time.Now().Add(-interval)
    for key, entry := range c.cache {
        if entry.createdAt.Before(timeThreshold) {
            delete(c.cache, key)
        }
    }
    defer c.mux.Unlock()
}
