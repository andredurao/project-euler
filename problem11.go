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

const Limit = 20

type Cursor struct {
	Row int
	Col int
}

func readMatrix() [][]int {
	file, _ := os.Open("problem11grid.txt")
	fileScanner := bufio.NewScanner(file)
	matrix := make([][]int, 20)
	index := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(line, " ")
		for _, v := range numbers {
			number, _ := strconv.Atoi(v)
			matrix[index] = append(matrix[index], number)
		}
		index++
	}

	return matrix
}

// TODO: CREATE Struct CURSOR
func seekByDirection(cursor Cursor, matrix [][]int) int {
	largest := 0
	// up right down left
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	fmt.Println("seekByDirection")
	for _, direction := range directions {
		result := directionHandler(cursor, matrix, direction)
		if result > largest {
			largest = result
		}
	}
	return largest
}

// TODO: CREATE Struct CURSOR
func directionHandler(cursor Cursor, matrix [][]int, direction []int) int {
	multiple := matrix[cursor.Row][cursor.Col]
	return multiple
}

func seek(matrix [][]int) {
	fmt.Println("seek")
	largest := 0
	for i := 0; i < Limit; i++ {
		for j := 0; j < Limit; j++ {
			// TODO: CREATE Struct CURSOR
			result := seekByDirection(Cursor{j, i}, matrix)
			if result > largest {
				largest = result
			}
		}
	}
}

func main() {
	fmt.Println("Problem 11")
	matrix := readMatrix()
	seek(matrix)
}
