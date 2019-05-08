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
	"math"
)

var p = fmt.Println

func digits(n int) (digits []int) {
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func validate(n int, power int) bool {
	sum := 0

	for _, i := range digits(n) {
		sum += int(math.Pow(float64(i), float64(power)))
		if sum > n {
			return false
		}
	}
	return sum == n
}

func limit(power int) int {
	maxDigitValue := int(math.Pow(9.0, float64(power)))
	return len(digits(maxDigitValue)) * maxDigitValue
}

func main() {
	p("Problem 30")
	total := 0
	for i := 2; i <= limit(5); i++ {
		if validate(i, 5) {
			total += i
		}
	}
	p(total)
}
