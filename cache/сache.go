package cache

// Cache is a basic in-memory storage for data.
type Cache struct {
	data map[string]interface{}
}

// New creates and returns a new Cache.
func New() Cache {
	// Create a variable result
	result := Cache{}

	// Initialize the storage field with an empty map
	result.data = make(map[string]interface{})

	return result
}

// Set adds a value.
func (c *Cache) Set(key string, value interface{}) error {
	if err := validateKey(key); err != nil {
		return err
	}
	c.data[key] = value
	return nil
}

// Get retrieves a value.
func (c *Cache) Get(key string) (interface{}, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}
	if err := c.checkExistence(key); err != nil {
		return nil, err
	}
	return c.data[key], nil
}

// Delete removes a value.
func (c *Cache) Delete(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}
	if err := c.checkExistence(key); err != nil {
		return err
	}
	delete(c.data, key)
	return nil
}

// Exists checks if a key exists.
func (c *Cache) Exists(key string) (bool, error) {
	if err := validateKey(key); err != nil {
		return false, err
	}
	_, result := c.data[key]
	return result, nil
}

// Keys returns a list of keys.
func (c *Cache) Keys() ([]string, error) {
	result := make([]string, 0, len(c.data))
	for key := range c.data {
		result = append(result, key)
	}
	return result, nil
}
