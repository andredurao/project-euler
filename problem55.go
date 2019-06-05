// Lychrel numbers

package main

import (
	"fmt"
	"math/big"
)

var p = fmt.Println

type BigInt struct {
	*big.Int
}

func (n *BigInt) Reverse() *BigInt {
	reverseNumber := NewBigInt(0)
	reverseNumber.SetString(reverse(n.String()), 10)
	return reverseNumber
}

// TODO Add NewBigInt into BigInt interface
func NewBigInt(value int64) *BigInt {
	number := new(big.Int)
	number.SetInt64(value)
	return &BigInt{number}
}

func (number *BigInt) Add(n2 *BigInt) {
	value := new(big.Int)
	value.SetString(number.String(), 10)
	value2 := new(big.Int)
	value2.SetString(n2.String(), 10)
	value = value.Add(value, value2)
	*number = BigInt{value}
}

func isPalindrome(n *BigInt) bool {
	number := n.String()
	center := int(len(number) / 2)
	equalNumbers := 0
	for i := 0; i < center; i++ {
		if number[i] == number[len(number)-(i+1)] {
			equalNumbers++
		}
	}
	return equalNumbers == center
}

func reverse(str string) string {
	if str != "" {
		return reverse(str[1:]) + str[:1]
	}
	return ""
}

func isLychrel(n int) bool {
	number := NewBigInt(int64(n))
	for i := 0; i < 50; i++ {
		if i > 0 && isPalindrome(number) {
			return false
		}
		number.Add(number.Reverse())
	}
	return true
}

func main() {
	p("Problem 55")
	total := 0
	for i := 100; i < 10000; i++ {
		if isLychrel(i) {
			total++
		}
	}
	p(total)
}
