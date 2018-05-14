// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?
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
	fmt.Println("Problem 07")
	// includes 2 already
	var primeCounter uint = 1
	// starts at 3
	var counter uint = 3
	var currentPrime uint
	// TODO: use defer of go funcs to use concurrence
	for primeCounter < 10001 {
		if isPrime(counter) {
			currentPrime = counter
			primeCounter++
		}
		// iterate through odd numbers only
		counter += 2
	}
	fmt.Println(currentPrime)
}
