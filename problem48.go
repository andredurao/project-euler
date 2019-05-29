// Self powers
// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

func main() {
	p("Problem 48")
	total := big.NewInt(0)
	var i int64
	for i = 1; i <= 1000; i++ {
		item := big.NewInt(i)
		exp := item.Exp(item, item, nil)
		total = total.Add(total, exp)
	}
	number := total.String()
	p(number[(len(number) - 10):(len(number))])
}
