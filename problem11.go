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
const WindowSize = 4

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

func seekByDirection(cursor Cursor, matrix [][]int) int {
	largest := 0
	// up right down left
	directions := []Cursor{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {1, 1}, {-1, -1}}
	for _, direction := range directions {
		result := directionHandler(cursor, matrix, direction)
		if result > largest {
			largest = result
		}
	}
	return largest
}

func directionHandler(cursor Cursor, matrix [][]int, direction Cursor) int {
	multiple := 1
	//fmt.Println("---")
	//fmt.Println(multiple)
	var ary [4]int
	for index := 0; index < WindowSize; index++ {
		//fmt.Println(matrix[cursorCopy.Row][cursorCopy.Col])
		cursor.Row += direction.Row
		cursor.Col += direction.Col
		if cursor.Row >= Limit || cursor.Col >= Limit || cursor.Row < 0 || cursor.Col < 0 {
			ary[index] = 0
		} else {
			ary[index] = matrix[cursor.Row][cursor.Col]
		}
		multiple *= ary[index]
	}
	fmt.Println(ary[0], "\t", ary[1], "\t", ary[1], "\t", ary[1], "\t", multiple)
	return multiple
}

func seek(matrix [][]int) {
	largest := 0
	for i := 0; i < Limit; i++ {
		for j := 0; j < Limit; j++ {
			result := seekByDirection(Cursor{j, i}, matrix)
			if result > largest {
				largest = result
			}
		}
	}
	fmt.Println(largest)
}

func main() {
	fmt.Println("Problem 11")
	matrix := readMatrix()
	seek(matrix)
}
