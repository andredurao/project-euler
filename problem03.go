// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

// set an array of factors and iterate through items checking factor of the array
func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	divCount := 0
	for i := 1; i < int(max); i++ {
		if int(n)%i == 0 {
			divCount++
			if divCount > 1 {
				return false
			}
		}
	}
	return divCount == 1
}

func PrimeDividerOnList(s []int, n int) int {
	// iterate on splice checking if n is a multiple of some item
	for _, v := range s {
		if n%v == 0 {
			return v
		}
	}
	return 0
}

func FindNextPrime(primes []int, num int) int {
	firstVal := primes[len(primes)-1] + 1
	for i := firstVal; i <= num; i++ {
		if num%i == 0 && isPrime(i) {
			return i
		}
	}
	return -1
}

// TODO: There's no need to use a slice here, just repeat dividing until remainder != 0
func main() {
	var primes []int
	primes = append(primes, 2)
	num := 600851475143

	for num > 1 {
		primeDivider := PrimeDividerOnList(primes, num)
		if primeDivider != 0 {
			num = num / primeDivider
		} else {
			primes = append(primes, FindNextPrime(primes, num))
		}
	}

	fmt.Println(primes)
}
