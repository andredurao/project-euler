// Starting with the number 1 and moving to the right in a clockwise direction a 5
// by 5 spiral is formed as follows:
// 21 22 23 24 25
// 20  7  8  9 10
// 19  6  1  2 11
// 18  5  4  3 12
// 17 16 15 14 13
// It can be verified that the sum of the numbers on the diagonals is 101.
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed
// in the same way?

package main

import (
	"fmt"
)

func generateMatrix(size int) ([][]int, error) {
	var matrix [][]int

	if size%2 == 0 {
		return matrix, fmt.Errorf("Size %d is not even", size)
	}

	matrix = make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	return matrix, nil
}

func main() {
	fmt.Println("Problem 28")
	matrix, err := generateMatrix(5)
	if err == nil {
		fmt.Println(matrix)
	} else {
		fmt.Println(err)
	}
}
