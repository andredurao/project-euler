/*
Cubic permutations

Problem 62
The cube, 41063625 (345^3), can be permuted to produce two other cubes:
  56623104 (384^3) and 66430125 (405^3).

In fact, 41063625 is the smallest cube which has exactly three permutations
of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

var p = fmt.Println

// uint 64 max 18446744073709551615 (2^64)
// floor((2^64) ^ 1/3) = 2642245
// starting at 1k limit
func smallestCubeOfGroup() uint64 {
	cubes := make(map[string][]uint64)
	i := 0
	for true {
		i += 1
		cube := uint64(i * i * i)
		key := sortedDigits(cube)
		_, found := cubes[key]
		if found {
			cubes[key] = append(cubes[key], cube)
			if len(cubes[key]) == 5 {
				return cubes[key][0]
			}
		} else {
			cubes[key] = make([]uint64, 1)
			cubes[key][0] = cube
		}
	}
	return uint64(0)
}

func sortedDigits(cube uint64) string {
	str := strconv.FormatUint(cube, 10)
	digits := []rune(str)
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})
	return string(digits)
}

func main() {
	p("Problem 62")
	result := smallestCubeOfGroup()
	p(result)
}
