## Go Data Types

Go is a statically typed language, meaning every variable has a specific type that is known at compile time. Go provides a rich set of built-in data types that can be categorized into basic types, composite types, and other special types.

### Type Categories

Go's data types can be grouped as:
- **Basic Types** - Numeric, boolean, and string types
- **Composite Types** - Arrays, slices, maps, structs
- **Reference Types** - Pointers, functions, interfaces, channels
- **Custom Types** - User-defined types and type aliases

## Basic Data Types

### Numeric Types

#### Integer Types

Go provides both signed and unsigned integer types in various sizes:

```go
package main

import (
    "fmt"
    "math"
    "unsafe"
)

func main() {
    // Signed integers
    var a int8 = 127                    // -128 to 127
    var b int16 = 32000                 // -32,768 to 32,767
    var c int32 = 2000000000            // -2,147,483,648 to 2,147,483,647
    var d int64 = 9000000000000000000   // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
    var e int = 42                      // Platform-dependent (32 or 64 bits)
    
    // Unsigned integers
    var f uint8 = 255                   // 0 to 255 (also called byte)
    var g uint16 = 65535                // 0 to 65,535
    var h uint32 = 4000000000           // 0 to 4,294,967,295
    var i uint64 = 18000000000000000000 // 0 to 18,446,744,073,709,551,615
    var j uint = 42                     // Platform-dependent
    
    fmt.Printf("int8: %d, size: %d bytes\n", a, unsafe.Sizeof(a))
    fmt.Printf("int64: %d, size: %d bytes\n", d, unsafe.Sizeof(d))
    fmt.Printf("uint8: %d, size: %d bytes\n", f, unsafe.Sizeof(f))
}
```

#### Floating-Point Types

```go
func floatingPointTypes() {
    var f32 float32 = 3.14159
    var f64 float64 = 2.718281828459045
    
    // Scientific notation
    var scientific float64 = 1.23e-4  // 0.000123
    
    // Special values
    var posInf float64 = math.Inf(1)   // +Infinity
    var negInf float64 = math.Inf(-1)  // -Infinity
    var nan float64 = math.NaN()       // Not a Number
    
    fmt.Printf("float32: %.5f\n", f32)
    fmt.Printf("float64: %.15f\n", f64)
    fmt.Printf("Scientific: %e\n", scientific)
    fmt.Printf("Infinity: %f\n", posInf)
    fmt.Printf("NaN: %f\n", nan)
}
```

#### Complex Types

```go
func complexTypes() {
    var c64 complex64 = 3 + 4i
    var c128 complex128 = 1.5 + 2.5i
    
    // Complex number operations
    realPart := real(c128)
    imagPart := imag(c128)
    
    fmt.Printf("complex64: %v\n", c64)
    fmt.Printf("complex128: %v\n", c128)
    fmt.Printf("Real part: %.2f\n", realPart)
    fmt.Printf("Imaginary part: %.2f\n", imagPart)
}
```

### Boolean Type

```go
func booleanType() {
    var isTrue bool = true
    var isFalse bool = false
    
    // Boolean operations
    result := isTrue && isFalse  // false
    result2 := isTrue || isFalse // true
    result3 := !isTrue           // false
    
    fmt.Printf("true && false = %t\n", result)
    fmt.Printf("true || false = %t\n", result2)
    fmt.Printf("!true = %t\n", result3)
}
```

### String Type

```go
func stringType() {
    // String literals
    str1 := "Hello, World!"
    str2 := `This is a raw string literal
    that can span multiple lines
    and preserves formatting`
    
    // String operations
    length := len(str1)
    char := str1[0]  // First character (byte)
    
    // String concatenation
    combined := str1 + " " + "Go!"
    
    // String formatting
    formatted := fmt.Sprintf("Length: %d, First char: %c", length, char)
    
    fmt.Printf("String: %s\n", str1)
    fmt.Printf("Length: %d\n", length)
    fmt.Printf("Combined: %s\n", combined)
    fmt.Printf("Formatted: %s\n", formatted)
}
```

### Byte and Rune Types

