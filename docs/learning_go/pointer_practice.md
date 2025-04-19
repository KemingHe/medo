# Go Pointer Practice: Pass-by-Value

## Code Snippet

Consider the following Go code:

```go
package main

import "fmt"

func main() {
  name := "bill"

  // namePointer holds the memory address of 'name'
  namePointer := &name

  // Print the memory address *of the pointer variable itself* in main
  fmt.Println(&namePointer)
  printPointer(namePointer)
}

// printPointer receives a *copy* of the pointer
func printPointer(namePointer *string) {
  // Print the memory address *of the function's copy* of the pointer variable
  fmt.Println(&namePointer)
}
```

## Question

Will the memory addresses printed by the two `fmt.Println(&namePointer)` calls be the same or different? Why?

## Answer

The memory addresses printed will be **different**.

## Explanation: Pass-by-Value

Go is strictly **pass-by-value**. When you pass *any* argument to a function, including a pointer, Go makes a **copy** of that argument for the function to use.

1. **In `main`**:
    * `name := "bill"`: Creates a string variable `name` at some memory address (e.g., `0x100`).
    * `namePointer := &name`: Creates a pointer variable `namePointer` (type `*string`). This variable itself resides at its own address (e.g., `0x200`) and holds the value `0x100` (the address of `name`).
    * `fmt.Println(&namePointer)` prints the address of the `namePointer` variable in `main`, which is `0x200`.

2. **In `printPointer`**:
    * When `printPointer(namePointer)` is called, Go copies the *value* inside `main`'s `namePointer` (which is `0x100`).
    * A *new* pointer variable, also named `namePointer` but local to `printPointer`, is created. This new variable resides at a *different* memory address (e.g., `0x300` on the function's stack).
    * This new variable also holds the value `0x100`.
    * `fmt.Println(&namePointer)` inside the function prints the address of *this function's local copy* of the pointer variable, which is `0x300`.

**Key Takeaway:** Passing a pointer allows a function to modify the original value the pointer *points to*, but the function receives a *copy of the pointer variable itself*. The pointer variable inside the function and the pointer variable outside the function are distinct entities residing at different memory locations.

## Memory Visualization (Step-by-Step)

1. **After `name := "bill"` in `main`:**
    A variable `name` is created.

    | Scope | Variable | Address | Value |
    | :--- | :--- | :--- | :--- |
    | main | `name` | `0x100` | `"bill"` |

2. **After `namePointer := &name` in `main`:**
    A pointer variable `namePointer` is created. It stores the address of `name`.

    | Scope | Variable | Address | Value |
    | :--- | :--- | :--- | :--- |
    | main | `name` | `0x100` | `"bill"` |
    | main | `namePointer` | `0x200` | `0x100` (Address of `name`) |

3. **During `fmt.Println(&namePointer)` in `main`:**
    The address of `main`'s `namePointer` variable (`0x200`) is printed.

    | Scope | Variable | Address | Value |
    | :--- | :--- | :--- | :--- |
    | main | `name` | `0x100` | `"bill"` |
    | main | `namePointer` | `0x200` | `0x100` (Address of `name`) |

4. **When `printPointer(namePointer)` is called (Inside `printPointer`):**
    Go copies the *value* of `main`'s `namePointer` (`0x100`) into a new variable, also named `namePointer`, local to the function. This *new* variable gets its own address.

    | Scope | Variable | Address | Value |
    | :--- | :--- | :--- | :--- |
    | main | `name` | `0x100` | `"bill"` |
    | main | `namePointer` | `0x200` | `0x100` (Address of `name`) |
    | printPointer | `namePointer` | `0x300` (New Addr) | `0x100` (Copied Address of `name`) |

5. **During `fmt.Println(&namePointer)` in `printPointer`:**
    The address of the *function's local copy* of `namePointer` (`0x300`) is printed.

    | Scope | Variable | Address | Value |
    | :--- | :--- | :--- | :--- |
    | main | `name` | `0x100` | `"bill"` |
    | main | `namePointer` | `0x200` | `0x100` (Address of `name`) |
    | printPointer | `namePointer` | `0x300` (New Addr) | `0x100` (Copied Address of `name`) |
