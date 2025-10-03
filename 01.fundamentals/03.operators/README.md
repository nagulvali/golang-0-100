## Go Operators

Operators are symbols that perform operations on operands (variables, constants, or expressions). Go provides a comprehensive set of operators for arithmetic, comparison, logical operations, bit manipulation, and more.

### Operator Categories

Go operators can be categorized as:
- **Arithmetic Operators** - Mathematical operations
- **Comparison Operators** - Value comparison
- **Logical Operators** - Boolean logic
- **Bitwise Operators** - Bit-level operations
- **Assignment Operators** - Value assignment
- **Miscellaneous Operators** - Address, pointer, channel operations

### Operator Precedence

Operators have different precedence levels. Higher precedence operators are evaluated first. When operators have the same precedence, they are evaluated left-to-right (left-associative).

## Arithmetic Operators

Arithmetic operators perform mathematical operations on numeric operands.

### Basic Arithmetic Operations

```go
package main

import "fmt"

func main() {
    a := 15
    b := 4
    
    // Addition
    sum := a + b
    fmt.Printf("%d + %d = %d\n", a, b, sum) // 15 + 4 = 19
    
    // Subtraction
    diff := a - b
    fmt.Printf("%d - %d = %d\n", a, b, diff) // 15 - 4 = 11
    
    // Multiplication
    product := a * b
    fmt.Printf("%d * %d = %d\n", a, b, product) // 15 * 4 = 60
    
    // Division (integer division for integers)
    quotient := a / b
    fmt.Printf("%d / %d = %d\n", a, b, quotient) // 15 / 4 = 3
    
    // Modulus (remainder)
    remainder := a % b
    fmt.Printf("%d %% %d = %d\n", a, b, remainder) // 15 % 4 = 3
}
```

### Floating-Point Arithmetic

```go
func floatingPointExample() {
    x := 10.5
    y := 3.2
    
    fmt.Printf("%.2f + %.2f = %.2f\n", x, y, x+y) // 10.50 + 3.20 = 13.70
    fmt.Printf("%.2f - %.2f = %.2f\n", x, y, x-y) // 10.50 - 3.20 = 7.30
    fmt.Printf("%.2f * %.2f = %.2f\n", x, y, x*y) // 10.50 * 3.20 = 33.60
    fmt.Printf("%.2f / %.2f = %.2f\n", x, y, x/y) // 10.50 / 3.20 = 3.28
}
```

### Important Notes

- **Integer Division**: For integers, division truncates toward zero (no rounding)
- **Modulus with Negatives**: The sign of the result matches the dividend
- **Type Consistency**: Operands must be of compatible types
- **Overflow**: Integer overflow wraps around (two's complement)

```go
func divisionExamples() {
    // Integer division truncates
    fmt.Println(7 / 3)   // 2 (not 2.33)
    fmt.Println(-7 / 3)  // -2 (not -2.33)
    
    // Modulus with negatives
    fmt.Println(7 % 3)   // 1
    fmt.Println(-7 % 3)  // -1 (sign matches dividend)
    fmt.Println(7 % -3)  // 1
    
    // Floating-point division
    fmt.Println(7.0 / 3.0) // 2.3333333333333335
}
```



## Comparison Operators

Comparison operators (also called relational operators) compare two values and return a boolean result (`true` or `false`).

### Comparison Operations

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    c := 10
    
    // Equality
    fmt.Printf("%d == %d: %t\n", a, b, a == b) // 10 == 5: false
    fmt.Printf("%d == %d: %t\n", a, c, a == c) // 10 == 10: true
    
    // Inequality
    fmt.Printf("%d != %d: %t\n", a, b, a != b) // 10 != 5: true
    fmt.Printf("%d != %d: %t\n", a, c, a != c) // 10 != 10: false
    
    // Less than
    fmt.Printf("%d < %d: %t\n", a, b, a < b)   // 10 < 5: false
    fmt.Printf("%d < %d: %t\n", b, a, b < a)   // 5 < 10: true
    
    // Greater than
    fmt.Printf("%d > %d: %t\n", a, b, a > b)   // 10 > 5: true
    fmt.Printf("%d > %d: %t\n", b, a, b > a)   // 5 > 10: false
    
    // Less than or equal
    fmt.Printf("%d <= %d: %t\n", a, c, a <= c) // 10 <= 10: true
    fmt.Printf("%d <= %d: %t\n", a, b, a <= b) // 10 <= 5: false
    
    // Greater than or equal
    fmt.Printf("%d >= %d: %t\n", a, c, a >= c) // 10 >= 10: true
    fmt.Printf("%d >= %d: %t\n", a, b, a >= b) // 10 >= 5: true
}
```

### String Comparison

```go
func stringComparison() {
    str1 := "apple"
    str2 := "banana"
    str3 := "apple"
    
    fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2) // false
    fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3) // true
    fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)   // true (lexicographic)
}
```

### Floating-Point Comparison

```go
func floatComparison() {
    // Be careful with floating-point comparisons
    a := 0.1 + 0.2
    b := 0.3
    
    fmt.Printf("%.17f == %.17f: %t\n", a, b, a == b) // false due to precision
    
    // Use epsilon for floating-point comparison
    epsilon := 1e-9
    fmt.Printf("|%.17f - %.17f| < %e: %t\n", a, b, epsilon, abs(a-b) < epsilon)
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}
```

### Type-Specific Notes

- **Integers**: Direct comparison of values
- **Floats**: Be aware of precision issues
- **Strings**: Lexicographic (dictionary) order
- **Booleans**: `true` > `false`
- **Pointers**: Compare memory addresses



## Logical Operators

Logical operators work with boolean values and implement boolean logic. They are essential for combining conditions and controlling program flow.

### Logical Operations

```go
package main

