package main

import (
	"fmt"
	"math"
)

func main () {

	// Numeric types:

	// signed integers
	var a int8 = 127 											// -128 to 127
	// var a int8 = 128 									// Error: 02.data-types/challenges/main.go:9:15: cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)
	var b int16 = 32000   								// -32768 to 32767
	var c int32 = 2147483647							// -2,14,74,83,648 to 2,14,74,83,647
	var d int64 = 9223372036854775807   	// - 9.2233720369E18 to (9.2233720369E18 - 1)
	var e int = 9223372036854775807				// Platform dependent (64 arch is 64 bit)

	fmt.Println(a, b, c, d, e)


	// Un-signed integers
	var f uint8 = 255										// 0 to 255 (also called byte)
	var g uint16 = 65535 								// 0 to 65535
	var h uint32 = 4000000000						// 0 to 4,294,967,295
	var i uint64 = 18000000000000000000 // 0 to 18e18
	var j uint = 18000000000000000000		// platform dependent (current arch is 64 so uint = uint64)

	fmt.Println(f, g, h, i, j)


	// Floating point numbers
	var f32 float32 = 3.14159
	var f64 float64 = 2.7937138261873918043793719431937971

	// scientific
	var scientific float64 = 1.23e-50

	// special values
	var posInf float64 = math.Inf(1)		// +Infinity
	var negInf float64 = math.Inf(-1)		// -Infinity
	var nan float64	= math.NaN()				// Not a number


	fmt.Println(f32, f64, scientific, posInf, negInf, nan)


	// Complex Numbers
	var c64 complex64 = 3 + 4i
	var c128 complex128 = 1.5 + 2.5i

	// complex number operations
	realPart := real(c128)
	imagPart := imag(c128)
	

	fmt.Println(c64, c128, realPart, imagPart)


	// Booleans
	var isTrue bool = true
	var isFalse bool = false

	// boolean operations
	result := isTrue && isFalse
	result1 := isTrue || isFalse
	result2 := !isFalse

	fmt.Println(isTrue, isFalse, result, result1, result2)


	// String Type
	var name string = "vali"
	str1 := "Hello, world !"
	str2 := `This is a raw string literal`

	// string operations
	length := len(str1)
	char := str1[0]

	// string concatination
	combined := str1 + " " + "Go !"

	// string formatting
	fmt.Printf("Length %d, First Char %c\n", length, char)								// Prints formatted version
	formatted := fmt.Sprintf("Length %d, First Char %c\n", length, char)	// returns formatted version

	fmt.Println(name)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(combined)
	fmt.Println(formatted)

	// Bytes and runes
	// byte is an alias for uint8
	var bee byte = 65 			// ASCII 'A'
	
	// rune is an alias for int32
	var ru rune = '♥'		// Unicode heart symbol
	var ru2 rune = 65 		// ASCII 'A'

	fmt.Println(bee, ru, ru2)

	// string to rune slice
	str := "Hello, 世界"
	runes := []rune(str)
	bytes := []byte(str)

	fmt.Println(runes, bytes)


	// Composite data types
	// Arrays: are fixed in size
	fmt.Println("")

	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [2]string = [2]string{"hello", "world"}
	
	// short declaration
	arr3 := [4]float64{1.1, 2.2, 3.3, 4.4}

	// array with ellipsis (compiler determines size)
	arr4 := [...]int{10, 20, 30, 40, 50}

	// partial initialization (remaining elements are zero values)
	arr5 := [5]int{1, 2}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)


	// Slices
	fmt.Println("\n\nSlices: \n")
	
	var s1 []int	// delcaration
	s1 = []int{1, 2, 3, 3, 4, 5} // initialization

	// short form
	s2 := []int{1, 2, 3}

	// create slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:2] // modifing slice modifies array


	// using make
	s4 := make([]int, 5)	// length 5, capacity 5
	s5 := make([]int, 5, 10) // length 5, capacity 10

	// slice operations
	s2 = append(s2, 6, 7, 8)

	// copy slices
	s6  := make([]int, len(s2))
	copy(s6, s2)


	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)


	// Maps
	fmt.Println("\n\nMaps: \n")

	//declaration and initialization
	var m1 map[string]int
	m1 = make(map[string]int)

	//short declaration
	m2 := map[string]int{
		"apple": 5,
		"banana": 3,
		"orange": 2,
	}

	// map operations
	m2["orange"] = 10
	
	fmt.Println(m1)
	fmt.Println(m2)


	// struct
	// multiple types


	type Address struct {
		Street string
		City string
		Zip string
	}

	type Person struct {
		Name string
		Age int
		Address Address
	}


	p1 := Person{	
		Name: "vali",
		Age: 29,
		Address: Address{
			Street: "Hyderabad",
			City: "Hyderabad",
			Zip: "502032",
		},
	}


	// anonymous structs
	p2 := struct {
		Name string
		Age int
	}{
		Name: "xyz",
		Age: 29,
	}

	fmt.Println(p1)
	fmt.Println(p2)


	// Pointers
	xx := 43
	var p *int = &xx	// &xx hold memory address, p *int reference to the address

	fmt.Println(xx)
	fmt.Println(p)
	fmt.Println(*p) // *p dereference the memory address to the value

	// modify the value through pointer
	*p = 200
	fmt.Println(xx)

	// pointer to pointer
	var pp **int = &p
	
	fmt.Println(pp)
	fmt.Println(**pp)

	// nil pointer
	var nilPtr *int
	fmt.Println(nilPtr)
	if nilPtr == nil {
		fmt.Println("nilPtr is nil")
	}


	// Functions:
	// Functions are first class types in go

	// Function type declaration
	var add func(int, int) int

	// assign function to a variable
	add = func(a int, b int) int {
		return a + b
	}

	// use function
	result23 := add(3, 4)
	fmt.Println(result23)

	calculate := func(op func(int, int) int, a, b int) int {
		return op(a, b)
	}

	multiply := func(a, b int) int { return a*b }
	fmt.Println(calculate(multiply, 2, 3))


	// Interfaces:
	// Interfaces define method signatures
	







}
