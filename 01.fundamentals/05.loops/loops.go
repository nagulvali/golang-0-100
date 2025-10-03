package main

import "fmt"

// for -> only construct in go for looping

func main() {
	//while loop
	i :=1 
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}


	// infinite loops
	for {
		
		var input string
		fmt.Scan(&input)

		if input == "quit" {
			break
		}

		fmt.Printf("Entered input is %v \n", input)
		fmt.Println("Enter \"quit\" to quit the input")

	}


	// classic for loop
	for i := 0; i<=5; i++ {
		fmt.Println(i)
	}

	for i,j := 0, 5; i < j; i, j = i+1, j-1 {

		fmt.Println(i, j)

	}


	// range
	for i := range 11 {
		fmt.Println(i)
	}


}