import "fmt"

func main() {
    a := true
    b := false
    
    // Logical AND (&&)
    fmt.Printf("%t && %t = %t\n", a, a, a && a) // true && true = true
    fmt.Printf("%t && %t = %t\n", a, b, a && b) // true && false = false
    fmt.Printf("%t && %t = %t\n", b, b, b && b) // false && false = false
    
    // Logical OR (||)
    fmt.Printf("%t || %t = %t\n", a, a, a || a) // true || true = true
    fmt.Printf("%t || %t = %t\n", a, b, a || b) // true || false = true
    fmt.Printf("%t || %t = %t\n", b, b, b || b) // false || false = false
    
    // Logical NOT (!)
    fmt.Printf("!%t = %t\n", a, !a) // !true = false
    fmt.Printf("!%t = %t\n", b, !b) // !false = true
}
```

### Short-Circuit Evaluation

Go's logical operators use short-circuit evaluation:

```go
func shortCircuitExample() {
    // && short-circuits: if left side is false, right side is not evaluated
    fmt.Println("Short-circuit AND:")
    fmt.Println(false && someExpensiveFunction()) // someExpensiveFunction() not called
    
    // || short-circuits: if left side is true, right side is not evaluated
    fmt.Println("Short-circuit OR:")
    fmt.Println(true || someExpensiveFunction()) // someExpensiveFunction() not called
}

func someExpensiveFunction() bool {
    fmt.Println("Expensive function called!")
    return true
}
```

### Practical Examples

```go
func practicalExamples() {
    age := 25
    hasLicense := true
    hasInsurance := false
    
    // Complex conditions
    canDrive := age >= 18 && hasLicense
    needsInsurance := age >= 18 && hasLicense && !hasInsurance
    
    fmt.Printf("Can drive: %t\n", canDrive)           // true
    fmt.Printf("Needs insurance: %t\n", needsInsurance) // true
    
    // Input validation
    username := "john"
    password := ""
    
    isValid := len(username) > 0 && len(password) > 0
    fmt.Printf("Valid login: %t\n", isValid) // false
}
```

### Truth Tables

```go
func truthTables() {
    fmt.Println("AND Truth Table:")
    fmt.Printf("true  && true  = %t\n", true && true)   // true
    fmt.Printf("true  && false = %t\n", true && false)  // false
    fmt.Printf("false && true  = %t\n", false && true)  // false
    fmt.Printf("false && false = %t\n", false && false) // false
    
    fmt.Println("\nOR Truth Table:")
    fmt.Printf("true  || true  = %t\n", true || true)   // true
    fmt.Printf("true  || false = %t\n", true || false)  // true
    fmt.Printf("false || true  = %t\n", false || true)  // true
    fmt.Printf("false || false = %t\n", false || false) // false
    
    fmt.Println("\nNOT Truth Table:")
    fmt.Printf("!true  = %t\n", !true)  // false
    fmt.Printf("!false = %t\n", !false) // true
}
```

### Operator Precedence

```go
func precedenceExample() {
    a := true
    b := false
    c := true
    
    // ! has highest precedence, then &&, then ||
    result1 := !a && b || c
    result2 := !(a && b) || c
    result3 := !a && (b || c)
    
    fmt.Printf("!a && b || c = %t\n", result1)     // false
    fmt.Printf("!(a && b) || c = %t\n", result2)   // true
    fmt.Printf("!a && (b || c) = %t\n", result3)   // false
}
```


## Bitwise Operators

Bitwise operators perform operations on individual bits of integer operands. They are essential for low-level programming, flag manipulation, and performance-critical code.

### Basic Bitwise Operations

```go
package main

