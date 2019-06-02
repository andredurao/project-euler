// Permuted multiples
// It can be seen that the number, 125874, and its double, 251748, contain
// exactly the same digits, but in a different order.
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
// contain the same digits.
package main

import (
	"fmt"
	"sort"
	"strconv"
)

var p = fmt.Println

func digits(number int) (list []int) {
	strNumber := strconv.Itoa(number)
	for _, runeDigit := range []rune(strNumber) {
		digit := int(runeDigit - '0')
		list = append(list, digit)
	}
	return
}

func sliceEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func validate(value int) bool {
	factors := []int{2, 3, 4, 5, 6}
	numberDigits := digits(value)
	sort.Sort(sort.IntSlice(numberDigits))
	for _, factor := range factors {
		valueDigits := digits(value * factor)
		sort.Sort(sort.IntSlice(valueDigits))
		if !sliceEqual(numberDigits, valueDigits) {
			return false
		}
	}
	return true
}

func seek() (value int) {
	for {
		value++
		if validate(value) {
			return value
		}
	}
}

func main() {
	p("Problem 52")
	p(seek())
}
