package main

import (
	"fmt"
	"slices"
)

// slice -> dymanic length declire
// most used construct in go
func main() {
	// uninitialized slice is nill

	// var nums []int
	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums))   // 0

	var nums = make([]int, 2, 5)
	// fmt.Println(nums == nil) // false
	fmt.Println(nums)      // [0 0]
	fmt.Println(cap(nums)) // 5
	fmt.Println(len(nums)) // 2

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)

	fmt.Println(nums)      // [0 0 1 2 3 4]
	fmt.Println(cap(nums)) // 10

	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	nums = append(nums, 8)
	nums = append(nums, 9)
	nums = append(nums, 10)

	fmt.Println(nums)      // [0 0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println(cap(nums)) // 20

	clear(nums)
	nums = append(nums, 10)

	fmt.Println(nums)
	fmt.Println(len(nums)) // 12
	fmt.Println(cap(nums)) // 20

	// copy slice
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println(nums2)

	// Slice operator
	var nums5 = []int{1, 2, 3}

	fmt.Println(nums5[0:2]) // [1 2]
	fmt.Println(nums5[:2])  // [1 2]
	fmt.Println(nums5[0:1]) // [1]
	fmt.Println(nums5[0:])  // [1 2 3]
	fmt.Println(nums5[1:])  // [2 3]

	// slice
	var nums6 = []int{5, 6}
	var nums7 = []int{5, 6}
	fmt.Println(slices.Equal(nums6, nums7)) // true
	nums6 = append(nums6, 2)
	fmt.Println(slices.Equal(nums6, nums7)) // false

	var nums8 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums8)
}
