# Interfaces
- An interface is like a contract or a blueprint in programming. It tells you what methods (functions) a class must have, but not how those methods should work.
- An interface in Go is a type that specifies a set of method signatures (behavior), but does not implement them. 
- Any type (struct or custom type) that implements those methods automatically satisfies the interface.
- Any type that provides implementations for all the methods in an interface is said to "implement" that interface, implicitly (no explicit declaration needed).

> **Analogy:** </br>
Imagine a remote control. Every TV remote has buttons like Power, Volume Up, Volume Down, etc. The interface is the list of these buttons. How each TV reacts when you press "Power" is up to the TV (the class), but every TV must have those buttons if it claims to support the remote (the interface).

- Interfaces define what should be done, not how.
- Classes that use (implement) an interface must provide the actual code for the methods in the interface.

## Interfaces in go
- In Go, an interface is a type that lists methods without providing their code.
- You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.
- Interface statisfaction is implicit in go
    - You don’t need to declare that a type implements an interface. If the methods match, it’s automatic.


## Implementation
- In Go, to implement an interface, a type must define all methods declared by the interface.
- Implementation is implicit, meaning no keyword (such as implements) is needed.

```golang
package main

import (
	"fmt"
	"math"
)

// This is like an abstract defines which methods object should contain to satisfy this type
type Shape interface {
	Area() float64
	Perimeter() float64
}

// structs are like classes/blueprints of objects
// Circle type
// Circle type has both the method signatures Area() and Perimeter so it satisfies with the Shape interface
// so shape type varible can hold and implement Circle
type Circle struct {
	radius int
}

// function with struct reciver are methods of that struct/object
func (c Circle) Area() float64 {
	return float64(math.Pi) * float64(c.radius*c.radius)
}

func (c Circle) Perimeter() float64 {
	// return permiter
	return 2 * math.Pi * float64(c.radius)
}

// Rectable type
// Rectangle custom type has both the method signatures Area() and Permiter()
// so Interface shape can hold Rectangle value as well
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() float64 {
	return float64(r.length) * float64(r.width)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * float64(r.length*r.width)
}

func main() {

	// shape type interface can hold and implement both Circle and Rectangle
	// with this pattern we can implement polymorphism like this.
	var s Shape

	s = Circle{radius: 4}
	fmt.Println("C Area: ", s.Area())
	fmt.Println("C Perimeter: ", s.Perimeter())

	s = Rectangle{length: 20, width: 20}
	fmt.Println("R Area: ", s.Area())
	fmt.Println("R Perimeter: ", s.Perimeter())

}


```

## Syntax & Example
- Syntax
```golang
type InterfaceName interface {
    Method1(param1 type1, ...) returnType1
    Method2(param2 type2, ...) returnType2
    // ... more methods
}
```
- Basic Example
```golang
type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}
func (c Cat) Speak() string {
    return "Meow!"
}

type Speaker interface {
    Speak() string
}

func MakeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    d := Dog{}
    c := Cat{}
    MakeItSpeak(d) // Output: Woof!
    MakeItSpeak(c) // Output: Meow!
}
```
- Both Dog and Cat implement the Speak() method.
- Both are now considered Speakers.
- The function MakeItSpeak can accept any Speaker.
- The interface Speaker satisfies both Dog and Cat types, so those values can be stored in the interface and can call the only interface method


## Purpose
- Abstraction: Hide implementation details, expose only behavior.
- Polymorphism: Write code that works with different types in a uniform way.
- Decoupling: Reduce dependencies between components.
- Testing: Easily mock dependencies for unit tests.

### Abstraction
- Hide implementation details, expose only behavior.
```golang
package main

import "fmt"

// Interface defines the contract
type PaymentMethod interface {
    Pay(amount float64) string
}

// Concrete implementation 1
type CreditCard struct {
    CardNumber string
}
func (c CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid $%.2f using Credit Card %s", amount, c.CardNumber)
}

// Concrete implementation 2
type PayPal struct {
    Email string
}
func (p PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, p.Email)
}

// Function uses abstraction
func Checkout(pm PaymentMethod, amount float64) {
    fmt.Println(pm.Pay(amount))
}

func main() {
    cc := CreditCard{CardNumber: "1234-5678"}
    pp := PayPal{Email: "user@example.com"}
    Checkout(cc, 100.0)
    Checkout(pp, 200.0)
}
```
- Checkout doesn’t care how payment is made, just that it can call Pay.
- You can add new payment methods without changing Checkout.

