package main

import "fmt"

// a normal function can take single type
// what if we would like to pass sting or any other time instead int
func printSlice(items []int) {
	for _, value := range items {
		fmt.Println(value)
	}
}

// we can have generic type which can take any or optional types as input param for function
func printSlice1[T any](items []T) {
	for _, value := range items {
		fmt.Println(value)
	}
}


func main() {

	names := []string{"john", "karol", "jack"}
	printSlice1(names)

}
