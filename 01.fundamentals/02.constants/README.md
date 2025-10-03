# Go Constants

Constants in Go are immutable values that are known at compile time. They provide a way to define values that should not change during program execution, making your code more maintainable and less error-prone.

## What You'll Learn

- Constant declaration and initialization
- Different types of constants
- Constant expressions and calculations
- Typed vs untyped constants
- Constant blocks and iota
- Best practices and common patterns
- When to use constants vs variables

## Key Concepts

- **Constant**: An immutable value known at compile time
- **Declaration**: Creating a constant with a specific name and value
- **Expression**: Constants can be calculated from other constants
- **Typed vs Untyped**: Constants can have explicit types or be untyped
- **iota**: A special constant generator for creating sequences

## Basic Constant Declaration

### Syntax

```go
const <name> <type> = <value>
const <name> = <value>  // Type inferred
```

### Examples

```go
const Pi float64 = 3.14159
const Name string = "Go"
const MaxUsers int = 1000
const IsActive bool = true
```

## Different Declaration Methods

### 1. Explicit Type Declaration

```go
const Pi float64 = 3.14159
const Name string = "Go"
const MaxUsers int = 1000
const IsActive bool = true
```

**When to use:**
- When you want to be explicit about the type
- When the type might not be obvious from the value
- For better code documentation

### 2. Type Inference

```go
const Pi = 3.14159        // float64
const Name = "Go"         // string
const MaxUsers = 1000     // int
const IsActive = true     // bool
```

**When to use:**
- When the type is obvious from the value
- For cleaner, more readable code
- When you want Go to infer the type

### 3. Multiple Constants

```go
// Individual declarations
const Pi = 3.14159
const E = 2.71828
const Name = "Go"

// Grouped declarations
const (
    Pi = 3.14159
    E = 2.71828
    Name = "Go"
)
```

## Constant Expressions

Constants can be calculated from other constants and literals:

```go
const (
    Pi = 3.14159
    Radius = 5.0
    Area = Pi * Radius * Radius  // 78.53975
    Circumference = 2 * Pi * Radius  // 31.4159
)

// Mathematical expressions
const (
    SecondsPerMinute = 60
    MinutesPerHour = 60
    SecondsPerHour = SecondsPerMinute * MinutesPerHour  // 3600
)

// String concatenation
const (
    Greeting = "Hello"
    Name = "World"
    Message = Greeting + ", " + Name + "!"  // "Hello, World!"
)
```

## Typed vs Untyped Constants

### Typed Constants

```go
const Pi float64 = 3.14159
const MaxUsers int = 1000
const Name string = "Go"
```

**Characteristics:**
- Have an explicit type
- Cannot be mixed with other types without conversion
- More restrictive but safer

### Untyped Constants

```go
const Pi = 3.14159
const MaxUsers = 1000
const Name = "Go"
```

**Characteristics:**
- No explicit type
- Can be used with any compatible type
- More flexible but less type-safe

### Type Flexibility Example

```go
const untyped = 42
const typed int = 42

func main() {
    var i int = untyped     // OK: untyped constant
    var f float64 = untyped // OK: untyped constant
    var s string = untyped  // Error: cannot convert
    
    var i2 int = typed      // OK: same type
    var f2 float64 = typed  // Error: type mismatch
}
```

## The iota Constant Generator

`iota` is a special constant that generates a sequence of values:

### Basic iota Usage

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### iota with Expressions

```go
const (
    _ = iota                    // 0 (ignored)
    KB = 1 << (10 * iota)      // 1024
    MB                          // 1048576
    GB                          // 1073741824
    TB                          // 1099511627776
)

// Custom starting value
const (
    First = iota + 1    // 1
    Second              // 2
    Third               // 3
)
```

### iota in Different Blocks

```go
const (
    A = iota    // 0
    B           // 1
    C           // 2
)

const (
    D = iota    // 0 (resets in new block)
    E           // 1
    F           // 2
)
```

## Practical Examples

### Example 1: Application Configuration

```go
package main

import "fmt"

// Application constants
const (
    AppName    = "MyApp"
    Version    = "1.0.0"
    MaxUsers   = 1000
    Timeout    = 30
    DebugMode  = false
)

// HTTP status codes
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// File permissions
const (
    ReadPermission  = 0o444
    WritePermission = 0o644
    ExecPermission  = 0o755
)

func main() {
    fmt.Printf("App: %s v%s\n", AppName, Version)
    fmt.Printf("Max Users: %d\n", MaxUsers)
    fmt.Printf("Timeout: %d seconds\n", Timeout)
    fmt.Printf("Debug Mode: %t\n", DebugMode)
    
    fmt.Printf("HTTP Status OK: %d\n", StatusOK)
    fmt.Printf("Read Permission: %o\n", ReadPermission)
}
```

