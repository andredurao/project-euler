// A perfect number is a number for which the sum of its proper divisors is exactly equal
// to the number.

// For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28,
// which means that 28 is a perfect number.
// A number n is called deficient if the sum of its proper divisors is less than n and it
// is called abundant if this sum exceeds n.
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number
// that can be written as the sum of two abundant numbers is 24.

// By mathematical analysis, it can be shown that all integers greater than 28123 can be
// written as the sum of two abundant numbers. However, this upper limit cannot be reduced
// any further by analysis even though it is known that the greatest number that cannot be
// expressed as the sum of two abundant numbers is less than this limit.
// Find the sum of all the positive integers which cannot be written as the sum of two
// abundant numbers.
package main

import (
	"fmt"
	"math"
)

var abundantList []int
var sumMap map[int]bool

func main() {
	fmt.Println("Problem 23")
	abundantList = buildAbundantList()
	fmt.Println(len(abundantList))
	sumMap = buildSumOfAbundants()
	checkSumOfAbundantsNumbers()
}

func checkSumOfAbundantsNumbers() {
	var total uint64
	for i := 0; i < 28123; i++ {
		_, value := sumMap[i]
		if !value {
			total += uint64(i)
		}
	}
	fmt.Println(total)
}

func divisors(n int) []int {
	m := make(map[int]struct{})
	appendInMap(m, 1)
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			appendInMap(m, i)
			appendInMap(m, n/i)
		}
	}
	return keys(m)
}

func sum(items []int) int {
	total := 0
	for _, v := range items {
		total += v
	}
	return total
}

func abundant(value int) bool {
	return sum(divisors(value)) > value
}

func buildAbundantList() []int {
	var list []int
	for i := 1; i < 28123; i++ {
		if abundant(i) {
			list = append(list, i)
		}
	}
	return list
}

func buildSumOfAbundants() map[int]bool {
	sumMap := make(map[int]bool)
	for i := 0; i < len(abundantList); i++ {
		for j := i; j < len(abundantList); j++ {
			sum := abundantList[i] + abundantList[j]
			if sum <= 28123 {
				sumMap[sum] = true
			} else {
				break
			}
		}
	}
	return sumMap
}

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
