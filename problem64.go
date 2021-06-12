/*
All square roots are periodic when written as continued fractions and can be written in the form:
Check: https://projecteuler.net/problem=64

Exactly four continued fractions, for N <= 13, have an odd period.
How many continued fractions for N <= 10000 have an odd period?
*/

package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

// Initially I was following the solution proposed in sympy lib for method
// continued_fraction_periodic. but then I've found other references explaining
// about the stop point when a reaches 2 times the amount of the integer root
// Initially I've tried writing my own "isRepeating" method but that didn't
// worked as expected

func continuedFraction(n int) []int {
	terms := make([]int, 0)
	root := int(math.Sqrt(float64(n)))
	if int(root*root) == n {
		return terms
	}
	a := root
	numerator := 0
	denominator := 1

	for a != (2 * root) {
		numerator = denominator*a - numerator
		denominator = (n - numerator*numerator) / denominator
		a = (root + numerator) / denominator
		// p(a, numerator, denominator, period)
		terms = append(terms, a)
	}
	return terms
}

func main() {
	p("Problem 64")
	total := 0
	for i := 0; i <= 10000; i++ {
		terms := continuedFraction(i)
		if len(terms)%2 == 1 {
			total += 1
			p(i)
		}
	}
	p("total=", total)
}