### Example 2: Mathematical Constants

```go
package main

import "fmt"

// Mathematical constants
const (
    Pi = 3.141592653589793
    E = 2.718281828459045
    Phi = 1.618033988749895  // Golden ratio
    Sqrt2 = 1.4142135623730951
)

// Unit conversions
const (
    MetersPerKilometer = 1000
    CentimetersPerMeter = 100
    MillimetersPerCentimeter = 10
    MillimetersPerMeter = CentimetersPerMeter * MillimetersPerCentimeter
)

// Temperature conversions
const (
    FreezingPointC = 0
    BoilingPointC = 100
    FreezingPointF = 32
    BoilingPointF = 212
)

func main() {
    fmt.Printf("Pi: %.15f\n", Pi)
    fmt.Printf("E: %.15f\n", E)
    fmt.Printf("Golden Ratio: %.15f\n", Phi)
    
    fmt.Printf("Millimeters per meter: %d\n", MillimetersPerMeter)
    
    fmt.Printf("Water freezes at %d°C or %d°F\n", FreezingPointC, FreezingPointF)
    fmt.Printf("Water boils at %d°C or %d°F\n", BoilingPointC, BoilingPointF)
}
```

### Example 3: Enum-like Constants

```go
package main

import "fmt"

// User roles
const (
    RoleGuest = iota
    RoleUser
    RoleModerator
    RoleAdmin
    RoleSuperAdmin
)

// User status
const (
    StatusInactive = iota
    StatusActive
    StatusSuspended
    StatusBanned
)

// Priority levels
const (
    PriorityLow = iota + 1
    PriorityMedium
    PriorityHigh
    PriorityCritical
)

func main() {
    fmt.Printf("User roles: Guest=%d, User=%d, Moderator=%d, Admin=%d, SuperAdmin=%d\n",
        RoleGuest, RoleUser, RoleModerator, RoleAdmin, RoleSuperAdmin)
    
    fmt.Printf("User status: Inactive=%d, Active=%d, Suspended=%d, Banned=%d\n",
        StatusInactive, StatusActive, StatusSuspended, StatusBanned)
    
    fmt.Printf("Priority levels: Low=%d, Medium=%d, High=%d, Critical=%d\n",
        PriorityLow, PriorityMedium, PriorityHigh, PriorityCritical)
}
```

### Example 4: Bit Flags

```go
package main

import "fmt"

// Permission flags
const (
    ReadPermission = 1 << iota  // 1
    WritePermission             // 2
    ExecutePermission           // 4
    DeletePermission            // 8
    AdminPermission             // 16
)

// File types
const (
    FileTypeText = 1 << iota    // 1
    FileTypeImage               // 2
    FileTypeVideo               // 4
    FileTypeAudio               // 8
    FileTypeDocument            // 16
)

func main() {
    // Check permissions
    userPermissions := ReadPermission | WritePermission
    fmt.Printf("User permissions: %d\n", userPermissions)
    
    if userPermissions&ReadPermission != 0 {
        fmt.Println("User has read permission")
    }
    
    if userPermissions&ExecutePermission != 0 {
        fmt.Println("User has execute permission")
    } else {
        fmt.Println("User does not have execute permission")
    }
    
    // File type checking
    fileType := FileTypeText | FileTypeDocument
    fmt.Printf("File type: %d\n", fileType)
    
    if fileType&FileTypeText != 0 {
        fmt.Println("File is a text file")
    }
}
```

## Constants vs Variables

### Constants

```go
const (
    Pi = 3.14159
    MaxUsers = 1000
    AppName = "MyApp"
)
```

**Characteristics:**
- Immutable (cannot be changed)
- Known at compile time
- Can be used in type declarations
- More efficient (no memory allocation)

### Variables

```go
var (
    currentUser string = "admin"
    sessionCount int = 0
    isLoggedIn bool = false
)
```

**Characteristics:**
- Mutable (can be changed)
- Known at runtime
- Memory allocated
- More flexible

## Best Practices

### 1. Use Descriptive Names

```go
// Good
const (
    MaxRetries = 3
    TimeoutSeconds = 30
    DefaultPort = 8080
)

// Less preferred
const (
    MR = 3
    TS = 30
    DP = 8080
)
```

### 2. Group Related Constants

```go
// Good
const (
    // HTTP status codes
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
    
    // Timeouts
    ReadTimeout = 30
    WriteTimeout = 30
    ConnectTimeout = 10
)
```

### 3. Use iota for Sequences

