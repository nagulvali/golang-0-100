package main

import "fmt"

func main() {
	// Arrays are fixed length in golang
	var nums [4]int
	
	// Print array length
	fmt.Println(len(nums))

	// Initially zeroed values added if no values are added after declaration
	fmt.Println(nums)

	// Add elements to the array
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(nums)

	// we can not add more than 4 numbers as arrays is already fixed
	// nums[4] = 5
	// fmt.Println(nums)



	// BOOLEANS
	var vals [5]bool
	fmt.Println(vals)

	vals[2] = true
	
	fmt.Println(vals)


	// STRINGS
	var names [4]string
	names[1] = "Nagul"
	names[2] = "Vali"

	fmt.Println(names)



	// Notes:
	// We can also initialize the array with values in declaration itself


	var nums1 [3]int = [3]int{1, 2, 3,}
	fmt.Println(nums1)

	// short form
	nums2 := [3]int{1,2,3}
	fmt.Println(nums2)

	// 2d arrays
	var nums3 [2][2]int = [2][2]int{{2,3}, {4,5}}
	fmt.Println(nums3)




}