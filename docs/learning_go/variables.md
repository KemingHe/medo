# Variables in Go

Go offers two ways to declare variables: `:=` (short) and `var` (standard).

## 1. Short Declaration (`:=`)

**Use Case:** Preferred for declaring & initializing variables **inside functions**. It's idiomatic, concise, and infers type.

```go
// Inside a function:
message := "Hello" // Type string inferred
count := 10       // Type int inferred
```

**Cannot be used at package level.**

## 2. Standard Declaration (`var`)

**Use Case:**

- Package-level variables.
- Declaring a variable before assignment (gets zero value).
- Explicitly setting a type if inference isn't desired (less common).

```go
// Package level:
var defaultName string = "Guest"

func main() {
    var score int      // Zero value (0) assigned
    var ratio float64 // Zero value (0.0) assigned
    score = 100
    ratio = 0.5
    _ = ratio // Use variables
    _ = defaultName
}
```

## 3. Blank Identifier (`_`)

The blank identifier (`_`) is used to signal that a variable or value is intentionally unused. This is useful in several scenarios:

- **Ignoring function return values:** If a function returns multiple values but you only need some of them.

  ```go
  value, _ := someFunction() // Ignore the second return value
  ```

- **Ignoring loop variables:** When you only need the index or the value in a `for...range` loop.

  ```go
  // Inside a function:
  for i, _ := range items { /* use only index i */ }
  for _, item := range items { /* use only value item */ }
  ```

- **Checking for type assertion success:**

  ```go
  _, ok := interfaceValue.(ConcreteType) // Check if assertion ok, ignore the value
  ```

- **Side effects of imports:** Importing a package solely for its initialization side effects (e.g., registering a database driver).

  ```go
  import _ "github.com/go-sql-driver/mysql"
  ```

Using `_` tells the Go compiler you've acknowledged the variable/value but don't intend to use it, preventing "declared but not used" errors.

## Basic Types & Zero Values

Go is statically typed. Uninitialized variables get a "zero value".

| Type | Example | Zero Value | Description |
| :-- | :-- | :-- | :-- |
| `bool` | `true` | `false` | Boolean |
| `string` | `"Hi"` | `""` | UTF-8 text |
| `int` | `10` | `0` | Signed integer (32/64b, platform-dep.) |
| `int8` - `int64` | `127` | `0` | Signed integers (fixed 8-64 bit) |
| `uint` | `10` | `0` | Unsigned integer (32/64b, platform-dep.) |
| `uint8` - `uint64` | `255` | `0` | Unsigned integers (fixed 8-64 bit) |
| `uintptr` | *(varies)* | `0` | Integer large enough to store pointer address |
| `byte` | `'a'` | `0` | Alias for `uint8` |
| `rune` | `'好'` | `0` | Alias for `int32` (Unicode code point) |
| `float32`, `float64` | `3.14` | `0.0` | Floating-point (IEEE-754 32/64b) |
| `complex64`, `complex128` | `1+2i` | `0+0i` | Complex numbers (float32/64 parts) |

> [!NOTE]
>
> - Default numeric types: `int`, `float64`, `complex128`.
> - Composite types (arrays, slices, maps, structs, channels, etc.) are not listed here.

## Understanding `nil`

`nil` is a pre-declared identifier representing the zero value for several reference types:

- Pointers (`*int`, `*MyStruct`, etc.)
- Interfaces (`error`, `io.Reader`, etc.)
- Slices (`[]string`, `[]byte`, etc.)
- Maps (`map[string]int`, etc.)
- Channels (`chan int`, etc.)
- Function types (`func()`, `func(int) string`, etc.)

It signifies "no value" or an uninitialized state for these types. For other types like numbers, bools, and strings, the zero value is `0`, `false`, or `""` respectively, **not** `nil`.

A key point is that `nil` is *typed*. Although it represents zero, it takes on the specific type required by its context. For example, if a function returns `([]string, error)`:

```go
func processData(invalid bool) ([]string, error) {
    if invalid {
        // Here, nil is returned for the []string type.
        // It represents a nil slice, not just an untyped nil.
        return nil, fmt.Errorf("invalid input")
    }
    // ... process valid data ...
    return []string{"data"}, nil // Return actual slice and nil error
}
```

Therefore, you can compare a slice (or map, pointer, etc.) directly to `nil` to check if it's uninitialized:

```go
someSlice := []int{}
if someSlice == nil { // This comparison is valid
    fmt.Println("Slice is nil")
}
```

## Summary: When to Use Which

- **`:=`**: Inside functions for combined declaration + initialization.
- **`var`**: At package level, or inside functions when needing zero value / delayed initialization / explicit type.
