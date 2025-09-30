# GoLang Data Types

Go has several built-in data types, which can be grouped as follows:

## Basic Data Types

| Type         | Description                        | Example                |
|--------------|------------------------------------|------------------------|
| `int`        | Integer (platform dependent size)  | `var x int = 10`       |
| `int8`       | 8-bit integer                      | `var y int8 = 127`     |
| `int16`      | 16-bit integer                     | `var z int16 = 32000`  |
| `int32`      | 32-bit integer                     | `var a int32 = 100000` |
| `int64`      | 64-bit integer                     | `var b int64 = 1e12`   |
| `uint`       | Unsigned integer                   | `var u uint = 42`      |
| `uint8`      | 8-bit unsigned integer (byte)      | `var b uint8 = 255`    |
| `uint16`     | 16-bit unsigned integer            | `var c uint16 = 65535` |
| `uint32`     | 32-bit unsigned integer            | `var d uint32 = 1e9`   |
| `uint64`     | 64-bit unsigned integer            | `var e uint64 = 1e18`  |
| `float32`    | 32-bit floating point              | `var f float32 = 3.14` |
| `float64`    | 64-bit floating point              | `var g float64 = 2.718`|
| `complex64`  | Complex numbers (float32 parts)    | `var c complex64 = 1+2i`|
| `complex128` | Complex numbers (float64 parts)    | `var d complex128 = 2+3i`|
| `byte`       | Alias for `uint8`                  | `var b byte = 255`     |
| `rune`       | Alias for `int32` (Unicode code point) | `var r rune = '♥'` |
| `bool`       | Boolean                            | `var ok bool = true`   |
| `string`     | String of characters               | `var s string = "Go!"` |

## Composite Data Types

- **Array**: Fixed-size sequence of elements of the same type  
  `var arr [3]int = [3]int{1, 2, 3}`

- **Slice**: Dynamically-sized, flexible view into arrays  
  `var s []int = []int{1, 2, 3}`

- **Map**: Key-value pairs  
  `var m map[string]int = map[string]int{"a": 1, "b": 2}`

- **Struct**: Collection of fields  
  ```
  type Person struct {
      Name string
      Age  int
  }
  ```

## Other Types

- **Pointer**: Holds the memory address of a value  
  `var p *int = &x`

- **Function**: Functions are first-class types  
  `var fn func(int) int`

- **Interface**: Specifies method signatures  
  `type Reader interface { Read(p []byte) (n int, err error) }`

- **Channel**: Used for communication between goroutines  
  `var ch chan int = make(chan int)`


> Go is a statically typed language, so the type of every variable is known at compile time.


---
## Type Conversions in Go

Go is a statically typed language, so you often need to convert (or "cast") values from one type to another explicitly. This is called **type conversion** (not "type casting" as in some other languages).

The general syntax is:

```go
T(x)
```

where `T` is the target type and `x` is the value/expression.

### Numeric conversions

- Converting between integer sizes/signs is explicit. Narrowing may overflow (wrap) and float→int truncates toward zero.

```go
var i32 int32 = 1000
i64 := int64(i32)     // widening: ok
u8  := uint8(300)     // wraps modulo 256 => 44
f   := float64(3)     // 3 -> 3.0
i   := int(3.9)       // truncates toward zero => 3
```

### Strings, bytes, and runes

- `string` ↔ `[]byte`: UTF-8 bytes
- `string` ↔ `[]rune`: Unicode code points

```go
s := "hé"
b := []byte(s)   // UTF-8 bytes
r := []rune(s)   // runes: 'h', 'é'

s1 := string(b)  // from UTF-8 bytes
s2 := string(r)  // from runes
```

### Booleans

- No implicit numeric↔bool conversions. Use comparisons.

```go
n := 5
b := n != 0
```

### Custom defined types vs aliases

```go
type MyInt int     // defined type
type ID = int      // alias

var m MyInt = 10
var n2 int = int(m)   // explicit conversion required
var id ID = 42         // alias: no conversion needed where int is expected
```

### Interfaces and type assertions

- Convert interface values to concrete types using assertions.

```go
var x any = 42
v, ok := x.(int)
if ok {
    _ = v // v is int
}

switch v := x.(type) {
case int:
    // use v as int
case string:
    // use v as string
}
```

### Parsing and formatting (strings)

- Use `strconv` for string↔number conversions.

```go
import "strconv"

i3, _ := strconv.Atoi("123")          // 123
s3 := strconv.FormatInt(255, 16)       // "ff"
f64, _ := strconv.ParseFloat("3.14", 64)
```

### Common pitfalls

- Narrowing integer conversions can overflow/wrap.
- Float→int conversions truncate (they do not round).
- `string([]byte{0xff})` may produce replacement runes when printed if invalid UTF-8.


