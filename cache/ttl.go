package cache

import "time"

// LaunchCleaner periodically removes expired cache data.
func (c *Cache) LaunchCleaner(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			for key, item := range c.data {
				if isExpired(item) {
					delete(c.data, key)
				}
			}
		}
	}()
}
