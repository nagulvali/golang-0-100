/*
	print below values as both binary and hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5, 6
*/

package main

import "fmt"

func main() {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", a, a, a)
	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", b, b, b)
	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", c, c, c)
	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", d, d, d)
	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", e, e, e)
	fmt.Printf("a: %d as binary: %b, as hexadecimal: %x\n", f, f, f)
}
