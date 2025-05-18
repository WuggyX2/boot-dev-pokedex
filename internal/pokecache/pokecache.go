package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	values   map[string]cacheEntry
	mut      *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		interval: interval,
		values:   make(map[string]cacheEntry),
		mut:      &sync.Mutex{},
	}
	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.values[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	entry, ok := c.values[key]

	if ok {
		return entry.val, ok
	}

	return nil, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	defer ticker.Stop()

	for {
		select {
		case currentTime := <-ticker.C:
			c.mut.Lock()

			for key, value := range c.values {
				newTime := value.createdAt.Add(c.interval)
				if newTime.Before(currentTime) {
					delete(c.values, key)
				}
			}
			c.mut.Unlock()
		}

	}
}
