// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"math"
)

func isPrime(n uint) bool {
	max := uint(math.Ceil(math.Sqrt(float64(n))))
	var i uint
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Problem 10")
	// includes 2 already
	var i uint = 3
	var sum uint = 2
	for i < 2000000 {
		if isPrime(i) {
			sum += i
		}
		i += 2
	}
	fmt.Println(sum)
}
