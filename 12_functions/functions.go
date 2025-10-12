package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func add2(a, b int) int { // a, b datatype are same
	return a + b
}

func getLanguages() (string, string, bool) { // multiple return values
	return "Go Lang", "JavaScript", true
}

func procesIt(fn func()) func(b int) int {
	// Function call
	fn()
	x := 23
	return func(b int) int { // function return
		return x * b
	}
}

func main() {
	// int value sum
	result := add(45, 54)
	fmt.Println(result)
	fmt.Println(add2(23, 34))

	fmt.Println(getLanguages())

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	// function in function
	fn := func() {
		fmt.Println("Multification")
	}

	fmt.Println(procesIt(fn)(34))
}
