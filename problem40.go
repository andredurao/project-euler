// Champernowne's constant
// An irrational decimal fraction is created by concatenating the positive
// integers:

// 0.123456789101112131415161718192021...
//              ^

// It can be seen that the 12th digit of the fractional part is 1.

// If (d)n represents the nth digit of the fractional part, find the value of
// the following expression.

// d(1) * d(10) * d(100) * d(1000) * d(10000) * d(100000) * d(1000000)
package main

import (
	"fmt"
	"math"
	"strconv"
)

var p = fmt.Println

type Range struct {
	Min int
	Max int
}

// Returns the min and max values for a given class, ex:
// rangeInClass(3) => 100:999
// rangeInClass(4) => 1000:9999
func rangeInClass(class int) (classRange Range) {
	if class == 0 {
		return Range{0, 0}
	}
	classRange.Min = int(math.Pow(10, float64(class-1)))
	classRange.Max = int(math.Pow(10, float64(class))) - 1
	return
}

func classShift(class int) int {
	acc := 0
	for i := 0; i < class; i++ {
		currentPage := rangeInClass(i)
		delta := currentPage.Max - currentPage.Min + 1
		acc += (delta * i)
	}
	return acc
}

func rangeOfDigitsInClass(class int) Range {
	return Range{Min: (classShift(class) + 1), Max: classShift(class + 1)}
}

func rangeOfNumber(n int) (Range, int) {
	class := 1
	currentRange := rangeOfDigitsInClass(class)
	for !(n >= currentRange.Min && n <= currentRange.Max) {
		class++
		currentRange = rangeOfDigitsInClass(class)
	}
	return currentRange, class
}

func valueOfNumber(n int) (value int) {
	page, class := rangeOfNumber(n)
	// p(page, class)
	slot := (n - page.Min) / class
	// p("slot", slot)
	classNumbers := rangeInClass(class)
	// p("classNumbers", classNumbers)
	number := classNumbers.Min + slot
	// p("number", number)
	digitIndex := (n - page.Min) % class
	// p("digitIndex", digitIndex)
	strValue := strconv.Itoa(number)
	value, _ = strconv.Atoi(string(strValue[digitIndex]))
	return
}

func main() {
	p("Problem 40")
	values := [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}
	total := 1
	for _, value := range values {
		total *= valueOfNumber(value)
	}
	p(total)
}
