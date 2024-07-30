package cache

import "errors"

// validateKey checks if the key is valid.
func validateKey(key string) error {
	if key == "" {
		return errors.New("key empty")
	}
	return nil
}

// checkExistence checks if the key exists in the cache.
func (c *Cache) checkExistence(key string) error {
	if _, exists := c.data[key]; !exists {
		return errors.New("key not found")
	}
	return nil
}
