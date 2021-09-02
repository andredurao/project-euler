// Consider quadratic Diophantine equations of the form: x^2 - Dy^2 = 1
// For example, when D=13, the minimal solution in x is 649^2 – 13×180^2 = 1.
// It can be assumed that there are no solutions in positive integers when D is square.
// By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:
// 3^2 – 2×2^2 = 1
// 2^2 – 3×1^2 = 1
// 9^2 – 5×4^2 = 1
// 5^2 – 6×2^2 = 1
// 8^2 – 7×3^2 = 1
// Hence, by considering minimal solutions in x for D ≤ 7, the largest x is obtained when D=5.
// Find the value of D ≤ 1000 in minimal solutions of x for which the largest value of x is obtained.

// References:
// 1: https://en.wikipedia.org/wiki/Pell%27s_equation
// 2: https://en.wikipedia.org/wiki/Continued_fraction#Square_roots (again)
// 3: https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
package main

import (
	"fmt"
	"math"
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
	gcd := Gcd(fraction.Numerator, fraction.Denominator)
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

func fractionExpansion(n int, size int, originalTerms []int) *Fraction {
	size--
	result := NewFraction(big.NewInt(int64(0)), big.NewInt(1))
	terms := expandSeries(n, originalTerms)
	for i := size - 1; i >= 0; i-- {
		term := NewFraction(big.NewInt(int64(terms[i])), big.NewInt(1))
		result = result.Sum(term)
		result.Invert()
	}
	sqrt := int64(math.Sqrt(float64(n)))
	isqrt := NewFraction(big.NewInt(sqrt), big.NewInt(1))
	result = isqrt.Sum(result)
	result.Simplify()
	return result
}

func expandSeries(n int, terms []int) []int {
	var result []int
	for len(result) < n {
		result = append(result, terms...)
	}
	return result
}

func continuedFraction(n int) []int {
	terms := make([]int, 0)
	root := int(math.Sqrt(float64(n)))
	if int(root*root) == n {
		return terms
	}
	a := root
	numerator := 0
	denominator := 1

	for a != (2 * root) {
		numerator = denominator*a - numerator
		denominator = (n - numerator*numerator) / denominator
		a = (root + numerator) / denominator
		terms = append(terms, a)
	}
	return terms
}

func solve(d int) (*big.Int, *big.Int) {
	// For square n, there is no solution except (1, 0)
	if isPerfectSquareRoot(d) {
		return big.NewInt(0), big.NewInt(1)
	}
	terms := continuedFraction(d)
	one := big.NewInt(int64(1))
	for i := 1; i <= 100; i++ {
		fraction := fractionExpansion(d, i, terms)
		x, y := fraction.Numerator, fraction.Denominator
		x2 := big.NewInt(int64(x.Int64()))
		y2 := big.NewInt(int64(y.Int64()))
		x2.Mul(x, x)
		y2.Mul(y, y)
		// x^2 - Dy^2 = 1
		result := x2.Sub(x2, y2.Mul(y2, big.NewInt(int64(d))))

		if result.Cmp(one) == 0 {
			return x, y
		}
	}
	return big.NewInt(0), big.NewInt(0)
}

func isPerfectSquareRoot(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func main() {
	fmt.Println("Problem 66")
	largest_x := big.NewInt(0)
	largest_d := 0
	// compare with https://en.wikipedia.org/wiki/Pell%27s_equation#Solutions
	for d := 2; d <= 1000; d++ {
		x, _ := solve(d)
		if x.Cmp(largest_x) > 0 {
			largest_x = x
			largest_d = d
		}
	}

	p("Result = ", largest_d)
}
