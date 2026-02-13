package cache

import "sync"

type Cache struct {
	entries map[string]Entry
	mu      sync.RWMutex
}

type Entry struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
}

func New() *Cache {
	return &Cache{
		entries: make(map[string]Entry),
	}
}

func (c *Cache) Get(key string) (Entry, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.entries[key]
	return value, ok
}

func (c *Cache) Set(key string, value Entry) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = value
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string]Entry)
}
