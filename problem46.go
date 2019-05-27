// Goldbach's other conjecture
// It was proposed by Christian Goldbach that every odd composite number can be
// written as the sum of a prime and twice a square.
// 9 =  7  + 2×1^2
// 15 = 7  + 2×2^2
// 21 = 3  + 2×3^2
// 25 = 7  + 2×3^2
// 27 = 19 + 2×2^2
// 33 = 31 + 2×1^2
// It turns out that the conjecture was false.
// What is the smallest odd composite that cannot be written as the sum of a
// prime and twice a square?

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

func ensurePrimesSet(n int, primes *[]int) {
	primesSetLength := len(*primes) - 1
	lastPrime := []int(*primes)[primesSetLength]
	if lastPrime < n {
		for {
			next := nextPrime(lastPrime)
			*primes = append(*primes, next)
			primesSetLength = len(*primes) - 1
			lastPrime = []int(*primes)[primesSetLength]
			if lastPrime >= n {
				return
			}
		}
	}
}

func decompose(n int, primes *[]int) (a, b int) {
	ensurePrimesSet(n, primes)
	for i := 0; []int(*primes)[i] < n; i++ {
		prime := []int(*primes)[i]
		j := 1
		for {
			result := prime + (2 * (j * j))
			if result < n {
				j++
			} else if result == n {
				return prime, j
			} else {
				break
			}
		}
	}
	return -1, -1
}
func main() {
	p("Problem 46")
	primes := []int{2, 3, 5, 7}
	oddNumber := 9
	for {
		a, b := 0, 0
		if isPrime(oddNumber) {
			primes = append(primes, oddNumber)
		} else {
			a, b = decompose(oddNumber, &primes)
		}
		if a == -1 && b == -1 {
			break
		}
		oddNumber += 2
	}
	p(oddNumber)
}
