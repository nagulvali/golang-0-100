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
