package main

import "fmt"

func main() {
	// age := 13

	// if age >= 18 {
	// 	fmt.Println("Persion is an adult")
	// } else {
	// 	fmt.Println("Persion is not an adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Persion is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Persion is teenager")
	// } else {
	// 	fmt.Println("Persion is a kid")
	// }

	// var role = "admin"
	// hasPermision := false

	// if role == "admin" || hasPermision {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }

	// if role == "admin" && hasPermision {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }

	if age := 15; age >= 18 {
		fmt.Println("Persion is an adult", age)
	} else if age >= 12 {
		fmt.Println("Persion is a teenager", age)
	} else {
		fmt.Println("Persion is not an adult", age)
	}
}
