// In England the currency is made up of pound, £, and pence, p,
// and there are eight coins in general circulation:

// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
// It is possible to make £2 in the following way:

// 1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
// How many different ways can £2 be made using any number of coins?

package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

func sumOfCoinPurse(values []int) (total int) {
	if len(values) == 0 {
		return 0
	}
	for _, coin := range values {
		total += coin
	}
	return
}

func slice2String(slice []int) (text string) {
	for _, value := range slice {
		text += strconv.Itoa(value)
	}
	return
}

func pack(coinValues []int, coinPurse []int, startIndex int, limit int, combinationsMap map[string]struct{}) {
	sum := sumOfCoinPurse(coinPurse)
	if sum < limit {
		for i := startIndex; i < len(coinValues); i++ {
			pack(coinValues, append(coinPurse, coinValues[i]), i, limit, combinationsMap)
		}
	} else if sum == limit {
		combinationsMap[slice2String(coinPurse)] = struct{}{}
	}
}

// This problem resembles the Knapsack problem
// An alternate solution is trying to solve it using golang concurrency
// instead of recursivity

func main() {
	p("Problem 31")
	coinValues := []int{200, 100, 50, 20, 10, 5, 2, 1}
	coinPurse := []int{}
	combinationsMap := make(map[string]struct{})
	pack(coinValues, coinPurse, 0, 200, combinationsMap)
	p(len(combinationsMap))
}
