// A unit fraction contains 1 in the numerator.
// The decimal representation of the unit fractions with denominators 2 to 10 are given:
// 1/2  = 0.5            1/3  = 0.(3)       1/4  = 0.25
// 1/5  = 0.2            1/6  = 0.1(6)      1/7  = 0.(142857)
// 1/8  = 0.125          1/9  = 0.(1)       1/10 = 0.1
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle.
// It can be seen that 1/7 has a 6-digit recurring cycle.

// Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its
// decimal fraction part.
// https://en.wikipedia.org/wiki/Repeating_decimal
package main

import (
	"fmt"
)

func recurringCycle(numerator int, denominator int) (repetend string) {
	remaindersMap := make(map[int]int)

	remainder := numerator % denominator

	//starts at false because it's first iteration
	repeatingRemainder := false

	for remainder != 0 && !repeatingRemainder {
		remaindersMap[remainder] = len(repetend)
		remainder *= 10
		currentDigit := remainder / denominator
		repetend = fmt.Sprintf("%s%d", repetend, currentDigit)
		remainder %= denominator
		_, repeatingRemainder = remaindersMap[remainder]
	}

	if remainder == 0 {
		repetend = ""
	} else {
		repetend = repetend[remaindersMap[remainder]:]
	}

	return
}

func main() {
	fmt.Println("Problem 26")
	var maxRecurringCycleLength, maxD int

	for d := 2; d < 1000; d++ {
		currentRecurringCycle := recurringCycle(1, d)
		currentRecurringCycleLength := len(currentRecurringCycle)
		if currentRecurringCycleLength > maxRecurringCycleLength {
			maxRecurringCycleLength = currentRecurringCycleLength
			maxD = d
		}
	}
	fmt.Println("max d = ", maxD, " length = ", maxRecurringCycleLength)
}
