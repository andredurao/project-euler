/*
Check: https://projecteuler.net/problem=65
The first ten terms in the sequence of convergents for e are:
2, 3, 8/3, 11/4, 19/7, 87/32, 106/39, 193/71, 1264/465, 1457/536
The sum of digits in the numerator of the 10th convergent is 1+4+5+7 = 17.
Find the sum of digits in the numerator of the 100th convergent of the continued fraction for e.

I've written a simplifier for fractions of the continued fraction of √2 in problem 57 already

refs:
- https://en.wikipedia.org/wiki/Euler%27s_continued_fraction_formula
- https://mathworld.wolfram.com/eContinuedFraction.html
*/

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
	// p("fraction", fraction)
	gcd := Gcd(fraction.Numerator, fraction.Denominator)
	// p("fraction", fraction)
	// p("gcd", gcd)
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
	terms := generateTerms(n)
	result := NewFraction(big.NewInt(int64(0)), big.NewInt(1))
	for i := n - 1; i >= 0; i-- {
		term := NewFraction(big.NewInt(int64(terms[i])), big.NewInt(1))
		result = result.Sum(term)
		result.Invert()
	}
	two := NewFraction(big.NewInt(2), big.NewInt(1))
	result = two.Sum(result)
	result.Simplify()
	return result
}

func digits(n *big.Int) int {
	return len(n.String())
}

func generateTerms(length int) []int {
	// [1, 2, 1, 1, 4, 1, 1, 6, 1, 1, 8, ...]
	result := make([]int, length)
	result[0] = 1
	result[1] = 2
	for i := 2; i < length; i++ {
		x := i - 1
		if x%3 == 0 {
			result[i] = ((x / 3) + 1) * 2
		} else {
			result[i] = 1
		}
	}
	return result
}

func main() {
	p("Problem 57")

	expansion := fractionExpansion(100)
	p("100th fraction =", expansion)

	zero := big.NewInt(0)
	ten := big.NewInt(10)
	modulus := big.NewInt(0)
	numerator := expansion.Numerator
	total := 0
	for numerator.Cmp(zero) > 0 {
		numerator.DivMod(numerator, ten, modulus)
		total += int(modulus.Int64())
	}
	p(total)
}
