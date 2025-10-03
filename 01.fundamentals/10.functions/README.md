# Go Functions

Functions in Go are first-class citizens that allow you to organize code into reusable, modular blocks. They provide a way to encapsulate logic, accept parameters, and return values. Go functions support multiple return values, variadic parameters, and can be assigned to variables or passed as arguments.

## What You'll Learn

- Function declaration and definition
- Parameters and return values
- Multiple return values
- Variadic functions
- Anonymous functions and closures
- Function types and first-class functions
- Method receivers
- Function composition and higher-order functions
- Best practices and common patterns
- Error handling with functions

## Key Concepts

- **Function**: A reusable block of code that performs a specific task
- **Parameters**: Input values passed to a function
- **Return Values**: Output values returned by a function
- **Signature**: The function's name, parameters, and return types
- **First-class**: Functions can be assigned to variables and passed as arguments
- **Closure**: A function that captures variables from its surrounding scope

## Basic Function Declaration

### Syntax

```go
func <name>(<parameters>) <return-type> {
    // function body
    return <value>
}
```

### Examples

```go
// Simple function
func greet() {
    fmt.Println("Hello, World!")
}

// Function with parameters
func add(a, b int) int {
    return a + b
}

// Function with multiple parameters
func multiply(x, y int) int {
    return x * y
}

// Function with different parameter types
func formatName(firstName, lastName string) string {
    return firstName + " " + lastName
}
```

## Function Parameters

### Basic Parameters

```go
// Single parameter
func square(x int) int {
    return x * x
}

// Multiple parameters of same type
func add(a, b int) int {
    return a + b
}

// Multiple parameters of different types
func greet(name string, age int) {
    fmt.Printf("Hello %s, you are %d years old\n", name, age)
}

// Mixed parameter types
func processData(id int, name string, active bool) {
    fmt.Printf("ID: %d, Name: %s, Active: %t\n", id, name, active)
}
```

### Parameter Types

```go
// Value parameters (copied)
func modifyValue(x int) {
    x = 100 // This doesn't affect the original
}

// Pointer parameters (reference)
func modifyPointer(x *int) {
    *x = 100 // This affects the original
}

// Slice parameters (reference)
func modifySlice(s []int) {
    s[0] = 100 // This affects the original slice
}

// Map parameters (reference)
func modifyMap(m map[string]int) {
    m["key"] = 100 // This affects the original map
}
```

### Default Parameters (Not Supported)

Go doesn't support default parameters. Use function overloading or struct parameters instead:

```go
// Not supported in Go
// func greet(name string, greeting string = "Hello") { }

// Alternative: Use struct parameters
type GreetOptions struct {
    Name     string
    Greeting string
}

func greetWithOptions(opts GreetOptions) {
    if opts.Greeting == "" {
        opts.Greeting = "Hello" // Default value
    }
    fmt.Printf("%s, %s!\n", opts.Greeting, opts.Name)
}
```

## Return Values

### Single Return Value

```go
func add(a, b int) int {
    return a + b
}

func isEven(n int) bool {
    return n%2 == 0
}

func getMessage() string {
    return "Hello, World!"
}
```

### Multiple Return Values

```go
// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return // Naked return
}

// Mixed return types
func processUser(id int) (name string, age int, active bool) {
    // Simulate database lookup
    if id == 1 {
        return "Alice", 25, true
    }
    return "Unknown", 0, false
}
```

### Named Return Values

```go
// Named return values
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return // Naked return
}

// Named return values with different types
func getUserInfo(id int) (name string, age int, err error) {
    if id <= 0 {
        err = fmt.Errorf("invalid user ID")
        return
    }
    
    // Simulate database lookup
    name = "Alice"
    age = 25
    return
}
```

## Variadic Functions

### Basic Variadic Functions

```go
// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Variadic function with other parameters
func printInfo(prefix string, values ...interface{}) {
    fmt.Printf("%s: ", prefix)
    for _, value := range values {
        fmt.Printf("%v ", value)
    }
    fmt.Println()
}

// Variadic function with mixed types
func processItems(items ...string) {
    for i, item := range items {
        fmt.Printf("Item %d: %s\n", i+1, item)
    }
}
```

### Using Variadic Functions

```go
func main() {
    // Call with no arguments
    fmt.Println(sum()) // 0
    
    // Call with one argument
    fmt.Println(sum(5)) // 5
    
    // Call with multiple arguments
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
    
    // Call with slice
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(numbers...)) // 15
    
    // Print info
    printInfo("Debug", "value1", "value2", 42, true)
}
```

## Anonymous Functions and Closures

