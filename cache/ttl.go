package cache

import "time"

// LaunchCleaner periodically removes expired cache data.
func (c *Cache) LaunchCleaner(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			c.mu.Lock() // Заблокировать доступ для записи
			for key, object := range c.data {
				if isExpired(object) {
					delete(c.data, key)
				}
			}
			c.mu.Unlock() // Разблокировать доступ
		}
	}()
}
