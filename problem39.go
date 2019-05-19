// If p is the perimeter of a right angle triangle with integral length sides,
// {a,b,c}, there are exactly three solutions for p = 120.
// {20,48,52}, {24,45,51}, {30,40,50}
// For which value of p ≤ 1000, is the number of solutions maximised?
package main

import (
	"fmt"
)

type Triangle struct {
	A int
	B int
	C int
}

func solutions(perimeter int) (solution []Triangle) {
	// A has p-2 at most
	for a := (perimeter - 2); a > 2; a-- {
		for b := (perimeter - a - 1); b > 1; b-- {
			c := perimeter - a - b
			if b > c && a*a == b*b+c*c {
				solution = append(solution, Triangle{a, b, c})
			}
		}
	}
	return
}

func main() {
	max, maxP := 0, 0
	for p := 1000; p >= 3; p-- {
		currentSolutions := solutions(p)
		if len(currentSolutions) > max {
			max, maxP = len(currentSolutions), p
		}
	}
	fmt.Println("MAX", maxP)
}
