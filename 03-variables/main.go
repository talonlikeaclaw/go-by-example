package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// multiple at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go can infer type
	var d = true
	fmt.Println(d)

	// declarations without values are assigned "zero-values"
	var e int
	fmt.Println(e)

	// := is shorthand for declaring/initializing a variable
	f := "apple"
	fmt.Println(f)
}
