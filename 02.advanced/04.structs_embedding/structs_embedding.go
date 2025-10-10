package main

import "fmt"

// Notes:
// In other programming languages we use objects inheritance


type person struct {
	name string
	age int
	height float32
}


type user struct {
	person
	autoAvatar bool
}



func main() {

	// person1 := person{
	// 	name: "john",
	// 	age: 52,
	// 	height: 6.3,
	// }

	user1 := user{
		autoAvatar: true,
		// person: person1, // this can be added inline as well
		person: person{
			name: "catherine",
			age: 27,
			height: 5.8,
		},

	}

	fmt.Println(user1)
	fmt.Println(user1.person.name)
}