import "fmt"

func main() {
    a := 12  // 1100 in binary
    b := 10  // 1010 in binary
    
    fmt.Printf("a = %d (%04b)\n", a, a)
    fmt.Printf("b = %d (%04b)\n", b, b)
    fmt.Println()
    
    // Bitwise AND (&)
    result := a & b
    fmt.Printf("%d & %d = %d (%04b)\n", a, b, result, result) // 8 (1000)
    
    // Bitwise OR (|)
    result = a | b
    fmt.Printf("%d | %d = %d (%04b)\n", a, b, result, result) // 14 (1110)
    
    // Bitwise XOR (^)
    result = a ^ b
    fmt.Printf("%d ^ %d = %d (%04b)\n", a, b, result, result) // 6 (0110)
    
    // Bitwise NOT (^) - unary operator
    result = ^a
    fmt.Printf("^%d = %d (%04b)\n", a, result, result) // -13 (in two's complement)
}
```

### Shift Operators

```go
func shiftOperators() {
    x := 8  // 1000 in binary
    
    fmt.Printf("Original: %d (%04b)\n", x, x)
    
    // Left shift (<<) - multiply by powers of 2
    fmt.Printf("%d << 1 = %d (%04b)\n", x, x<<1, x<<1)   // 16 (10000)
    fmt.Printf("%d << 2 = %d (%04b)\n", x, x<<2, x<<2)   // 32 (100000)
    fmt.Printf("%d << 3 = %d (%04b)\n", x, x<<3, x<<3)   // 64 (1000000)
    
    // Right shift (>>) - divide by powers of 2
    fmt.Printf("%d >> 1 = %d (%04b)\n", x, x>>1, x>>1)   // 4 (0100)
    fmt.Printf("%d >> 2 = %d (%04b)\n", x, x>>2, x>>2)   // 2 (0010)
    fmt.Printf("%d >> 3 = %d (%04b)\n", x, x>>3, x>>3)   // 1 (0001)
}
```

### Bit Clear Operator (&^)

```go
func bitClearOperator() {
    a := 15  // 1111 in binary
    b := 5   // 0101 in binary
    
    // &^ (AND NOT) - clears bits in 'a' where 'b' has 1
    result := a &^ b
    fmt.Printf("%d &^ %d = %d\n", a, b, result)
    fmt.Printf("%04b &^ %04b = %04b\n", a, b, result) // 1010
    
    // Equivalent to: a & (^b)
    result2 := a & (^b)
    fmt.Printf("%d & (^%d) = %d\n", a, b, result2) // Same result
}
```

### Practical Bit Manipulation

```go
func practicalBitManipulation() {
    // Check if bit is set
    num := 12  // 1100
    bit := 2   // Check 3rd bit (0-indexed)
    
    isSet := (num & (1 << bit)) != 0
    fmt.Printf("Bit %d in %d is set: %t\n", bit, num, isSet) // true
    
    // Set a bit
    num |= (1 << 1)  // Set bit 1
    fmt.Printf("After setting bit 1: %d (%04b)\n", num, num) // 14 (1110)
    
    // Clear a bit
    num &^= (1 << 2)  // Clear bit 2
    fmt.Printf("After clearing bit 2: %d (%04b)\n", num, num) // 10 (1010)
    
    // Toggle a bit
    num ^= (1 << 0)  // Toggle bit 0
    fmt.Printf("After toggling bit 0: %d (%04b)\n", num, num) // 11 (1011)
}
```

### Flag Manipulation

```go
const (
    Read    = 1 << iota  // 1 (001)
    Write                 // 2 (010)
    Execute               // 4 (100)
)