```go
func byteAndRuneTypes() {
    // byte is an alias for uint8
    var b byte = 65  // ASCII 'A'
    
    // rune is an alias for int32 (Unicode code point)
    var r rune = '♥'  // Unicode heart symbol
    var r2 rune = 65  // ASCII 'A'
    
    // String to rune slice
    str := "Hello, 世界"
    runes := []rune(str)
    bytes := []byte(str)
    
    fmt.Printf("byte: %c (%d)\n", b, b)
    fmt.Printf("rune: %c (%d)\n", r, r)
    fmt.Printf("String: %s\n", str)
    fmt.Printf("Runes: %v\n", runes)
    fmt.Printf("Bytes: %v\n", bytes)
}
```

## Composite Data Types

Composite types are built from basic types and can hold multiple values.

### Arrays

Arrays are fixed-size sequences of elements of the same type:

```go
func arrayTypes() {
    // Array declaration and initialization
    var arr1 [3]int = [3]int{1, 2, 3}
    var arr2 [5]string = [5]string{"a", "b", "c", "d", "e"}
    
    // Short declaration
    arr3 := [4]float64{1.1, 2.2, 3.3, 4.4}
    
    // Array with ellipsis (compiler determines size)
    arr4 := [...]int{10, 20, 30, 40, 50}
    
    // Partial initialization (remaining elements are zero values)
    arr5 := [5]int{1, 2}  // [1, 2, 0, 0, 0]
    
    // Accessing elements
    fmt.Printf("arr1[0]: %d\n", arr1[0])
    fmt.Printf("arr4 length: %d\n", len(arr4))
    fmt.Printf("arr5: %v\n", arr5)
}
```

### Slices

Slices are dynamically-sized, flexible views into arrays:

```go
func sliceTypes() {
    // Slice declaration
    var slice1 []int
    slice2 := []int{1, 2, 3, 4, 5}
    
    // Create slice from array
    arr := [5]int{1, 2, 3, 4, 5}
    slice3 := arr[1:4]  // [2, 3, 4]
    
    // Using make
    slice4 := make([]int, 5)        // length 5, capacity 5
    slice5 := make([]int, 5, 10)    // length 5, capacity 10
    
    // Slice operations
    slice2 = append(slice2, 6, 7, 8)
    fmt.Printf("slice2: %v, length: %d, capacity: %d\n", 
               slice2, len(slice2), cap(slice2))
    
    // Copy slices
    slice6 := make([]int, len(slice2))
    copy(slice6, slice2)
    fmt.Printf("copied slice: %v\n", slice6)
}
```

### Maps

Maps are key-value pairs with fast lookup:

```go
func mapTypes() {
    // Map declaration and initialization
    var m1 map[string]int
    m1 = make(map[string]int)
    
    // Short declaration
    m2 := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }
    
    // Map operations
    m2["grape"] = 12
    fmt.Printf("banana count: %d\n", m2["banana"])
    
    // Check if key exists
    if count, exists := m2["kiwi"]; exists {
        fmt.Printf("kiwi count: %d\n", count)
    } else {
        fmt.Println("kiwi not found")
    }
    
    // Delete key
    delete(m2, "orange")
    fmt.Printf("map after deletion: %v\n", m2)
    
    // Iterate over map
    for fruit, count := range m2 {
        fmt.Printf("%s: %d\n", fruit, count)
    }
}
```

### Structs

Structs are collections of named fields:

```go
// Define struct type
  type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address
}

type Address struct {
    Street string
    City   string
    Zip    string
}

func structTypes() {
    // Struct initialization
    p1 := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            Zip:    "10001",
        },
    }
    
    // Anonymous struct
    p2 := struct {
      Name string
      Age  int
    }{
        Name: "Bob",
        Age:  25,
    }
    
    // Access fields
    fmt.Printf("Person: %s, Age: %d\n", p1.Name, p1.Age)
    fmt.Printf("Address: %s, %s %s\n", p1.Address.Street, 
               p1.Address.City, p1.Address.Zip)
    
    // Pointer to struct
    p3 := &Person{Name: "Charlie", Age: 35}
    fmt.Printf("Pointer person: %s\n", p3.Name)
}
```

## Reference Types

### Pointers

Pointers hold memory addresses of values:

```go
func pointerTypes() {
    x := 42
    var p *int = &x  // p points to x
    
    fmt.Printf("x = %d\n", x)
    fmt.Printf("p = %p\n", p)
    fmt.Printf("*p = %d\n", *p)  // dereference
    
    // Modify value through pointer
    *p = 100
    fmt.Printf("x after modification: %d\n", x)
    
    // Pointer to pointer
    var pp **int = &p
    fmt.Printf("**pp = %d\n", **pp)
    
    // Nil pointer
    var nilPtr *int
    if nilPtr == nil {
        fmt.Println("nilPtr is nil")
    }
}
```

