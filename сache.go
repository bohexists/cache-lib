package cache

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
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get retrieves a value from the cache.
func (c *Cache) Get(key string) interface{} {

	value := c.data[key]
	return value
}

// Delete removes a value from the cache.
func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
