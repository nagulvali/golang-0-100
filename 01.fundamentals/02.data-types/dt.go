package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Data Types:")

	// Booleans
	var isActive = true
	fmt.Println(isActive)

	// Integers
	var age = 25
	fmt.Println(age)

	// Floats
	var pi = 3.14
	fmt.Println(pi)

	// Strings
	var name = "John"
	fmt.Println(name)

	// Arrays
	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

	// Slices
	var fruits = []string{"apple", "banana", "cherry"}
	fmt.Println(fruits)

	// Maps
	var person = map[string]int{"age": 25}
	fmt.Println(person)

	// Structs
	var user = struct {
		name string
		age int
	}{name: "John", age: 25}
	fmt.Println(user)

	// Pointers
	var ptr *int
	fmt.Println(ptr)

	// Functions
	var fn = func() {
		fmt.Println("Hello, World!")
	}
	fmt.Println(fn)

	// Interfaces
	var iface interface{} = 42
	fmt.Println(iface)

	// Channels
	var ch chan int
	fmt.Println(ch)

	// Errors
	var err error = errors.New("error")
	fmt.Println(err)

}