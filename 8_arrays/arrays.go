package main

import "fmt"

//  numbered sequence of specific length
func main() {
	// var nums [4]int

	// nums[0] = 1
	// fmt.Println(nums[0])

	// // array print
	// fmt.Println(nums)

	// // array length
	// fmt.Println(len(nums))

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var names [3]string
	// names[2] = "12"
	// fmt.Println(names) // value not define then show empty value ""

	// nums := [3]int{1, 2, 3}

	// fmt.Println(nums)

	nums := [2][2]int{{1, 2}, {3, 4}}
	nums[0][1] = 5
	fmt.Println(nums)

	// - fixed size, that is predictable
	// - memory optimization
	// - Contant time access

}
