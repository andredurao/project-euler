// Problem 14

//The following iterative sequence is defined for the set of positive integers:
//
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//
//Using the rule above and starting with 13, we generate the following sequence:
//
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
//Although it has not been proved yet (Collatz Problem), it is thought that all starting
//numbers finish at 1.
//
//Which starting number, under one million, produces the longest chain?
//
//NOTE: Once the chain starts the terms are allowed to go above one million.
package main

import (
	"fmt"
)

func Chain(items []int) []int {
	n := items[len(items)-1]
	nn := 0

	if n == 1 {
		return items
	} else {
		if n%2 == 0 {
			nn = n / 2
		} else {
			nn = 3*n + 1
		}
		items = append(items, nn)
		items = Chain(items)
	}
	return items
}

func main() {
	fmt.Println("Problem 14")
	chainSize := 0
	item := 1
	for i := 1; i < 1000000; i++ {
		list := Chain([]int{i})
		if len(list) > chainSize {
			chainSize = len(list)
			item = i
		}
	}
	fmt.Println(item, " -> ", chainSize)
}
