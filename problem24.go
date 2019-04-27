// A permutation is an ordered arrangement of objects. For example, 3124 is one possible
// permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed
// numerically or alphabetically, we call it lexicographic order. The lexicographic
// permutations of 0, 1 and 2 are:

// 012   021   102   120   201   210

// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
package main

import (
	"fmt"
	"sort"
)

// one implementation of https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generatePermutations(array []rune, n int, result *[]string) {
	if n == 1 {
		*result = append(*result, string(array))
	} else {
		for i := 0; i < n; i++ {
			generatePermutations(array, n-1, result)
			if n%2 == 0 {
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}

// The usage of Heap's permutation algorithm is not the ideal solution for this problem
// I will try to use a sorted permutation algorithm like Steinhaus and instead of append
// items the program could stop at N
func main() {
	fmt.Println("Problem 24")
	var permutations []string
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	generatePermutations(digits, len(digits), &permutations)
	sort.Sort(sort.StringSlice(permutations))
	fmt.Println(permutations[1000000-1])
}
