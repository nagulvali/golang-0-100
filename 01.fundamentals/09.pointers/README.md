# Go Pointers

Pointers in Go are variables that store the memory address of another variable. They provide a way to reference and manipulate data indirectly, enabling efficient memory management and the ability to modify data in place.

## What You'll Learn

- Pointer declaration and initialization
- Address-of and dereference operators
- Pointer types and zero values
- Passing pointers to functions
- Pointer arithmetic (limited in Go)
- Pointers with structs and arrays
- Pointer safety and nil pointers
- Best practices and common patterns
- When to use pointers vs values

## Key Concepts

- **Pointer**: A variable that stores a memory address
- **Address**: The memory location where a value is stored
- **Dereference**: Accessing the value at a memory address
- **Reference**: Creating a pointer to a variable
- **Nil Pointer**: A pointer that doesn't point to any memory location

## Basic Pointer Operations

### Address-of Operator (&)

```go
var x int = 42
var p *int = &x  // p points to x

fmt.Printf("Value of x: %d\n", x)        // 42
fmt.Printf("Address of x: %p\n", &x)     // 0xc000014080
fmt.Printf("Value of p: %p\n", p)        // 0xc000014080
fmt.Printf("Value at p: %d\n", *p)       // 42
```

### Dereference Operator (*)

```go
var x int = 42
var p *int = &x

// Modify value through pointer
*p = 100

fmt.Printf("x is now: %d\n", x)  // 100
```

## Pointer Declaration and Initialization

### Basic Declaration

```go
// Declare a pointer
var p *int

// Initialize with address
var x int = 42
p = &x

// Short declaration
var x int = 42
p := &x
```

### Different Pointer Types

```go
// Integer pointer
var intPtr *int

// String pointer
var strPtr *string

// Struct pointer
type Person struct {
    Name string
    Age  int
}
var personPtr *Person

// Slice pointer
var slicePtr *[]int

// Map pointer
var mapPtr *map[string]int
```

### Zero Value

```go
var p *int
fmt.Printf("Zero value of pointer: %v\n", p)  // <nil>
fmt.Printf("Is nil: %t\n", p == nil)          // true
```

## Pointer Operations

### Basic Operations

```go
func main() {
    var x int = 42
    var p *int = &x
    
    // Get address
    fmt.Printf("Address of x: %p\n", &x)
    fmt.Printf("Value of p: %p\n", p)
    
    // Get value through pointer
    fmt.Printf("Value of x: %d\n", x)
    fmt.Printf("Value at p: %d\n", *p)
    
    // Modify through pointer
    *p = 100
    fmt.Printf("x after modification: %d\n", x)  // 100
    
    // Check if pointer is nil
    if p != nil {
        fmt.Println("Pointer is not nil")
    }
}
```

### Pointer Comparison

```go
func main() {
    var x, y int = 42, 42
    var p1, p2 *int = &x, &y
    
    // Compare values
    fmt.Printf("x == y: %t\n", x == y)           // true
    
    // Compare pointers
    fmt.Printf("p1 == p2: %t\n", p1 == p2)       // false (different addresses)
    fmt.Printf("p1 == &x: %t\n", p1 == &x)       // true
    fmt.Printf("*p1 == *p2: %t\n", *p1 == *p2)   // true (same values)
}
```

## Pointers with Functions

### Passing Pointers to Functions

```go
// Function that modifies value through pointer
func increment(p *int) {
    *p++
}

// Function that returns a pointer
func getPointer(x int) *int {
    return &x
}

func main() {
    var x int = 10
    
    // Pass pointer to function
    increment(&x)
    fmt.Printf("x after increment: %d\n", x)  // 11
    
    // Get pointer from function
    p := getPointer(42)
    fmt.Printf("Value at pointer: %d\n", *p)  // 42
}
```

### Returning Pointers

