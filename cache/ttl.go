package cache

import "time"

// LaunchCleaner periodically removes expired cache data.
func (c *Cache) LaunchCleaner(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			for key, object := range c.data {
				if isExpired(object) {
					delete(c.data, key)
				}
			}
		}
	}()
}