func flagManipulation() {
    permissions := Read | Write  // 3 (011)
    
    fmt.Printf("Permissions: %d (%03b)\n", permissions, permissions)
    
    // Check if read permission
    if permissions&Read != 0 {
        fmt.Println("Read permission granted")
    }
    
    // Check if execute permission
    if permissions&Execute != 0 {
        fmt.Println("Execute permission granted")
    } else {
        fmt.Println("Execute permission denied")
    }
    
    // Add execute permission
    permissions |= Execute
    fmt.Printf("After adding execute: %d (%03b)\n", permissions, permissions)
    
    // Remove write permission
    permissions &^= Write
    fmt.Printf("After removing write: %d (%03b)\n", permissions, permissions)
}
```

### Bit Counting and Manipulation

```go
func bitCounting() {
    num := 13  // 1101 in binary
    
    // Count set bits (population count)
    count := 0
    temp := num
    for temp != 0 {
        count += temp & 1
        temp >>= 1
    }
    fmt.Printf("Number of set bits in %d: %d\n", num, count) // 3
    
    // Check if number is power of 2
    isPowerOfTwo := num != 0 && (num&(num-1)) == 0
    fmt.Printf("%d is power of 2: %t\n", num, isPowerOfTwo) // false
    
    // Get the rightmost set bit
    rightmostSetBit := num & -num
    fmt.Printf("Rightmost set bit in %d: %d\n", num, rightmostSetBit) // 1
}
```

### Important Notes

- **Operand Types**: Bitwise operators work on integer types only
- **Sign Extension**: Right shift on signed integers extends the sign bit
- **Overflow**: Left shift can cause overflow
- **Performance**: Bitwise operations are very fast
- **Readability**: Use constants for bit positions to improve code clarity

## Assignment Operators

Assignment operators assign values to variables. Go provides both simple assignment and compound assignment operators that combine an operation with assignment.

### Simple Assignment

```go
package main

import "fmt"

