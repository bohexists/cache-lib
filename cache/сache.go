package cache

import "time"

// Cache is a basic in-memory storage for data.
type Cache struct {
	data map[string]cacheObject
}

// cacheObject struct to store value and expired in Cache.
type cacheObject struct {
	value   interface{}
	expired int64
}

// New creates new Cache.
func New() *Cache {
	// Create a variable result
	result := &Cache{}

	// Initialize the storage
	result.data = make(map[string]cacheObject)

	return result
}

// Set adds a value.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
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
	err := validateKey(key)
	if err != nil {
		return nil, err
	}
	result := c.data[key]

	return result, nil
}

// Delete removes a value.
func (c *Cache) Delete(key string) error {
	err := validateKey(key)
	if err != nil {
		return err
	}
	delete(c.data, key)
	return nil
}

// Exists checks if a key exists.
func (c *Cache) Exists(key string) bool {
	err := validateKey(key)
	if err != nil {
		return false
	}
	_, result := c.data[key]
	return result
}

// Keys returns a list of keys.
func (c *Cache) Keys() ([]string, error) {
	result := make([]string, 0, len(c.data))
	for key := range c.data {
		result = append(result, key)
	}
	return result, nil
}
