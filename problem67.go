// By starting at the top of the triangle below and moving to adjacent numbers on the
// row below, the maximum total from top to bottom is 23.
//
//    3
//   7 4
//  2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom in triangle.txt, a 15K text file containing a triangle with one-hundred rows.
//
// NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem,
// as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty
// billion years to check them all. There is an efficient algorithm to solve it.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Println

func readMatrix() [][]int {
	var matrix [][]int
	file, _ := os.Open("p067_triangle.txt")
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(line, " ")
		array := make([]int, 0)
		for _, value := range numbers {
			int_value, _ := strconv.Atoi(value)
			array = append(array, int_value)
		}
		matrix = append(matrix, array)
	}
	return matrix
}

// Creates a copy of the original filled with zeroes
func initializeResultMatrix(original [][]int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(original); i++ {
		result = append(result, make([]int, len(original[i])))
	}
	return result
}

// Since there's no way to brute force this problem I'm gonna try to implement
// a kind of memoization
func solve(matrix [][]int) int {
	resultMatrix := initializeResultMatrix(matrix)
	resultMatrix[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		row := matrix[i]
		bufferRow := resultMatrix[i-1]
		for j := 0; j < len(row); j++ {
			// left, right are results of the previous as in pascal's triangle
			// left values = current value + previous(left most)
			if j == 0 {
				resultMatrix[i][j] = row[0] + bufferRow[0]
			}
			// right values = current value + previous(right most)
			if j == len(row)-1 {
				resultMatrix[i][j] = row[j] + bufferRow[len(bufferRow)-1]
			}
			// inbetween cases: sum with the largest of the two possible parents
			//     59      |       59     |
			//   73 41     |    132 100   |
			//  52 40 09   |   184 ? 109  | -> first case choose the largest of the possible
			// 26 53 06 34 |  ?? ?? ?? ?? |    in [ 132, 100 ]
			if j > 0 && j < len(row)-1 {
				largest := 0
				if bufferRow[j-1] > bufferRow[j] {
					largest = bufferRow[j-1]
				} else {
					largest = bufferRow[j]
				}
				resultMatrix[i][j] = largest + row[j]
			}
		}
	}
	largest := 0
	for _, item := range resultMatrix[len(resultMatrix)-1] {
		if item > largest {
			largest = item
		}
	}
	return largest
}

func main() {
	fmt.Println("Problem 67")
	matrix := readMatrix()
	result := solve(matrix)
	p(result)
}
