package cache

import "errors"

// validateKey checks if the key is valid.
func validateKey(key string) error {
	if key == "" {
		return errors.New("key empty")
	}
	return nil
}
