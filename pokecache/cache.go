package pokecache

import (
	"sync"
	"time"
)

type cacheEntry[T any] struct {
	lastUsed time.Time
	value    T
}

type Cache[T any] struct {
	data map[string]cacheEntry[T]
	mux  *sync.Mutex
	stop chan bool
}

func NewCache[T any](retentionTime time.Duration) Cache[T] {
	cache := Cache[T]{
		data: make(map[string]cacheEntry[T]),
		mux:  &sync.Mutex{},
		stop: make(chan bool),
	}
	cache.repeatCleanupLoop(retentionTime)
	return cache
}

func (c *Cache[T]) Close() {
	c.stop <- true
}

func (c *Cache[T]) Get(key string) (*T, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.data[key]
	if ok {
		entry.lastUsed = time.Now()
		c.data[key] = entry
		return &entry.value, ok
	}
	return nil, ok
}

func (c *Cache[T]) Set(key string, value T) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry := cacheEntry[T]{value: value, lastUsed: time.Now()}
	c.data[key] = entry
}

func (c *Cache[T]) repeatCleanupLoop(retentionTime time.Duration) {
	ticker := time.NewTicker(retentionTime)

	go func() {
		for {
			select {
			case <-c.stop:
				ticker.Stop()
				return
			case <-ticker.C:
				c.runCleanup(retentionTime)
			}
		}
	}()
}

func (c *Cache[T]) runCleanup(retentionTime time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	now := time.Now()
	for key, value := range c.data {
		if value.lastUsed.Sub(now).Abs() > retentionTime {
			delete(c.data, key)
		}
	}
}
