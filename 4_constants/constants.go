package main

import "fmt"

const age int = 30

var age2 int = 12

// age3:= 23 must be declare on any inside the function

func main() {
	const name string = "Go Lang"

	// name = "JavaScript" not again initial

	fmt.Println(name)

	fmt.Println(age, age2)

	const (
		host = "localhost"
		port = 5000
	)

	// port = 3000 not again declare const value

	fmt.Println(host, port)
}
