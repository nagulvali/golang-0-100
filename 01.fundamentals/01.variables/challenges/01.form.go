package main

import "fmt"

func main() {

	var (
		username string
		age int
		height float64
		weight int
	)

	fmt.Println("Enter Username: ")
	fmt.Scan(&username)
	fmt.Println("Enter Age: ")
	fmt.Scan(&age)
	fmt.Println("Enter Height: ")
	fmt.Scan(&height)
	fmt.Println("Enter Weight: ")
	fmt.Scan(&weight)

	fmt.Println("=====================")
	fmt.Println("User Details:")
	fmt.Println("Username: ", username)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
	fmt.Println("Weight: ", weight)
	fmt.Println("---------------------")
	fmt.Println("")

}