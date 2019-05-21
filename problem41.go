// Pandigital prime
// We shall say that an n-digit number is pandigital if it makes use of all the
// digits 1 to n exactly once.
// For example, 2143 is a 4-digit pandigital and is also prime.
// What is the largest n-digit pandigital prime that exists?
package main

import (
	"fmt"
	"math"
	"strconv"
)

var p = fmt.Println

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generatePermutations(array []rune, n int, permutations *[]string) {
	if n == 1 {
		*permutations = append(*permutations, string(array))
	} else {
		for i := 0; i < n; i++ {
			generatePermutations(array, n-1, permutations)
			if n%2 == 0 {
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}

func main() {
	p("Problem 41")
	// 💡: iterate through the permutations of 1..[5:9];
	// checking the largest prime
	digits := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	max := 0
	for i := 5; i <= 9; i++ {
		var permutations []string
		generatePermutations(digits[0:i], i, &permutations)
		for _, value := range permutations {
			number, _ := strconv.Atoi(value)
			if isPrime(number) && number > max {
				max = number
			}
		}
	}
	p(max)
}
