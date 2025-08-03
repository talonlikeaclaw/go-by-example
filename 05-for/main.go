package main

import "fmt"

func main() {
	// basic for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic (not preferred)
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// iteration over range (preferred)
	for i := range 3 {
		fmt.Println("range", i)
	}

	// without condition = infinite unless break/return
	for {
		fmt.Println("loop")
		break // break out of loop
	}

	for n := range 6 {
		if n%2 == 0 {
			 continue // continue to next iteration
		}
		fmt.Println(n)
	}
}
