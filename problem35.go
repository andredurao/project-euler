// The number, 197, is called a circular prime because all rotations of the
// digits: 197, 971, and 719, are themselves prime.
// There are thirteen such primes below 100:
// 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
// How many circular primes are there below one million?

package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

func digits(n int) (digits []int) {
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}
	return
}

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func aryToI(ary []int) (value int) {
	for _, digit := range ary {
		value *= 10
		value += digit
	}
	return
}

func isCircular(n int, primesMap map[int]bool) (result bool) {
	elementsMap := make(map[int]struct{})
	elementsMap[n] = struct{}{}
	number := digits(n)
	for i := 0; i < len(number); i++ {
		number = append(number[1:], number[0])
		elementsMap[aryToI(number)] = struct{}{}
	}

	primesCount := 0
	for value := range elementsMap {
		valuePrime, found := primesMap[value]
		if !found {
			valuePrime = isPrime(value)
			primesMap[value] = valuePrime
		}
		if valuePrime {
			primesCount++
		}
	}
	result = primesCount == len(elementsMap)
	// could cache all values from elementsMap if result is true

	return
}

func main() {
	p("Problem 35")
	primesMap := make(map[int]bool)

	count := 0
	for i := 1; i < 1000000; i++ {
		if isCircular(i, primesMap) {
			count++
		}
	}
	p(count)
}
