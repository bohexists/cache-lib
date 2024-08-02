package cache

import (
	"sync"
	"time"
)

// Cache is a basic in-memory storage for data.
type Cache struct {
	data map[string]cacheObject
	mu   sync.RWMutex // Mutex to ensure thread-safety
}

// cacheObject struct to store value and expiration in Cache.
type cacheObject struct {
	value   interface{}
	expired int64
}

// New creates and returns a new Cache.
func New() *Cache {
	// Create a variable result
	result := &Cache{}

	// Initialize the storage
	result.data = make(map[string]cacheObject)

	return result
}

// Set adds a value.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	// Lock for writing
	c.mu.Lock()
	defer c.mu.Unlock()

	err := validateKey(key)
	if err != nil {
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

	err := validateKey(key)
	if err != nil {
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

	err := validateKey(key)
	if err != nil {
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

	err := validateKey(key)
	if err != nil {
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
