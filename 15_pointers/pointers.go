package main

import "fmt"

func changeNum(num int) {
	num = 65
	fmt.Println("In change Number", num)
}

func changeMainNum(num *int) {
	*num = 23
	fmt.Println("In chage Number", *num)
}

func main() {

	num := 1

	changeNum(num)

	fmt.Println("After changeNum in main", num) // 1

	fmt.Println("Memory address", &num) // memory address

	// Actual pointer use
	changeMainNum(&num)
	fmt.Println("After chage num in main", num) // now 23
}
