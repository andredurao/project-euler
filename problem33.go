// The fraction 49/98 is a curious fraction, as an inexperienced mathematician
// in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which
// is correct, is obtained by cancelling the 9s.

// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

// There are exactly four non-trivial examples of this type of fraction, less
// than one in value, and containing two digits in the numerator and denominator.

// If the product of these four fractions is given in its lowest common terms,
// find the value of the denominator.

package main

import (
	"fmt"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func digits(n int) (digits []int) {
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func simplify(fraction Fraction) (simplifiedFraction Fraction) {
	simplifiedFraction = fraction
	numeratorDigits := digits(fraction.Numerator)
	denominatorDigits := digits(fraction.Denominator)
	repeatedDigit := -1
	for _, digit := range numeratorDigits {
		if denominatorDigits[0] == digit || denominatorDigits[1] == digit {
			repeatedDigit = digit
		}
	}
	if repeatedDigit >= 0 {
		for _, digit := range numeratorDigits {
			if digit != repeatedDigit {
				simplifiedFraction.Numerator = digit
			}
		}
		for _, digit := range denominatorDigits {
			if digit != repeatedDigit {
				simplifiedFraction.Denominator = digit
			}
		}
	}
	return
}

var p = fmt.Println

func rangesForDigits(digits int) (min int, max int) {
	for i := 0; i < digits; i++ {
		min *= 10
		min++
	}
	max = min * 9
	return
}

func validFraction(fraction Fraction) bool {
	fractionValue := float64(fraction.Numerator) / float64(fraction.Denominator)
	if fractionValue >= 1.0 {
		return false
	}
	simplifiedFraction := simplify(fraction)
	simplifiedValue := float64(simplifiedFraction.Numerator) / float64(simplifiedFraction.Denominator)
	return fraction.Numerator%10 != 0 &&
		fractionValue == simplifiedValue &&
		simplifiedFraction.Numerator != fraction.Numerator
}

func main() {
	p("Problem 33")
	minValue, maxValue := rangesForDigits(2)
	for numerator := minValue; numerator <= maxValue; numerator++ {
		for denominator := minValue; denominator <= maxValue; denominator++ {
			fraction := Fraction{numerator, denominator}
			if validFraction(fraction) {
				p(fraction, simplify(fraction))
			}
		}
	}
}

// 1/4 1/5 2/5 4/8
// 1/4 * 1/5 * 2/5 * 1/2 => 2/200 => 1/100
