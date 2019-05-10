// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

// Find the sum of all numbers which are equal
// to the sum of the factorial of their digits.

// Note: as 1! = 1 and 2! = 2 are not sums they are not included.

package main

import (
	"fmt"
)

var p = fmt.Println

func factorial(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

func digits(n int) (digits []int) {
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func validateFactorialDestructure(n int, mapFactorials map[int]int) bool {
	total := 0
	for _, digit := range digits(n) {
		value, found := mapFactorials[digit]
		if found {
			total += value
		} else {
			value = factorial(digit)
			total += value
			mapFactorials[digit] = value
		}
	}
	return total == n
}

func main() {
	p("Problem 34")
	mapFactorials := make(map[int]int)
	for i := 3; i < 99999; i++ {
		if validateFactorialDestructure(i, mapFactorials) {
			p(i)
		}
	}
}