```go
// Function that returns a pointer to a new value
func newInt(x int) *int {
    return &x
}

// Function that returns a pointer to existing value
func getAddress(x *int) *int {
    return x
}

func main() {
    // Create new integer
    p1 := newInt(100)
    fmt.Printf("New integer: %d\n", *p1)
    
    // Get address of existing variable
    var x int = 200
    p2 := getAddress(&x)
    fmt.Printf("Address of x: %d\n", *p2)
    
    // Modify through returned pointer
    *p2 = 300
    fmt.Printf("x is now: %d\n", x)  // 300
}
```

## Pointers with Structs

### Basic Struct Pointers

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create struct
    person := Person{Name: "Alice", Age: 30}
    
    // Create pointer to struct
    p := &person
    
    // Access fields through pointer
    fmt.Printf("Name: %s\n", (*p).Name)  // Explicit dereference
    fmt.Printf("Age: %d\n", p.Age)       // Implicit dereference (Go allows this)
    
    // Modify fields through pointer
    p.Name = "Bob"
    p.Age = 35
    
    fmt.Printf("Updated person: %+v\n", person)  // {Name:Bob Age:35}
}
```

### Struct Methods with Pointers

```go
type Counter struct {
    value int
}

// Method with value receiver
func (c Counter) GetValue() int {
    return c.value
}

// Method with pointer receiver
func (c *Counter) Increment() {
    c.value++
}

// Method with pointer receiver
func (c *Counter) SetValue(v int) {
    c.value = v
}

func main() {
    counter := Counter{value: 0}
    
    // Call methods
    fmt.Printf("Initial value: %d\n", counter.GetValue())  // 0
    
    counter.Increment()
    fmt.Printf("After increment: %d\n", counter.GetValue())  // 1
    
    counter.SetValue(10)
    fmt.Printf("After set: %d\n", counter.GetValue())  // 10
}
```

## Pointers with Arrays and Slices

### Array Pointers

```go
func main() {
    // Array
    arr := [3]int{1, 2, 3}
    
    // Pointer to array
    p := &arr
    
    // Access elements through pointer
    fmt.Printf("First element: %d\n", (*p)[0])  // 1
    fmt.Printf("Second element: %d\n", p[1])    // 2 (Go allows this)
    
    // Modify elements through pointer
    p[0] = 100
    fmt.Printf("Modified array: %v\n", arr)  // [100 2 3]
}
```

### Slice Pointers

```go
func main() {
    // Slice
    slice := []int{1, 2, 3, 4, 5}
    
    // Pointer to slice
    p := &slice
    
    // Access elements through pointer
    fmt.Printf("First element: %d\n", (*p)[0])  // 1
    fmt.Printf("Length: %d\n", len(*p))         // 5
    
    // Modify elements through pointer
    (*p)[0] = 100
    fmt.Printf("Modified slice: %v\n", slice)  // [100 2 3 4 5]
    
    // Append through pointer
    *p = append(*p, 6)
    fmt.Printf("After append: %v\n", slice)  // [100 2 3 4 5 6]
}
```

## Pointer Safety

### Nil Pointer Checks

```go
func safeDereference(p *int) {
    if p != nil {
        fmt.Printf("Value: %d\n", *p)
    } else {
        fmt.Println("Pointer is nil")
    }
}

func main() {
    var p *int
    safeDereference(p)  // Pointer is nil
    
    var x int = 42
    p = &x
    safeDereference(p)  // Value: 42
}
```

### Nil Pointer Dereference

```go
func main() {
    var p *int
    
    // This will cause a panic
    // fmt.Printf("Value: %d\n", *p)  // Panic: runtime error: invalid memory address or nil pointer dereference
    
    // Safe way
    if p != nil {
        fmt.Printf("Value: %d\n", *p)
    } else {
        fmt.Println("Pointer is nil")
    }
}
```

## Practical Examples

### Example 1: Swap Function

```go
// Swap two integers using pointers
func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 10, 20
    fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
    
    swap(&x, &y)
    fmt.Printf("After swap: x=%d, y=%d\n", x, y)
}
```

### Example 2: Linked List Node

```go
type Node struct {
    Value int
    Next  *Node
}

