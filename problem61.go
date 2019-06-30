/*
Cyclical figurate numbers

Triangle, square, pentagonal, hexagonal, heptagonal, and octagonal numbers are all figurate (polygonal) numbers and are generated by the following formulae:

Triangle   P3,n=n(n+1)/2   1, 3, 6,  10, 15, ...
Square     P4,n=n*n        1, 4, 9,  16, 25, ...
Pentagonal P5,n=n(3n−1)/2  1, 5, 12, 22, 35, ...
Hexagonal  P6,n=n(2n−1)    1, 6, 15, 28, 45, ...
Heptagonal P7,n=n(5n−3)/2  1, 7, 18, 34, 55, ...
Octagonal  P8,n=n(3n−2)    1, 8, 21, 40, 65, ...
The ordered set of three 4-digit numbers: 8128, 2882, 8281, has three interesting properties.

The set is cyclic, in that the last two digits of each number is the first two digits of the next number (including the last number with the first).
Each polygonal type: triangle (P(3,127)=8128), square (P(4,91)=8281), and pentagonal (P(5,44)=2882), is represented by a different number in the set.
This is the only set of 4-digit numbers with this property.
Find the sum of the only ordered set of six cyclic 4-digit numbers for which each polygonal type: triangle, square, pentagonal, hexagonal, heptagonal, and octagonal, is represented by a different number in the set.
*/

package main

import (
	"fmt"
)

var p = fmt.Println

type PolygonalNumber struct {
	value   int // the number itself
	polygon int // the polygon type (number of edges)
}

// ValuesSet will be composed by of parts of a number
type PolygonalValuesSet struct {
	prefixes []*PolygonalNumber
	suffixes []*PolygonalNumber
}

func triangle(n int) int {
	return (n * (n + 1)) / 2
}

func square(n int) int {
	return n * n
}

func pentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func hexagonal(n int) int {
	return (n * (2*n - 1))
}

func heptagonal(n int) int {
	return (n * (5*n - 3)) / 2
}

func octagonal(n int) int {
	return (n * (3*n - 2))
}

func generateTriangleNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < ((x (x + 1))/2) < 10000
	// positive values => (-1+sqrt(7993))/2 < x < (-1+sqrt(8889)*3)/2
	for i := 45; i < 141; i++ {
		values = append(values, triangle(i))
	}
	return
}

func generateSquareNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < x^2 < 10000
	// positive values => sqrt(999) < x < sqrt(10000)
	for i := 32; i < 100; i++ {
		values = append(values, square(i))
	}
	return
}

func generatePentagonalNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < x(3x−1)/2 < 10000
	// positive values => (1+sqrt(23977))/6 < x < (1+sqrt(240001))/6
	for i := 26; i < 82; i++ {
		values = append(values, pentagonal(i))
	}
	return
}

func generateHexagonalNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < 2x^2 - x < 10000
	// positive values => (1+sqrt(7993))/4 < x < (1+3*sqrt(8889))/4
	for i := 23; i < 71; i++ {
		values = append(values, hexagonal(i))
	}
	return
}

func generateHeptagonalNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < x(5x−3)/2 < 10000
	// positive values => (3+3*sqrt(4441))/10 < x < (3+sqrt(400009))/10
	for i := 21; i < 64; i++ {
		values = append(values, heptagonal(i))
	}
	return
}

func generateOctagonalNumbers() (values []int) {
	// the 4 digit values should be in the range: 999 < 3x^2 - 2x < 10000
	// positive values => (1+sqrt(2998))/3 < x < (1+sqrt(30001))/3
	for i := 19; i < 59; i++ {
		values = append(values, octagonal(i))
	}
	return
}

func split(value int) (prefix, suffix int) {
	prefix = value / 100
	suffix = value % 100
	return
}

func populatePolygonalNumbers() (map[int]*PolygonalValuesSet, []*PolygonalNumber) {
	functions := make(map[int]func() []int)
	functions[3] = generateTriangleNumbers
	functions[4] = generateSquareNumbers
	functions[5] = generatePentagonalNumbers
	functions[6] = generateHexagonalNumbers
	functions[7] = generateHeptagonalNumbers
	functions[8] = generateOctagonalNumbers

	result := make(map[int]*PolygonalValuesSet)
	values := make([]*PolygonalNumber, 0)

	for polygonValue, polygonalFunction := range functions {
		for _, value := range polygonalFunction() {
			polygonalNumber := &PolygonalNumber{value, polygonValue}
			values = append(values, polygonalNumber)
			prefix, suffix := split(value)
			_, found := result[prefix]
			if !found {
				valuesSet := &PolygonalValuesSet{}
				valuesSet.prefixes = append(valuesSet.prefixes, polygonalNumber)
				result[prefix] = valuesSet
			} else {
				valuesSet := result[prefix]
				valuesSet.prefixes = append(valuesSet.prefixes, polygonalNumber)
			}
			_, found = result[suffix]
			if !found {
				valuesSet := &PolygonalValuesSet{}
				valuesSet.suffixes = append(valuesSet.suffixes, polygonalNumber)
				result[suffix] = valuesSet
			} else {
				valuesSet := result[suffix]
				valuesSet.suffixes = append(valuesSet.suffixes, polygonalNumber)
			}
		}
	}
	return result, values
}

func seekSiblings(path []*PolygonalNumber, polygonalValues map[int]*PolygonalValuesSet, values []*PolygonalNumber) {
	item := path[len(path)-1]
	_, suffix := split(item.value)
	siblings := polygonalValues[suffix]
	polygonsSet := make(map[int]struct{})
	for _, currItem := range path {
		polygonsSet[currItem.polygon] = struct{}{}
	}

	for _, sibling := range siblings.prefixes {
		_, found := polygonsSet[sibling.polygon]
		if !found {
			newPath := append(path, sibling)
			if len(newPath) == 6 {
				// check first and last items
				lastItem := newPath[len(newPath)-1]
				_, lastSuffix := split(lastItem.value)
				firstItem := newPath[0]
				firstPrefix, _ := split(firstItem.value)
				if lastSuffix == firstPrefix {
					p("--")
					for _, polygonalNumber := range newPath {
						p(polygonalNumber.value, polygonalNumber.polygon)
					}
				}
			} else {
				seekSiblings(append(path, sibling), polygonalValues, values)
			}
		}
	}
}

func main() {
	p("Problem 61")
	// All the printed sequences where items of the same, because it's cyclical
	polygonalValues, values := populatePolygonalNumbers()

	for _, polygonalNumber := range values {
		path := make([]*PolygonalNumber, 0)
		seekSiblings(
			append(path, polygonalNumber),
			polygonalValues,
			values,
		)
	}
}
