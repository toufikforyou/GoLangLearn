package main

import (
	"fmt"
)

func main() {
	// Switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Friday, time.Saturday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's workday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its int", t)
		case string:
			fmt.Println("its string", t)
		case bool:
			fmt.Println("its bool", t)
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(false)
}
