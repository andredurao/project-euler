// Let d(n) be defined as the sum of proper divisors of n
// (numbers less than n which divide evenly into n).
//
// If d(a) = b and d(b) = a, where a ≠ b,
// then a and b are an amicable pair and each of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
// therefore d(220) = 284.
// The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.
package main

import (
	"fmt"
	"math"
	_ "os"
	_ "strconv"
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

func divisors(n int) []int {
	items := make(map[int]struct{})
	items[1] = struct{}{}

	limit := int(math.Sqrt(float64(n)) + 1)
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			appendInMap(items, i)
			appendInMap(items, n/i)
		}
	}
	return keys(items)
}

func sum(items []int) int {
	total := 0
	for _, v := range items {
		total += v
	}
	return total
}

func createSetOfSums() map[int][]int {
	items := make(map[int][]int)
	for i := 1; i < 10000; i++ {
		sumOfDivisors := sum(divisors(i))
		items[sumOfDivisors] = append(items[sumOfDivisors], i)
	}
	return items
}

func main() {
	fmt.Println("Problem 21")
	items := createSetOfSums()
	total := 0
	for _, item := range items {
		if len(item) > 1 {
			total += sum(item)
		}
	}
	fmt.Println(total)
}