### Anonymous Functions

```go
func main() {
    // Anonymous function assigned to variable
    add := func(a, b int) int {
        return a + b
    }
    
    result := add(5, 3)
    fmt.Printf("Result: %d\n", result)
    
    // Anonymous function called immediately
    result2 := func(x, y int) int {
        return x * y
    }(4, 6)
    
    fmt.Printf("Result 2: %d\n", result2)
}
```

### Closures

```go
// Function that returns a closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Function that returns a closure with parameters
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Create counter
    c1 := counter()
    c2 := counter()
    
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2
    fmt.Println(c2()) // 1 (separate closure)
    fmt.Println(c1()) // 3
    
    // Create multiplier
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5)) // 10
    fmt.Println(triple(5)) // 15
}
```

## Function Types

### Function Type Declaration

```go
// Function type
type MathOperation func(int, int) int

// Function that accepts function type
func calculate(a, b int, op MathOperation) int {
    return op(a, b)
}

// Functions that match the type
func add(x, y int) int {
    return x + y
}

func multiply(x, y int) int {
    return x * y
}

func main() {
    result1 := calculate(5, 3, add)
    result2 := calculate(5, 3, multiply)
    
    fmt.Printf("Addition: %d\n", result1)    // 8
    fmt.Printf("Multiplication: %d\n", result2) // 15
}
```

### Function as Parameters

```go
// Function that takes a function as parameter
func processNumbers(numbers []int, processor func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = processor(num)
    }
    return result
}

// Processor functions
func square(x int) int {
    return x * x
}

func double(x int) int {
    return x * 2
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    squared := processNumbers(numbers, square)
    doubled := processNumbers(numbers, double)
    
    fmt.Printf("Original: %v\n", numbers)
    fmt.Printf("Squared: %v\n", squared)
    fmt.Printf("Doubled: %v\n", doubled)
}
```

## Method Receivers

### Value Receivers

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Value receiver method
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    
    fmt.Printf("Area: %.2f\n", rect.Area())       // 15.00
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter()) // 16.00
}
```

### Pointer Receivers

```go
type Counter struct {
    value int
}

// Pointer receiver method
func (c *Counter) Increment() {
    c.value++
}

// Pointer receiver method
func (c *Counter) GetValue() int {
    return c.value
}

// Pointer receiver method
func (c *Counter) Reset() {
    c.value = 0
}

func main() {
    counter := Counter{value: 0}
    
    counter.Increment()
    counter.Increment()
    
    fmt.Printf("Value: %d\n", counter.GetValue()) // 2
    
    counter.Reset()
    fmt.Printf("Value after reset: %d\n", counter.GetValue()) // 0
}
```

### Mixed Receivers

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver (read-only)
func (p Person) GetInfo() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Pointer receiver (modifying)
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Pointer receiver (modifying)
func (p *Person) SetName(name string) {
    p.Name = name
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    
    fmt.Println(person.GetInfo()) // Alice is 25 years old
    
    person.SetAge(30)
    person.SetName("Bob")
    
    fmt.Println(person.GetInfo()) // Bob is 30 years old
}
```

## Practical Examples

### Example 1: Calculator Functions

```go
package main

import (
    "fmt"
    "math"
)

// Basic arithmetic functions
func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Advanced functions
func power(base, exponent float64) float64 {
    return math.Pow(base, exponent)
}

func squareRoot(x float64) (float64, error) {
    if x < 0 {
        return 0, fmt.Errorf("square root of negative number")
    }
    return math.Sqrt(x), nil
}

// Function that takes operation as parameter
func calculate(a, b float64, operation func(float64, float64) float64) float64 {
    return operation(a, b)
}

func main() {
    x, y := 10.0, 3.0
    
    fmt.Printf("Addition: %.2f\n", add(x, y))
    fmt.Printf("Subtraction: %.2f\n", subtract(x, y))
    fmt.Printf("Multiplication: %.2f\n", multiply(x, y))
    
    if result, err := divide(x, y); err != nil {
        fmt.Printf("Division error: %v\n", err)
    } else {
        fmt.Printf("Division: %.2f\n", result)
    }
    
    fmt.Printf("Power: %.2f\n", power(x, y))
    
    if result, err := squareRoot(x); err != nil {
        fmt.Printf("Square root error: %v\n", err)
    } else {
        fmt.Printf("Square root: %.2f\n", result)
    }
    
    // Using function as parameter
    fmt.Printf("Using function parameter: %.2f\n", calculate(x, y, add))
}
```

### Example 2: String Processing Functions

