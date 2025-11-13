# Golang variables
- Variables in Golang are used to store data values. 
- Each variable has a specific type that determines the kind of data it can hold.
- To declare a variable, you can use the `var` keyword followed by the variable name and its type.
```golang
var age int
var name string
```

- you can declare first and initialize later:
```golang
var x, y, z int
x = 1
y = 2
z = 3
```

- You can also declare and initialize a variable in a single line:
```go
var age int = 25
var name string = "John"
```

- Go also supports type inference, allowing you to declare a variable without specifying its type explicitly
- Use this short declaration syntax `:=` to declare and initialize a variable
- Short declaration is often generally used in golang unless there is a specific need to use `var` keyword
  - Use var if you want to declare a variable at the package level
  - Use var if you want to declare a variable without initializing it immediately
  - Use var if you want to declare multiple variables of different types in a single statement
```golang
age := 25
name := "John"
```

- You can declare multiple variables in a single line:
```golang
var x, y, z int = 1, 2, 3
```


## Zero values
- When a variable is declared without an explicit initial value, it is assigned a default value known as the "zero value" for its type.
- The zero values for some common types are:
```go
var i int        // zero value is 0
var f float64    // zero value is 0.0
var b bool       // zero value is false
var s string     // zero value is ""    
```
- Below is the zero values for types
  - `false` for booleans
  - `0` for numeric types
  - `""` (empty string) for strings
  - `nil` for pointers, functions, interfaces, slices, channels, and maps


## Blank Identifier
- The blank identifier `_` is used to ignore values in Go.
- It can be used when you want to declare a variable but do not intend to use it.
- For example:
```go
 _ := 10  // Ignoring the value 10
```

## Using Printf for decimal & hexadecimal values
- You can use `fmt.Printf` to format and print variables in different formats.
- For decimal format, use `%d` and for hexadecimal format, use `%x`
```golang
package main
import "fmt"

func main() {
    num := 255
    fmt.Printf("Decimal: %d\n", num)      // Output: Decimal: 255
    fmt.Printf("Hexadecimal: %x\n", num)   // Output: Hexadecimal: ff
}
```