### Functions

Functions are first-class types in Go:

```go
func functionTypes() {
    // Function type declaration
    var add func(int, int) int
    
    // Assign function to variable
    add = func(a, b int) int {
        return a + b
    }
    
    // Use function
    result := add(3, 4)
    fmt.Printf("3 + 4 = %d\n", result)
    
    // Function as parameter
    calculate := func(op func(int, int) int, a, b int) int {
        return op(a, b)
    }
    
    multiply := func(x, y int) int { return x * y }
    fmt.Printf("3 * 4 = %d\n", calculate(multiply, 3, 4))
}
```

### Interfaces

Interfaces define method signatures:

```go
// Define interface
type Writer interface {
    Write([]byte) (int, error)
}

type StringWriter struct {
    buffer string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
    sw.buffer += string(data)
    return len(data), nil
}

func interfaceTypes() {
    var w Writer
    sw := &StringWriter{}
    
    w = sw  // StringWriter implements Writer interface
    
    data := []byte("Hello, Interface!")
    n, err := w.Write(data)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Wrote %d bytes\n", n)
    }
    
    // Type assertion
    if sw, ok := w.(*StringWriter); ok {
        fmt.Printf("Buffer content: %s\n", sw.buffer)
    }
}
```

### Channels

Channels are used for communication between goroutines:

```go
func channelTypes() {
    // Channel declaration
    var ch1 chan int
    ch1 = make(chan int)
    
    // Buffered channel
    ch2 := make(chan string, 3)
    
    // Send-only channel
    var sendOnly chan<- int = make(chan int)
    
    // Receive-only channel
    var receiveOnly <-chan int = make(chan int)
    
    // Channel operations
    go func() {
        ch2 <- "Hello"
        ch2 <- "World"
        ch2 <- "Go"
        close(ch2)
    }()
    
    // Receive from channel
    for msg := range ch2 {
        fmt.Printf("Received: %s\n", msg)
    }
    
    fmt.Printf("Channel types: %T, %T\n", sendOnly, receiveOnly)
}
```

## Custom Types

### Type Definitions

```go
// Define new type
type UserID int
type UserName string

// Define struct type
type User struct {
    ID   UserID
    Name UserName
}

func customTypes() {
    var id UserID = 123
    var name UserName = "john_doe"
    
    // Type conversion required
    regularInt := int(id)
    regularString := string(name)
    
    fmt.Printf("User ID: %d\n", id)
    fmt.Printf("User Name: %s\n", name)
    fmt.Printf("Regular int: %d\n", regularInt)
    fmt.Printf("Regular string: %s\n", regularString)
    
    // Create user
    user := User{
        ID:   id,
        Name: name,
    }
    fmt.Printf("User: %+v\n", user)
}
```

### Type Aliases

```go
// Type alias (Go 1.9+)
type MyInt = int
type MyString = string

func typeAliases() {
    var x MyInt = 42
    var s MyString = "hello"
    
    // No conversion needed - aliases are the same type
    var regularInt int = x
    var regularString string = s
    
    fmt.Printf("MyInt: %d, regular int: %d\n", x, regularInt)
    fmt.Printf("MyString: %s, regular string: %s\n", s, regularString)
}
```

## Zero Values

Every type in Go has a zero value (default value when declared but not initialized):

```go
func zeroValues() {
    var i int
    var f float64
    var b bool
    var s string
    var arr [3]int
    var slice []int
    var m map[string]int
    var p *int
    var ch chan int
    
    fmt.Printf("int zero value: %d\n", i)           // 0
    fmt.Printf("float64 zero value: %f\n", f)       // 0.000000
    fmt.Printf("bool zero value: %t\n", b)          // false
    fmt.Printf("string zero value: '%s'\n", s)      // ""
    fmt.Printf("array zero value: %v\n", arr)       // [0 0 0]
    fmt.Printf("slice zero value: %v\n", slice)     // []
    fmt.Printf("map zero value: %v\n", m)           // map[]
    fmt.Printf("pointer zero value: %v\n", p)       // <nil>
    fmt.Printf("channel zero value: %v\n", ch)      // <nil>
}
```