func main() {
    // Create nodes
    node1 := &Node{Value: 1}
    node2 := &Node{Value: 2}
    node3 := &Node{Value: 3}
    
    // Link nodes
    node1.Next = node2
    node2.Next = node3
    
    // Traverse list
    current := node1
    for current != nil {
        fmt.Printf("Node value: %d\n", current.Value)
        current = current.Next
    }
}
```

### Example 3: Function that Modifies Slice

```go
// Function that modifies slice in place
func doubleSlice(s *[]int) {
    for i := range *s {
        (*s)[i] *= 2
    }
}

// Function that appends to slice
func appendToSlice(s *[]int, value int) {
    *s = append(*s, value)
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Original: %v\n", numbers)
    
    // Double all values
    doubleSlice(&numbers)
    fmt.Printf("After doubling: %v\n", numbers)
    
    // Append new value
    appendToSlice(&numbers, 10)
    fmt.Printf("After append: %v\n", numbers)
}
```

### Example 4: Configuration Manager

```go
type Config struct {
    Host     string
    Port     int
    Debug    bool
    Timeout  int
}

type ConfigManager struct {
    config *Config
}

func NewConfigManager() *ConfigManager {
    return &ConfigManager{
        config: &Config{
            Host:    "localhost",
            Port:    8080,
            Debug:   false,
            Timeout: 30,
        },
    }
}

func (cm *ConfigManager) GetConfig() *Config {
    return cm.config
}

func (cm *ConfigManager) UpdateHost(host string) {
    cm.config.Host = host
}

func (cm *ConfigManager) UpdatePort(port int) {
    cm.config.Port = port
}

func main() {
    manager := NewConfigManager()
    
    // Get configuration
    config := manager.GetConfig()
    fmt.Printf("Initial config: %+v\n", *config)
    
    // Update configuration
    manager.UpdateHost("example.com")
    manager.UpdatePort(9090)
    
    fmt.Printf("Updated config: %+v\n", *config)
}
```

## Pointer Arithmetic (Limited)

Go doesn't support pointer arithmetic like C/C++. However, you can:

### Safe Pointer Operations

```go
func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    
    // Get pointer to first element
    p := &arr[0]
    fmt.Printf("First element: %d\n", *p)
    
    // Get pointer to second element
    p2 := &arr[1]
    fmt.Printf("Second element: %d\n", *p2)
    
    // You cannot do: p++ or p + 1
    // This is not allowed in Go
}
```

### Using unsafe Package (Advanced)

```go
import "unsafe"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    
    // Get pointer to first element
    p := &arr[0]
    
    // Use unsafe to get pointer to second element
    p2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(int(0))))
    
    fmt.Printf("First element: %d\n", *p)   // 1
    fmt.Printf("Second element: %d\n", *p2) // 2
    
    // Note: This is unsafe and not recommended for general use
}
```

## Best Practices

### 1. Use Pointers for Large Data Structures

```go
// Good - use pointer for large struct
type LargeStruct struct {
    Data [1000]int
    // ... many fields
}

func processLargeStruct(s *LargeStruct) {
    // Process without copying
}

// Less preferred - copying large struct
func processLargeStructCopy(s LargeStruct) {
    // Expensive copy operation
}
```

### 2. Use Pointers for Mutating Functions

```go
// Good - function that modifies data
func incrementCounter(c *Counter) {
    c.value++
}

// Less preferred - function that doesn't modify
func getCounterValue(c Counter) int {
    return c.value
}
```

### 3. Check for Nil Pointers

```go
// Good - always check for nil
func safeProcess(p *Data) {
    if p != nil {
        // Process data
    }
}

// Bad - potential nil pointer dereference
func unsafeProcess(p *Data) {
    // Process data without checking
}
```

### 4. Use Value Receivers When Possible

```go
type Point struct {
    X, Y int
}

