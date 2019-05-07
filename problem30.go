// Surprisingly there are only three numbers that can be written as the sum of
// fourth powers of their digits:

// 1634 = 1^4 + 6^4 + 3^4 + 4^4
// 8208 = 8^4 + 2^4 + 0^4 + 8^4
// 9474 = 9^4 + 4^4 + 7^4 + 4^4
// As 1 = 14 is not a sum it is not included.

// The sum of these numbers is 1634 + 8208 + 9474 = 19316.

// Find the sum of all the numbers that can be written as the sum of fifth
// powers of their digits.
package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

func splitDigits(n uint64) (digits []int) {
	for n > 0 {
		digit := int(n % 10)
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func sumOfDigits(digits []int) (sum *big.Int) {
	for item := range digits {
		exp := new(big.Int)
		value := new(big.Int)
		value.SetInt64(int64(item))
		five := new(big.Int)
		five.SetInt64(int64(5))
		exp = exp.Exp(value, five, nil)
		sum.Add(sum, five)
	}
	return
}

func main() {
	p("Problem 30")
	digits := splitDigits(1234)
	p(digits)
	p(sumOfDigits(digits))
}
