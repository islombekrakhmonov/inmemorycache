package cashe

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

// Cache represents the in-memory cache
type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

// NewCache creates a new instance of Cache
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

// Set adds an item to the cache with an optional expiration time (in seconds)
func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).Unix()
	}
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: exp,
	}
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 && time.Now().Unix() > item.Expiration {
		// Item has expired
		delete(c.items, key)
		return nil, false
	}

	return item.Value, true
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Cleanup removes expired items from the cache
func (c *Cache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().Unix()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			delete(c.items, k)
		}
	}
}