## Type Size and Memory

```go
import "unsafe"

func typeSizes() {
    fmt.Printf("int8 size: %d bytes\n", unsafe.Sizeof(int8(0)))
    fmt.Printf("int16 size: %d bytes\n", unsafe.Sizeof(int16(0)))
    fmt.Printf("int32 size: %d bytes\n", unsafe.Sizeof(int32(0)))
    fmt.Printf("int64 size: %d bytes\n", unsafe.Sizeof(int64(0)))
    fmt.Printf("int size: %d bytes\n", unsafe.Sizeof(int(0)))
    fmt.Printf("float32 size: %d bytes\n", unsafe.Sizeof(float32(0)))
    fmt.Printf("float64 size: %d bytes\n", unsafe.Sizeof(float64(0)))
    fmt.Printf("bool size: %d bytes\n", unsafe.Sizeof(bool(false)))
    fmt.Printf("string size: %d bytes\n", unsafe.Sizeof(string("")))
    fmt.Printf("pointer size: %d bytes\n", unsafe.Sizeof(&int(0)))
}
```


## Type Conversions in Go

Go is a statically typed language, so you often need to convert values from one type to another explicitly. This is called **type conversion** (not "type casting" as in some other languages).

### Basic Syntax

The general syntax for type conversion is:

```go
T(x)
```

where `T` is the target type and `x` is the value/expression.

### Numeric Conversions

```go
func numericConversions() {
    // Integer conversions
var i32 int32 = 1000
    i64 := int64(i32)     // widening: safe
    u8 := uint8(300)      // narrowing: wraps modulo 256 => 44
    
    // Float conversions
    f := float64(3)       // int to float: 3 -> 3.0
    i := int(3.9)         // float to int: truncates toward zero => 3
    
    // Between different integer types
    var x int = 42
    var y int32 = int32(x)
    var z uint64 = uint64(y)
    
    fmt.Printf("int32 to int64: %d -> %d\n", i32, i64)
    fmt.Printf("300 to uint8: %d -> %d\n", 300, u8)
    fmt.Printf("3.9 to int: %.1f -> %d\n", 3.9, i)
    fmt.Printf("int to int32 to uint64: %d -> %d -> %d\n", x, y, z)
}
```

### String Conversions

```go
func stringConversions() {
    // String to byte slice and rune slice
    s := "héllo"
    bytes := []byte(s)    // UTF-8 bytes
    runes := []rune(s)    // Unicode code points
    
    // Byte slice and rune slice to string
    s1 := string(bytes)   // from UTF-8 bytes
    s2 := string(runes)   // from runes
    
    fmt.Printf("Original string: %s\n", s)
    fmt.Printf("As bytes: %v\n", bytes)
    fmt.Printf("As runes: %v\n", runes)
    fmt.Printf("From bytes: %s\n", s1)
    fmt.Printf("From runes: %s\n", s2)
    
    // Individual character conversions
    char := 'A'
    byteVal := byte(char)
    runeVal := rune(char)
    
    fmt.Printf("Character 'A': %c\n", char)
    fmt.Printf("As byte: %d\n", byteVal)
    fmt.Printf("As rune: %d\n", runeVal)
}
```

### Boolean Conversions

```go
func booleanConversions() {
    // Go doesn't allow implicit numeric to boolean conversion
    // Use explicit comparisons instead
    
    n := 5
    b1 := n != 0          // true (non-zero)
    b2 := n > 0           // true (positive)
    
    // Boolean to string (using strconv)
    import "strconv"
    str := strconv.FormatBool(true)  // "true"
    
    // String to boolean
    boolVal, err := strconv.ParseBool("true")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Parsed boolean: %t\n", boolVal)
    }
    
    fmt.Printf("n != 0: %t\n", b1)
    fmt.Printf("n > 0: %t\n", b2)
    fmt.Printf("bool to string: %s\n", str)
}
```

### Custom Type Conversions

```go
// Define custom types
type MyInt int
type UserID int

// Type alias
type ID = int

func customTypeConversions() {
    // Custom defined types require explicit conversion
var m MyInt = 10
    var regularInt int = int(m)  // explicit conversion required
    
    // Type aliases don't require conversion
    var id ID = 42
    var regularInt2 int = id     // no conversion needed
    
    // Between different custom types
    var uid UserID = 123
    var mi MyInt = MyInt(uid)    // explicit conversion required
    
    fmt.Printf("MyInt to int: %d -> %d\n", m, regularInt)
    fmt.Printf("ID to int: %d -> %d\n", id, regularInt2)
    fmt.Printf("UserID to MyInt: %d -> %d\n", uid, mi)
}
```

