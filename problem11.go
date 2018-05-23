// In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
// The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
// What is the greatest product of four adjacent numbers in the same direction
// (up, down, left, right, or diagonally) in the 20×20 grid?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix() [][]int {
	file, _ := os.Open("problem11grid.txt")
	fileScanner := bufio.NewScanner(file)
	matrix := make([][]int, 20)
	index := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		numbers := strings.Split(line, " ")
		for _, v := range numbers {
			number, _ := strconv.Atoi(v)
			matrix[index] = append(matrix[index], number)
		}
		fmt.Println(matrix[index])
		index++
	}

	return matrix
}

func main() {
	fmt.Println("Problem 11")
	readMatrix()
}
