/*
Prime pair sets
The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
*/

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

// Checking if the qty of k-permutations for n is evaluated: n! / (n-k)!
// ex: P(n,k) => P(5,2) = 5! / (5-2)! = 20
func permutations(primes []int) (values []int) {
	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primes); j++ {
			if i != j {
				strValue := strconv.Itoa(primes[i])
				strValue += strconv.Itoa(primes[j])
				intValue, _ := strconv.Atoi(strValue)
				values = append(values, intValue)
			}
		}
	}
	return
}

func generatePrimes(limit int) (map[int]struct{}, []int) {
	primesList := make([]int, 0)
	primesMap := make(map[int]struct{})
	prime := 3
	for prime < limit {
		primesMap[prime] = struct{}{}
		primesList = append(primesList, prime)
		prime = nextPrime(prime)
	}
	return primesMap, primesList
}

func isSetComposedOfPrimes(primes []int, primesMap map[int]struct{}) bool {
	for _, prime := range primes {
		_, found := primesMap[prime]
		if !found {
			if !isPrime(prime) {
				return false
			}
			primesMap[prime] = struct{}{}
		}
	}
	return true
}

// WIP - Start looking for a pair of values that are in the same conditions
func seekPairSets(primesList []int, primesMap map[int]struct{}) {
	// start at 3,5 because permutations with 2 won't be valid
	for i := 0; i < (len(primesList) - 1); i++ {
		for j := i + 1; j < len(primesList); j++ {
			currentSet := []int{primesList[i], primesList[j]}
			primes := permutations(currentSet)
			if isSetComposedOfPrimes(primes, primesMap) {
				for k := j + 1; k < len(primesList); k++ {
					currentSet := []int{primesList[i], primesList[j], primesList[k]}
					primes := permutations(currentSet)
					if isSetComposedOfPrimes(primes, primesMap) {
						for l := k + 1; l < len(primesList); l++ {
							currentSet := []int{primesList[i], primesList[j], primesList[k], primesList[l]}
							primes := permutations(currentSet)
							if isSetComposedOfPrimes(primes, primesMap) {
								p(currentSet)
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	p("Problem 60")
	primesMap, primesList := generatePrimes(1000)
	p(len(primesList))
	p(len(primesMap))
	seekPairSets(primesList, primesMap)
}
