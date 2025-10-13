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


// Generics in struct
type Stack 



func main() {

	nums := []int{1, 2, 3, 4, 9, 8, 7}
	names := []string{"john", "karol", "jack"}
	printSlice1(nums)
	printSlice1(names)

}
