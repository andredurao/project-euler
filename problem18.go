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
// Find the maximum total from top to bottom of the triangle below:
//
// NOTE: As there are only 16384 routes, it is possible to solve this problem by trying
// every route. However, Problem 67, is the same challenge with a triangle containing
// one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)

package main

import (
	"fmt"
)

var max int = 0

func initialize() [][]uint8 {
	var matrix [][]uint8
	//	matrix = append(matrix, []uint8{3})
	//	matrix = append(matrix, []uint8{7, 4})
	//	matrix = append(matrix, []uint8{2, 4, 6})
	//	matrix = append(matrix, []uint8{8, 5, 9, 3})
	matrix = append(matrix, []uint8{75})
	matrix = append(matrix, []uint8{95, 64})
	matrix = append(matrix, []uint8{17, 47, 82})
	matrix = append(matrix, []uint8{18, 35, 87, 10})
	matrix = append(matrix, []uint8{20, 04, 82, 47, 65})
	matrix = append(matrix, []uint8{19, 01, 23, 75, 03, 34})
	matrix = append(matrix, []uint8{88, 02, 77, 73, 07, 63, 67})
	matrix = append(matrix, []uint8{99, 65, 04, 28, 06, 16, 70, 92})
	matrix = append(matrix, []uint8{41, 41, 26, 56, 83, 40, 80, 70, 33})
	matrix = append(matrix, []uint8{41, 48, 72, 33, 47, 32, 37, 16, 94, 29})
	matrix = append(matrix, []uint8{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14})
	matrix = append(matrix, []uint8{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57})
	matrix = append(matrix, []uint8{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48})
	matrix = append(matrix, []uint8{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31})
	matrix = append(matrix, []uint8{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23})
	return matrix
}

func walk(i int, j int, path string, sum int, matrix [][]uint8) string {
	path += fmt.Sprintf(" %d", matrix[i][j])
	sum += int(matrix[i][j])
	if i < len(matrix)-1 {
		walk(i+1, j, path, sum, matrix)
		walk(i+1, j+1, path, sum, matrix)
	} else {
		if sum > max {
			max = sum
		}
		fmt.Println(path, " = ", sum)
	}
	return path
}

func main() {
	matrix := initialize()
	fmt.Println("Problem 18")

	fmt.Println(matrix)
	walk(0, 0, "", 0, matrix)
	fmt.Println("max = ", max)
}
