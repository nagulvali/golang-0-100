# Go Miscellaneous Functions and Utilities

This document covers Go's built-in functions, common utility functions, and miscellaneous features that are essential for Go programming. These functions provide fundamental operations for working with different data types and performing common tasks.

## What You'll Learn

- Built-in functions (len, cap, make, new, append, copy, delete, close)
- Type conversion functions
- String manipulation functions
- Mathematical functions
- Time and date functions
- Error handling functions
- Reflection functions
- Utility functions and patterns
- Best practices and common use cases

## Key Concepts

- **Built-in Functions**: Functions provided by Go runtime
- **Type Conversion**: Converting between different types
- **String Operations**: Common string manipulation tasks
- **Mathematical Operations**: Basic math functions
- **Time Operations**: Working with dates and times
- **Error Handling**: Managing errors effectively
- **Reflection**: Runtime type inspection
- **Utilities**: Common helper functions

## Built-in Functions

### len() - Length Function

The `len()` function returns the length of various types:

```go
// Arrays and slices
arr := [5]int{1, 2, 3, 4, 5}
slice := []int{1, 2, 3, 4, 5}
fmt.Println(len(arr))   // 5
fmt.Println(len(slice)) // 5

// Strings
str := "Hello, World!"
fmt.Println(len(str))   // 13

// Maps
m := map[string]int{"a": 1, "b": 2}
fmt.Println(len(m))     // 2

// Channels
ch := make(chan int, 5)
fmt.Println(len(ch))    // 0 (number of queued elements)
```

### cap() - Capacity Function

The `cap()` function returns the capacity of slices, arrays, and channels:

```go
// Slices
slice := make([]int, 5, 10)
fmt.Println(cap(slice)) // 10

// Arrays
arr := [5]int{1, 2, 3, 4, 5}
fmt.Println(cap(arr))   // 5

// Channels
ch := make(chan int, 5)
fmt.Println(cap(ch))    // 5
```

### make() - Memory Allocation

The `make()` function allocates and initializes slices, maps, and channels:

```go
// Slices
slice := make([]int, 5, 10) // length 5, capacity 10
slice2 := make([]int, 5)    // length 5, capacity 5

// Maps
m := make(map[string]int)
m2 := make(map[string]int, 100) // hint for initial capacity

// Channels
ch := make(chan int)        // unbuffered channel
ch2 := make(chan int, 5)    // buffered channel with capacity 5
```

### new() - Memory Allocation

The `new()` function allocates memory for a type and returns a pointer:

```go
// Allocate memory for int
p := new(int)
*p = 42
fmt.Println(*p) // 42

// Allocate memory for struct
type Person struct {
    Name string
    Age  int
}
person := new(Person)
person.Name = "Alice"
person.Age = 30
```

### append() - Slice Append

The `append()` function appends elements to slices:

```go
// Basic append
slice := []int{1, 2, 3}
slice = append(slice, 4, 5, 6)
fmt.Println(slice) // [1 2 3 4 5 6]

// Append one slice to another
slice1 := []int{1, 2, 3}
slice2 := []int{4, 5, 6}
slice1 = append(slice1, slice2...)
fmt.Println(slice1) // [1 2 3 4 5 6]

// Append to nil slice
var slice []int
slice = append(slice, 1, 2, 3)
fmt.Println(slice) // [1 2 3]
```

### copy() - Slice Copy

The `copy()` function copies elements from one slice to another:

```go
// Copy slices
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
n := copy(dst, src)
fmt.Println(dst) // [1 2 3]
fmt.Println(n)   // 3 (number of elements copied)

// Copy with different lengths
src2 := []int{1, 2}
dst2 := make([]int, 5)
n2 := copy(dst2, src2)
fmt.Println(dst2) // [1 2 0 0 0]
fmt.Println(n2)   // 2
```

### delete() - Map Delete

The `delete()` function removes key-value pairs from maps:

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

// Delete a key
delete(m, "b")
fmt.Println(m) // map[a:1 c:3]

// Delete non-existent key (safe)
delete(m, "nonexistent")
fmt.Println(m) // map[a:1 c:3]
```

### close() - Channel Close

The `close()` function closes channels:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

// Reading from closed channel
for val := range ch {
    fmt.Println(val) // 1, 2
}

// Check if channel is closed
val, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}
```

## Type Conversion Functions

### Type Assertions

Type assertions allow you to extract the concrete value from an interface:

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)
fmt.Println(s) // hello

// Safe type assertion
s2, ok := i.(string)
if ok {
    fmt.Println(s2) // hello
}

