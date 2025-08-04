package main

import "fmt"

// function that takes two ints and returns an int
func plus(a int, b int) int {
	return a + b
}

// consecutive parameters of same type may omit type
func plusPlus(a, b, c int) int {
	return a + b + c // explicit returns are necessary
}

func main() {
	// call a function
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
