# Go Method Receivers

Attach functions to types using a **receiver** to create **methods**.

## Declaration & Calling

```go
type MyType string

// Method with receiver 'm' of type 'MyType'
func (m MyType) Print() { // Value receiver (operates on copy)
  println(m)
}

// Method with pointer receiver
func (m *MyType) Update(newVal string) { // Pointer receiver (operates on original)
  *m = MyType(newVal)
}

func main() {
  var val MyType = "Hello"
  val.Print() // Calls Print method on val

  val.Update("World") // Go automatically uses (&val).Update("World")
  val.Print()         // Prints "World"
}
```

## Best Practices

- **Naming:** Use short, idiomatic, single-letter receiver names (e.g., `t` for `MyType`). Be consistent across methods for the same type.
- **Avoid `this`/`self`:** Not idiomatic Go.
- **Value vs. Pointer Receiver:**
  - **Value (`func (t MyType) M()`):** Method operates on a copy. Use when modification is not needed.
  - **Pointer (`func (t *MyType) M()`):** Method operates on the original value (via pointer). Use for modification or to avoid copying large structs.
  - **Consistency:** If any method needs a pointer receiver, often all methods for that type use pointer receivers for consistency.
- **Receiver vs. Function:**
  - Use **methods (receivers)** when the function's primary role is to operate on or access the internal state of a specific instance of that type (e.g., `user.SetEmail()`). This groups behavior with the data it relates to.
  - Use **regular functions** for operations that don't conceptually "belong" to a single instance or type, or act as utility functions (e.g., `NewUserFromConfig(cfg)`).

  > **Case Study: `deal(deck, int) (deck, deck, error)`**
  >
  > Consider a function to deal cards from a `deck` type:
  >
  > ```go
  > type deck []string
  >
  > func deal(d deck, handSize int) (deck, deck, error) {
  >   // ... validation ...
  >   return d[:handSize], d[handSize:], nil // Returns slices, doesn't modify original d
  > }
  > ```
  >
  > **Why a regular function?**
  > 1. **Clarity:** The signature `func deal(d deck, ...)` clearly shows it takes a deck as *input* and produces two *new* decks (the hand and the remainder) as *output*.
  > 2. **Immutability:** It doesn't modify the original `deck` (`d`) passed in. A standalone function naturally signals this non-mutating, transformative behavior (input -> output).
  > 3. **Idiomatic:** While methods are common, using functions for transformations that produce new values without side effects is perfectly idiomatic Go.
  >
  > **When would a method be better?** If the function *mutated* the original deck (e.g., removed cards from it), a pointer receiver method (`func (d *deck) deal(...)`) would be more appropriate, signaling that the operation modifies the receiver instance.

> [!NOTE]
> Go automatically handles dereferencing/referencing for method calls (e.g., `val.Update()` works even though `Update` has a pointer receiver and `val` is a value).
