// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there
// are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
// how many letters would be used?

// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two)
// contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of
// "and" when writing out numbers is in compliance with British usage

// execute : go get github.com/divan/num2words
package main

import (
	"fmt"
	"github.com/divan/num2words"
	"os"
	"regexp"
	"strconv"
)

func sum(value int) {
	var total int = 0
	re := regexp.MustCompile(`\w`)
	for i := 1; i <= value; i++ {
		number := num2words.ConvertAnd(i)
		total += len(re.FindAllString(number, -1))
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Problem 17")
	value, _ := strconv.Atoi(os.Args[1])
	sum(value)
}
