# Go Variables

Variables in Go are containers for storing data values. They provide a way to name and reference memory locations where values are stored.

## What You'll Learn

- Variable declaration and initialization
- Different ways to declare variables
- Type inference and explicit typing
- Variable scope and lifetime
- Zero values
- Constants vs variables
- Best practices and common patterns

## Key Concepts

- **Variable**: A name for a memory location where a value is stored
- **Declaration**: Creating a variable with a specific name and type
- **Initialization**: Giving a variable its first value
- **Scope**: Where in the code a variable can be accessed
- **Zero Value**: Default value for uninitialized variables

## Basic Variable Declaration

### Syntax

```go
var <name> <type> = <value>
```

### Examples

```go
var name string = "vali"
var age int = 30
var isActive bool = true
var price float64 = 99.99
```

## Different Declaration Methods

### 1. Explicit Type Declaration

```go
var name string = "Go"
var age int = 13
var isAwesome bool = true
var pi float64 = 3.14159
```

**When to use:**
- When you want to be explicit about the type
- When the type might not be obvious from the value
- In function parameters and return types

### 2. Type Inference

```go
var name = "Go"        // string
var age = 13           // int
var isAwesome = true   // bool
var pi = 3.14159       // float64
```

**When to use:**
- When the type is obvious from the value
- For cleaner, more readable code
- When you want Go to infer the type

### 3. Short Declaration (Inside Functions Only)

```go
func main() {
    name := "Go"        // string
    age := 13           // int
    isAwesome := true   // bool
    pi := 3.14159       // float64
}
```

**When to use:**
- Inside functions for local variables
- When you want concise syntax
- For most local variable declarations

### 4. Multiple Variable Declaration

```go
// Same type
var a, b, c int = 1, 2, 3

// Different types
var (
    name string = "Go"
    age  int    = 13
    pi   float64 = 3.14159
)

// Short declaration
func main() {
    name, age, pi := "Go", 13, 3.14159
}
```

## Zero Values

In Go, variables declared without initialization get their "zero value":

```go
var i int           // 0
var f float64       // 0.0
var b bool          // false
var s string        // ""
var p *int          // nil
var slice []int     // nil
var m map[string]int // nil
```

## Variable Scope

### Package Level Variables

```go
package main

import "fmt"

// Package-level variables
var globalName string = "Global"
var globalAge int = 25

func main() {
    fmt.Println(globalName, globalAge)
}
```

### Function Level Variables

```go
func example() {
    // Function-level variables
    var localName string = "Local"
    var localAge int = 30
    
    fmt.Println(localName, localAge)
}
```

### Block Level Variables

```go
func example() {
    if true {
        // Block-level variables
        var blockName string = "Block"
        fmt.Println(blockName)
    }
    // blockName is not accessible here
}
```

## Constants vs Variables

### Constants

```go
const (
    PI = 3.14159
    E = 2.71828
    Name = "Go"
)

// Typed constants
const MaxInt int = 2147483647
```

### Variables

```go
var (
    currentUser string = "admin"
    sessionCount int = 0
    isLoggedIn bool = false
)
```

## Practical Examples

### Example 1: User Information

```go
package main

import "fmt"

func main() {
    // Different ways to declare variables
    var firstName string = "John"
    var lastName = "Doe"
    age := 30
    email := "john.doe@example.com"
    
    // Multiple variables
    var (
        city    string = "New York"
        country string = "USA"
        zipCode int    = 10001
    )
    
    fmt.Printf("Name: %s %s\n", firstName, lastName)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Email: %s\n", email)
    fmt.Printf("Location: %s, %s %d\n", city, country, zipCode)
}
```

### Example 2: Calculator

```go
package main

import "fmt"

func main() {
    // Input values
    var num1, num2 float64 = 10.5, 3.2
    
    // Calculations
    sum := num1 + num2
    difference := num1 - num2
    product := num1 * num2
    quotient := num1 / num2
    
    // Display results
    fmt.Printf("Numbers: %.2f and %.2f\n", num1, num2)
    fmt.Printf("Sum: %.2f\n", sum)
    fmt.Printf("Difference: %.2f\n", difference)
    fmt.Printf("Product: %.2f\n", product)
    fmt.Printf("Quotient: %.2f\n", quotient)
}
```

### Example 3: Configuration

```go
package main

import "fmt"

// Package-level configuration
var (
    appName    string = "MyApp"
    version    string = "1.0.0"
    debugMode  bool   = false
    maxUsers   int    = 1000
    serverPort int    = 8080
)

func main() {
    // Local configuration
    var (
        databaseURL string = "localhost:5432"
        apiKey      string = "secret-key"
        timeout     int    = 30
    )
    
    fmt.Printf("App: %s v%s\n", appName, version)
    fmt.Printf("Debug: %t\n", debugMode)
    fmt.Printf("Max Users: %d\n", maxUsers)
    fmt.Printf("Server Port: %d\n", serverPort)
    fmt.Printf("Database: %s\n", databaseURL)
    fmt.Printf("API Key: %s\n", apiKey)
    fmt.Printf("Timeout: %d seconds\n", timeout)
}
```

