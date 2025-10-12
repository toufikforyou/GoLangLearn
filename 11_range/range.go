package main

import "fmt"

// iterating over data structures
func main() {
	nums := []int{10, 20, 30, 40, 50}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(i, num)
	}
	fmt.Println("Sum =", sum)

	m := map[string]string{"fname": "MH", "lname": "TOUFIK", "age": "tweenty"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, string(c)) // use only c then return unicode value for this char
	}
}
