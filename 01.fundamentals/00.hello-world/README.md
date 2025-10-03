# Hello World in Go

Welcome to your first Go program! This tutorial will walk you through the classic "Hello World" program and explain the fundamental concepts of Go programming.

## What You'll Learn

- Basic Go program structure
- Package declarations
- Import statements
- Functions and the main function
- Comments in Go
- String formatting and interpolation
- Raw string literals

## Program Overview

Let's examine the `hello-world.go` file:

```go
package main

import "fmt"

// Single Line Comment

/*
COMMENT:
Multi Line Comment

Fmt package: https://pkg.go.dev/fmt@go1.25.1
Escape Chars: https://go.dev/ref/spec
*/

func main() {
    fmt.Println("Hello World ! 😊 ✨")

    // String Interpolation
    var name = "vali"
    var age = 30
    fmt.Printf("%s is %d years old\n", name, age)

    // Raw String Litteral
    fmt.Println(`
        This is formatted raw string litteral
        This is printed as is written with spaces &
        Tabs, etc.
    `)
}
```

## Breaking Down the Code

### 1. Package Declaration

```go
package main
```

**What it means:**
- Every Go file must start with a `package` declaration
- `main` is a special package name that tells Go this is an executable program
- Other packages are typically named after their directory or purpose

**Why it matters:**
- Go programs are organized into packages
- The `main` package is required for executable programs
- It contains the entry point (`main` function)

### 2. Import Statement

```go
import "fmt"
```

**What it means:**
- `import` brings in external packages
- `"fmt"` is Go's formatting package (similar to `printf` in C)
- It provides functions for formatted I/O

**Common fmt functions:**
- `fmt.Println()` - prints with a newline
- `fmt.Printf()` - formatted printing
- `fmt.Print()` - prints without newline

### 3. Comments

```go
// Single Line Comment

/*
Multi Line Comment
Can span multiple lines
*/
```

**Types of comments:**
- `//` - Single line comments
- `/* */` - Multi-line comments
- Comments are ignored by the compiler
- Use comments to explain your code

**Best practices:**
- Write comments that explain "why", not "what"
- Keep comments up to date with code changes
- Use clear, concise language

### 4. The Main Function

```go
func main() {
    // Your code goes here
}
```

**What it means:**
- `func` declares a function
- `main` is the special function name that serves as the program entry point
- `()` indicates no parameters
- `{}` contains the function body

**Key points:**
- Every executable Go program must have a `main` function
- Execution starts from the `main` function
- Only one `main` function per program

### 5. Printing to Console

```go
fmt.Println("Hello World ! 😊 ✨")
```

**What it does:**
- Prints the string to the console
- Automatically adds a newline at the end
- Supports Unicode characters (emojis, special characters)

**Output:**
```
Hello World ! 😊 ✨
```

### 6. Variables and String Interpolation

```go
var name = "vali"
var age = 30
fmt.Printf("%s is %d years old\n", name, age)
```

**Variable declaration:**
- `var` keyword declares a variable
- `name` and `age` are variable names
- Go infers the type from the assigned value
- `name` is a string, `age` is an integer

**String formatting:**
- `%s` - placeholder for strings
- `%d` - placeholder for integers
- `\n` - newline character
- Arguments are passed after the format string

**Output:**
```
vali is 30 years old
```

### 7. Raw String Literals

```go
fmt.Println(`
    This is formatted raw string litteral
    This is printed as is written with spaces &
    Tabs, etc.
`)
```

**What it means:**
- Backticks (`` ` ``) create raw string literals
- Preserves all formatting, spaces, and tabs
- No escape sequences are processed
- Useful for multi-line strings

**Output:**
```
    This is formatted raw string litteral
    This is printed as is written with spaces &
    Tabs, etc.
```

## Running the Program

### Method 1: Using `go run`

```bash
go run hello-world.go
```

**What happens:**
- Compiles and runs the program in one step
- No executable file is created
- Good for testing and development

### Method 2: Using `go build` and `go run`

```bash
# Compile the program
go build hello-world.go

# Run the executable
./hello-world        # On Unix/Linux/macOS
hello-world.exe      # On Windows
```

**What happens:**
- `go build` creates an executable file
- The executable can be run directly
- Good for distribution and production

## Expected Output

When you run the program, you should see:

```
Hello World ! 😊 ✨
vali is 30 years old

    This is formatted raw string litteral
    This is printed as is written with spaces &
    Tabs, etc.
```

## Common Variations

### 1. Different Ways to Declare Variables

```go
// Method 1: var with type
var name string = "Go"

// Method 2: var with type inference
var name = "Go"

// Method 3: Short declaration (inside functions only)
name := "Go"

// Method 4: Multiple variables
var (
    name string = "Go"
    age  int    = 13
)
```

### 2. Different Print Functions

```go
// Print with newline
fmt.Println("Hello", "World")  // Hello World

// Print without newline
fmt.Print("Hello", "World")    // HelloWorld

// Formatted print
fmt.Printf("Hello %s\n", "World")  // Hello World

// Print to string
result := fmt.Sprintf("Hello %s", "World")
```

### 3. Escape Sequences

```go
fmt.Println("Hello\nWorld")     // Newline
fmt.Println("Hello\tWorld")     // Tab
fmt.Println("Hello\"World\"")   // Quotes
fmt.Println("Hello\\World")     // Backslash
```

## Troubleshooting

### Common Errors

1. **Missing package declaration:**
   ```
   Error: expected 'package', found 'import'
   ```
   **Solution:** Add `package main` at the top

2. **Missing import:**
   ```
   Error: undefined: fmt
   ```
   **Solution:** Add `import "fmt"`

3. **Missing main function:**
   ```
   Error: runtime.main_main·f: function main is undeclared in the main package
   ```
   **Solution:** Add a `main` function

4. **Wrong file extension:**
   ```
   Error: cannot find package
   ```
   **Solution:** Use `.go` extension for Go files

### File Structure

Make sure your file structure looks like this:

```
01.fundamentals/
└── 00.hello-world/
    ├── README.md
    └── hello-world.go
```

## Next Steps

Now that you've written your first Go program, you're ready to explore:

1. **[Variables](../01.variables)** - Learn about different data types and variable declarations
2. **[Data Types](../02.data-types)** - Understand Go's type system
3. **[Operators](../03.operators)** - Learn about arithmetic, logical, and other operators
4. **[Conditional Statements](../04.conditional-statements)** - Control program flow with if/else and switch
5. **[Loops](../05.loops)** - Repeat code with for loops and range

## Key Takeaways

- Every Go program needs a `package main` declaration
- Import statements bring in external functionality
- The `main` function is the program entry point
- `fmt` package provides formatted I/O functions
- Go supports both single-line and multi-line comments
- Variables can be declared with `var` or short declaration syntax
- String formatting uses placeholders like `%s` and `%d`
- Raw string literals preserve formatting exactly

## Practice Exercises

1. **Modify the program** to print your own name and age
2. **Add more variables** and print them using different formatting
3. **Try different escape sequences** in your strings
4. **Create a raw string literal** with your own multi-line text
5. **Experiment with different print functions** (`Print`, `Println`, `Printf`)

## Additional Resources

- [Go Tour - Hello World](https://tour.golang.org/welcome/1)
- [fmt package documentation](https://pkg.go.dev/fmt)
- [Go by Example - Hello World](https://gobyexample.com/hello-world)
- [Effective Go - Comments](https://golang.org/doc/effective_go.html#commentary)

Congratulations! You've written your first Go program. Welcome to the world of Go programming! 🎉
