package main

import "fmt"

// Notes:
// Go follow pass by value to the functions
// which means in function parameters a disctict value copy passed to the function as parameter
// that is the reason below code will not update the num variable value

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changeNum: ", num)
// }

// func main() {
// 	num := 1

// 	changeNum(num)

// 	fmt.Println("After changeNum: ", num)

// }


func changeNum(num *int) {

	*num += 5


	fmt.Println("In changeNum: ", *num)
	
}


func main() {

	num := 1

	fmt.Println("Initial value: ", num)

	changeNum(&num)

	fmt.Println("After changeNum: ", num)
}


