// Notes:
// slices are dynamic
// most used construct in go
// useful methods

package main

import "fmt"


func main() {

	// declaring slice
	var nums []int
	
	// uninitialized slice is nil
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))
	
	// add some values
	nums = []int{1, 2, 3, 4, 5, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))


	// another way to initialize slices
	// we can add capacity for slice in this, with zero values
	// var <name> = make(<type>, capacity, initial capacity)
	var nums1 = make([]int, 2)
	// capacity: max nums can fit, it is not fixed

	fmt.Println(nums1)
	fmt.Println(len(nums1))
	fmt.Println(nums1 == nil) // zeroed values slice is not nil

	// apend




}
