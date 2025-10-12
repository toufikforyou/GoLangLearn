package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// // while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// // infinite loop
	// for {
	// 	fmt.Println("1")
	// }

	// classic for loop
	// for i := 0; i < 3; i++ {

	// 	if i == 1 {
	// 		continue
	// 	}

	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// range
	for i := range 10 {
		fmt.Println(i)
	}
}
