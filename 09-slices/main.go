package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice equals nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// use make to create slice with non-zero length
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// set and get like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// use append to return slice with new values at end
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// use copy to make new slice with values from existing slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// use [low:high] syntax to get precise slice (high not inclusive)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slice up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slice up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize slice in single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices package contains useful utility functions
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// slices can be multi-dimensional with varying inner slice lengths
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