```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

// String processing functions
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func countWords(s string) int {
    words := strings.Fields(s)
    return len(words)
}

func capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}

func removeSpaces(s string) string {
    return strings.ReplaceAll(s, " ", "")
}

// Function that processes string with multiple operations
func processString(s string, operations []func(string) string) string {
    result := s
    for _, op := range operations {
        result = op(result)
    }
    return result
}

func main() {
    text := "hello world"
    
    fmt.Printf("Original: %s\n", text)
    fmt.Printf("Reversed: %s\n", reverse(text))
    fmt.Printf("Word count: %d\n", countWords(text))
    fmt.Printf("Capitalized: %s\n", capitalize(text))
    fmt.Printf("No spaces: %s\n", removeSpaces(text))
    
    // Process with multiple operations
    operations := []func(string) string{
        capitalize,
        removeSpaces,
    }
    
    result := processString(text, operations)
    fmt.Printf("Processed: %s\n", result)
}
```

### Example 3: Data Processing Functions

```go
package main

import "fmt"

type Student struct {
    Name  string
    Grade int
}

// Functions for processing student data
func filterStudents(students []Student, predicate func(Student) bool) []Student {
    var result []Student
    for _, student := range students {
        if predicate(student) {
            result = append(result, student)
        }
    }
    return result
}

func mapStudents(students []Student, mapper func(Student) Student) []Student {
    result := make([]Student, len(students))
    for i, student := range students {
        result[i] = mapper(student)
    }
    return result
}

func reduceStudents(students []Student, reducer func(int, Student) int, initial int) int {
    result := initial
    for _, student := range students {
        result = reducer(result, student)
    }
    return result
}

// Predicate functions
func isHighGrade(student Student) bool {
    return student.Grade >= 80
}

func isLowGrade(student Student) bool {
    return student.Grade < 60
}

// Mapper functions
func addBonus(student Student) Student {
    return Student{
        Name:  student.Name,
        Grade: student.Grade + 5,
    }
}

// Reducer functions
func sumGrades(total int, student Student) int {
    return total + student.Grade
}

func main() {
    students := []Student{
        {Name: "Alice", Grade: 85},
        {Name: "Bob", Grade: 72},
        {Name: "Carol", Grade: 95},
        {Name: "Dave", Grade: 58},
        {Name: "Eve", Grade: 91},
    }
    
    fmt.Println("All students:")
    for _, student := range students {
        fmt.Printf("  %s: %d\n", student.Name, student.Grade)
    }
    
    // Filter high-grade students
    highGradeStudents := filterStudents(students, isHighGrade)
    fmt.Println("\nHigh-grade students (>= 80):")
    for _, student := range highGradeStudents {
        fmt.Printf("  %s: %d\n", student.Name, student.Grade)
    }
    
    // Add bonus to all students
    bonusStudents := mapStudents(students, addBonus)
    fmt.Println("\nStudents with bonus (+5):")
    for _, student := range bonusStudents {
        fmt.Printf("  %s: %d\n", student.Name, student.Grade)
    }
    
    // Calculate total grades
    totalGrades := reduceStudents(students, sumGrades, 0)
    fmt.Printf("\nTotal grades: %d\n", totalGrades)
    
    // Calculate average
    average := float64(totalGrades) / float64(len(students))
    fmt.Printf("Average grade: %.2f\n", average)
}
```

### Example 4: Configuration and Validation

