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

## Summary: When to Use Which

- **`:=`**: Inside functions for combined declaration + initialization.
- **`var`**: At package level, or inside functions when needing zero value / delayed initialization / explicit type.