```go
// Good
const (
    StateIdle = iota
    StateRunning
    StatePaused
    StateStopped
)

// Less preferred
const (
    StateIdle = 0
    StateRunning = 1
    StatePaused = 2
    StateStopped = 3
)
```

### 4. Prefer Untyped Constants

```go
// Good
const Pi = 3.14159

// Less preferred (unless you need specific type)
const Pi float64 = 3.14159
```

### 5. Use Constants for Magic Numbers

```go
// Good
const MaxBufferSize = 1024
const DefaultTimeout = 30

func processData(data []byte) {
    if len(data) > MaxBufferSize {
        // Handle error
    }
}

// Less preferred
func processData(data []byte) {
    if len(data) > 1024 {  // Magic number
        // Handle error
    }
}
```

## Common Patterns

### 1. Configuration Constants

```go
const (
    // Database
    DBHost = "localhost"
    DBPort = 5432
    DBName = "myapp"
    
    // API
    APIVersion = "v1"
    APIPrefix = "/api/" + APIVersion
    
    // Security
    JWTSecret = "your-secret-key"
    TokenExpiry = 3600  // seconds
)
```

### 2. Error Codes

```go
const (
    ErrCodeSuccess = iota
    ErrCodeNotFound
    ErrCodeUnauthorized
    ErrCodeForbidden
    ErrCodeInternalError
)
```

### 3. Feature Flags

```go
const (
    FeatureNewUI = iota
    FeatureBetaAPI
    FeatureAdvancedSearch
    FeatureAnalytics
)
```

### 4. Unit Conversions

```go
const (
    // Time
    SecondsPerMinute = 60
    MinutesPerHour = 60
    HoursPerDay = 24
    DaysPerWeek = 7
    
    // Derived constants
    SecondsPerHour = SecondsPerMinute * MinutesPerHour
    SecondsPerDay = SecondsPerHour * HoursPerDay
    SecondsPerWeek = SecondsPerDay * DaysPerWeek
)
```

## Common Mistakes

### 1. Trying to Modify Constants

```go
const Pi = 3.14159
Pi = 3.14  // Error: cannot assign to Pi
```

**Solution:**
```go
const Pi = 3.14159
// Use Pi as-is, or create a new constant
const PiApprox = 3.14
```

### 2. Using Variables for Constants

```go
// Less preferred
var MaxUsers = 1000

// Better
const MaxUsers = 1000
```

### 3. Inconsistent iota Usage

```go
// Inconsistent
const (
    A = iota    // 0
    B = 5       // 5
    C           // 5 (not 2!)
)

// Consistent
const (
    A = iota    // 0
    B = iota    // 1
    C = iota    // 2
)
```

### 4. Missing iota Reset

```go
// Confusing
const (
    A = iota    // 0
    B           // 1
)

const (
    C = iota    // 0 (resets!)
    D           // 1
)
```

## Advanced Topics

### 1. Constant Expressions with Functions

```go
const (
    // These work because they're evaluated at compile time
    MaxInt = 1<<31 - 1
    MinInt = -1 << 31
    MaxUint = 1<<32 - 1
)
```

### 2. String Constants

```go
const (
    Greeting = "Hello, World!"
    Template = "User: %s, Age: %d"
    Path = "/api/v1/users"
)
```

### 3. Boolean Constants

```go
const (
    DebugMode = true
    ProductionMode = false
    EnableLogging = true
)
```

## Exercises

### Exercise 1: Mathematical Constants
Create constants for common mathematical values and use them in calculations.

### Exercise 2: Application Configuration
Define constants for an application's configuration (database, API, security).

### Exercise 3: Enum-like Constants
Create constants for user roles, statuses, and priority levels using iota.

### Exercise 4: Bit Flags
Implement permission flags using bit operations and constants.

### Exercise 5: Unit Conversions
Create constants for unit conversions and use them in calculations.

## Next Steps

Now that you understand constants, explore:

1. **[Data Types](../02.data-types)** - Learn about different data types in Go
2. **[Operators](../03.operators)** - Perform operations on constants and variables
3. **[Conditional Statements](../04.conditional-statements)** - Control program flow
4. **[Loops](../05.loops)** - Repeat operations

## Key Takeaways

- Constants are immutable values known at compile time
- Use constants for values that should not change
- Constants can be typed or untyped
- iota generates sequences of values
- Constants are more efficient than variables
- Group related constants together
- Use descriptive names for constants
- Prefer untyped constants for flexibility

## Additional Resources

- [Go Tour - Constants](https://tour.golang.org/basics/15)
- [Effective Go - Constants](https://golang.org/doc/effective_go.html#constants)
- [Go by Example - Constants](https://gobyexample.com/constants)
- [Go Language Specification - Constants](https://golang.org/ref/spec#Constants)
