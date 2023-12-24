package server

import (
	"sync"
)

// Cache is a simple in-memory cache implementation that is safe for concurrent use by multiple goroutines.
// It is used to cache files in memory.

type Cache struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) GetAll() map[string][]byte {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = val
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}
