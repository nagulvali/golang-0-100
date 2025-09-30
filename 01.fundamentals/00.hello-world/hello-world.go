package main

import "fmt"

// Single Line Comment

/*
COMMENT:
Multi Line Comment

Fmt package: https://pkg.go.dev/fmt@go1.25.1
Escape Chars: https://go.dev/ref/spec
*/


func main() {
    fmt.Println("Hello World ! 😊 ✨")

    // String Interpolation
    var name = "vali"
    var age = 30
    fmt.Printf("%s is %d years old\n", name, age)

    // Raw String Litteral
    fmt.Println(`
        This is formatted raw string litteral
        This is printed as is written with spaces &
        Tabs, etc.
    `)
}