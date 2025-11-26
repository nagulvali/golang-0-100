package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = "hello world"
	// i = 20

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
