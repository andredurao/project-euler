/*
The 5-digit number, 16807=7^5, is also a fifth power.
Similarly, the 9-digit number, 134217728=8^9, is a ninth power.

How many n-digit positive integers exist which are also an nth power?
*/

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

func pow(x, y int64) *big.Int {
	bx := big.NewInt(x)
	by := big.NewInt(y)
	return bx.Exp(bx, by, nil)
}

func qtyOfDigits(pow *big.Int) int64 {
	return int64(len(pow.String()))
}

func checkPow(pow uint64) (result bool) {
	result = false
	return
}

func main() {
	p("Problem 63")

	total := 0
	var x, y, size int64
	// iterate from 1 ~ 99 to calculate x^y
	for x = 1; x < 99; x++ {
		for y = 1; y < 99; y++ {
			pow := pow(x, y)
			size = qtyOfDigits(pow)
			if size == y {
				total++
				// p(x, y, pow, size, total)
			}
			// If size(qty of digits of the result) > y move to next x
			if size > y {
				break
			}
		}
	}
	p("total =", total)
}