func main() {
    // Simple assignment
    var x int
    x = 10
    fmt.Printf("x = %d\n", x) // x = 10
    
    // Multiple assignment
    var a, b, c int
    a, b, c = 1, 2, 3
    fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c) // a=1, b=2, c=3
    
    // Swap values
    a, b = b, a
    fmt.Printf("After swap: a=%d, b=%d\n", a, b) // a=2, b=1
}
```

### Compound Assignment Operators

```go
func compoundAssignment() {
    x := 10
    fmt.Printf("Initial x: %d\n", x)
    
    // Addition assignment
    x += 5
    fmt.Printf("x += 5: %d\n", x) // 15
    
    // Subtraction assignment
    x -= 3
    fmt.Printf("x -= 3: %d\n", x) // 12
    
    // Multiplication assignment
    x *= 2
    fmt.Printf("x *= 2: %d\n", x) // 24
    
    // Division assignment
    x /= 4
    fmt.Printf("x /= 4: %d\n", x) // 6
    
    // Modulus assignment
    x %= 5
    fmt.Printf("x %%= 5: %d\n", x) // 1
}
```

### Bitwise Assignment Operators

```go
func bitwiseAssignment() {
    x := 12  // 1100 in binary
    fmt.Printf("Initial x: %d (%04b)\n", x, x)
    
    // Bitwise AND assignment
    x &= 10  // 1010
    fmt.Printf("x &= 10: %d (%04b)\n", x, x) // 8 (1000)
    
    // Bitwise OR assignment
    x |= 3   // 0011
    fmt.Printf("x |= 3: %d (%04b)\n", x, x) // 11 (1011)
    
    // Bitwise XOR assignment
    x ^= 5   // 0101
    fmt.Printf("x ^= 5: %d (%04b)\n", x, x) // 14 (1110)
    
    // Left shift assignment
    x <<= 1
    fmt.Printf("x <<= 1: %d (%04b)\n", x, x) // 28 (11100)
    
    // Right shift assignment
    x >>= 2
    fmt.Printf("x >>= 2: %d (%04b)\n", x, x) // 7 (0111)
}
```

### String Assignment

```go
func stringAssignment() {
    str := "Hello"
    fmt.Printf("Initial: %s\n", str)
    
    // String concatenation assignment
    str += " World"
    fmt.Printf("After +=: %s\n", str) // "Hello World"
    
    // String repetition (using a loop)
    str = "Go"
    for i := 0; i < 3; i++ {
        str += "!"
    }
    fmt.Printf("After repetition: %s\n", str) // "Go!!!"
}
```

### Type-Specific Assignment

```go
func typeSpecificAssignment() {
    // Integer types
    var i int = 10
    i += 5
    fmt.Printf("int: %d\n", i)
    
    // Floating-point types
    var f float64 = 3.14
    f *= 2.0
    fmt.Printf("float64: %.2f\n", f)
    
    // Boolean (only simple assignment)
    var b bool = true
    b = false
    fmt.Printf("bool: %t\n", b)
    
    // Complex numbers
    var c complex128 = 1 + 2i
    c += 3 + 4i
    fmt.Printf("complex128: %v\n", c)
}
```

### Assignment with Type Conversion

```go
func assignmentWithConversion() {
    // Implicit conversion (when types are compatible)
    var i int = 42
    var f float64 = float64(i)  // Explicit conversion
    fmt.Printf("int %d -> float64 %.1f\n", i, f)
    
    // Multiple assignment with conversion
    var x, y int = 10, 20
    var a, b float64 = float64(x), float64(y)
    fmt.Printf("Converted: %.1f, %.1f\n", a, b)
}
```

### Assignment Best Practices

```go
func assignmentBestPractices() {
    // Use short variable declaration when possible
    x := 10  // Preferred over var x int = 10
    
    // Use compound operators for clarity
    count := 0
    count++  // Preferred over count = count + 1
    count += 5  // Preferred over count = count + 5
    
    // Be careful with type conversions
    var i int = 42
    var f float64 = float64(i)  // Explicit conversion required
    
    // Use multiple assignment for swapping
    a, b := 1, 2
    a, b = b, a  // Swap values
    fmt.Printf("Swapped: a=%d, b=%d\n", a, b)
}
```

### Important Notes

- **Type Safety**: Go requires explicit type conversion for different types
- **Multiple Assignment**: Go supports multiple assignment and swapping
- **Short Declaration**: Use `:=` for declaration and assignment
- **Compound Operators**: More concise and often more efficient
- **Immutability**: Strings are immutable; concatenation creates new strings

## Miscellaneous Operators

Go provides several special operators for memory management, pointers, and channel operations.

### Address and Pointer Operators

```go
package main

import "fmt"

