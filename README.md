# Cache Library

Cache Library is a simple in-memory caching solution for Go, providing basic cache operations with TTL (time-to-live) functionality and multiple eviction strategies.

## Features

- **Set and Get**: Store and retrieve values with optional expiration times.
- **Delete**: Remove cache entries.
- **Check Existence**: Verify if a key exists in the cache.
- **List Keys**: Retrieve all keys stored in the cache.
- **Eviction Strategies**: Support for FILO (First-In-Last-Out), LRU (Least Recently Used), and FIFO (First-In-First-Out) eviction strategies.
- **TTL Management**: Automatic removal of expired cache entries.
- **Periodic Cleanup**: Launch a background cleaner to remove expired entries at regular intervals.

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
	// Create a new cache with specific configuration
	c := cache.New(cache.CacheConfig{
		MaxSize:      100,           // Maximum number of elements in cache
		DefaultTTL:   5 * time.Minute, // Default TTL for cache entries
		EvictionType: cache.LRU,     // Eviction strategy: FILO, LRU, or FIFO
	})

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

#### `func New(config CacheConfig) *Cache`

Creates and returns a new cache with the specified configuration.

#### `func (c *Cache) Set(key string, value interface{}, ttl ...time.Duration) error`

Sets a value in the cache with a specified TTL. If no TTL is provided, the default TTL is used.

#### `func (c *Cache) Get(key string) (interface{}, error)`

Retrieves a value from the cache. Returns `nil` if the key does not exist or is expired.

#### `func (c *Cache) Delete(key string) error`

Deletes a value from the cache.

#### `func (c *Cache) Exists(key string) bool`

Checks if a key exists in the cache.

#### `func (c *Cache) Keys() ([]string, error)`

Returns a list of all keys in the cache.

#### `func (c *Cache) Clear()`

Removes all objects from the cache.

#### `func (c *Cache) Size() int`

Returns the number of elements currently stored in the cache.

#### `func (c *Cache) LaunchCleaner(interval time.Duration)`

Launches a cleaner that periodically removes expired cache entries.

## Eviction Strategies

The library supports three eviction strategies:

- **FILO (First-In-Last-Out)**: Removes the oldest element when the cache is full.
- **LRU (Least Recently Used)**: Removes the least recently used element when the cache is full.
- **FIFO (First-In-First-Out)**: Removes the first element added to the cache when the cache is full.

To specify an eviction strategy, set the `EvictionType` field in the `CacheConfig` when creating a new cache:

```go
c := cache.New(cache.CacheConfig{
    MaxSize:      100,
    DefaultTTL:   5 * time.Minute,
    EvictionType: cache.LRU, // Choose between FILO, LRU, FIFO
})
```

