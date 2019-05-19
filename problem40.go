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

type Range struct {
	Min int
	Max int
}

// Returns the min and max values for a given class, ex:
// rangeInClass(3) => 100:999
// rangeInClass(4) => 1000:9999
func rangeInClass(class int) (classRange Range) {
	classRange.Min = int(math.Pow(10, float64(class-1)))
	classRange.Max = int(math.Pow(10, float64(class))) - 1
	return
}

func digitsQty(number int) int {
	// logResult := math.Log10(float64(number))
	// return int(math.Floor(logResult)) + 1
	return len(strconv.Itoa(number))
}

func shiftOfNumber(number int) (class int) {
	class, shift := 1, 0
	status := true
	for status {
		currentRange := rangeInClass(class)
		nextShift := shift + ((currentRange.Max - currentRange.Min + 1) * class)
		fmt.Println("nextShift", nextShift)
		if nextShift > number {
			status = false
		} else {
			shift = nextShift
			class++
		}
	}
	shiftStarter := number - shift + 1
	window := shiftStarter / class
	slide := shiftStarter % class
	l := rangeInClass(class-1).Max + window
	strl := string(strconv.Itoa(l)[slide])

	fmt.Println("class", class)
	fmt.Println("shift", shift)
	fmt.Println("window", window)
	fmt.Println("slide", slide)
	fmt.Println("l", l)
	fmt.Println("strl", strl)
	return 0
}

func main() {
	fmt.Println("Problem 40")
	// fmt.Println(rangeInClass(3))
	// fmt.Println(digitsQty(10))
	// fmt.Println(classOf(1))
	// fmt.Println(classOf(10))
	class := shiftOfNumber(100)
	fmt.Println(class)
	// fmt.Println(rangeInClass(3))
	// fmt.Println(rangeInClass(class))
	// fmt.Println(valueInPos(1))
}
