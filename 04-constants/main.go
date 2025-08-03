package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// declaring a constant
	const n = 500000000

	// constant expression
	const d = 3e20 / n
	fmt.Println(d)

	// explicit conversion
	fmt.Println(int64(d))

	// n is given a type based on context
	fmt.Println(math.Sin(n))
}
