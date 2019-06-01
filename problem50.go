// Consecutive prime sum
// The prime 41, can be written as the sum of six consecutive primes:
// 41 = 2 + 3 + 5 + 7 + 11 + 13
// This is the longest sum of consecutive primes that adds to a prime below 100
// The longest sum of consecutive primes below one-thousand that adds to a prime
// contains 21 terms, and is equal to 953.
// Which prime, below one-million, can be written as the sum of the most
// consecutive primes?

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

func generatePrimes(limit int) (primes []int) {
	primes = append(primes, 2)
	prime := 3
	for prime < limit {
		primes = append(primes, prime)
		prime = nextPrime(prime)
	}
	return
}

func consecutiveSumsCount(limit int, index int, primes []int) (count int, total int) {
	sum := 0
	for i := index; i < len(primes); i++ {
		sum += primes[i]
		if sum > limit {
			return
		}
		if isPrime(sum) && i > index {
			total = sum
			count = (i - index) + 1
		}
	}
	return
}
func seekPrime(limit int) (count int, max int) {
	primes := generatePrimes(limit)

	for i := 0; i < len(primes); i++ {
		qty, total := consecutiveSumsCount(limit, i, primes)
		if qty > count {
			count = qty
			max = total
		}
	}
	return
}

func main() {
	p("Problem 50")
	p(seekPrime(1000000))
}
