// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10
// without any remainder.

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?

package main

import "fmt"

func BruteForceMultipleNumber() int {
	number := 20
	for !IsDividedByAll(number) {
		number++
	}
	return number
}

func IsDividedByAll(number int) bool {
	// All the numbers from 1..10 can be infered, like: 12 -> 2 x 3 x 4
	// There's no need to check then
	divisors := [11]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	quantity := 0
	for _, value := range divisors {
		if number%value == 0 {
			quantity++
		}
	}
	// fmt.Println(number, quantity)
	return quantity == 11
}

func main() {
	// This problem could be solved only using simple Math, but ok let's golang it... :(

	number := BruteForceMultipleNumber()
	fmt.Println(number)
}