### Polymorphism
- Write code that works with different types in a uniform way.
```golang
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
    PrintArea(c) // Area: 78.50
    PrintArea(r) // Area: 24.00
}

```
- PrintArea works with any Shape.
- You can add more shapes (Triangle, Square, etc.) without changing PrintArea.

### Decoupling
- Reduce dependencies between components.
```golang
package main

import "fmt"

// Logger interface
type Logger interface {
    Log(message string)
}

// ConsoleLogger implementation
type ConsoleLogger struct{}
func (c ConsoleLogger) Log(message string) {
    fmt.Println("Console:", message)
}

// FileLogger implementation (simulated)
type FileLogger struct{}
func (f FileLogger) Log(message string) {
    fmt.Println("File:", message)
}

// Service depends on Logger, not on a specific implementation
type Service struct {
    logger Logger
}

func (s *Service) DoSomething() {
    s.logger.Log("Doing something important!")
}

func main() {
    cl := ConsoleLogger{}
    fl := FileLogger{}
    s1 := Service{logger: cl}
    s2 := Service{logger: fl}
    s1.DoSomething() // Console: Doing something important!
    s2.DoSomething() // File: Doing something important!
}

```
- Service can use any logger.
- You can swap loggers without changing Service.

### Testing (Mocking)
- Easily mock dependencies for unit tests.
```golang

package main

import (
    "fmt"
    "errors"
)

// Database interface
type Database interface {
    Save(data string) error
}

// Real implementation
type RealDB struct{}
func (r RealDB) Save(data string) error {
    fmt.Println("Saving to real database:", data)
    return nil
}

// Mock implementation for testing
type MockDB struct{}
func (m MockDB) Save(data string) error {
    if data == "" {
        return errors.New("no data")
    }
    fmt.Println("Pretend saving:", data)
    return nil
}

// Service uses Database interface
type Service struct {
    db Database
}

func (s *Service) Store(data string) error {
    return s.db.Save(data)
}

func main() {
    // In production
    realService := Service{db: RealDB{}}
    realService.Store("production data")

    // In testing
    mockService := Service{db: MockDB{}}
    err := mockService.Store("")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```
- Service can use a real or mock database.
- In tests, you can use MockDB to simulate errors or behaviors.


## Use Cases

### Generic/Empty Interface/Dynamic Values
- An empty interface satisfies all the types in the go
- So it can hold any type of value.


```golang
func PrintAnything(val interface{}) {
    fmt.Println(val)
}

func main() {
    PrintAnything(123)
    PrintAnything("hello")
    PrintAnything([]int{1, 2, 3})
}
```

- We can store dynamic values in it
```golang
import (
    "fmt"
)

func main() {
    var i interface{} // 'i' is an empty interface, can hold any value

    i = 42
    fmt.Printf("i holds value: %v, type: %T\n", i, i)

    i = "hello"
    fmt.Printf("i holds value: %v, type: %T\n", i, i)

    i = 3.14
    fmt.Printf("i holds value: %v, type: %T\n", i, i)
}
```






### Interface Embedding (Composition)
- Interfaces can include other interfaces, allowing you to build complex behaviors from simple ones.
```golang

type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}
```
- TODO: add complex example for this

### Nil Interfaces and Interface Values
- An interface value is nil only if both the type and value are nil.
- A non-nil interface with a nil underlying value can cause bugs.

```golang
var r Reader // nil interface
fmt.Println(r == nil) // true

var r2 Reader = (*MyReader)(nil)
fmt.Println(r2 == nil) // false (type is set, value is nil)
```

### Type Assertion
- Type assertion allows extracting the underlying concrete type of an interface.

```golang
var s Speaker = Dog{}
dog, ok := s.(Dog)
if ok {
    fmt.Println("It's a Dog:", dog.Speak())
}
```

