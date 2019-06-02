// Prime digit replacements
// By replacing the 1st digit of the 2-digit number *3, it turns out that six of
// the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.
// By replacing the 3rd and 4th digits of 56**3 with the same digit, this
// 5-digit number is the first example having seven primes among the ten
// generated numbers, yielding the family:
// 56003, 56113, 56333, 56443, 56663, 56773, and 56993.
// Consequently 56003, being the first member of this family, is the smallest
// prime with this property.
// Find the smallest prime which, by replacing part of the number
// (not necessarily adjacent digits) with the same digit, is part of an eight
// prime value family.
package main

import (
	"fmt"
	"math"
	"strconv"
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

func generatePrimes(limit int) map[int]struct{} {
	primes := make(map[int]struct{})
	prime := 3
	for prime < limit {
		primes[prime] = struct{}{}
		prime = nextPrime(prime)
	}
	return primes
}

func uniqueDigits(number int) (int, int) {
	result := make(map[rune]struct{})
	strNumber := strconv.Itoa(number)
	for _, digit := range []rune(strNumber) {
		result[digit] = struct{}{}
	}
	return len(strNumber), len(result)
}

func filteredPrimes(primes map[int]struct{}, qtyOfRepeatedDigits int) map[int]struct{} {
	result := make(map[int]struct{})
	for prime := range primes {
		digits, uniqueCount := uniqueDigits(prime)
		if digits-uniqueCount >= qtyOfRepeatedDigits {
			result[prime] = struct{}{}
		}
	}
	return result
}

func generatePermutations(array []rune, n int, permutations *[]string) {
	if n == 1 {
		*permutations = append(*permutations, string(array))
	} else {
		for i := 0; i < n; i++ {
			generatePermutations(array, n-1, permutations)
			if n%2 == 0 {
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}

func familyOfNumbers(prime int, replacementString string) (family []int) {
	return
}

func seekNumber(familySize int, primes map[int]struct{}) {
	// base case 3rd and 4rd digits
}

func main() {
	p("Problem 51")
	primes := generatePrimes(1000000)
	p(len(primes))
	newPrimes := filteredPrimes(primes, 3)
	p(len(newPrimes))

	// primes := generatePrimes(100)
	// seekNumber(7, primes)
	// p(familyOfNumbers(56663, "--**-"))
	// p(filteredPrimes(primes, 2))
	// p(uniqueDigits(12345559))
}
