package main

import (
	"fmt"
	"time"
)

// Notes:
// struct in golang is a composite/custom type
// golang do not have classes to create objects, with structs we can create objects
// with structs we can store, any type key value pairs
// syntax:
// type <name> struct {
// 		name string
//    age int
// }

// order struct
type order struct {
	id int
	amount float32
	status string
	createdAt time.Time
}

// Notes:
// In object orienged programming, we do instantiating.
// In other programming languages we use constructors for instantiation
// In go we write a normal function with "new" prefix as naming convention

func newOrder(id int, amount float32, status string) *order {
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}

	// we return pointer
	return &myOrder
}


// we can add a method to struct 
// we need to pass struct as reciever to the function

// go convention: use pointer if you would like to modify the struct
// if you are just using struct parameter not modifying it just use reciver without pointer
func (o *order) changeStatus(status string) {
	o.status = status
}

// as we are just getting amount without modifying it we don't need to use pointer here.
func (o order) getAmount() (float32) {
	return o.amount
}



func main() {

	// we need to instantiate the struct to add values
	// syntax:
	var myOrder1 order = order{
		id: 123,
		amount: 50.00,
		status: "recieved",
		createdAt: time.Now(),
	}

	fmt.Println(myOrder1)

	// short syntax
	// if no value is set, it's default value would be zeroed value
	myOrder2 := order{
		id: 456,
		amount: 37.50,
		status: "inProgress",
	}

	fmt.Println(myOrder2)

	// we can add values after instantiation
	myOrder2.createdAt = time.Now()

	fmt.Println(myOrder2)



	myOrder1.changeStatus("confirmed")
	fmt.Println(myOrder1.status)
	fmt.Println(myOrder1.getAmount())



	// create a new order using contsructor function
	// we don't need to de-reference the pointer with structs, 
	// struct automatically takes care of it similar to method recievers
	myOrder3 := newOrder(789, 78.9, "shipping")
	fmt.Println(myOrder3.getAmount())


	// Create nameless structs
	// If we are using struct only one time, we will not use it further
	// then we can create a struct

	language := struct{
		name string
		isGood bool
	} {
			"golang",
			true,
	}

	fmt.Println(language)
	
}