// Type assertion with type switch
switch v := i.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Integer: %d\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### Type Conversions

Explicit type conversions in Go:

```go
// Numeric conversions
var i int = 42
var f float64 = float64(i)
var s string = string(i) // Note: converts to Unicode character

// String conversions
str := "123"
num, err := strconv.Atoi(str)
if err == nil {
    fmt.Println(num) // 123
}

// Boolean conversions
b := true
s := strconv.FormatBool(b)
fmt.Println(s) // true
```

## String Manipulation Functions

### String Package Functions

```go
import "strings"

// String comparison
str1 := "hello"
str2 := "HELLO"
fmt.Println(strings.EqualFold(str1, str2)) // true (case-insensitive)

// String searching
text := "Hello, World!"
fmt.Println(strings.Contains(text, "World"))     // true
fmt.Println(strings.HasPrefix(text, "Hello"))    // true
fmt.Println(strings.HasSuffix(text, "!"))        // true
fmt.Println(strings.Index(text, "World"))        // 7

// String manipulation
fmt.Println(strings.ToUpper("hello"))            // HELLO
fmt.Println(strings.ToLower("HELLO"))            // hello
fmt.Println(strings.TrimSpace("  hello  "))      // hello
fmt.Println(strings.Trim("hello", "ho"))         // ell

// String splitting and joining
words := strings.Split("hello world", " ")
fmt.Println(words) // [hello world]
joined := strings.Join(words, "-")
fmt.Println(joined) // hello-world

// String replacement
text2 := "Hello, World!"
replaced := strings.Replace(text2, "World", "Go", 1)
fmt.Println(replaced) // Hello, Go!
```

### strconv Package Functions

```go
import "strconv"

// String to number conversions
str := "123"
i, err := strconv.Atoi(str)
if err == nil {
    fmt.Println(i) // 123
}

f, err := strconv.ParseFloat("3.14", 64)
if err == nil {
    fmt.Println(f) // 3.14
}

// Number to string conversions
num := 123
str2 := strconv.Itoa(num)
fmt.Println(str2) // 123

float := 3.14
str3 := strconv.FormatFloat(float, 'f', 2, 64)
fmt.Println(str3) // 3.14

// Boolean conversions
b, err := strconv.ParseBool("true")
if err == nil {
    fmt.Println(b) // true
}

str4 := strconv.FormatBool(true)
fmt.Println(str4) // true
```

## Mathematical Functions

### math Package Functions

```go
import "math"

// Basic math functions
fmt.Println(math.Abs(-5))        // 5
fmt.Println(math.Max(10, 20))    // 20
fmt.Println(math.Min(10, 20))    // 10
fmt.Println(math.Pow(2, 3))      // 8
fmt.Println(math.Sqrt(16))       // 4
fmt.Println(math.Ceil(3.2))      // 4
fmt.Println(math.Floor(3.8))     // 3
fmt.Println(math.Round(3.6))     // 4

// Trigonometric functions
fmt.Println(math.Sin(math.Pi/2)) // 1
fmt.Println(math.Cos(0))         // 1
fmt.Println(math.Tan(math.Pi/4)) // 1

// Logarithmic functions
fmt.Println(math.Log(math.E))    // 1
fmt.Println(math.Log10(100))     // 2
fmt.Println(math.Log2(8))        // 3

// Constants
fmt.Println(math.Pi)             // 3.141592653589793
fmt.Println(math.E)              // 2.718281828459045
```

### math/rand Package Functions

```go
import (
    "math/rand"
    "time"
)

// Seed random number generator
rand.Seed(time.Now().UnixNano())

// Generate random numbers
fmt.Println(rand.Int())          // Random integer
fmt.Println(rand.Intn(100))      // Random integer 0-99
fmt.Println(rand.Float64())      // Random float 0.0-1.0

// Generate random string
const charset = "abcdefghijklmnopqrstuvwxyz"
b := make([]byte, 10)
for i := range b {
    b[i] = charset[rand.Intn(len(charset))]
}
randomString := string(b)
fmt.Println(randomString)
```

## Time and Date Functions

### time Package Functions

