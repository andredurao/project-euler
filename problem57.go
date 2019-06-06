// Powerful digit sum
// Square root convergents

package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

type Fraction struct {
	Numerator   uint64
	Denominator uint64
}

func NewFraction(numerator, denominator uint64) *Fraction {
	return &Fraction{numerator, denominator}
}

func (fraction Fraction) String() (result string) {
	result = fmt.Sprintf("%d / %d", fraction.Numerator, fraction.Denominator)
	return
}

func Gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (fraction *Fraction) Simplify() {
	gcd := Gcd(fraction.Numerator, fraction.Denominator)
	fraction.Numerator /= gcd
	fraction.Denominator /= gcd
}

func (a *Fraction) Sum(b *Fraction) *Fraction {
	d := (a.Denominator * b.Denominator) / Gcd(a.Denominator, b.Denominator)
	n := ((d / a.Denominator) * a.Numerator) + ((d / b.Denominator) * b.Numerator)
	return NewFraction(n, d)
}

func (a *Fraction) Invert() {
	a.Numerator, a.Denominator = a.Denominator, a.Numerator
}

func fractionExpansion(n int) *Fraction {
	n--
	one := NewFraction(1, 1)
	two := NewFraction(2, 1)
	fraction := NewFraction(1, 2)
	sum := two.Sum(fraction)
	for i := 1; i < n; i++ {
		sum.Invert()
		sum = two.Sum(sum)
	}
	sum.Invert()
	sum = one.Sum(sum)
	// sum.Simplify()
	return sum
}

func digits(n uint64) int {
	str := strconv.FormatUint(n, 10)
	return (len(str))
}

func main() {
	p("Problem 57")

	total := 0

	for i := 2; i <= 1000; i++ {
		expansion := fractionExpansion(i)
		if digits(expansion.Numerator) > digits(expansion.Denominator) {
			total++
		}
	}
	p(total)
}
