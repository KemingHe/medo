# Maps in Go

Maps are key-value data structures that provide fast lookups, insertions, and deletions. Unlike arrays and slices, maps are unordered collections.

## Core Concepts

```go
// Declaration
var scores map[string]int      // Declares a nil map (zero value)
scores = make(map[string]int)  // Initialize with make()

// Literal initialization
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#4bf745",
}
```

## Why Maps Matter

- **Reference Type**: Maps are reference types, similar to slices
- **Fast Lookups**: O(1) average time complexity for retrievals
- **Dynamic Size**: Automatically grow as you add elements
- **Type Safety**: Both keys and values have consistent types throughout the map

## Operations

```go
// Adding or updating a value
users := make(map[string]int)
users["alice"] = 25    // Add new key-value pair
users["alice"] = 26    // Update existing key

// Reading values (two ways)
age := users["alice"]  // Direct access returns zero value if key doesn't exist

age, exists := users["bob"]  // The "comma ok" idiom
if exists {
    fmt.Println("Bob's age:", age)
} else {
    fmt.Println("Bob not found")
}

// Deleting a key
delete(users, "alice")
```

## Iterating Over Maps

```go
scores := map[string]int{"alice": 98, "bob": 87, "carol": 92}

// Iteration order is not guaranteed!
for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}

// Iterate over keys only
for name := range scores {
    fmt.Println(name)
}
```

## Common Patterns

### Maps as Sets

```go
seen := make(map[string]bool)
seen["hello"] = true

// Check if element exists
if seen["hello"] {
    fmt.Println("Already seen 'hello'")
}
```

### Nested Maps

```go
// Map of maps for more complex data
userProfiles := map[string]map[string]string{
    "alice": {
        "email": "alice@example.com",
        "role":  "admin",
    },
}

// Access nested data
aliceEmail := userProfiles["alice"]["email"]
```

## Important Notes

1. **Not Thread-Safe**: Concurrent writes require synchronization (use `sync.RWMutex` or channels)
2. **Reference Semantics**: When maps are passed to functions, they pass a reference to the same underlying data
3. **No Size Limit**: Maps can grow to accommodate as many items as memory allows
4. **Keys Must Be Comparable**: Only types that support the `==` operator can be used as map keys
5. **Keys Are Unique**: Adding with an existing key overwrites the previous value

## Performance Considerations

- **Pre-sizing**: Use `make(map[K]V, capacity)` when you know approximate size
- **Memory Overhead**: Maps trade some memory efficiency for speed
- **Zero Values**: Accessing a nonexistent key returns the zero value of the value type
