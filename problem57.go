// Powerful digit sum
// Square root convergents

package main

import (
	"fmt"
)

var p = fmt.Println

type Fraction struct {
	Numerator   int
	Denominator int
}

func NewFraction(numerator, denominator int) *Fraction {
	return &Fraction{numerator, denominator}
}

func (fraction Fraction) String() (result string) {
	result = fmt.Sprintf("%d / %d", fraction.Numerator, fraction.Denominator)
	return
}

func Gcd(a, b int) int {
	for b != 0 {
		p(a, b)
		a, b = b, a%b
	}
	return a
}

func (fraction *Fraction) Simplify() {
	gcd := Gcd(fraction.Numerator, fraction.Denominator)
	fraction.Numerator /= gcd
	fraction.Denominator /= gcd
}

func main() {
	p("Problem 57")
	p(Gcd(67, 198))
	fraction := NewFraction(66, 198)
	p(fraction)
	fraction.Simplify()
	p(fraction)
}
