// Euler discovered the remarkable quadratic formula:

// n2+n+41
// It turns out that the formula will produce 40 primes for the consecutive
// integer values 0≤n≤39.
// However, when n=40,402+40+41=40(40+1)+41 is divisible by 41, and certainly
// when n=41,412+41+41 is clearly divisible by 41.

// The incredible formula n2−79n+1601 was discovered, which produces 80 primes
// for the consecutive values 0≤n≤79.
// The product of the coefficients, −79 and 1601, is −126479.

// Considering quadratics of the form: n2+an+b, where |a|<1000 and |b|≤1000
// where |n| is the modulus/absolute value of n
// e.g. |11|=11 and |−4|=4
// Find the product of the coefficients, a and b, for the quadratic expression
// that produces the maximum number of primes for consecutive values of n,
// starting with n=0.

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func populatePrimesList(list *[]int, max int) {
	*list = append(*list, 2)
	for i := 3; i <= max; i += 2 {
		if isPrime(i) {
			*list = append(*list, i)
		}
	}
}

func main() {
	// b must be prime, because the consecutive values start with n=0
	primesList := make([]int, 0)

	fmt.Println("Problem 27")
	populatePrimesList(&primesList, 1000)
	fmt.Println(primesList)
}