```golang
package main

import (
    "fmt"
)

func main() {
    var i interface{}
    i = "hello world"

    // Type assertion to extract the string value
    s, ok := i.(string)
    if ok {
        fmt.Printf("Extracted string: %s\n", s)
    } else {
        fmt.Println("i does not hold a string")
    }

    // Type assertion to extract an int value (will fail)
    n, ok := i.(int)
    if ok {
        fmt.Printf("Extracted int: %d\n", n)
    } else {
        fmt.Println("i does not hold an int")
    }
}

```

#### Why do we need type assertions
- Interfaces Hide Concrete Types:
	- When you use interfaces, you lose information about the concrete type stored inside. 
	- Type assertion lets you recover that information.

- Perform Type-Specific Operations:
	- Sometimes, you need to use methods or properties specific to the underlying type, which are not available on the interface itself.


#### Common scenarios for using type assertion
##### Working with generic data structure
- Suppose you have a function that accepts interface{} to allow any type, but you need to process the value differently based on its actual type.
```golang
func printValue(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

##### Handling Values from Channels or Maps of Interface Type
- When you receive values from a channel or map of type interface{}, you often need to assert their actual type to use them.

```golang
var ch chan interface{}
value := <-ch
strValue, ok := value.(string)
if ok {
    // Use strValue as a string
}
```

##### Implementing Functions That Accept Any Type
- For example, logging, serialization, or error handling functions often accept interface{} and then use type assertion to handle different types appropriately.
```golang
func logValue(i interface{}) {
    if err, ok := i.(error); ok {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Value:", i)
    }
}
```

##### Accessing Methods Not in the Interface
- If you have an interface variable but need to call a method that’s not part of the interface, you can use type assertion to access the concrete type.

```golang
type Animal interface {
    Speak()
}

type Dog struct{}
func (d Dog) Speak() { fmt.Println("Woof!") }
func (d Dog) Fetch() { fmt.Println("Fetching...") }

var a Animal = Dog{}
if dog, ok := a.(Dog); ok {
    dog.Fetch() // Access Dog-specific method
}
```


### Type switch 
- Handle multiple possible types stored in an interface.
```golang
func Describe(i interface{}) {
    switch v := i.(type) {
    case Dog:
        fmt.Println("Dog:", v.Speak())
    case Cat:
        fmt.Println("Cat:", v.Speak())
    default:
        fmt.Println("Unknown type")
    }
}
```
- A type switch can be used to compare the dynamic type of an interface against multiple types.

```golang
switch variable.(type) {
case type1:
    // Code for type1
case type2:
    // Code for type2
default:
    // Code if no types match
}

```

```golang
package main

import (
	"fmt"
	"math"
)
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Function to determine the type of shape and calculate area
func calculateArea(shape interface{}) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle area: %.2f\n", s.area())
	case Rectangle:
		fmt.Printf("Rectangle area: %.2f\n", s.area())
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{length: 4, width: 3}

	calculateArea(c)
	calculateArea(r)
}
```

## Interfaces implemented by pointer reciever
- Interfaces in Go specify method sets.
	- Any type that implements all the methods in an interface (with the correct signatures) satisfies that interface.
- Pointer receiver methods are methods with a receiver of the form `func (p *Type) Method()`.
- If an interface requires a method, and a type implements that method with a pointer receiver, only the pointer to that type implements the interface—not the value type.

```golang
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d *Dog) Speak() {
    fmt.Println("Woof!")
}

func main() {
    var s Speaker
    d := Dog{}
    // s = d      // ERROR: Dog does not implement Speaker (Speak method has pointer receiver)
    s = &d        // OK: *Dog implements Speaker
    s.Speak()     // Output: Woof!
}
```

## Interface Reflection
- You can use the reflect package to inspect interfaces at runtime (advanced).

```golang
import "reflect"

func PrintType(i interface{}) {
    fmt.Println(reflect.TypeOf(i))
}
```


## Best practices
- **Accept Interfaces, return structs:** Functions should accept interfaces for flexibility, return concrete types for clarity.
- **Small interfaces are better:** Prefer interfaces with one or two methods (e.g., io.Reader).
- **Don’t overuse interfaces:** Use interfaces only when you need abstraction or multiple implementations.

## Common pitfalls
- **Nil interface vs. interface holding nil value:** Can cause unexpected behavior in error handling.
- **Interface Pollution:** Don’t create interfaces just for the sake of it.
- **Interface Segregation:** Keep interfaces small and focused.