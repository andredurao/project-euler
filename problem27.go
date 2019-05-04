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

// FunctionMap refers to the coefficents of a function
type FunctionMap struct {
	A int
	B int
}

func isPrime(n int64) bool {
	max := int64(math.Ceil(math.Sqrt(float64(n))))
	var i int64
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Populate an array of int with prime values ≤ max
func populatePrimesList(list *[]int, max int) {
	*list = append(*list, 2)
	for i := 3; i <= max; i += 2 {
		if isPrime(int64(i)) {
			*list = append(*list, i)
		}
	}
}

func consecutivePrimes(a int, b int, primesMap map[int64]struct{}) (primes []int64) {
	var i, value int64
	for {
		value = i*i + int64(a)*i + int64(b)
		value = int64(math.Abs(float64(value)))
		_, primeFound := primesMap[value]
		if primeFound {
			primes = append(primes, value)
		} else if isPrime(value) {
			primesMap[value] = struct{}{}
			primes = append(primes, value)
		} else {
			break
		}
		i++
	}
	return
}

func main() {
	// b must be prime, because n starts with 0
	primesList := make([]int, 0)
	// map of primes to cache primes
	primesMap := make(map[int64]struct{})
	// map of the primes found for the coefficients A and B
	consecutivePrimesMap := make(map[FunctionMap][]int64)
	// the coefficients with most consecutive primes
	var maxFunctionMap FunctionMap

	fmt.Println("Problem 27")

	populatePrimesList(&primesList, 1000)

	for a := -999; a < 1000; a++ {
		for _, b := range primesList {
			coefficients := FunctionMap{a, b}
			consecutivePrimesMap[coefficients] = consecutivePrimes(a, b, primesMap)
		}
	}

	maxLen := 0
	for k, v := range consecutivePrimesMap {
		if len(v) > maxLen {
			maxLen = len(v)
			maxFunctionMap = k
		}
	}
	fmt.Println(maxFunctionMap)
	fmt.Println(len(consecutivePrimesMap[maxFunctionMap]))
	fmt.Println(maxFunctionMap.A * maxFunctionMap.B)
}
