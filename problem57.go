// Powerful digit sum
// Square root convergents

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

func NewFraction(n, d *big.Int) *Fraction {
	return &Fraction{n, d}
}

func (fraction Fraction) String() (result string) {
	result = fmt.Sprintf(
		"%s / %s",
		fraction.Numerator.String(),
		fraction.Denominator.String(),
	)
	return
}

func Gcd(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	newA := new(big.Int)
	newB := new(big.Int)
	newA.Set(a)
	newB.Set(b)
	for newB.Cmp(zero) != 0 {
		newA, newB = newB, newA.Mod(newA, newB)
	}
	return newA
}

func (fraction *Fraction) Simplify() {
	p("fraction", fraction)
	gcd := Gcd(fraction.Numerator, fraction.Denominator)
	p("fraction", fraction)
	p("gcd", gcd)
	fraction.Numerator = fraction.Numerator.Div(fraction.Numerator, gcd)
	fraction.Denominator = fraction.Denominator.Div(fraction.Denominator, gcd)
}

func (a *Fraction) Sum(b *Fraction) *Fraction {
	d := new(big.Int)
	d.Set(a.Denominator)
	d.Mul(a.Denominator, b.Denominator)
	d.Div(d, Gcd(a.Denominator, b.Denominator))

	na := new(big.Int)
	na.Set(d)
	na.Div(d, a.Denominator)
	na.Mul(na, a.Numerator)

	nb := new(big.Int)
	nb.Set(d)
	nb.Div(d, b.Denominator)
	nb.Mul(nb, b.Numerator)

	return NewFraction(na.Add(na, nb), d)
}

func (a *Fraction) Invert() {
	a.Numerator, a.Denominator = a.Denominator, a.Numerator
}

func fractionExpansion(n int) *Fraction {
	n--
	one := NewFraction(big.NewInt(1), big.NewInt(1))
	two := NewFraction(big.NewInt(2), big.NewInt(1))
	fraction := NewFraction(big.NewInt(1), big.NewInt(2))
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

func digits(n *big.Int) int {
	return len(n.String())
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
