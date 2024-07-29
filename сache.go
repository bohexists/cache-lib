package cache

// Cache is a basic in-memory storage for data.
type Cache struct {
	data map[string]interface{}
}

// New creates and returns a new Cache.
func New() Cache {
	// Create a variable result of type Cache
	result := Cache{}

	// Initialize the storage field with an empty map
	result.data = make(map[string]interface{})

	// Return the new Cache
	return result
}
