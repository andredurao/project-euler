// The Fibonacci sequence is defined by the recurrence relation:
// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:
// F1 = 1 // F2 = 1 // F3 = 2
// F4 = 3 // F5 = 5 // F6 = 8
// F7 = 13 // F8 = 21 // F9 = 34
// F10 = 55 // F11 = 89 // F12 = 144
// The 12th term, F12, is the first term to contain three digits.

// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
package main

import (
	"fmt"
	"math/big"
)

func digits(value *big.Int) (count int) {
	count = len(value.String())
	return
}

func main() {
	fmt.Println("Problem 25")
	value1 := big.NewInt(1)
	value2 := big.NewInt(1)
	// Starts at 3 because the items 1 and 2 are 1, 1
	index := 3
	for {
		total := new(big.Int)
		total.Add(value1, value2)
		if digits(total) == 1000 {
			break
		}
		value1 = value2
		value2 = total
		index++
	}
	fmt.Println("index = ", index)
}
