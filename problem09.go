// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
//a2 + b2 = c2
//For example, 32 + 42 = 9 + 16 = 25 = 52.
//
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.
package main

import (
	"fmt"
)

func triplet(a int, b int, c int) bool {
	validSequence := ((a + b + c) == 1000) && (a < b && b < c)
	if !validSequence {
		return false
	}
	// fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
	return a*a+b*b == c*c
}

func main() {
	fmt.Println("Problem 09")

	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			for k := 1; k < 1000; k++ {
				if triplet(i, j, k) {
					fmt.Printf("TRIPLET a=%d b=%d c=%d\n", i, j, k)
				}
			}
		}
	}
}
