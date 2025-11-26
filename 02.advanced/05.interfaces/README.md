# Interfaces
- In Go, an interface is a type that lists methods without providing their code.
- You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.

- Syntax:
```golang

type InterfaceName interface {
    Method1() returnType
    Method2() returnType
}
```

- Example:
```golang
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

- In this example, Shape is an interface that requires any type(custom type with any name) implementing it to have Area and Perimeter methods.

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

> Queries ?
> in above example i am not able to use pointer reciver for function methods why ?


## Dynamic Values
- An interface can hold any value
- The value and its type are stored dynamically.

```golang
package main

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

## Type Assertion
- Type assertion allows extracting the underlying type of an interface.

- Syntax:
```golang
value := interfaceVariable.(ConcreteType)
```

- Example:
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

### Why do we need type assertions
- Interfaces Hide Concrete Types:
	- When you use interfaces, you lose information about the concrete type stored inside. 
	- Type assertion lets you recover that information.

- Perform Type-Specific Operations:
	- Sometimes, you need to use methods or properties specific to the underlying type, which are not available on the interface itself.

### Common scenarios for using type assertion
#### Working with generic data structure
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

#### Handling Values from Channels or Maps of Interface Type
- When you receive values from a channel or map of type interface{}, you often need to assert their actual type to use them.

```golang
var ch chan interface{}
value := <-ch
strValue, ok := value.(string)
if ok {
    // Use strValue as a string
}
```

#### Implementing Functions That Accept Any Type
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

#### Accessing Methods Not in the Interface
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


## Type Switch
- A type switch can be used to compare the dynamic type of an interface against multiple types.
- Syntax:
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
- Example:

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
