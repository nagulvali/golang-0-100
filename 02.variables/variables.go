package main

import "fmt"

func main() {
	// Variable declaration
	var name string
	var age int

	// If variables are not initialized, they hold the zero value of their type
	fmt.Println(name, age) // Output: "" 0

	// Initializing variables
	name = "Alice"
	age = 30

	fmt.Println(name, age)

	// Short variable declaration
	city := "New York"
	salary := 75000.50

	fmt.Println(city, salary)

	// Blank identifier
	// Even if blank identifier is not used, the code compiles without error
	_, height := "Bob", 180
	fmt.Println(height) // Output: 180
}
