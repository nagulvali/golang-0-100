# Go Closures

## Table of Contents
1. [Introduction](#introduction)
2. [What is a Closure?](#what-is-a-closure)
3. [How Closures Work](#how-closures-work)
4. [Basic Examples](#basic-examples)
5. [Common Use Cases](#common-use-cases)
6. [Advanced Patterns](#advanced-patterns)
7. [Closures vs Regular Functions](#closures-vs-regular-functions)
8. [Common Pitfalls](#common-pitfalls)
9. [Best Practices](#best-practices)
10. [Real-World Applications](#real-world-applications)

---

## Introduction

Closures are one of the most powerful features in Go (and many other programming languages). They enable elegant solutions to complex problems and are fundamental to understanding functional programming concepts in Go.

A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.

---

## What is a Closure?

A **closure** is:
- A function that captures and remembers variables from its surrounding scope
- A combination of a function and the lexical environment within which that function was declared
- A way to create "stateful" functions that maintain data between calls

**Key Characteristics:**
- Functions are first-class values in Go (can be assigned, passed around, and returned)
- Anonymous functions can access variables declared in their enclosing scope
- The captured variables persist for the lifetime of the closure
- Each closure instance maintains its own copy of captured variables

---

## How Closures Work

When you create a closure:

1. **Declaration**: An outer function declares local variables
2. **Inner Function**: An inner (anonymous) function references those variables
3. **Return**: The outer function returns the inner function
4. **Capture**: The inner function "closes over" the outer function's variables
5. **Persistence**: Those variables remain accessible even after the outer function returns

```go
func outer() func() int {
    x := 0                    // Variable in outer scope
    
    return func() int {       // Anonymous function (closure)
        x++                   // References outer variable
        return x
    }
}
// When outer() returns, x is captured by the closure
// and persists across multiple calls
```

**Memory Model:**
- Captured variables are allocated on the heap (not stack)
- They survive as long as any closure referencing them exists
- Go's garbage collector handles cleanup automatically

---

## Basic Examples

### Example 1: Counter (State Management)

```go
func counter() func() int {
    count := 0
    
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
    fmt.Println(increment())  // 3
    
    // Create a new counter - independent state
    increment2 := counter()
    fmt.Println(increment2()) // 1
}
```

**Explanation:**
- Each call to `counter()` creates a new closure with its own `count` variable
- The `count` variable persists between calls to `increment()`
- Multiple closures maintain independent state

### Example 2: Accumulator

```go
func accumulator(start int) func(int) int {
    sum := start
    
    return func(value int) int {
        sum += value
        return sum
    }
}

func main() {
    acc := accumulator(10)
    
    fmt.Println(acc(5))   // 15
    fmt.Println(acc(3))   // 18
    fmt.Println(acc(7))   // 25
}
```

### Example 3: String Builder

```go
func stringBuilder() func(string) string {
    var result string
    
    return func(s string) string {
        result += s
        return result
    }
}

func main() {
    builder := stringBuilder()
    
    fmt.Println(builder("Hello"))       // Hello
    fmt.Println(builder(" "))           // Hello 
    fmt.Println(builder("World"))       // Hello World
}
```

### Example 4: Multiplier Generator

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

---

## Common Use Cases

### 1. Encapsulation and Data Privacy

Closures provide a way to create private variables that can't be accessed directly:

```go
func newAccount(initialBalance float64) (deposit func(float64), withdraw func(float64), balance func() float64) {
    bal := initialBalance
    
    deposit = func(amount float64) {
        bal += amount
        fmt.Printf("Deposited: $%.2f\n", amount)
    }
    
    withdraw = func(amount float64) {
        if amount > bal {
            fmt.Println("Insufficient funds")
            return
        }
        bal -= amount
        fmt.Printf("Withdrawn: $%.2f\n", amount)
    }
    
    balance = func() float64 {
        return bal
    }
    
    return
}

func main() {
    deposit, withdraw, balance := newAccount(100.0)
    
    fmt.Printf("Balance: $%.2f\n", balance())  // $100.00
    deposit(50)                                 // Deposited: $50.00
    fmt.Printf("Balance: $%.2f\n", balance())  // $150.00
    withdraw(30)                                // Withdrawn: $30.00
    fmt.Printf("Balance: $%.2f\n", balance())  // $120.00
}
```

### 2. Function Factories

Create specialized functions from generic templates:

```go
func makeGreeter(greeting string) func(string) string {
    return func(name string) string {
        return greeting + ", " + name + "!"
    }
}

func main() {
    sayHello := makeGreeter("Hello")
    sayHi := makeGreeter("Hi")
    
    fmt.Println(sayHello("Alice"))  // Hello, Alice!
    fmt.Println(sayHi("Bob"))       // Hi, Bob!
}
```

### 3. Callback Functions with Context

Pass state along with callback functions:

```go
func processItems(items []int, callback func(int)) {
    for _, item := range items {
        callback(item)
    }
}

func main() {
    sum := 0
    count := 0
    
    // Closure captures sum and count
    processItems([]int{1, 2, 3, 4, 5}, func(n int) {
        sum += n
        count++
    })
    
    fmt.Printf("Sum: %d, Count: %d, Average: %.2f\n", 
        sum, count, float64(sum)/float64(count))
    // Output: Sum: 15, Count: 5, Average: 3.00
}
```

### 4. Memoization (Caching)

Cache expensive function results:

```go
func fibonacci() func(int) int {
    cache := make(map[int]int)
    
    var fib func(int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        
        if val, exists := cache[n]; exists {
            return val
        }
        
        result := fib(n-1) + fib(n-2)
        cache[n] = result
        return result
    }
    
    return fib
}

func main() {
    fib := fibonacci()
    
    fmt.Println(fib(10))  // 55 (computed)
    fmt.Println(fib(10))  // 55 (cached)
    fmt.Println(fib(15))  // 610
}
```

### 5. Middleware and Decorators

Wrap functions with additional behavior:

```go
func logger(fn func(string) string) func(string) string {
    return func(s string) string {
        fmt.Printf("Calling function with: %s\n", s)
        result := fn(s)
        fmt.Printf("Function returned: %s\n", result)
        return result
    }
}

func toUpper(s string) string {
    return strings.ToUpper(s)
}

func main() {
    wrappedFunc := logger(toUpper)
    result := wrappedFunc("hello")
    // Output:
    // Calling function with: hello
    // Function returned: HELLO
}
```

---

## Advanced Patterns

### 1. Closure with Multiple Return Values

```go
func createOperations(initial int) (increment func(), decrement func(), get func() int) {
    value := initial
    
    increment = func() {
        value++
    }
    
    decrement = func() {
        value--
    }
    
    get = func() int {
        return value
    }
    
    return
}

func main() {
    inc, dec, get := createOperations(10)
    
    inc()
    inc()
    fmt.Println(get())  // 12
    
    dec()
    fmt.Println(get())  // 11
}
```

### 2. Closure with Method Receivers

```go
type Counter struct {
    value int
}

func (c *Counter) Incrementer(by int) func() int {
    return func() int {
        c.value += by
        return c.value
    }
}

func main() {
    c := &Counter{value: 0}
    
    inc5 := c.Incrementer(5)
    inc10 := c.Incrementer(10)
    
    fmt.Println(inc5())   // 5
    fmt.Println(inc10())  // 15
    fmt.Println(inc5())   // 20
}
```

### 3. Generator Pattern

```go
func numberGenerator(start, end int) func() (int, bool) {
    current := start - 1
    
    return func() (int, bool) {
        current++
        if current > end {
            return 0, false
        }
        return current, true
    }
}

func main() {
    gen := numberGenerator(1, 5)
    
    for {
        num, ok := gen()
        if !ok {
            break
        }
        fmt.Println(num)  // 1, 2, 3, 4, 5
    }
}
```

### 4. Partial Application (Currying)

```go
func add(a int) func(int) func(int) int {
    return func(b int) func(int) int {
        return func(c int) int {
            return a + b + c
        }
    }
}

func main() {
    result := add(1)(2)(3)
    fmt.Println(result)  // 6
    
    // Partial application
    add1 := add(1)
    add1and2 := add1(2)
    
    fmt.Println(add1and2(3))  // 6
    fmt.Println(add1and2(5))  // 8
}
```

### 5. Defer with Closures

```go
func process() {
    x := 10
    
    // Closure captures x by reference
    defer func() {
        fmt.Println("Deferred x:", x)  // Prints: Deferred x: 20
    }()
    
    x = 20
    fmt.Println("Current x:", x)  // Prints: Current x: 20
}
```

---

## Closures vs Regular Functions

| Aspect | Regular Functions | Closures |
|--------|------------------|----------|
| **State** | Stateless (no memory between calls) | Stateful (remembers captured variables) |
| **Variables** | Only accesses parameters and local variables | Can access outer scope variables |
| **Independence** | Each call is completely independent | State persists across calls |
| **Memory** | Stack-allocated (usually) | Captured variables heap-allocated |
| **Use Case** | Pure computations, utilities | State management, callbacks, factories |

**Example Comparison:**

```go
// Regular function - no state
func regularIncrement(x int) int {
    return x + 1
}

// Closure - maintains state
func closureIncrement() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func main() {
    // Regular function
    fmt.Println(regularIncrement(0))  // 1
    fmt.Println(regularIncrement(0))  // 1 (always starts from 0)
    
    // Closure
    inc := closureIncrement()
    fmt.Println(inc())  // 1
    fmt.Println(inc())  // 2 (remembers previous state)
}
```

---

## Common Pitfalls

### 1. Loop Variable Capture

**Problem:**

```go
// WRONG: All closures reference the same variable
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)  // All print 3!
    })
}

for _, f := range funcs {
    f()  // Output: 3, 3, 3
}
```

**Solution 1: Use a parameter**

```go
funcs := []func(){}
for i := 0; i < 3; i++ {
    i := i  // Create a new variable in each iteration
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}

for _, f := range funcs {
    f()  // Output: 0, 1, 2
}
```

**Solution 2: Pass as parameter**

```go
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func(val int) func() {
        return func() {
            fmt.Println(val)
        }
    }(i))
}

for _, f := range funcs {
    f()  // Output: 0, 1, 2
}
```

### 2. Goroutine and Closure Capture

**Problem:**

```go
// WRONG: Race condition
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // Unpredictable output
    }()
}
time.Sleep(time.Second)
```

**Solution:**

```go
// CORRECT: Pass value to goroutine
for i := 0; i < 5; i++ {
    go func(val int) {
        fmt.Println(val)  // Output: 0, 1, 2, 3, 4 (in any order)
    }(i)
}
time.Sleep(time.Second)
```

### 3. Memory Leaks

Be careful with long-lived closures capturing large data structures:

```go
// Potentially problematic
func createHandler(data []byte) func() {
    // If data is large (e.g., 100MB), it stays in memory
    // as long as the closure exists
    return func() {
        fmt.Println(len(data))
    }
}

// Better: Only capture what you need
func createHandler(dataLen int) func() {
    return func() {
        fmt.Println(dataLen)
    }
}
```

### 4. Unintended Modifications

```go
// Be aware: captured variables can be modified
func example() {
    x := 10
    
    modify := func() { x = 20 }
    print := func() { fmt.Println(x) }
    
    print()    // 10
    modify()
    print()    // 20 (x was modified!)
}
```

---

## Best Practices

### 1. Keep Closures Simple

```go
// Good: Simple, clear purpose
func counter(start int) func() int {
    count := start
    return func() int {
        count++
        return count
    }
}

// Avoid: Too complex, hard to understand
func complexClosure() func(int, string, bool) (int, string, error) {
    // Multiple captured variables, complex logic...
    // Consider using a struct instead
}
```

### 2. Document Captured Variables

```go
// Good: Clear documentation
// newRateLimiter returns a function that checks if an action
// is allowed based on the rate limit. It captures maxRequests
// and maintains request count internally.
func newRateLimiter(maxRequests int) func() bool {
    requestCount := 0
    
    return func() bool {
        if requestCount >= maxRequests {
            return false
        }
        requestCount++
        return true
    }
}
```

### 3. Consider Concurrency Safety

```go
// Unsafe for concurrent use
func unsafeCounter() func() int {
    count := 0
    return func() int {
        count++  // Race condition!
        return count
    }
}

// Safe for concurrent use
func safeCounter() func() int {
    var mu sync.Mutex
    count := 0
    
    return func() int {
        mu.Lock()
        defer mu.Unlock()
        count++
        return count
    }
}
```

### 4. Use Descriptive Names

```go
// Good: Clear names
makeAdder := func(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

// Avoid: Unclear names
f := func(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}
```

### 5. Limit Lifetime of Large Captures

```go
// Good: Only capture what's needed
func processSummary(records []Record) func() string {
    summary := calculateSummary(records)  // Create summary
    // records is no longer captured
    
    return func() string {
        return summary
    }
}
```

---

## Real-World Applications

### 1. HTTP Middleware

```go
func authMiddleware(secret string) func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            if token != secret {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            next(w, r)
        }
    }
}

func main() {
    auth := authMiddleware("secret-token")
    
    http.HandleFunc("/protected", auth(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Protected resource")
    }))
}
```

### 2. Configuration Builder

```go
type Config struct {
    Host string
    Port int
    SSL  bool
}

func configBuilder() func(func(*Config)) *Config {
    cfg := &Config{
        Host: "localhost",
        Port: 8080,
        SSL:  false,
    }
    
    return func(option func(*Config)) *Config {
        option(cfg)
        return cfg
    }
}

func main() {
    builder := configBuilder()
    
    cfg := builder(func(c *Config) {
        c.Host = "example.com"
        c.Port = 443
        c.SSL = true
    })
    
    fmt.Printf("%+v\n", cfg)  // {Host:example.com Port:443 SSL:true}
}
```

### 3. Event System

```go
type EventSystem struct {
    handlers map[string][]func(interface{})
}

func NewEventSystem() *EventSystem {
    return &EventSystem{
        handlers: make(map[string][]func(interface{})),
    }
}

func (es *EventSystem) On(event string, handler func(interface{})) func() {
    es.handlers[event] = append(es.handlers[event], handler)
    
    // Return an unsubscribe function (closure)
    return func() {
        handlers := es.handlers[event]
        for i, h := range handlers {
            if &h == &handler {
                es.handlers[event] = append(handlers[:i], handlers[i+1:]...)
                break
            }
        }
    }
}

func (es *EventSystem) Emit(event string, data interface{}) {
    for _, handler := range es.handlers[event] {
        handler(data)
    }
}

func main() {
    es := NewEventSystem()
    
    // Subscribe with a closure
    count := 0
    unsubscribe := es.On("click", func(data interface{}) {
        count++
        fmt.Printf("Click event #%d: %v\n", count, data)
    })
    
    es.Emit("click", "Button 1")  // Click event #1: Button 1
    es.Emit("click", "Button 2")  // Click event #2: Button 2
    
    unsubscribe()  // Unsubscribe using closure
    es.Emit("click", "Button 3")  // No output
}
```

### 4. Retry Logic with Exponential Backoff

```go
func retryWithBackoff(maxRetries int, initialDelay time.Duration) func(func() error) error {
    return func(operation func() error) error {
        delay := initialDelay
        
        for attempt := 0; attempt < maxRetries; attempt++ {
            err := operation()
            if err == nil {
                return nil
            }
            
            if attempt < maxRetries-1 {
                fmt.Printf("Attempt %d failed, retrying in %v...\n", attempt+1, delay)
                time.Sleep(delay)
                delay *= 2  // Exponential backoff
            }
        }
        
        return fmt.Errorf("operation failed after %d attempts", maxRetries)
    }
}

func main() {
    retry := retryWithBackoff(3, 1*time.Second)
    
    attempt := 0
    err := retry(func() error {
        attempt++
        if attempt < 3 {
            return fmt.Errorf("temporary error")
        }
        return nil
    })
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Success!")
    }
}
```

---

## Summary

**Key Takeaways:**

1. **Closures = Function + Environment**: They capture and remember variables from their surrounding scope

2. **State Management**: Perfect for maintaining state without using global variables or structs

3. **First-Class Functions**: Go treats functions as first-class citizens, enabling powerful functional programming patterns

4. **Memory Implications**: Captured variables live on the heap and persist as long as the closure exists

5. **Common Uses**:
   - Function factories
   - Data encapsulation
   - Callbacks with context
   - Middleware patterns
   - Memoization
   - Event handlers

6. **Watch Out For**:
   - Loop variable capture
   - Goroutine race conditions
   - Memory leaks with large captures
   - Concurrent access without synchronization

7. **Best Practices**:
   - Keep closures simple and focused
   - Document captured variables
   - Consider thread safety
   - Use descriptive names
   - Only capture what you need

Closures are a fundamental concept in Go that enable elegant, functional solutions to many programming problems. Master them, and you'll write more expressive and maintainable code!

---

## Further Reading

- [Go by Example: Closures](https://gobyexample.com/closures)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Go Blog: First Class Functions in Go](https://go.dev/blog/first-class-functions-in-go-and-new-go)

## Practice Exercises

Try implementing these to master closures:

1. **Rate Limiter**: Create a closure that limits how many times a function can be called per second
2. **Debouncer**: Implement a debounce function that delays execution until after a period of inactivity
3. **Once Function**: Create a function that can only be executed once (similar to `sync.Once`)
4. **Pipeline Builder**: Build a data processing pipeline using closures
5. **State Machine**: Implement a simple state machine using closures