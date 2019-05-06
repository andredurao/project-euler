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

// Pos Row, Col
type Pos struct {
	Row int
	Col int
}

var p = fmt.Println

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

func layerItemsCount(layer int) int {
	if layer == 1 {
		return 1
	}
	rows := (2 * layer) - 1
	return 4*(rows-2) + 4
}

func initialValue(layer int) int {
	if layer == 1 {
		return 1
	}
	count := 0
	for i := 1; i < layer; i++ {
		count += layerItemsCount(i)
	}
	return count + 1
}

// initialPos = right from to the top right element from PREVIOUS layer
func initialPos(matrix [][]int, layer int) (pos Pos) {
	center := len(matrix) / 2
	if layer == 1 {
		pos = Pos{center, center}
	} else {
		pos.Col = center + layer - 1
		pos.Row = center - layer + 2
		//[77] 1->(3,3), 2->(3,4), 3->(2,5), 4->(1,6)
		//[99] 1->(4,4), 2->(4,5), 3->(3,6), 4->(2,7), 5->(1,8)
	}
	return
}

func fillSpiralLayer(matrix *[][]int, layer int) {
	cursor := initialPos(*matrix, layer)
	value := initialValue(layer)
	center := len(*matrix) / 2
	if layer == 1 {
		(*matrix)[cursor.Row][cursor.Col] = value
	} else {
		size := layer*2 - 1
		// walk down from initialPos
		bottomRow := center + size/2
		for i := cursor.Row; i <= bottomRow; i++ {
			cursor.Row = i
			(*matrix)[cursor.Row][cursor.Col] = value
			value++
		}
		// walk left from bottomPos
		leftCol := cursor.Col - size
		for i := cursor.Col - 1; i > leftCol; i-- {
			cursor.Col = i
			(*matrix)[cursor.Row][cursor.Col] = value
			value++
		}
		// walk up from leftPos
		topRow := cursor.Row - size
		for i := cursor.Row - 1; i > topRow; i-- {
			cursor.Row = i
			(*matrix)[cursor.Row][cursor.Col] = value
			value++
		}
		// walk right from topRow
		rightCol := cursor.Col + size
		for i := cursor.Col + 1; i < rightCol; i++ {
			cursor.Col = i
			(*matrix)[cursor.Row][cursor.Col] = value
			value++
		}
	}
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		p(matrix[i])
	}
}

func main() {
	p("Problem 28")
	size := 1001
	matrix, err := generateMatrix(size)
	if err != nil {
		p(err)
	}
	for i := 1; i <= (size+1)/2; i++ {
		fillSpiralLayer(&matrix, i)
	}
	// printMatrix(matrix)
	sum := 0
	for i := 0; i < size; i++ {
		sum += matrix[i][i]
		sum += matrix[i][size-1-i]
	}
	p(sum - 1) // -1 because the center (1) exists in both diagonals
}
