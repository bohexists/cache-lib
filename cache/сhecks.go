package cache

import (
	"errors"
	"time"
)

// validateKey checks if the key is valid.
func validateKey(key string) error {
	if key == "" {
		return errors.New("key empty")
	}
	return nil
}

// isExpired checks if the object is expired.
func isExpired(object cacheObject) bool {
	return time.Now().UnixNano() > object.expired
}
