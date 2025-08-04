package main

import (
	"fmt"
	"maps"
)

func main() {
	// create an empty map with make(map[key-type]val-type)
	m := make(map[string]int)

	// set key/value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get a value from key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if key doesn't exist a zero-value is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// get the length of map
	fmt.Println("len:", len(m))

	// delete removes key/value pairs from map
	delete(m, "k2")
	fmt.Println("map:", m)

	// remove all key/valeu pairs from a map
	clear(m)
	fmt.Println("map:", m)

	// check whether a key exists, _ ignores the value
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize on same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// the maps package contains useful utility functions for maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