```go
import "time"

// Current time
now := time.Now()
fmt.Println(now) // 2023-12-07 10:30:45.123456789 +0000 UTC

// Time components
fmt.Println(now.Year())    // 2023
fmt.Println(now.Month())   // December
fmt.Println(now.Day())     // 7
fmt.Println(now.Hour())    // 10
fmt.Println(now.Minute())  // 30
fmt.Println(now.Second())  // 45

// Time formatting
fmt.Println(now.Format("2006-01-02 15:04:05")) // 2023-12-07 10:30:45
fmt.Println(now.Format(time.RFC3339))          // 2023-12-07T10:30:45Z

// Time parsing
t, err := time.Parse("2006-01-02", "2023-12-07")
if err == nil {
    fmt.Println(t) // 2023-12-07 00:00:00 +0000 UTC
}

// Time arithmetic
future := now.Add(24 * time.Hour)
past := now.Add(-24 * time.Hour)
duration := future.Sub(past)
fmt.Println(duration) // 48h0m0s

// Sleep
time.Sleep(1 * time.Second)
```

## Error Handling Functions

### Error Creation and Handling

```go
import "errors"

// Create custom errors
err := errors.New("something went wrong")
fmt.Println(err) // something went wrong

// Error wrapping (Go 1.13+)
import "fmt"

func doSomething() error {
    return fmt.Errorf("operation failed: %w", err)
}

// Check error types
if errors.Is(err, someError) {
    fmt.Println("Error is of expected type")
}

// Unwrap errors
if unwrapped := errors.Unwrap(err); unwrapped != nil {
    fmt.Println("Unwrapped error:", unwrapped)
}
```

### panic() and recover()

```go
// Panic function
func riskyFunction() {
    panic("something went wrong")
}

// Recover function
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    riskyFunction()
}

// Example usage
safeFunction() // Recovered from panic: something went wrong
```

## Reflection Functions

### reflect Package Functions

```go
import "reflect"

// Get type information
var x int = 42
t := reflect.TypeOf(x)
fmt.Println(t) // int

// Get value information
v := reflect.ValueOf(x)
fmt.Println(v) // 42

// Check if value is settable
if v.CanSet() {
    v.SetInt(100)
}

// Get kind of type
fmt.Println(t.Kind()) // int

// Get struct field information
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
v2 := reflect.ValueOf(p)
t2 := reflect.TypeOf(p)

for i := 0; i < v2.NumField(); i++ {
    field := t2.Field(i)
    value := v2.Field(i)
    fmt.Printf("%s: %v\n", field.Name, value)
}
```

## Utility Functions and Patterns

### Common Utility Functions

```go
// Min and Max functions
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Clamp function
func Clamp(value, min, max int) int {
    if value < min {
        return min
    }
    if value > max {
        return max
    }
    return value
}

// Contains function for slices
func Contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

// Unique function for slices
func Unique(slice []int) []int {
    keys := make(map[int]bool)
    result := []int{}
    
    for _, item := range slice {
        if !keys[item] {
            keys[item] = true
            result = append(result, item)
        }
    }
    return result
}
```

### String Utility Functions

```go
// String utilities
func IsEmpty(s string) bool {
    return len(s) == 0
}

func IsBlank(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

### Slice Utility Functions

```go
// Slice utilities
func Filter(slice []int, predicate func(int) bool) []int {
    result := []int{}
    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

func Map(slice []int, mapper func(int) int) []int {
    result := make([]int, len(slice))
    for i, item := range slice {
        result[i] = mapper(item)
    }
    return result
}

func Reduce(slice []int, initial int, reducer func(int, int) int) int {
    result := initial
    for _, item := range slice {
        result = reducer(result, item)
    }
    return result
}
```

## Practical Examples

### Example 1: Data Processing Pipeline

```go
package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    // Process CSV data
    csvData := "1,2,3,4,5\n6,7,8,9,10"
    lines := strings.Split(csvData, "\n")
    
    for _, line := range lines {
        numbers := strings.Split(line, ",")
        sum := 0
        
        for _, numStr := range numbers {
            if num, err := strconv.Atoi(strings.TrimSpace(numStr)); err == nil {
                sum += num
            }
        }
        
        fmt.Printf("Line sum: %d\n", sum)
    }
}
```

### Example 2: Configuration Management

```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Config struct {
    Host    string
    Port    int
    Debug   bool
    Timeout int
}

func parseConfig(configStr string) (*Config, error) {
    config := &Config{}
    lines := strings.Split(configStr, "\n")
    
    for _, line := range lines {
        if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
            continue
        }
        
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }
        
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        
        switch key {
        case "host":
            config.Host = value
        case "port":
            if port, err := strconv.Atoi(value); err == nil {
                config.Port = port
            }
        case "debug":
            if debug, err := strconv.ParseBool(value); err == nil {
                config.Debug = debug
            }
        case "timeout":
            if timeout, err := strconv.Atoi(value); err == nil {
                config.Timeout = timeout
            }
        }
    }
    
    return config, nil
}

