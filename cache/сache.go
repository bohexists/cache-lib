package cache

import (
	"sync"
	"time"
)

// CacheConfig holds the configuration for the Cache.
type CacheConfig struct {
	MaxSize int // Maximum number of Object in the Cache.
}

// Cache is a basic in-memory storage for data.
type Cache struct {
	data    map[string]cacheObject
	mu      sync.RWMutex
	maxSize int
}

// cacheObject struct to store value and expiration in Cache.
type cacheObject struct {
	key     string
	value   interface{}
	expired int64
}

// New creates and returns a new Cache.
func New(сonfig CacheConfig) *Cache {
	// Create a variable result
	return &Cache{
		data:    make(map[string]cacheObject),
		maxSize: сonfig.MaxSize,
	}
}

// Set adds a value.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := validateKey(key); err != nil {
		return err
	}

	if err := checkCacheSize(c.data, c.maxSize); err != nil {
		return err
	}

	expired := time.Now().Add(ttl).UnixNano()
	c.data[key] = cacheObject{
		value:   value,
		expired: expired,
	}
	return nil
}

// Get retrieves a value.
func (c *Cache) Get(key string) (interface{}, error) {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := validateKey(key); err != nil {
		return nil, err
	}

	result, exists := c.data[key]
	if !exists || isExpired(result) {
		return nil, nil // or a specific error indicating the key does not exist or is expired
	}

	return result.value, nil
}

// Delete removes a value.
func (c *Cache) Delete(key string) error {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := validateKey(key); err != nil {
		return err
	}
	delete(c.data, key)
	return nil
}

// Exists checks if a key exists.
func (c *Cache) Exists(key string) bool {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	if err := validateKey(key); err != nil {
		return false
	}
	_, exists := c.data[key]
	return exists
}

// Keys returns a list of keys.
func (c *Cache) Keys() ([]string, error) {
	// Lock for reading
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]string, 0, len(c.data))
	for key := range c.data {
		result = append(result, key)
	}
	return result, nil
}
