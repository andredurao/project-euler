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
)

func divisors(n int) []int {
	items := []int{1}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			items = append(items, i)
			items = append(items, n/i)
		}
	}
	return items
}

func sum(items []int) int {
	total := 0
	for _, v := range items {
		total += v
	}
	return total
}

func sumsMap() map[int]int {
	items := make(map[int]int)
	for i := 1; i < 10000; i++ {
		items[i] = sum(divisors(i))
	}
	return items
}

func main() {
	fmt.Println("Problem 21")
	items := sumsMap()
	total := 0
	for k, v := range items {
		if items[v] == k && k != v {
			total += v
		}
	}
	fmt.Println(total)
}
