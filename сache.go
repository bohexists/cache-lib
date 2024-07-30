package cache

import (
	"errors"
)

// Cache is a basic in-memory storage for data.
type Cache struct {
	data map[string]interface{}
}

// New creates and returns a new Cache.
func New() Cache {
	// Create a variable result of type Cache
	result := Cache{}

	// Initialize the storage field with an empty map
	result.data = make(map[string]interface{})

	// Return the new Cache
	return result
}

// Set adds a value to the cache.
func (c *Cache) Set(key string, value interface{}) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	c.data[key] = value
	return nil
}

// Get retrieves a value from the cache.
func (c *Cache) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("key cannot be empty")
	}
	value, exists := c.data[key]
	if !exists {
		return nil, errors.New("key not found in cache")
	}
	return value, nil
}

// Delete removes a value from the cache.
func (c *Cache) Delete(key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if _, exists := c.data[key]; !exists {
		return errors.New("key not found in cache")
	}
	delete(c.data, key)
	return nil
}

// Exists checks if a key exists in the cache.
func (c *Cache) Exists(key string) (bool, error) {
	if key == "" {
		return false, errors.New("key cannot be empty")
	}
	_, result := c.data[key]
	return result, nil
}

// Keys returns a list of all keys in the cache.
func (c *Cache) Keys() ([]string, error) {
	if len(c.data) == 0 {
		return nil, errors.New("cache is empty")
	}
	result := make([]string, 0, len(c.data))
	for key := range c.data {
		result = append(result, key)
	}
	return result, nil
}
