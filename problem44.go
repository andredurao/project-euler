// Pentagon numbers
// Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2.
// The first ten pentagonal numbers are:
// 1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...
// It can be seen that P4 + P7 = 22 + 70 = 92 = P8.
// However, their difference, 70 − 22 = 48, is not pentagonal.
// Find the pair of pentagonal numbers, Pj and Pk, for which their sum and
// difference are pentagonal and D = |Pk − Pj| is minimised;
// what is the value of D?
package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

func roots(a int, b int, c int) (solutions []float64) {
	delta := (b * b) - 4*(a*c)
	x1 := ((float64(b) * -1.0) + math.Sqrt(float64(delta))) / float64(2*a)
	x2 := ((float64(b) * -1.0) - math.Sqrt(float64(delta))) / float64(2*a)
	solutions = append(solutions, x1)
	solutions = append(solutions, x2)

	return
}

func isIntegral(value float64) bool {
	return value == float64(int(value))
}

func isPentagonalNumber(n int) bool {
	solutions := roots(3, -1, (-2 * n))
	for _, value := range solutions {
		if value > 0 && isIntegral(value) {
			return true
		}
	}
	return false
}

func pentagonalNumber(n int) int {
	return ((3 * n * n) - n) / 2
}

func main() {
	p("Problem 44")
	p(isPentagonalNumber(5))
	for i := 1; i < 10000; i++ {
		for j := i; j <= 10000; j++ {
			pi := pentagonalNumber(i)
			pj := pentagonalNumber(j)
			sum := pi + pj
			diff := pi - pj
			if diff < 0 {
				diff *= -1
			}
			if isPentagonalNumber(sum) && isPentagonalNumber(diff) {
				p("diff -> ", diff)
			}
		}
	}
}