// Good - value receiver for small struct
func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Less preferred - pointer receiver for simple operations
func (p *Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

## Common Patterns

### 1. Factory Functions

```go
type User struct {
    ID   int
    Name string
}

// Factory function that returns pointer
func NewUser(id int, name string) *User {
    return &User{
        ID:   id,
        Name: name,
    }
}

func main() {
    user := NewUser(1, "Alice")
    fmt.Printf("User: %+v\n", *user)
}
```

### 2. Optional Parameters

```go
type Config struct {
    Host    string
    Port    int
    Timeout int
}

// Function with optional parameters using pointers
func NewConfig(host string, port *int, timeout *int) *Config {
    config := &Config{
        Host: host,
        Port: 8080, // default
        Timeout: 30, // default
    }
    
    if port != nil {
        config.Port = *port
    }
    
    if timeout != nil {
        config.Timeout = *timeout
    }
    
    return config
}

func main() {
    // Use default port and timeout
    config1 := NewConfig("localhost", nil, nil)
    
    // Specify custom port
    port := 9090
    config2 := NewConfig("example.com", &port, nil)
    
    fmt.Printf("Config 1: %+v\n", *config1)
    fmt.Printf("Config 2: %+v\n", *config2)
}
```

### 3. Error Handling with Pointers

```go
type Result struct {
    Value int
    Error error
}

func divide(a, b int) *Result {
    if b == 0 {
        return &Result{Error: fmt.Errorf("division by zero")}
    }
    return &Result{Value: a / b}
}

func main() {
    result := divide(10, 2)
    if result.Error != nil {
        fmt.Printf("Error: %v\n", result.Error)
    } else {
        fmt.Printf("Result: %d\n", result.Value)
    }
}
```

## Common Mistakes

### 1. Nil Pointer Dereference

```go
// Error
var p *int
fmt.Printf("Value: %d\n", *p) // Panic

// Correct
var p *int
if p != nil {
    fmt.Printf("Value: %d\n", *p)
}
```

### 2. Returning Pointer to Local Variable

```go
// Problematic
func getPointer() *int {
    x := 42
    return &x // Returns pointer to local variable
}

// Better
func getPointer() *int {
    x := 42
    return &x // Go's escape analysis handles this
}
```

### 3. Modifying Immutable Values

```go
// This doesn't work as expected
func modifyString(s *string) {
    *s = "modified"
}

func main() {
    str := "original"
    modifyString(&str)
    fmt.Println(str) // "modified" - this actually works
}
```

### 4. Pointer to Interface

```go
// Problematic
var p *interface{} = &someInterface

// Better
var p interface{} = someInterface
```

## Exercises

### Exercise 1: Swap Function
Implement a function that swaps two integers using pointers.

### Exercise 2: Linked List
Create a simple linked list using pointers and implement basic operations.

### Exercise 3: Configuration Manager
Build a configuration manager that uses pointers to manage application settings.

### Exercise 4: Array Operations
Implement functions that modify arrays in place using pointers.

### Exercise 5: Optional Parameters
Create a function that accepts optional parameters using pointers.

## Next Steps

Now that you understand pointers, explore:

1. **[Functions](../10.functions)** - Learn about function parameters and return values
2. **[Structs](../11.structs)** - Define custom data types
3. **[Interfaces](../12.interfaces)** - Define behavior contracts
4. **[Methods](../13.methods)** - Add behavior to types

## Key Takeaways

- Pointers store memory addresses
- Use `&` to get address, `*` to dereference
- Pointers have zero value of `nil`
- Always check for nil pointers before dereferencing
- Use pointers for large data structures and mutating functions
- Go doesn't support pointer arithmetic
- Pointers are reference types
- Use value receivers when possible

## Additional Resources

- [Go Tour - Pointers](https://tour.golang.org/moretypes/1)
- [Effective Go - Pointers](https://golang.org/doc/effective_go.html#pointers)
- [Go by Example - Pointers](https://gobyexample.com/pointers)
- [Go Language Specification - Pointers](https://golang.org/ref/spec#Pointer_types)
