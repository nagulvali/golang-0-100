# Go Types
- Go is a statically typed language, which means that every variable has a specific type that is known at compile time. 
- This allows for better performance and early detection of type-related errors.
- Categories
  - Basic Types
    - Numbers
      - Integer Types
        - `int`, `int8`, `int16`, `int32`, `int64`
        - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
      - Floating-Point Types
        - `float32`, `float64`
      - Complex Types
        - `complex64`, `complex128`
    - Booleans
    - Strings
  - Aggregate Types
    - Array
    - Slice
    - Struct
  - Reference Types
    - Pointer
    - Function
    - Interface
    - Map
    - Channel
  - Interface Types
    - Empty Interface

    
## Basic Types
### Numbers
#### Integer Types
- In Go language, both signed and unsigned integers are available in four different sizes
- The signed int is represented by int and the unsigned integer is represented by uint.
- Signed Integers: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned Integers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- `int` and `uint` are platform-dependent types, meaning their size can vary based on the architecture (32-bit or 64-bit) of the system.
- `rune` is an alias for `int32` and is used to represent a Unicode code point.
- `uintptr` is an unsigned integer type that is large enough to hold the bit pattern of any pointer.
- Example:
```go
package main
import (
    "fmt"
)

func main() {
    var a int = -10
    var b uint = 20
    var c int8 = 30
    var d uint16 = 4000
    var e rune = 'A' // Unicode code point for 'A'

    fmt.Println("Signed Integer a:", a)
    fmt.Println("Unsigned Integer b:", b)
    fmt.Println("Signed Integer c (int8):", c)
    fmt.Println("Unsigned Integer d (uint16):", d)
    fmt.Println("Rune e:", e)
}
```


#### Floating-Point Types
- Go provides two floating-point types: `float32` and `float64`.
- `float32` is a single-precision floating-point type, while `float64` is a double-precision floating-point type.
- `float64` is the default type for floating-point literals in Go.
- Floating-point numbers are used to represent real numbers and can store fractional values.
- They are commonly used in scientific calculations, graphics, and other applications that require a high degree of precision.
- There are three literal styles for floating-point numbers in Go:
  - Decimal: e.g., `123.456`, `0.001`, `3.14e10`
  - Exponential: e.g., `13e4`, `5E8`
  - Mixed: e.g., `1.23e+4`, `0.001E2`
- Example:
```go
package main
import (
    "fmt"
)

func main() {
    var f1 float32 = 3.14
    var f2 float64 = 2.718281828459045

    fmt.Println("Float32 f1:", f1)
    fmt.Println("Float64 f2:", f2)
}
```


#### Complex Types
- Go provides two complex number types: `complex64` and `complex128`.
- `complex64` is a complex number type that uses `float32` for both its real and imaginary parts.
- `complex128` is a complex number type that uses `float64` for both its real and imaginary parts.
- Complex numbers are used in various fields such as engineering, physics, and signal processing.
- They are particularly useful in applications that involve oscillations, waves, and other phenomena that can be represented in the complex plane.
- Complex numbers can be created using the built-in `complex` function or by using the `i` suffix for the imaginary part.
- There are few built-in functions in complex numbers:
  - complex(real, imag) - creates a complex number from real and imaginary parts
  - real(c) - returns the real part of complex number c
  - imag(c) - returns the imaginary part of complex number c
- Example:
```go
package main
import (
    "fmt"
)

func main() {
    var c1 complex64 = complex(3, 4)       // Using complex function
    var c2 complex128 = 5 + 6i              // Using 'i' suffix

    fmt.Println("Complex Number c1:", c1)   // Output: (3+4i)
    fmt.Println("Complex Number c2:", c2)   // Output: (5+6i)

    fmt.Println("Real part of c1:", real(c1))   // Output: 3
    fmt.Println("Imaginary part of c1:", imag(c1)) // Output: 4
}
```

### Booleans
- The boolean type in Go is represented by the `bool` keyword.
- It can take one of two values: `true` or `false`.
- Booleans are commonly used in conditional statements and logical operations.
- The values of type boolean are not converted implicitly or explicitly to any other types.
- Example:
```go
package main
import (
    "fmt"
)

func main() {
    var str1 string = "Hello"
    var str2 string = "World"
    var str3 string = "Hello"

    fmt.Println("str1 == str2:", str1 == str2) // Output: false
    fmt.Println("str1 == str3:", str1 == str3) // Output: true
}
```

### Strings
- The string type in Go is represented by the `string` keyword.
- A string is a sequence of bytes that represents text.
- Strings in Go are immutable, meaning that once a string is created, its contents cannot be changed.
- Strings can be created using double quotes `""` for regular strings and backticks `` for raw string literals.
- Strings can be concatenated using plus(+) operator
- Example:
```go
package main
import (
    "fmt"
)
func main() {
    var str1 string = "Hello, "
    var str2 string = "World!"
    var str3 string = str1 + str2

    fmt.Println("Concatenated String:", str3) // Output: Hello, World!
}
```


