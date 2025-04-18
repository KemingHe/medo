# Naming Conventions in Go

Go's naming conventions are crucial for readability and determining visibility. Clarity and conciseness are key.

## 1. Exported vs. Unexported (Visibility)

The **first letter's case** determines if an identifier (variable, constant, type, function, struct field) is visible outside its package.

- **`PascalCase` (Starts Uppercase): Exported**
  - Accessible from other packages. Defines the public API.

  ```go
  package mypkg

  var MaxConnections int // Exported variable
  type User struct {     // Exported type
      Name string        // Exported field
  }
  func CalculateTotal() {} // Exported function
  ```

- **`camelCase` (Starts Lowercase): Unexported**
  - Only accessible within the *same* package. Internal implementation detail.

  ```go
  package mypkg

  var maxConnections int // Unexported variable
  type user struct {     // Unexported type
      name string        // Unexported field
  }
  func calculateTotal() {} // Unexported function
  ```

## 2. Package Names

- **Short, concise, lowercase, single word.**
- Avoid underscores (`_`) or `mixedCaps`.
- Name matches the directory (e.g., code in `net/http` is `package http`).
- Consumers use the package name (e.g., `http.Get(...)`).
- Avoid generic names like `util` or `common`.

## 3. File Names

- Typically **`snake_case.go`** (e.g., `http_client.go`, `deck_test.go`).
- File names don't affect the Go package itself but should describe content.
- Test files **must** end with `_test.go`.

## 4. Variable, Function, Method, Type Names

- Use `PascalCase` or `camelCase` based on desired visibility (see rule #1).
- Names should be descriptive but reasonably concise.
- Multi-word names use capitalization: `maxRetries`, `ReadTimeout`, `ServeHTTP`.
- **No underscores** (`_`) in these names (except the blank identifier).
- Short names (e.g., `i`, `buf`, `r`) are idiomatic for local variables with small scope.

## 5. Constants

- Follow variable naming rules (`PascalCase`/`camelCase`).
- **Do NOT use `ALL_CAPS_SNAKE_CASE`** like in some other languages.

  ```go
  // Correct:
  const DefaultTimeout = 5 // Exported
  const defaultTimeout = 5 // Unexported

  // Incorrect:
  // const DEFAULT_TIMEOUT = 5
  ```

## 6. Interface Names

- Single-method interfaces often use the method name + `er` suffix: `io.Reader`, `fmt.Stringer`.
- Multi-method interfaces use descriptive nouns: `net.Conn`, `http.Handler`.

## 7. Acronyms

- Treat acronyms (`URL`, `ID`, `HTTP`) as whole words in names. Keep case consistent.

  ```go
  // Examples:
  ServeHTTP(...)
  productID
  parseJSON(...)
  MyURL
  customerID
  ```

  (Linters often prefer `ID` over `Id`).

## 8. Getters

- Usually omit the `Get` prefix.

  ```go
  type User struct { name string }

  // Getter for name field
  func (u *User) Name() string { // Not GetName()
      return u.name
  }
  ```
