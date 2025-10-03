## Go Conditional Statements

Go provides several conditional constructs for controlling program flow: `if`, `if-else`, `switch`, and `select`. These statements allow you to make decisions and handle different execution paths based on conditions.

### If Statement

The basic `if` statement evaluates a boolean expression and executes code if the condition is true:

```go
// Basic if statement
age := 18
if age >= 18 {
    fmt.Println("You are an adult")
}

// If with initialization
if err := someFunction(); err != nil {
    fmt.Printf("Error occurred: %v\n", err)
}

// Multiple conditions
if age >= 18 && age < 65 {
    fmt.Println("Working age")
}
```

### If-Else Statement

The `if-else` statement provides an alternative execution path:

```go
// Basic if-else
age := 16
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}

// If-else with initialization
if score := getScore(); score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

### If-Else If Chain

Multiple conditions can be chained using `else if`:

```go
temperature := 25
if temperature > 30 {
    fmt.Println("It's hot")
} else if temperature > 20 {
    fmt.Println("It's warm")
} else if temperature > 10 {
    fmt.Println("It's cool")
} else {
    fmt.Println("It's cold")
}
```

### Switch Statement

The `switch` statement provides a cleaner way to handle multiple conditions:

#### Basic Switch

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}
```

#### Switch with Expression

```go
score := 85
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
case score >= 70:
    fmt.Println("Grade: C")
case score >= 60:
    fmt.Println("Grade: D")
default:
    fmt.Println("Grade: F")
}
```

#### Switch with Type Assertion

```go
func processValue(x interface{}) {
    switch v := x.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

#### Switch with Fallthrough

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Monday blues")
    fallthrough
case "Tuesday", "Wednesday", "Thursday":
    fmt.Println("Work day")
case "Friday":
    fmt.Println("Friday!")
    fallthrough
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
}
```

### Select Statement

The `select` statement is used for channel operations and allows a goroutine to wait on multiple channel operations:

#### Basic Select

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(1 * time.Second)
    ch1 <- "from ch1"
}()

go func() {
    time.Sleep(2 * time.Second)
    ch2 <- "from ch2"
}()

select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received:", msg2)
}
```

#### Select with Default

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message received")
}
```

