// Take the number 192 and multiply it by each of 1, 2, and 3:
// 192 × 2 = 384
// 192 × 3 = 576
// By concatenating each product we get the 1 to 9 pandigital, 192384576.
// We will call 192384576 the concatenated product of 192 and (1,2,3)

// The same can be achieved by starting with 9 and multiplying by 1..5,
// giving the pandigital, 918273645,
// which is the concatenated product of 9 and (1,2,3,4,5).

// What is the largest 1 to 9 pandigital 9-digit number that can be formed as
// the concatenated product of an integer with (1,2, ... , n) where n > 1?
package main

import (
	"fmt"
	"strconv"
)

func isPandigital(number string) bool {
	digitsMap := make(map[rune]struct{})
	if len(number) != 9 {
		return false
	}
	for _, digit := range number {
		_, found := digitsMap[digit]
		if found || digit == '0' {
			return false
		}
		digitsMap[digit] = struct{}{}
	}
	return true
}

func buildNumber(base int) (number string) {
	number = strconv.Itoa(base)
	for i := 2; len(number) < 9; i++ {
		number += strconv.Itoa(base * i)
	}
	return
}

func main() {
	fmt.Println("Problem 38")
	// The largest base number will have 4 digits at most
	max := 0
	for i := 1; i < 10000; i++ {
		number := buildNumber(i)
		if isPandigital(number) {
			total, _ := strconv.Atoi(number)
			if total > max {
				max = total
			}
		}
	}
	fmt.Println("MAX", max)
}
