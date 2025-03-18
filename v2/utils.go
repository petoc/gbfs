package gbfs

import (
	"encoding/json"
	"sync"
)

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

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func (e *wrapError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.msg)
}

// NewError Wrap error
func NewError(msg string, err error) error {
	if err != nil {
		msg = msg + err.Error()
	}
	return &wrapError{msg, err}
}
