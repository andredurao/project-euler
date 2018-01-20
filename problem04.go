// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(num int) bool {
	word := strconv.Itoa(num)
	center := int(len(word) / 2)
	equalNumbers := 0
	for i := 0; i < center; i++ {
		if word[i] == word[len(word)-(i+1)] {
			equalNumbers++
		}
	}
	return equalNumbers == center
}

func main() {
	palindrome := 0

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j
			if IsPalindrome(product) && product > palindrome {
				palindrome = product
			}
		}
	}

	fmt.Printf("largest palindrome = %d\n", palindrome)
}
