// Sub-string divisibility
// The number, 1406357289, is a 0 to 9 pandigital number because it is made up of
// each of the digits 0 to 9 in some order, but it also has a rather interesting
// sub-string divisibility property.
// Let d1 be the 1st digit, d2 be the 2nd digit, and so on.
// In this way, we note the following:

// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
// Find the sum of all 0 to 9 pandigital numbers with this property.
package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

func isPandigital(number string) bool {
	digitsMap := make(map[rune]struct{})
	if len(number) != 10 || number[0] == '0' {
		return false
	}
	for _, digit := range number {
		_, found := digitsMap[digit]
		if found {
			return false
		}
		digitsMap[digit] = struct{}{}
	}
	return true
}

func checkSubStrings(number string, primesSet []PrimeStruct) bool {
	if !isPandigital(number) {
		return false
	}
	for _, primeStruct := range primesSet {
		currentNumber, _ := strconv.Atoi(number[primeStruct.StartPos : primeStruct.StartPos+3])
		if currentNumber%primeStruct.Prime != 0 {
			return false
		}
	}
	return true
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

type PrimeStruct struct {
	Prime    int
	StartPos int
}

func main() {
	var total int64

	p("Problem 43")
	// 💡: iterate through the permutations digits array
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	primesSet := []PrimeStruct{{2, 1}, {3, 2}, {5, 3}, {7, 4}, {11, 5}, {13, 6}, {17, 7}}
	var permutations []string
	generatePermutations(digits, len(digits), &permutations)
	for _, value := range permutations {
		if checkSubStrings(value, primesSet) {
			currentValue, _ := strconv.Atoi(value)
			total += int64(currentValue)
		}
	}
	p(total)
}
