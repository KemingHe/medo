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

> [!NOTE]
> Go automatically handles dereferencing/referencing for method calls (e.g., `val.Update()` works even though `Update` has a pointer receiver and `val` is a value).