```go
package main

import (
    "fmt"
    "strings"
)

type Config struct {
    Host     string
    Port     int
    Database string
    Debug    bool
}

// Validation functions
type Validator func(string) error

func validateHost(host string) error {
    if host == "" {
        return fmt.Errorf("host cannot be empty")
    }
    if strings.Contains(host, " ") {
        return fmt.Errorf("host cannot contain spaces")
    }
    return nil
}

func validatePort(port int) error {
    if port < 1 || port > 65535 {
        return fmt.Errorf("port must be between 1 and 65535")
    }
    return nil
}

func validateDatabase(db string) error {
    if db == "" {
        return fmt.Errorf("database name cannot be empty")
    }
    if len(db) < 3 {
        return fmt.Errorf("database name must be at least 3 characters")
    }
    return nil
}

// Configuration builder
type ConfigBuilder struct {
    config Config
    errors []error
}

func NewConfigBuilder() *ConfigBuilder {
    return &ConfigBuilder{
        config: Config{
            Host:     "localhost",
            Port:     8080,
            Database: "default",
            Debug:    false,
        },
    }
}

func (cb *ConfigBuilder) SetHost(host string) *ConfigBuilder {
    if err := validateHost(host); err != nil {
        cb.errors = append(cb.errors, err)
    } else {
        cb.config.Host = host
    }
    return cb
}

func (cb *ConfigBuilder) SetPort(port int) *ConfigBuilder {
    if err := validatePort(port); err != nil {
        cb.errors = append(cb.errors, err)
    } else {
        cb.config.Port = port
    }
    return cb
}

func (cb *ConfigBuilder) SetDatabase(db string) *ConfigBuilder {
    if err := validateDatabase(db); err != nil {
        cb.errors = append(cb.errors, err)
    } else {
        cb.config.Database = db
    }
    return cb
}

func (cb *ConfigBuilder) SetDebug(debug bool) *ConfigBuilder {
    cb.config.Debug = debug
    return cb
}

func (cb *ConfigBuilder) Build() (Config, error) {
    if len(cb.errors) > 0 {
        return Config{}, fmt.Errorf("validation errors: %v", cb.errors)
    }
    return cb.config, nil
}

func main() {
    // Valid configuration
    config, err := NewConfigBuilder().
        SetHost("example.com").
        SetPort(9090).
        SetDatabase("myapp").
        SetDebug(true).
        Build()
    
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Valid config: %+v\n", config)
    }
    
    // Invalid configuration
    invalidConfig, err := NewConfigBuilder().
        SetHost("").
        SetPort(70000).
        SetDatabase("ab").
        Build()
    
    if err != nil {
        fmt.Printf("Validation errors: %v\n", err)
    } else {
        fmt.Printf("Invalid config: %+v\n", invalidConfig)
    }
}
```

## Function Composition

### Basic Composition

```go
// Function composition
func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

// Simple functions
func addOne(x int) int {
    return x + 1
}

func multiplyByTwo(x int) int {
    return x * 2
}

func square(x int) int {
    return x * x
}

func main() {
    // Compose functions
    addOneThenMultiply := compose(multiplyByTwo, addOne)
    squareThenAddOne := compose(addOne, square)
    
    fmt.Printf("addOneThenMultiply(5): %d\n", addOneThenMultiply(5)) // (5+1)*2 = 12
    fmt.Printf("squareThenAddOne(3): %d\n", squareThenAddOne(3))     // (3*3)+1 = 10
}
```

### Higher-Order Functions

```go
// Higher-order function that returns a function
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Higher-order function that takes a function and returns a function
func createProcessor(processor func(int) int) func([]int) []int {
    return func(numbers []int) []int {
        result := make([]int, len(numbers))
        for i, num := range numbers {
            result[i] = processor(num)
        }
        return result
    }
}

func main() {
    // Create multiplier functions
    double := createMultiplier(2)
    triple := createMultiplier(3)
    
    fmt.Printf("Double 5: %d\n", double(5))   // 10
    fmt.Printf("Triple 5: %d\n", triple(5))   // 15
    
    // Create processor functions
    squareProcessor := createProcessor(func(x int) int { return x * x })
    addOneProcessor := createProcessor(func(x int) int { return x + 1 })
    
    numbers := []int{1, 2, 3, 4, 5}
    
    squared := squareProcessor(numbers)
    added := addOneProcessor(numbers)
    
    fmt.Printf("Original: %v\n", numbers)
    fmt.Printf("Squared: %v\n", squared)
    fmt.Printf("Added one: %v\n", added)
}
```

## Best Practices

### 1. Function Naming

```go
// Good - descriptive names
func calculateTotalPrice(items []Item) float64 { }
func validateUserInput(input string) error { }
func processPayment(amount float64, method string) error { }

// Less preferred - unclear names
func calc(items []Item) float64 { }
func validate(input string) error { }
func process(amount float64, method string) error { }
```

### 2. Function Size

```go
// Good - small, focused functions
func validateEmail(email string) error {
    if email == "" {
        return fmt.Errorf("email cannot be empty")
    }
    if !strings.Contains(email, "@") {
        return fmt.Errorf("email must contain @")
    }
    return nil
}

// Less preferred - large, complex functions
func processUser(user User) error {
    // 50+ lines of validation, processing, database operations, etc.
}
```

### 3. Error Handling

```go
// Good - explicit error handling
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Less preferred - panic for expected errors
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

### 4. Parameter Design

```go
// Good - use structs for many parameters
type CreateUserRequest struct {
    Name     string
    Email    string
    Password string
    Role     string
}

func CreateUser(req CreateUserRequest) (*User, error) {
    // Implementation
}

