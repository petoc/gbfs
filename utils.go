package gbfs

import "sync"

// Cache ...
type Cache interface {
	Get(k string) (v Feed, ok bool)
	Set(k string, v Feed)
}

// InMemoryCache ...
type InMemoryCache struct {
	sync.RWMutex
	m map[string]Feed
}

// Get ...
func (c *InMemoryCache) Get(k string) (v Feed, ok bool) {
	c.RLock()
	v, ok = c.m[k]
	c.RUnlock()
	return
}

// Set ...
func (c *InMemoryCache) Set(k string, v Feed) {
	c.Lock()
	c.m[k] = v
	c.Unlock()
	return
}

// NewInMemoryCache ...
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		m: make(map[string]Feed),
	}
}

func indexInSlice(n string, h []string) int {
	for k, v := range h {
		if n == v {
			return k
		}
	}
	return -1
}

func inSlice(n string, h []string) bool {
	return indexInSlice(n, h) > -1
}