#### Select with Timeout

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}
```

#### Select for Non-blocking Operations

```go
select {
case ch <- "message":
    fmt.Println("Message sent")
default:
    fmt.Println("Channel is full")
}
```

### Common Patterns

#### Error Handling

```go
// Pattern 1: Check error first
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Pattern 2: Handle different error types
if err := processFile(); err != nil {
    switch e := err.(type) {
    case *os.PathError:
        fmt.Printf("Path error: %s\n", e.Path)
    case *os.PermissionError:
        fmt.Println("Permission denied")
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

#### Nil Checks

```go
// Check for nil pointer
if ptr != nil {
    fmt.Println("Pointer is not nil")
}

// Check for nil slice
if len(slice) > 0 {
    fmt.Println("Slice is not empty")
}

// Check for nil map
if value, exists := myMap["key"]; exists {
    fmt.Printf("Value: %v\n", value)
}
```

#### Range with Conditions

```go
// Filter elements
for _, item := range items {
    if item.IsValid() {
        process(item)
    }
}

// Find first match
for i, item := range items {
    if item.Matches(criteria) {
        fmt.Printf("Found at index %d\n", i)
        break
    }
}
```

### Advanced Patterns

#### Guard Clauses

```go
func processUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }
    
    if user.ID == "" {
        return errors.New("user ID is empty")
    }
    
    if user.Email == "" {
        return errors.New("user email is empty")
    }
    
    // Main logic here
    return nil
}
```

#### Early Returns

```go
func validateInput(input string) (bool, error) {
    if input == "" {
        return false, errors.New("input cannot be empty")
    }
    
    if len(input) < 3 {
        return false, errors.New("input too short")
    }
    
    if !isValidFormat(input) {
        return false, errors.New("invalid format")
    }
    
    return true, nil
}
```

#### Switch with Complex Conditions

```go
func categorizeAge(age int) string {
    switch {
    case age < 0:
        return "invalid"
    case age < 13:
        return "child"
    case age < 20:
        return "teenager"
    case age < 65:
        return "adult"
    default:
        return "senior"
    }
}
```

### Best Practices

1. **Use guard clauses**: Check for error conditions first and return early
2. **Prefer switch over long if-else chains**: More readable and potentially more efficient
3. **Use type switches**: For handling different types in interfaces
4. **Handle all cases**: Always include a `default` case in switches when appropriate
5. **Use meaningful conditions**: Make boolean expressions clear and readable
6. **Avoid deep nesting**: Use early returns to reduce nesting levels
7. **Use select for channel operations**: Don't use if-else for channel operations

### Common Pitfalls

#### Assignment vs Comparison

```go
// Wrong: assignment instead of comparison
if x = 5 {  // This assigns 5 to x, always true
    fmt.Println("This will always execute")
}

// Correct: comparison
if x == 5 {  // This compares x with 5
    fmt.Println("x equals 5")
}
```

#### Missing Break in Switch

```go
// Go automatically breaks, but be explicit for clarity
switch day {
case "Monday":
    fmt.Println("Monday")
    // break is implicit
case "Tuesday":
    fmt.Println("Tuesday")
    // break is implicit
}

// Use fallthrough when you want to continue
switch day {
case "Monday":
    fmt.Println("Monday")
    fallthrough  // Continue to next case
case "Tuesday":
    fmt.Println("Tuesday")
}
```

#### Nil Pointer Dereference

```go
// Wrong: potential nil pointer dereference
if ptr != nil && ptr.value > 0 {  // Good
    fmt.Println("Valid pointer")
}

// Wrong: could panic
if ptr.value > 0 && ptr != nil {  // Bad - could panic
    fmt.Println("This might panic")
}
```

### Performance Considerations

#### Switch vs If-Else

```go
// Switch is often more efficient for multiple conditions
switch value {
case 1, 2, 3:
    // Handle cases 1, 2, 3
case 4, 5, 6:
    // Handle cases 4, 5, 6
default:
    // Handle other cases
}

// If-else chain for complex conditions
if complexCondition1() {
    // Handle condition 1
} else if complexCondition2() {
    // Handle condition 2
} else {
    // Handle default case
}
```

#### Select Performance

```go
// Use buffered channels to avoid blocking
ch := make(chan int, 1)

// Non-blocking send
select {
case ch <- value:
    // Successfully sent
default:
    // Channel is full, handle accordingly
}
```

### Quick Reference

```go
// Basic if
if condition { }

// If-else
if condition { } else { }

// If-else if
if condition1 { } else if condition2 { } else { }

// Switch
switch value {
case val1:
case val2, val3:
default:
}

// Switch with expression
switch {
case condition1:
case condition2:
default:
}

// Type switch
switch v := x.(type) {
case int:
case string:
default:
}

// Select
select {
case <-ch1:
case ch2 <- value:
case <-time.After(duration):
default:
}
```

### Examples

#### File Processing with Error Handling

```go
func processFile(filename string) error {
    if filename == "" {
        return errors.New("filename cannot be empty")
    }
    
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Process file
    return nil
}
```

#### HTTP Status Code Handling

```go
func handleResponse(statusCode int) {
    switch statusCode {
    case 200:
        fmt.Println("Success")
    case 404:
        fmt.Println("Not Found")
    case 500:
        fmt.Println("Internal Server Error")
    default:
        fmt.Printf("Unknown status: %d\n", statusCode)
    }
}
```

#### Channel Communication

```go
func worker(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        select {
        case results <- processJob(job):
            // Job processed successfully
        case <-time.After(5 * time.Second):
            fmt.Println("Job timeout")
        }
    }
}
```
