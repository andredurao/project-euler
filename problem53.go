// Combinatoric selections
// There are exactly ten ways of selecting three from five, 12345:
// 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
// In combinatorics, we use the notation, 5C3=10.
// In general, nCr=(n!)/(r!(n−r)!), where r≤n
// It is not until n=23, that a value exceeds one-million: 23C10=1144066.
// How many, not necessarily distinct, values of nCr for 1≤n≤100, are greater than
// one-million?

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

func factorial(n int64) *big.Int {
	result := big.NewInt(0)
	result.MulRange(1, n)
	return result
}

// nCr=(n!)/(r!(n−r)!); r≤n
func combinations(n, r int64) *big.Int {
	numerator := factorial(n)
	denominator := big.NewInt(0).Mul(factorial(r), factorial(n-r))
	return numerator.Div(numerator, denominator)
}

func main() {
	p("Problem 53")
	var n, r int64
	total := 0
	million := big.NewInt(1000000)
	for r = 1; r <= 100; r++ {
		for n = r; n <= 100; n++ {
			result := combinations(n, r)
			if result.Cmp(million) > 0 {
				total++
			}
		}
	}
	p(total)
}
