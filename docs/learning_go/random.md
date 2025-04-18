# Random Number Generation in Go (`math/rand`)

Go's `math/rand` package provides pseudo-random number generation.

## Default Behavior (Go 1.20+)

**Automatic Seeding:**

- Starting from **Go 1.20**, the global random number generator (used by functions like `rand.Intn`, `rand.Float64`, `rand.Shuffle`) is **automatically seeded** using a secure random source the first time it's needed.
- **No explicit seeding required** for typical use: You generally don't need to call `rand.Seed(time.Now().UnixNano())` anymore just to get different results on each program run.

```go
package main

import (
  "fmt"
  "math/rand" // Uses global generator
)

func main() {
  // Go 1.20+: Different output each run without explicit seeding
  fmt.Println(rand.Intn(100))
  fmt.Println(rand.Intn(100))
}

```

## When to Use `rand.Seed`

Despite automatic seeding of the global source, you might still use `rand.Seed` (or seed a custom source) in specific scenarios:

1. **Deterministic Results:**
    - For **testing, simulations, or reproducibility**, you need the *same* sequence of "random" numbers every time. Seed the global generator with a *fixed* value.

    ```go
    rand.Seed(42) // Use a constant seed
    fmt.Println(rand.Intn(100)) // Always prints the same first number
    fmt.Println(rand.Intn(100)) // Always prints the same second number
    ```

2. **Custom `rand.Source`:**
    - If you create your own random source and generator (e.g., for concurrent use without locking the global source), you must seed it yourself.

    ```go
    // Create a new source seeded with a specific value
    source := rand.NewSource(12345)
    // Create a generator using that source
    r := rand.New(source)

    fmt.Println(r.Intn(100)) // Uses the custom generator 'r'
    ```

> [!NOTE]
>
> For cryptographically secure random numbers (e.g., for keys, tokens), always use the `crypto/rand` package instead of `math/rand`.
>
> ```go
> // Example using crypto/rand
> package main
> 
> import (
>   "crypto/rand"
>   "encoding/base64"
>   "fmt"
> )
> 
> func main() {
>   // Generate 16 cryptographically secure random bytes
>   b := make([]byte, 16)
>   _, err := rand.Read(b)
>   if err != nil {
>     panic(err) // Handle errors appropriately in real applications
>   }
> 
>   // Often encoded for use (e.g., Base64 for tokens)
>   token := base64.URLEncoding.EncodeToString(b)
>   fmt.Println("Secure Token:", token)
> }
> ```