func main() {
    x := 42
    fmt.Printf("Value of x: %d\n", x)
    
    // Address operator (&) - gets the memory address
    ptr := &x
    fmt.Printf("Address of x: %p\n", ptr)
    fmt.Printf("Value at address: %d\n", *ptr)
    
    // Dereference operator (*) - gets the value at the address
    *ptr = 100
    fmt.Printf("New value of x: %d\n", x) // x is now 100
}
```

### Pointer Operations

```go
func pointerOperations() {
    // Declare and initialize
    a := 10
    b := 20
    
    // Get addresses
    ptrA := &a
    ptrB := &b
    
    fmt.Printf("a = %d, address = %p\n", a, ptrA)
    fmt.Printf("b = %d, address = %p\n", b, ptrB)
    
    // Modify values through pointers
    *ptrA = 30
    *ptrB = 40
    
    fmt.Printf("After modification:\n")
    fmt.Printf("a = %d\n", a) // 30
    fmt.Printf("b = %d\n", b) // 40
    
    // Pointer arithmetic (not directly supported in Go)
    // But you can compare pointers
    fmt.Printf("ptrA == ptrB: %t\n", ptrA == ptrB) // false
}
```

### Channel Operators

```go
func channelOperators() {
    // Create a channel
    ch := make(chan int, 2)
    
    // Send operator (<-)
    ch <- 42  // Send value 42 to channel
    ch <- 84  // Send value 84 to channel
    
    // Receive operator (<-)
    value1 := <-ch  // Receive value from channel
    value2 := <-ch  // Receive another value
    
    fmt.Printf("Received: %d, %d\n", value1, value2) // 42, 84
    
    // Close channel
    close(ch)
    
    // Check if channel is closed
    value, ok := <-ch
    fmt.Printf("Value: %d, Channel open: %t\n", value, ok) // 0, false
}
```

### Channel Direction

```go
func channelDirection() {
    // Send-only channel
    var sendOnly chan<- int = make(chan int)
    
    // Receive-only channel
    var receiveOnly <-chan int = make(chan int)
    
    // Bidirectional channel
    var bidirectional chan int = make(chan int)
    
    fmt.Printf("Send-only: %T\n", sendOnly)
    fmt.Printf("Receive-only: %T\n", receiveOnly)
    fmt.Printf("Bidirectional: %T\n", bidirectional)
}
```

### Type Assertion and Type Switch

```go
func typeAssertion() {
    var i interface{} = "hello"
    
    // Type assertion
    s := i.(string)
    fmt.Printf("String: %s\n", s)
    
    // Type assertion with ok
    s, ok := i.(string)
    fmt.Printf("String: %s, OK: %t\n", s, ok)
    
    // Type switch
    switch v := i.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Built-in Functions as Operators

```go
func builtinFunctions() {
    // len() - length of array, slice, string, map, channel
    arr := [5]int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3}
    str := "hello"
    m := map[string]int{"a": 1, "b": 2}
    
    fmt.Printf("Array length: %d\n", len(arr))   // 5
    fmt.Printf("Slice length: %d\n", len(slice)) // 3
    fmt.Printf("String length: %d\n", len(str))  // 5
    fmt.Printf("Map length: %d\n", len(m))       // 2
    
    // cap() - capacity of slice, array, channel
    fmt.Printf("Slice capacity: %d\n", cap(slice)) // 3
    
    // make() - create slices, maps, channels
    newSlice := make([]int, 5, 10)
    newMap := make(map[string]int)
    newChan := make(chan int, 5)
    
    fmt.Printf("New slice: %v\n", newSlice)
    fmt.Printf("New map: %v\n", newMap)
    fmt.Printf("New channel: %v\n", newChan)
    
    // new() - allocate memory for a type
    ptr := new(int)
    *ptr = 42
    fmt.Printf("New int pointer: %d\n", *ptr)
}
```

### Operator Precedence Summary

```go
func operatorPrecedence() {
    // From highest to lowest precedence:
    // 1. Unary operators: +, -, !, ^, *, &, <-
    // 2. Binary operators: *, /, %, <<, >>, &, &^
    // 3. Binary operators: +, -, |, ^
    // 4. Comparison operators: ==, !=, <, <=, >, >=
    // 5. Logical operators: &&
    // 6. Logical operators: ||
    
    a := 10
    b := 5
    c := 2
    
    // Example showing precedence
    result := a + b * c  // Equivalent to: a + (b * c)
    fmt.Printf("%d + %d * %d = %d\n", a, b, c, result) // 20
    
    // Use parentheses to change precedence
    result = (a + b) * c
    fmt.Printf("(%d + %d) * %d = %d\n", a, b, c, result) // 30
}
```

### Important Notes

- **Address Operator (&)**: Returns the memory address of a variable
- **Dereference Operator (*)**: Accesses the value at a memory address
- **Channel Operators (<-)**: Used for sending and receiving values
- **Type Assertion**: Converts interface{} to specific types
- **Built-in Functions**: len, cap, make, new are not operators but function-like
- **Precedence**: Use parentheses to clarify operator precedence
- **Memory Safety**: Go's garbage collector manages memory automatically


