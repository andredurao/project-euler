// Spiral primes

package main

import (
	"fmt"
	"math"
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
		return matrix, fmt.Errorf("Size %d must be an odd value", size)
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
	return 4 * (rows - 1)
}

func spiralInitialValue(layer int) int {
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
		pos.Row = center + layer - 2
	}
	return
}

func fillSpiralLayer(matrix *[][]int, layer int) {
	cursor := initialPos(*matrix, layer)
	value := spiralInitialValue(layer)
	center := len(*matrix) / 2
	if layer == 1 {
		(*matrix)[cursor.Row][cursor.Col] = value
	} else {
		size := layer*2 - 1
		// walk up from leftPos
		topRow := center - size/2
		for i := cursor.Row; i >= topRow; i-- {
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
		// walk down from initialPos
		bottomRow := cursor.Row + size
		for i := cursor.Row + 1; i < bottomRow; i++ {
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

func isPrime(n int) bool {
	max := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i = 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// returns the ratio of prime numbers in diagonals
func diagonalRatio(matrix [][]int, primesCache map[int]struct{}) float64 {
	size := len(matrix)
	count := (size * 2) - 1
	primesCount := 0
	for i := 0; i < size; i++ {
		// main diagonal
		diagonalValue := matrix[i][i]
		_, found := primesCache[diagonalValue]
		if found {
			primesCount++
		} else if isPrime(diagonalValue) {
			primesCache[diagonalValue] = struct{}{}
			primesCount++
		}
		// antidiagonal except center which is tested above
		antidiagonalValue := matrix[i][size-1-i]
		if antidiagonalValue != 1 {
			_, found := primesCache[antidiagonalValue]
			if found {
				primesCount++
			} else if isPrime(antidiagonalValue) {
				primesCache[antidiagonalValue] = struct{}{}
				primesCount++
			}
		}
	}
	return float64(primesCount) / float64(count)
}

func sampleMatrix(size int) (ratio float64) {
	primesCache := make(map[int]struct{})
	matrix, err := generateMatrix(size)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= (size+1)/2; i++ {
		fillSpiralLayer(&matrix, i)
	}
	printMatrix(matrix)
	ratio = diagonalRatio(matrix, primesCache)
	return
}

func initialValue(layer int) (start int) {
	if layer == 1 {
		return 1
	}
	innerArea := (layer - 2)
	innerArea *= innerArea
	return innerArea + (layer - 1)
}

func corners(layer int) (values []int) {
	start := initialValue(layer)
	delta := 0
	for i := 0; i < 4; i++ {
		values = append(values, start+delta)
		delta += layer - 1
	}
	return
}

func main() {
	p("Problem 58")
	ratio := 9.99
	itemsCount := 1
	size := 3
	primesCount := 0
	for ratio > 0.1000000 {
		newCorners := corners(size)
		// the last element should never been a prime
		// because it contains the area of the matrix (size²)
		for i := 0; i < 3; i++ {
			if isPrime(newCorners[i]) {
				primesCount++
			}
		}

		itemsCount += 4
		ratio = float64(primesCount) / float64(itemsCount)
		size += 2
	}
	size -= 2
	p(size)
}
