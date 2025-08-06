package main

import "fmt"

// (int, int) indicates this function returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// using multiple assignments
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use blank _ to ignore values
	_, c := vals()
	fmt.Println(c)
}