// Less preferred - too many parameters
func CreateUser(name, email, password, role, department, manager string) (*User, error) {
    // Implementation
}
```

## Common Patterns

### 1. Factory Functions

```go
type Logger struct {
    level string
    output string
}

func NewLogger(level, output string) *Logger {
    return &Logger{
        level:  level,
        output: output,
    }
}

func NewFileLogger(level string, filename string) *Logger {
    return &Logger{
        level:  level,
        output: filename,
    }
}

func NewConsoleLogger(level string) *Logger {
    return &Logger{
        level:  level,
        output: "console",
    }
}
```

### 2. Option Pattern

```go
type ServerConfig struct {
    Host    string
    Port    int
    Timeout int
    Debug   bool
}

type ServerOption func(*ServerConfig)

func WithHost(host string) ServerOption {
    return func(c *ServerConfig) {
        c.Host = host
    }
}

func WithPort(port int) ServerOption {
    return func(c *ServerConfig) {
        c.Port = port
    }
}

func WithTimeout(timeout int) ServerOption {
    return func(c *ServerConfig) {
        c.Timeout = timeout
    }
}

func WithDebug(debug bool) ServerOption {
    return func(c *ServerConfig) {
        c.Debug = debug
    }
}

func NewServer(options ...ServerOption) *ServerConfig {
    config := &ServerConfig{
        Host:    "localhost",
        Port:    8080,
        Timeout: 30,
        Debug:   false,
    }
    
    for _, option := range options {
        option(config)
    }
    
    return config
}
```

### 3. Middleware Pattern

```go
type Handler func(string) string

func loggingMiddleware(next Handler) Handler {
    return func(input string) string {
        fmt.Printf("Input: %s\n", input)
        result := next(input)
        fmt.Printf("Output: %s\n", result)
        return result
    }
}

func validationMiddleware(next Handler) Handler {
    return func(input string) string {
        if len(input) == 0 {
            return "Error: empty input"
        }
        return next(input)
    }
}

func processString(input string) string {
    return strings.ToUpper(input)
}

func main() {
    // Chain middlewares
    handler := loggingMiddleware(validationMiddleware(processString))
    
    result := handler("hello world")
    fmt.Printf("Final result: %s\n", result)
}
```

## Common Mistakes

### 1. Not Handling Errors

```go
// Problematic
func divide(a, b int) int {
    return a / b // Will panic if b is 0
}

// Correct
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### 2. Modifying Slice Parameters

```go
// Problematic - doesn't modify original slice
func appendToSlice(s []int, value int) {
    s = append(s, value) // This doesn't affect the original slice
}

// Correct - use pointer
func appendToSlice(s *[]int, value int) {
    *s = append(*s, value)
}
```

### 3. Returning Pointers to Local Variables

```go
// Problematic - returns pointer to local variable
func createUser() *User {
    user := User{Name: "Alice"}
    return &user // This is actually safe in Go due to escape analysis
}

// Better - be explicit about allocation
func createUser() *User {
    return &User{Name: "Alice"}
}
```

### 4. Ignoring Return Values

```go
// Problematic - ignoring error
divide(10, 0) // Error is ignored

// Correct - handle error
if result, err := divide(10, 0); err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Result: %d\n", result)
}
```

## Exercises

### Exercise 1: Calculator Functions
Implement a calculator with basic arithmetic operations and error handling.

### Exercise 2: String Processing
Create functions for string manipulation (reverse, capitalize, count words).

### Exercise 3: Data Processing
Implement map, filter, and reduce functions for processing data.

### Exercise 4: Configuration Builder
Build a configuration system using the builder pattern with validation.

### Exercise 5: Middleware Chain
Create a middleware system for processing requests with logging and validation.

## Next Steps

Now that you understand functions, explore:

1. **[Structs](../11.structs)** - Define custom data types
2. **[Interfaces](../12.interfaces)** - Define behavior contracts
3. **[Methods](../13.methods)** - Add behavior to types
4. **[Packages](../14.packages)** - Organize code into modules

## Key Takeaways

- Functions are first-class citizens in Go
- Use descriptive names and keep functions small
- Handle errors explicitly with multiple return values
- Use variadic functions for flexible parameter lists
- Closures capture variables from surrounding scope
- Function types enable higher-order programming
- Use value receivers for small, immutable data
- Use pointer receivers for large or mutable data
- Compose functions for complex behavior
- Follow the single responsibility principle

## Additional Resources

- [Go Tour - Functions](https://tour.golang.org/basics/4)
- [Effective Go - Functions](https://golang.org/doc/effective_go.html#functions)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go Language Specification - Functions](https://golang.org/ref/spec#Function_declarations)
