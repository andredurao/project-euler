// Prime permutations
// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases
// by 3330, is unusual in two ways:
// (i) each of the three terms are prime, and,
// (ii) each of the 4-digit numbers are permutations of one another.
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes,
// exhibiting this property, but there is one other 4-digit increasing sequence.
// What 12-digit number do you form by concatenating the three terms in this sequence?

package main

import (
	"fmt"
	"math"
	"sort"
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

// pre build a map of primes with 4 digits
func buildPrimesMap(primesMap map[int]struct{}) {
	prime := 999
	for {
		prime = nextPrime(prime)
		if prime < 10000 {
			primesMap[prime] = struct{}{}
		} else {
			break
		}
	}
}

func primePermutations(n int) (permutations []int) {
	digits := []rune(strconv.Itoa(n))
	numberPermutations := make([]string, 0)
	generatePermutations(digits, len(digits), &numberPermutations)
	permutationsMap := make(map[int]struct{})
	for _, value := range numberPermutations {
		number, _ := strconv.Atoi(value)
		if value[0] != '0' && isPrime(number) {
			permutationsMap[number] = struct{}{}
		}
	}
	for number := range permutationsMap {
		permutations = append(permutations, number)
	}

	return
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

func filterPrimes(primesMap map[int]struct{}) map[int][]int {
	result := make(map[int][]int)
	for prime := range primesMap {
		permutations := primePermutations(prime)
		sort.Sort(sort.IntSlice(permutations))

		if len(permutations) <= 2 {
			delete(primesMap, prime)
		} else {
			result[prime] = append(result[prime], permutations...)
		}
	}
	return result
}

func main() {
	p("Problem 49")
	primesMap := make(map[int]struct{})
	buildPrimesMap(primesMap)
	result := filterPrimes(primesMap)
	p(result)
	// p(primesMap)
	// p(primePermutations(1431))
}
