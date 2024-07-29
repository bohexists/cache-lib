# Cache Library
================================

## Description

This is a basic in-memory cache library for storing data in a Go application. It provides methods to set, get, and delete values.

## Installation

To install the cache library, use the following command:

```sh
go get -u github.com/yourusername/cache
```

Replace `yourusername` with your actual GitHub username or the repository URL where your library is hosted.

## Example #1

```go
package main

import (
    "fmt"
    "github.com/yourusername/cache" // Replace with the correct import path
)

func main() {
    // Create a new cache
    c := cache.New()

    // Set values in the cache
    c.Set("name", "John")
    c.Set("age", 30)

    // Get values from the cache
    name := c.Get("name")
    age := c.Get("age")

    fmt.Println("Name:", name) // Output: Name: John
    fmt.Println("Age:", age)   // Output: Age: 30

    // Delete a value from the cache
    c.Delete("name")
    name = c.Get("name")

    if name == nil {
        fmt.Println("Name not found") // Output: Name not found
    }
}
```

## Example #2

```go
package main

import (
    "fmt"
    "github.com/yourusername/cache" // Replace with the correct import path
)

func main() {
    // Create a new cache
    c := cache.New()

    // Set multiple values in the cache
    c.Set("city", "New York")
    c.Set("country", "USA")

    // Retrieve and print values
    city := c.Get("city")
    country := c.Get("country")

    fmt.Println("City:", city)       // Output: City: New York
    fmt.Println("Country:", country) // Output: Country: USA

    // Delete values from the cache
    c.Delete("city")
    c.Delete("country")

    // Attempt to retrieve deleted values
    city = c.Get("city")
    if city == nil {
        fmt.Println("City not found") // Output: City not found
    }

    country = c.Get("country")
    if country == nil {
        fmt.Println("Country not found") // Output: Country not found
    }
}
```

## API

### `New() Cache`
Creates and returns a new Cache instance.

### `Set(key string, value interface{})`
Adds a value to the cache with the specified key.

- **key**: The key for the cache entry.
- **value**: The value to store in the cache.

### `Get(key string) interface{}`
Retrieves a value from the cache for the specified key.

- **key**: The key for the cache entry.
- **returns**: The value stored in the cache, or `nil` if the key does not exist.

### `Delete(key string)`
Removes a value from the cache for the specified key.

- **key**: The key for the cache entry to be removed.
