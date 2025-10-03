## Go Loops

Go has only one loop construct: the `for` statement. It can be used in three forms: traditional C-style for loop, while-style loop, and infinite loop. Go also provides the `range` keyword for iterating over arrays, slices, strings, maps, and channels.

### Basic For Loop

The traditional for loop with initialization, condition, and post statement:

```go
// Basic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Multiple variables
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Printf("i=%d, j=%d\n", i, j)
}

// Using different increment
for i := 0; i < 10; i += 2 {
    fmt.Println(i) // 0, 2, 4, 6, 8
}
```

### While-Style Loop

Go doesn't have a separate `while` keyword. Use `for` with only a condition:

```go
// While-style loop
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}

// Reading input until condition
for {
    var input string
    fmt.Scanln(&input)
    if input == "quit" {
        break
    }
    fmt.Println("You entered:", input)
}
```

### Infinite Loop

Use `for` without any condition for an infinite loop:

```go
// Infinite loop (usually with break)
for {
    // Do something
    if someCondition {
        break
    }
}

// Common pattern with channels
for {
    select {
    case msg := <-ch:
        process(msg)
    case <-done:
        return
    }
}
```

### Range Loop

The `range` keyword provides a clean way to iterate over collections:

```go
// Arrays and slices
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Only values (use blank identifier for index)
for _, value := range numbers {
    fmt.Println(value)
}

// Only indices
for index := range numbers {
    fmt.Println(index)
}

// Strings (iterates over runes, not bytes)
text := "Hello, 世界"
for i, r := range text {
    fmt.Printf("Index: %d, Rune: %c\n", i, r)
}

// Maps
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Channels
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for value := range ch {
    fmt.Println(value)
}
```

### Loop Control Statements

#### Break Statement

Exits the innermost loop immediately:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i) // Prints 0, 1, 2, 3, 4
}

// Labeled break (breaks outer loop)
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

#### Continue Statement

Skips the rest of the current iteration and continues with the next:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Prints 1, 3, 5, 7, 9
}

// Labeled continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            continue outer
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### Nested Loops

```go
// Multiplication table
for i := 1; i <= 5; i++ {
    for j := 1; j <= 5; j++ {
        fmt.Printf("%2d ", i*j)
    }
    fmt.Println()
}

// Pattern printing
for i := 1; i <= 5; i++ {
    for j := 1; j <= i; j++ {
        fmt.Print("*")
    }
    fmt.Println()
}
```

### Common Loop Patterns

#### Iterating with Index and Value

```go
items := []string{"apple", "banana", "cherry"}
for i, item := range items {
    fmt.Printf("%d: %s\n", i, item)
}
```

#### Iterating Backwards

```go
// Using traditional for loop
for i := len(items) - 1; i >= 0; i-- {
    fmt.Printf("%d: %s\n", i, items[i])
}

// Using range (less efficient for large slices)
for i := range items {
    reverseIndex := len(items) - 1 - i
    fmt.Printf("%d: %s\n", reverseIndex, items[reverseIndex])
}
```

#### Conditional Iteration

```go
// Skip certain elements
for _, item := range items {
    if item == "banana" {
        continue
    }
    fmt.Println(item)
}

// Early termination
for _, item := range items {
    if item == "cherry" {
        break
    }
    fmt.Println(item)
}
```

#### Loop with Defer

```go
// Defer executes in LIFO order
for i := 0; i < 3; i++ {
    defer fmt.Printf("Defer %d\n", i)
    fmt.Printf("Loop %d\n", i)
}
// Output:
// Loop 0
// Loop 1
// Loop 2
// Defer 2
// Defer 1
// Defer 0
```

### Performance Considerations

#### Range vs Traditional For

```go
// Range (creates copies of values)
for _, value := range slice {
    // value is a copy
}

// Traditional for (accesses elements directly)
for i := 0; i < len(slice); i++ {
    // slice[i] is direct access
}

// Range with pointer (avoids copying)
for i := range slice {
    // &slice[i] gives pointer to element
}
```

#### Pre-allocating Slices

```go
// Inefficient: grows slice multiple times
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i*2)
}

// Efficient: pre-allocate capacity
result := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    result = append(result, i*2)
}

// Most efficient: set length and assign
result := make([]int, 1000)
for i := 0; i < 1000; i++ {
    result[i] = i * 2
}
```

### Error Handling in Loops

```go
// Process items with error handling
for i, item := range items {
    if err := processItem(item); err != nil {
        fmt.Printf("Error processing item %d: %v\n", i, err)
        continue // or break, depending on requirements
    }
    fmt.Printf("Successfully processed item %d\n", i)
}
```

### Loop with Goroutines

```go
// Process items concurrently
var wg sync.WaitGroup
for i, item := range items {
    wg.Add(1)
    go func(index int, value string) {
        defer wg.Done()
        processItem(value)
        fmt.Printf("Processed item %d\n", index)
    }(i, item)
}
wg.Wait()
```

### Best Practices

1. **Use range when possible**: It's more idiomatic and less error-prone
2. **Be careful with closures**: Capture loop variables properly in goroutines
3. **Consider performance**: Use traditional for loops for performance-critical code
4. **Handle errors gracefully**: Don't let one bad item stop the entire loop
5. **Use meaningful variable names**: `i, j, k` are fine for simple loops, but use descriptive names when appropriate
6. **Avoid infinite loops**: Unless you have a clear exit condition
7. **Use labeled breaks/continues**: For complex nested loop control

### Common Pitfalls

```go
// Pitfall 1: Modifying slice while iterating
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    if v%2 == 0 {
        slice = append(slice[:i], slice[i+1:]...) // Dangerous!
    }
}

// Pitfall 2: Closure capturing loop variable
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i) // Will print 3, 3, 3
    }()
}

// Correct way:
for i := 0; i < 3; i++ {
    go func(val int) {
        fmt.Println(val) // Will print 0, 1, 2
    }(i)
}

// Pitfall 3: Range over pointer to slice
slice := []int{1, 2, 3}
for i, v := range &slice { // This works but is unusual
    fmt.Println(i, v)
}
```

### Quick Reference

```go
// Basic for loop
for i := 0; i < n; i++ { }

// While-style
for condition { }

// Infinite loop
for { }

// Range over slice
for i, v := range slice { }

// Range over map
for k, v := range map { }

// Range over string
for i, r := range str { }

// Range over channel
for v := range ch { }

// Break and continue
for { if condition { break } }
for { if condition { continue } }

// Labeled break/continue
outer: for { for { break outer } }
```
