package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func printAny(value ...interface{}) {
	for _, val := range value {
		fmt.Print(val, " ")
	}
}

func main() {
	// Like a n number of parameters pass at function
	fmt.Println(1, 2, 3, 4, 5, 6, "5")

	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)

	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))

	// Custom print
	printAny("1", 112.234, 34532, 65, false, true, 1 == 2, false == true, "String")
}
