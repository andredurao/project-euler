//The number 3797 has an interesting property.
//Being prime itself, it is possible to continuously remove digits from left to
//right, and remain prime at each stage: 3797, 797, 97, and 7.
//Similarly we can work from right to left: 3797, 379, 37, and 3.

//Find the sum of the only eleven primes that are both truncatable from left to
//right and right to left.

// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPrime(n int64) bool {
	max := int64(math.Ceil(math.Sqrt(float64(n))))
	var i int64
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isTruncablePrime(n int64) bool {
	number := strconv.FormatInt(n, 10)
	for i := 0; i < len(number); i++ {
		leftStr := number[i:]
		if len(leftStr) > 0 {
			leftNumber, _ := strconv.ParseInt(leftStr, 10, 64)
			if !isPrime(leftNumber) {
				return false
			}
		}
		rightStr := number[:i]
		if len(rightStr) > 0 {
			rightNumber, _ := strconv.ParseInt(rightStr, 10, 64)
			if !isPrime(rightNumber) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Problem 37")

	var i, count, sum int64 = 13, 1, 0
	count = 1
	for count < 12 && i < 1000000 {
		if isTruncablePrime(i) {
			fmt.Println(count, i)
			count++
			sum += i
		}
		i += 2
	}
	fmt.Println(sum)
}
