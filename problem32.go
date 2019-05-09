// We shall say that an n-digit number is pandigital if it makes use of all the
// digits 1 to n exactly once;
// for example, the 5-digit number, 15234, is 1 through 5 pandigital.

// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
// multiplicand, multiplier, and product is 1 through 9 pandigital.

// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.

// HINT: Some products can be obtained in more than one way so be sure to only
// include it once in your sum.

package main

import (
	"fmt"
)

var p = fmt.Println

// the multiplicand, multiplier and product must have 9 digits altogheter
// m * n = p
// 0 ≤ log10(p) - log10(m) - log10(n) ≤ 1
// 9 ≤ 2 * ceil(log10(m)) + ceil(log10(n)) ≤ 10
// So to result in 9 digits the possible digits qty for m and n are:
// [1, 4] and [2, 3]; or m + n = 5
func digitsCountVariations() {
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			qty := 2 * (i + j)
			if qty == 9 || qty == 10 {
				p(i, j)
			}
		}
	}
}

func digits(n int) (digits []int) {
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func isPandigital(m int, n int) bool {
	digitsList := digits(m)
	digitsList = append(digitsList, digits(n)...)
	digitsList = append(digitsList, digits(m*n)...)
	if len(digitsList) != 9 {
		return false
	}
	digitsMap := make(map[int]struct{})
	for _, digit := range digitsList {
		_, repeated := digitsMap[digit]
		if repeated {
			return false
		}
		digitsMap[digit] = struct{}{}
	}
	total := 0
	for i := 1; i <= 9; i++ {
		_, found := digitsMap[i]
		if found {
			total++
		}
	}
	return total == 9
}

func rangesForDigits(digits int) (min int, max int) {
	for i := 0; i < digits; i++ {
		min *= 10
		min++
	}
	max = min * 9
	return
}

func main() {
	p("Problem 32")
	productsMap := make(map[int]struct{})
	for _, mDigits := range []int{1, 2} {
		mMin, mMax := rangesForDigits(mDigits)
		nMin, nMax := rangesForDigits(5 - mDigits)
		for m := mMin; m <= mMax; m++ {
			for n := nMin; n <= nMax; n++ {
				if isPandigital(m, n) {
					// p(m, n, (m * n))
					productsMap[m*n] = struct{}{}
				}
			}
		}
	}
	sum := 0
	for product := range productsMap {
		sum += product
	}
	p(sum)
}
