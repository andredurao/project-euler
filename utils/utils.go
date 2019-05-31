package main

import (
	"math"
	"sort"
	"strconv"
)

func contains(m map[int]struct{}, n int) bool {
	_, isPresent := m[n]
	return isPresent
}

func appendInMap(m map[int]struct{}, n int) {
	if !contains(m, n) {
		m[n] = struct{}{}
	}
}

func keys(m map[int]struct{}) []int {
	items := make([]int, len(m))
	i := 0
	for k, _ := range m {
		items[i] = k
		i++
	}
	return items
}

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

func sortedKeys(m map[int][]int) (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	return
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
	//compact result
	for _, key := range sortedKeys(result) {
		// keep only the permutations that refer to the smallest one
		if result[key][0] != key {
			delete(result, key)
		}
	}
	return result
}
