// The decimal number, 585 = 1001001001(binary), is palindromic in both bases.
// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
// (Please note that the palindromic number, in either base, may not include leading zeros.)

package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

// from Problem04
func isPalindrome(value int, base int) bool {
	number := strconv.FormatInt(int64(value), base)
	center := int(len(number) / 2)
	equalNumbers := 0
	for i := 0; i < center; i++ {
		if number[i] == number[len(number)-(i+1)] {
			equalNumbers++
		} else {
			return false
		}
	}
	return true
}

func main() {
	p("Problem 36")
	sum := 0
	for i := 1; i < 1000000; i++ {
		if isPalindrome(i, 10) && isPalindrome(i, 2) {
			sum += i
		}
	}
	p(sum)
}
