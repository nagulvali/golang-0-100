# Type Conveersions in Go
- Type conversion is the process of converting a value from one data type to another.
- Statically typed languages like C/C++, Java, provide the support for Implicit Type Conversion but Golang is different, as it doesn't support the Automatic Type Conversion or Implicit Type Conversion even if the data types are compatible.
  - The reason for this is the Strong Type System of the Golang which doesn't allow to do this.
- In Go, type conversion is done explicitly using the syntax: `T(v)` where `T` is the target type and `v` is the value to be converted.

## What is the need for Type Conversion?
- Well, if you need to take advantage of certain characteristics of data type hierarchies, then we have to change entities from one data type into another. 
- Example: You have a variable of type `int` and you want to perform floating-point division, you need to convert the `int` to `float64` first.
```go
package main

import (
    "fmt"
)

func main() {
    var a int = 5
    var b int = 2

    // Converting int to float64 for floating-point division
    var result float64 = float64(a) / float64(b)

    fmt.Println("Result of division:", result) // Output: 2.5
}
```

## Type conversion between numbers
- You can convert between different numeric types in Go using explicit type conversion.
- Table:
| From Type | To Type    | Example Conversion         |
|-----------|------------|----------------------------|
| int       | float64    | `float64(i)`               |
| float64   | int        | `int(f)`                   |
| int8      | int16      | `int16(i8)`                |
| uint8     | int        | `int(u8)`                  |
| float32   | float64    | `float64(f32)`             |
| complex64 | complex128 | `complex128(c64)`          |


## Type conversion from number to other types
- You can convert numeric types to strings and vice versa using the `strconv` package.
- Example:
```go
package main
import (
    "fmt"
    "strconv"
)
func main() {
    // Converting int to string
    var num int = 42
    str := strconv.Itoa(num)
    fmt.Println("String representation:", str) // Output: "42"

    // Converting string to int
    strNum := "123"
    intNum, err := strconv.Atoi(strNum)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
    } else {
        fmt.Println("Integer representation:", intNum) // Output: 123
    }
}
```

## Type conversion boolean to others
- In Go, boolean values can be converted to integers and strings.
- Example:
```go
package main
import (
    "fmt"
    "strconv"
)
func main() {
	// Converting boolean to int
    var b bool = true
    var intValue int
    if b {
        intValue = 1
    } else {
        intValue = 0
    }
    fmt.Println("Integer representation of boolean:", intValue) // Output: 1

    // Converting boolean to string
    strValue := strconv.FormatBool(b)
    fmt.Println("String representation of boolean:", strValue) // Output: "true"
}
```

## type conversion from numbers and string to boolean
- In Go, you can convert strings to boolean using the `strconv` package.
- Example:
```go
package main
import (
    "fmt"
    "strconv"
)
func main() {
	    // Converting string to boolean
    strTrue := "true"
    boolValue, err := strconv.ParseBool(strTrue)
    if err != nil {
        fmt.Println("Error converting string to boolean:", err)
    } else {
        fmt.Println("Boolean representation of string:", boolValue) // Output: true
    }

    strFalse := "false"
    boolValue, err = strconv.ParseBool(strFalse)
    if err != nil {
        fmt.Println("Error converting string to boolean:", err)
    } else {
        fmt.Println("Boolean representation of string:", boolValue) // Output: false
    }
}
```
- Note: Converting numbers directly to boolean is not supported in Go. You need to implement your own logic to handle such conversions based on your requirements.
- Example:
```go
package main
import (
    "fmt"
)
func main() {
    // Converting int to boolean
    var num int = 1
    var boolValue bool
    if num != 0 {
        boolValue = true
    } else {
        boolValue = false
    }
    fmt.Println("Boolean representation of int:", boolValue) // Output: true
}
```