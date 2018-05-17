// The four adjacent digits in the 1000-digit number that have the greatest
// product are 9 × 9 × 8 × 9 = 5832.
// Find the thirteen adjacent digits in the 1000-digit number that have the
// greatest product. What is the value of this product?

package main

import (
	"fmt"
	"os"
	"strconv"
)

func multiply(window []byte) uint64 {
	var multiple uint64 = 1
	for i := 0; i < len(window); i++ {
		digit, _ := strconv.Atoi(string(window[i]))
		multiple *= uint64(digit)
	}
	return multiple
}

func main() {
	fmt.Println("problem 08")

	file, err := os.Open("problem08number.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1000)
	file.Read(data)

	var largestMultiple uint64 = 0
	max := 1000
	windowSize := 13

	for index := 0; index < (max - windowSize); index++ {
		limit := index + windowSize
		window := data[index:limit]
		currentMultiple := multiply(window)
		if currentMultiple > largestMultiple {
			largestMultiple = currentMultiple
		}
	}

	fmt.Println(largestMultiple)
}
