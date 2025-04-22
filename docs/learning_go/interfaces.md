# Interfaces in Go

Interfaces define behavior through method signatures without implementation details. They enable polymorphism and decoupling in Go.

## Core Concepts

```go
// Interface declaration
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Types implementing the interface
type Rectangle struct {
    Width, Height float64
}

// Method implementation for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Function accepting any Shape
func PrintMeasurements(s Shape) {
    fmt.Printf("Area: %0.2f, Perimeter: %0.2f\n", s.Area(), s.Perimeter())
}
```

## Concrete Types vs Interface Types

In Go, the `type` keyword is used to define both concrete types and interface types:

```go
// Concrete type declaration
type Distance float64

// Interface type declaration
type Measurable interface {
    GetDistance() Distance
}
```

### Concrete Types

Concrete types specify exactly how data is structured and stored in memory. Examples include:

- Built-in types (`int`, `string`, `bool`)
- Composite types (`struct`, `array`, `slice`, `map`)
- User-defined types based on existing types (like `type Distance float64`)

Concrete types have a specific, known implementation.

### Interface Types

Interface types define only behavior through method signatures. They:

- Don't specify implementation details
- Represent a set of methods that a concrete type must implement
- Are satisfied implicitly by any concrete type that implements all required methods
- Allow for polymorphic behavior and decoupling

```go
// A variable of concrete type
car := Car{Position: 0}

// A variable of interface type, holding a concrete value
var movable Mover = &car  // Interface value contains both type and value information
```

When working with interfaces, you're dealing with two components:

1. The dynamic type (concrete type that satisfies the interface)
2. The value of that type

This distinction is important when considering type assertions, comparison, and the nil interface value.

## Implicit Implementation

Go uses **implicit interface satisfaction** - types satisfy interfaces automatically by implementing the required methods. No explicit declaration like `implements` is needed.

```go
// Circle also implements Shape without any declaration
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func main() {
    r := Rectangle{Width: 3, Height: 4}
    c := Circle{Radius: 5}
    
    // Both work with PrintMeasurements
    PrintMeasurements(r)
    PrintMeasurements(c)
}
```

## Value vs. Pointer Receivers with Interfaces

Methods with pointer receivers implement interfaces only for pointer values.

```go
type Mover interface {
    Move(distance float64)
}

type Car struct {
    Position float64
}

// Pointer receiver
func (c *Car) Move(distance float64) {
    c.Position += distance
}

func main() {
    carPtr := &Car{}      // Works: pointer satisfies Mover
    var m Mover = carPtr  // Valid
    
    car := Car{}          // DOESN'T WORK with Mover
    // var m2 Mover = car // Compile error: Car does not implement Mover
}
```

## Common Patterns

### The Empty Interface

```go
// Empty interface can hold any value
func PrintAny(v interface{}) {
    fmt.Println(v)
}

// Type assertions to access specific type behavior
func ProcessValue(v interface{}) {
    if str, ok := v.(string); ok {
        fmt.Println("String length:", len(str))
    } else if num, ok := v.(int); ok {
        fmt.Println("Number plus one:", num+1)
    }
}
```

### Type Switches

```go
func HandleValue(v interface{}) {
    switch x := v.(type) {
    case string:
        fmt.Println("String:", x)
    case int:
        fmt.Println("Integer:", x)
    case []byte:
        fmt.Println("Byte slice of length", len(x))
    default:
        fmt.Println("Unknown type")
    }
}
```

## Best Practices

1. **Keep interfaces small**: Prefer many small interfaces over large ones
2. **Interface as needed**: Define interfaces where you need them, not on the implementation types
3. **Accept interfaces, return structs**: Functions should accept interfaces and return concrete types
4. **Receiver consistency**: For a given type, use all pointer receivers or all value receivers for consistency

## Important Notes

1. **Zero Value**: The zero value of an interface is `nil` (no type, no value)
2. **Comparison**: Interface values are comparable when their dynamic types are comparable
3. **Embedding**: Interfaces can embed other interfaces
4. **Value vs. Pointer Receivers**: Choose based on these principles:
   - Use **pointer receivers** when methods modify the receiver
   - Use **pointer receivers** for large structs to avoid copies
   - Use **pointer receivers** for consistency when some methods require them
   - Use **value receivers** for immutable values or small structs where copies are cheap
5. **Not a Substitute for Generics**: Unlike languages with inheritance-based polymorphism, Go doesn't support generic types through interfaces. If coming from an OOP background, avoid trying to use interfaces to simulate generics or inheritance hierarchies. Go interfaces represent behavior only, not type hierarchies.

    ```go
    // Don't try to use interface{} as a generic container
    // This loses type safety and requires type assertions
    func ProcessItems(items []interface{}) {
        // Requires type assertions to do anything useful
    }
    
    // Instead, use concrete types or define interfaces with specific behaviors
    type Processor interface {
        Process()
    }
    
    func ProcessAll(items []Processor) {
        for _, item := range items {
            item.Process() // Type-safe behavior
        }
    }
    ```