### Interface Type Assertions

```go
func interfaceTypeAssertions() {
    var x interface{} = 42
    
    // Type assertion with ok
v, ok := x.(int)
if ok {
        fmt.Printf("x is int: %d\n", v)
    } else {
        fmt.Printf("x is not int\n")
}

    // Type assertion without ok (panics if wrong type)
    // v := x.(int)  // This would panic if x is not int
    
    // Type switch
switch v := x.(type) {
case int:
        fmt.Printf("x is int: %d\n", v)
case string:
        fmt.Printf("x is string: %s\n", v)
    case float64:
        fmt.Printf("x is float64: %f\n", v)
    default:
        fmt.Printf("x is unknown type: %T\n", v)
    }
    
    // Multiple type assertions
    var y interface{} = "hello"
    if str, ok := y.(string); ok {
        fmt.Printf("y is string: %s\n", str)
    }
}
```

### String Parsing and Formatting

```go
import "strconv"

func stringParsingAndFormatting() {
    // String to number conversions
    i, err := strconv.Atoi("123")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("String to int: %d\n", i)
    }
    
    // Number to string conversions
    s1 := strconv.Itoa(456)                    // "456"
    s2 := strconv.FormatInt(255, 16)           // "ff" (hex)
    s3 := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
    
    // Parse different number formats
    f64, err := strconv.ParseFloat("3.14", 64)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("String to float64: %f\n", f64)
    }
    
    // Parse with base
    i64, err := strconv.ParseInt("1010", 2, 64) // binary
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Binary '1010' to int: %d\n", i64)
    }
    
    fmt.Printf("Int to string: %s\n", s1)
    fmt.Printf("Int to hex: %s\n", s2)
    fmt.Printf("Float to string: %s\n", s3)
}
```

### Pointer Conversions

```go
func pointerConversions() {
    x := 42
    p := &x
    
    // Pointer to uintptr (unsafe)
    import "unsafe"
    addr := uintptr(unsafe.Pointer(p))
    
    // uintptr back to pointer (unsafe)
    p2 := (*int)(unsafe.Pointer(addr))
    
    fmt.Printf("Original value: %d\n", x)
    fmt.Printf("Pointer address: %p\n", p)
    fmt.Printf("Address as uintptr: %d\n", addr)
    fmt.Printf("Value through converted pointer: %d\n", *p2)
    
    // Note: unsafe.Pointer conversions are dangerous and should be used carefully
}
```

### Common Pitfalls and Best Practices

```go
func commonPitfalls() {
    // 1. Narrowing integer conversions can overflow
    var big int64 = 300
    var small int8 = int8(big)  // Wraps to 44
    fmt.Printf("300 as int8: %d\n", small)
    
    // 2. Float to int truncates (doesn't round)
    var f float64 = 3.9
    var i int = int(f)  // 3, not 4
    fmt.Printf("3.9 as int: %d\n", i)
    
    // 3. String to []byte creates a copy
    s := "hello"
    b := []byte(s)
    b[0] = 'H'  // Doesn't affect original string
    fmt.Printf("Original string: %s\n", s)
    fmt.Printf("Modified bytes: %s\n", string(b))
    
    // 4. Invalid UTF-8 in string conversion
    invalidBytes := []byte{0xff, 0xfe}
    invalidString := string(invalidBytes)
    fmt.Printf("Invalid UTF-8 string: %q\n", invalidString)
    
    // 5. Always check errors in parsing functions
    if val, err := strconv.Atoi("not-a-number"); err != nil {
        fmt.Printf("Parse error: %v\n", err)
    } else {
        fmt.Printf("Parsed value: %d\n", val)
    }
}
```

### Type Conversion Best Practices

1. **Always check errors** when using parsing functions
2. **Be aware of overflow** when converting to smaller integer types
3. **Use explicit conversions** for clarity and safety
4. **Prefer type assertions with ok** to avoid panics
5. **Use strconv package** for string↔number conversions
6. **Be careful with unsafe.Pointer** conversions
7. **Test edge cases** like overflow and invalid input


