# Cache Library

Cache Library is a simple in-memory caching solution for Go, providing basic cache operations with TTL (time-to-live) functionality.

## Features

- Set and get cache values with expiration times.
- Delete cache values.
- Check for the existence of cache keys.
- List all cache keys.
- Periodically clean up expired cache entries.

## Installation

To install the Cache Library, use the following command:

```sh
go get -u <module-name>
```

## Usage

Here's how you can use the Cache Library in your project:

### Example

```go
package main

import (
	"fmt"
	"time"

	"<module-name>/cache"
)

func main() {
	// Create a new cache
	c := cache.New()

	// Set a value in the cache with a TTL of 5 seconds
	err := c.Set("key1", "value1", 5*time.Second)
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}

	// Retrieve the value from the cache
	value, err := c.Get("key1")
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Println("Retrieved value:", value)

	// Check if the key exists
	exists := c.Exists("key1")
	fmt.Println("Key exists:", exists)

	// List all keys
	keys, err := c.Keys()
	if err != nil {
		fmt.Println("Error getting keys:", err)
		return
	}
	fmt.Println("All keys:", keys)

	// Launch the cleaner to remove expired cache entries every 10 seconds
	c.LaunchCleaner(10 * time.Second)

	// Wait for 6 seconds to let the key expire
	time.Sleep(6 * time.Second)

	// Try to retrieve the expired value
	value, err = c.Get("key1")
	if err != nil {
		fmt.Println("Error getting value:", err)
	} else {
		fmt.Println("Retrieved value:", value)
	}

	// Check if the key exists after expiration
	exists = c.Exists("key1")
	fmt.Println("Key exists:", exists)
}
```

## API

### Cache

#### `func New() *Cache`

Creates and returns a new cache.

#### `func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error`

Sets a value in the cache with a specified TTL.

#### `func (c *Cache) Get(key string) (interface{}, error)`

Retrieves a value from the cache.

#### `func (c *Cache) Delete(key string) error`

Deletes a value from the cache.

#### `func (c *Cache) Exists(key string) bool`

Checks if a key exists in the cache.

#### `func (c *Cache) Keys() ([]string, error)`

Returns a list of all keys in the cache.

#### `func (c *Cache) LaunchCleaner(interval time.Duration)`

Launches a cleaner that periodically removes expired cache entries.
