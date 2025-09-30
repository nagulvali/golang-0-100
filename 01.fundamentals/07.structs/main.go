package main

import "fmt"

type User struct {
    Name string
    Age int
}

func (u User) Greet() {

    fmt.Printf("Hello, my name is %s and i am %d years old\n", u.Name, u.Age)

}

func main() {
    user := User{ Name: "vali", Age: 29 }
    user.Greet()
}