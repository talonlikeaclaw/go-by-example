package main

import "fmt"

func main() {
	// create an array of 5 zero-valued ints
	var a [5]int
	fmt.Println("emp:", a)

	// set/get a value at an index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// get the length of an array
	fmt.Println("len:", len(a))

	// declare and initiazize array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// use ... to have compiler count elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// specify index with : to zero elements between
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// multi-dimensional arrays
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// create/initialize multi-dimensional array
	twoD = [2][3]int {
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
