// Distinct primes factors
// The first two consecutive numbers to have two distinct prime factors are:
// 14 = 2 × 7
// 15 = 3 × 5
// The first three consecutive numbers to have three distinct prime factors are:
// 644 = 2² × 7 × 23
// 645 = 3 × 5 × 43
// 646 = 2 × 17 × 19.
// Find the first four consecutive integers to have four distinct prime factors
// each. What is the first of these numbers?

package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int) (next int) {
	next = n + 2
	for !isPrime(next) {
		next += 2
	}
	return
}

func primeFactors(n int, primes *[]int) (primeDivisors []int) {
	number := n
	i := 0
	prime := (*primes)[i]
	for {
		if i == len(*primes) {
			*primes = append(*primes, nextPrime(prime))
		}
		prime = (*primes)[i]
		if number%prime == 0 {
			primeDivisors = append(primeDivisors, prime)
			number /= prime
		} else {
			i++
		}
		if number == 1 {
			return
		}
	}
}

func main() {
	p("Problem 47")
	primes := []int{2, 3, 5, 7}
	p(primeFactors(1579, &primes))
}
