package main

import "fmt"

func main() {

    // syntax
    // var <name> <type> = <value>


    // explicity type declaration
    // when to use:
    // when type is not obious from the value
    // In function parameters and return types
    // when you want to explicit about the type
    // 
    fmt.Println("Explicit Types: ")
    var name string = "vali"
    var age int = 30
    var isActive bool = true
    var height float64 = 5.7

    fmt.Printf("%v -> %T\n", name, name)
    fmt.Printf("%v -> %T\n", age, age)
    fmt.Printf("%v -> %T\n", isActive, isActive)
    fmt.Printf("%v -> %T\n", height, height)
    fmt.Println("")

    // Type inference
    // omit the type
    // when to use:
    // - when the type is obious from the value
    // - for cleaner and more readable code
    // - when you want to go to infer the type
    fmt.Println("Infered Types: ")
    var username = "nagulvali"
    var pinNumber = 12345
    var shoeSize = 10.5
    var isUserActive = true

   
    fmt.Printf("%v -> %T\n", username, username)
    fmt.Printf("%v -> %T\n", pinNumber, pinNumber)
    fmt.Printf("%v -> %T\n", shoeSize, shoeSize)
    fmt.Printf("%v -> %T\n", isUserActive, isUserActive)
    fmt.Println("")


    // Sort Declaration:
    // Use this in inside functions only
    // when to use:
    // - inside the function for local vars
    // - when you want to concise syntax
    // - for most local vars declarations
    fmt.Println("Short Declaration: ")
    avatarColor := "red"
    avatarAge := 3
    avatarXP := 50.6
    avatarActive := true

    fmt.Printf("%v -> %T\n", avatarColor, avatarColor)
    fmt.Printf("%v -> %T\n", avatarAge, avatarAge)
    fmt.Printf("%v -> %T\n", avatarXP, avatarXP)
    fmt.Printf("%v -> %T\n", avatarActive, avatarActive)
    fmt.Println("")


    // multiple vars declaration
    fmt.Println("Multiple declarations")
    var a, bee, c int = 1, 2, 3

    var (
        animalName string = "Crocodile"
        animalHealth int = 100
        animalXP = 570.50
        animalActive = true
    )

    // short multiple declarations

    x, y, z, w := "go", 23, 3.5, false

    fmt.Println(a, bee, c, x, y, z, w, animalName, animalHealth, animalXP, animalActive)
    fmt.Println("")




    //zero values
    fmt.Println("Zeroed Values: ")
    var (
        i int
        f float64
        s string
        b bool
        p *int
        slice []int
        m map[string]int
    )

    fmt.Printf("i -> %T -> %v\n", i, i)
    fmt.Printf("f -> %T -> %v\n", f, f)
    fmt.Printf("s -> %T -> %v\n", s, s)
    fmt.Printf("b -> %T -> %v\n", b, b)
    fmt.Printf("p -> %T -> %v\n", p, p)
    fmt.Printf("slice -> %T -> %v\n", slice, slice)
    fmt.Printf("m -> %T -> %v\n", m, m)
    fmt.Println("")

    
    //
    // Constants:
    // Contants value can not be changed, once initialized. they are immutable
    const (
        PI = 3.15
        E = 2.7
        G = 9.8
    )


}