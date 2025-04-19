# Pointers in Go

Go pointers hold the memory address of a variable. They are used to allow functions to modify the original variable, to avoid copying large amounts of data, and to represent optional values (using `nil`).

## Core Operations

1. **Address-Of Operator (`&`)**: Gets the memory address of a variable.

    ```go
    x := 10
    p := &x // p holds the memory address of x (p is of type *int)
    fmt.Println(p) // Output: 0x... (memory address)
    ```

2. **Pointer Type (`*T`)**: A variable of type `*T` holds the address of a variable of type `T`.

    ```go
    var p *int // Declares a pointer to an int; its zero value is nil
    x := 10
    p = &x // Assign the address of x to p
    ```

3. **Dereference Operator (`*`)**: Accesses the value stored at the memory address held by a pointer.

    ```go
    x := 10
    p := &x

    fmt.Println(*p) // Output: 10 (the value at the address p holds)

    // Modify the original variable x via the pointer p
    *p = 20
    fmt.Println(x) // Output: 20
    ```

## Pointers and Functions

Go is strictly **pass-by-value**. When a pointer is passed to a function, a **copy of the pointer (the memory address)** is passed.

- The function receives a *copy* of the memory address.
- This copy still points to the *original* variable.
- The function can use its copy of the pointer to *modify the original variable* by dereferencing (`*`).

```go
func increment(ptr *int) {
  // Dereference the pointer copy to modify the original value
  *ptr = *ptr + 1
}

func main() {
  count := 5
  increment(&count) // Pass the address of count
  fmt.Println(count) // Output: 6
}
```

(See [`pointer_practice.md`](./pointer_practice.md) for more detail on how pointer *addresses* themselves are copied).

## Pointers and Structs

Pointers are frequently used with structs, especially with methods.

```go
type Point struct {
  X, Y int
}

// Value Receiver: Operates on a *copy* of the Point.
// Cannot modify the original Point.
func (pt Point) Info() string {
  return fmt.Sprintf("X=%d, Y=%d", pt.X, pt.Y)
}

// Pointer Receiver: Operates on the *original* Point via its address.
// Can modify the original Point.
func (pt *Point) Move(dx, dy int) {
  pt.X += dx // Go automatically dereferences pt (*pt).X
  pt.Y += dy // Go automatically dereferences pt (*pt).Y
}

func main() {
  p1 := Point{X: 1, Y: 2}

  // Calling a Pointer Receiver Method:
  // Go provides a shortcut: you can call pointer receiver methods
  // directly on values. Go automatically passes the address (&p1).
  p1.Move(10, 20) 
  // Equivalent to (&p1).Move(10, 20)

  fmt.Println(p1.Info()) // Output: X=11, Y=22
}

```

**When to Use Pointer Receivers (`*T`) vs. Value Receivers (`T`):**

- Use **Pointer Receivers** (`*T`):
    1. When the method needs to **modify** the receiver.
    2. When the struct is large, to **avoid copying** it on each method call.
- Use **Value Receivers** (`T`):
    1. When the method **does not need to modify** the receiver.
    2. When the struct is small and cheap to copy, or you explicitly want the method to operate on a copy.

## Value Types vs. Reference Types

The key difference is what gets copied when passed to a function.

### Value Types

- **Examples:** `int`, `float`, `string`, `bool`, `struct`, `array`.
- **Behavior:** When passed to a function, the **entire value is copied**. The function operates on a completely separate copy.
- **Modification Effect:** Changes made to the copy inside the function **do not affect** the original variable.
- **To Modify Original:** Pass a pointer (`*T`) to the value type instead.

```go
// Example with string (value type)
func modifyString(s string) {
  s = "modified in function"
}

originalStr := "original"
modifyString(originalStr)
fmt.Println(originalStr) // Output: original

// Example with pointer to string
func modifyStringPtr(ps *string) {
  *ps = "modified via pointer"
}

originalStrPtr := "original"
modifyStringPtr(&originalStrPtr)
fmt.Println(originalStrPtr) // Output: modified via pointer
```

### Reference Types

- **Examples:** `slice`, `map`, `channel`, `func`, `interface`, `pointer` (`*T`).
- **Behavior:** When passed to a function, the **header/descriptor/pointer value is copied**. This copy still **refers to the same underlying data structure**.
- **Modification Effect:** Modifications made to the *underlying data* through the function's copy **are visible** outside the function.

```go
// Example with slice (reference type)
func modifySlice(sl []int) {
  if len(sl) > 0 {
    sl[0] = 99 // Modifies the underlying array
  }
}

originalSlice := []int{1, 2, 3}
modifySlice(originalSlice)
fmt.Println(originalSlice) // Output: [99 2 3]

// Example with map (reference type)
func modifyMap(m map[string]int) {
  m["existing"] = 100 // Modifies the underlying map data
  m["new"] = 200
}

originalMap := map[string]int{"existing": 1}
modifyMap(originalMap)
fmt.Println(originalMap) // Output: map[existing:100 new:200]
```

> [!NOTE]
>
> Even reference types are technically passed by value (the reference/header value itself is copied), but the practical effect often *feels* like pass-by-reference because the copy points to the same underlying data.

## `nil` Pointers

The zero value for any pointer type is `nil`. It indicates the pointer doesn't point to any valid memory location.

- Dereferencing a `nil` pointer causes a runtime panic.
- Always check if a pointer might be `nil` before dereferencing if its validity isn't guaranteed.

```go
var p *int // p is nil
if p != nil {
  fmt.Println(*p) // Safe: only dereference if not nil
}
// fmt.Println(*p) // PANIC: runtime error: invalid memory address or nil pointer dereference
```

## Operator Summary

| Operator | Name | Usage | Description |
| :--- | :--- | :--- | :--- |
| `&` | Address-Of | `&variable` | Gets the memory address of `variable` |
| `*` | Dereference | `*pointerVariable` | Gets the value stored at the address |
| `*` | Pointer Type | `var p *Type` | Declares `p` as a pointer to `Type` |
