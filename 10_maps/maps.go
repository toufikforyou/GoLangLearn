package main

import (
	"fmt"
	"maps"
)

// maps -> like a js object, key value structure
func main() {
	// creating map

	var m = make(map[string]string)
	m["name"] = "Go lang"
	fmt.Println(m)
	fmt.Println(m["name"])

	x := make(map[string]int)
	x["age"] = 23
	x["count"] = 5
	fmt.Println(x)
	fmt.Println(x["age"]) // 23

	fmt.Println(len(x)) // 2

	clear(x)
	fmt.Println(x)      // map[]
	fmt.Println(len(x)) // 0

	y := map[string]int{"price": 20, "phone": 2}

	fmt.Println(y)
	fmt.Println(len(y)) // 2

	// check key is define or not
	v, ok := y["price"]

	if ok {
		fmt.Println("all ok", v) // all ok 20
	} else {
		fmt.Println("Not ok", v) // Not ok 0
	}

	z := map[string]int{"price": 21, "phone": 2}

	fmt.Println(maps.Equal(y, z)) // false
	z["price"] = 20
	fmt.Println(maps.Equal(y, z)) // true
}