## Variable Naming Conventions

### Rules

1. **Must start with a letter or underscore**
2. **Can contain letters, digits, and underscores**
3. **Case sensitive**
4. **Cannot use Go keywords**

### Examples

```go
// Valid names
var userName string
var user_age int
var _private string
var UserName string
var user123 int

// Invalid names
// var 123user int     // starts with number
// var user-name string // contains hyphen
// var var string      // Go keyword
```

### Naming Styles

```go
// camelCase (preferred for variables)
var userName string
var userAge int
var isActive bool

// PascalCase (for exported variables)
var UserName string
var UserAge int
var IsActive bool

// snake_case (less common in Go)
var user_name string
var user_age int
var is_active bool
```

## Common Patterns

### 1. Variable Swapping

```go
func main() {
    a, b := 10, 20
    fmt.Printf("Before: a=%d, b=%d\n", a, b)
    
    // Swap values
    a, b = b, a
    fmt.Printf("After: a=%d, b=%d\n", a, b)
}
```

### 2. Multiple Return Values

```go
func getValues() (string, int, bool) {
    return "Go", 13, true
}

func main() {
    name, age, isAwesome := getValues()
    fmt.Printf("%s is %d years old and awesome: %t\n", name, age, isAwesome)
}
```

### 3. Blank Identifier

```go
func main() {
    // When you don't need all return values
    name, _, isAwesome := getValues()
    fmt.Printf("%s is awesome: %t\n", name, isAwesome)
}
```

## Best Practices

### 1. Use Short Declaration When Possible

```go
// Good
func main() {
    name := "Go"
    age := 13
}

// Less preferred
func main() {
    var name string = "Go"
    var age int = 13
}
```

### 2. Group Related Variables

```go
// Good
var (
    userName string = "admin"
    userAge  int    = 30
    userRole string = "administrator"
)

// Less preferred
var userName string = "admin"
var userAge int = 30
var userRole string = "administrator"
```

### 3. Use Meaningful Names

```go
// Good
var userCount int
var maxRetries int
var connectionTimeout time.Duration

// Less preferred
var c int
var m int
var t time.Duration
```

### 4. Initialize Variables When Possible

```go
// Good
var name string = "Go"
var count int = 0

// Less preferred (unless you specifically need zero values)
var name string
var count int
```

## Common Mistakes

### 1. Unused Variables

```go
func main() {
    var unused string = "This will cause an error"
    // Error: unused variable
}
```

**Solution:**
```go
func main() {
    var unused string = "This will cause an error"
    _ = unused // Use blank identifier
    // Or simply use the variable
    fmt.Println(unused)
}
```

### 2. Redeclaration

```go
func main() {
    name := "Go"
    name := "Python" // Error: no new variables on left side
}
```

**Solution:**
```go
func main() {
    name := "Go"
    name = "Python" // Assignment, not declaration
}
```

### 3. Short Declaration Outside Functions

```go
// Error: short declaration outside function
name := "Go"

// Correct
var name string = "Go"
```

## Type Conversion

```go
func main() {
    var i int = 42
    var f float64 = float64(i)
    var s string = string(i) // Note: converts to Unicode character
    
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float64: %.2f\n", f)
    fmt.Printf("string: %s\n", s)
}
```

## Exercises

### Exercise 1: Personal Information
Create variables for your name, age, city, and country. Print them in a formatted way.

### Exercise 2: Calculator Variables
Create variables for two numbers and calculate their sum, difference, product, and quotient.

### Exercise 3: Configuration
Create a configuration structure using package-level variables for an application.

### Exercise 4: Variable Swapping
Create two variables and swap their values without using a temporary variable.

## Next Steps

Now that you understand variables, explore:

1. **[Data Types](../02.data-types)** - Learn about different data types in Go
2. **[Operators](../03.operators)** - Perform operations on variables
3. **[Conditional Statements](../04.conditional-statements)** - Control program flow
4. **[Loops](../05.loops)** - Repeat operations

## Key Takeaways

- Variables store values in memory
- Go has multiple ways to declare variables
- Type inference makes code cleaner
- Short declaration is preferred for local variables
- Variables have zero values when uninitialized
- Scope determines where variables can be accessed
- Use meaningful names and group related variables
- Avoid unused variables and redeclaration errors

## Additional Resources

- [Go Tour - Variables](https://tour.golang.org/basics/8)
- [Effective Go - Variables](https://golang.org/doc/effective_go.html#variables)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Go Language Specification - Variables](https://golang.org/ref/spec#Variables)