func main() {
    configStr := `
host=localhost
port=8080
debug=true
timeout=30
`
    
    config, err := parseConfig(configStr)
    if err != nil {
        fmt.Printf("Error parsing config: %v\n", err)
        return
    }
    
    fmt.Printf("Config: %+v\n", config)
}
```

### Example 3: Error Handling with Utilities

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

func parseAndValidate(input string) (int, error) {
    if input == "" {
        return 0, errors.New("input cannot be empty")
    }
    
    num, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("invalid number: %w", err)
    }
    
    if num < 0 {
        return 0, errors.New("number must be positive")
    }
    
    if num > 1000 {
        return 0, errors.New("number must be less than 1000")
    }
    
    return num, nil
}

func main() {
    inputs := []string{"42", "-5", "abc", "1500", ""}
    
    for _, input := range inputs {
        result, err := parseAndValidate(input)
        if err != nil {
            fmt.Printf("Input '%s': Error - %v\n", input, err)
        } else {
            fmt.Printf("Input '%s': Success - %d\n", input, result)
        }
    }
}
```

## Best Practices

### 1. Use Built-in Functions Effectively

```go
// Good - use built-in functions
slice := make([]int, 0, 10)
slice = append(slice, 1, 2, 3)
fmt.Println(len(slice), cap(slice))

// Less preferred - manual implementation
var slice2 []int
slice2 = append(slice2, 1)
slice2 = append(slice2, 2)
slice2 = append(slice2, 3)
```

### 2. Handle Errors Properly

```go
// Good - proper error handling
num, err := strconv.Atoi("123")
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Println(num)

// Bad - ignoring errors
num, _ := strconv.Atoi("123")
fmt.Println(num)
```

### 3. Use Type Assertions Safely

```go
// Good - safe type assertion
if str, ok := i.(string); ok {
    fmt.Println(str)
}

// Bad - unsafe type assertion
str := i.(string) // May panic
```

### 4. Optimize String Operations

```go
// Good - use strings.Builder for concatenation
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("hello")
}
result := builder.String()

// Less preferred - string concatenation
result := ""
for i := 0; i < 1000; i++ {
    result += "hello" // Inefficient
}
```

## Common Pitfalls

### 1. String to Number Conversion

```go
// Problematic
var i int = 42
s := string(i) // Converts to Unicode character, not "42"

// Correct
s := strconv.Itoa(i) // "42"
```

### 2. Slice Append Behavior

```go
// Problematic - modifying original slice
slice1 := []int{1, 2, 3}
slice2 := slice1
slice2 = append(slice2, 4) // Modifies slice1 too

// Correct - create new slice
slice1 := []int{1, 2, 3}
slice2 := make([]int, len(slice1))
copy(slice2, slice1)
slice2 = append(slice2, 4)
```

### 3. Map Key Existence

```go
// Problematic - zero value confusion
m := map[string]int{"a": 0}
if m["b"] == 0 {
    fmt.Println("Key 'b' exists") // Wrong!
}

// Correct - check existence
if _, exists := m["b"]; exists {
    fmt.Println("Key 'b' exists")
}
```

## Exercises

### Exercise 1: String Utilities
Implement utility functions for string manipulation (reverse, capitalize, trim, etc.).

### Exercise 2: Slice Utilities
Implement utility functions for slice operations (filter, map, reduce, unique, etc.).

### Exercise 3: Configuration Parser
Create a configuration parser that reads key-value pairs from a string.

### Exercise 4: Data Validator
Implement a data validator that checks various data types and formats.

### Exercise 5: Error Wrapper
Create an error wrapper that provides additional context to errors.

## Next Steps

Now that you understand miscellaneous functions, explore:

1. **[Structs](../12.structs)** - Define custom data types
2. **[Interfaces](../13.interfaces)** - Define behavior contracts
3. **[Methods](../14.methods)** - Add behavior to types
4. **[Packages](../15.packages)** - Organize code into modules

## Key Takeaways

- Built-in functions provide essential operations for Go types
- Type conversions require explicit syntax in Go
- String manipulation is powerful with the strings package
- Mathematical operations are available in the math package
- Time operations are comprehensive in the time package
- Error handling is explicit and should be done properly
- Reflection allows runtime type inspection
- Utility functions can be created for common operations

## Additional Resources

- [Go Built-in Functions](https://golang.org/pkg/builtin/)
- [Go Standard Library](https://golang.org/pkg/)
- [Effective Go - Built-in Functions](https://golang.org/doc/effective_go.html#builtin)
- [Go by Example - Built-in Functions](https://gobyexample.com/builtin-functions)
