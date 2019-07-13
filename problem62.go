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
	"strconv"
)

var p = fmt.Println

// uint 64 max 18446744073709551615 (2^64)
// floor((2^64) ^ 1/3) = 2642245
// starting at 1k limit
func genCubes() map[string]struct{} {
	cubes := make(map[string]struct{})
	for i := 1; i <= 1000000; i++ {
		ui := uint64(i)
		ui = ui * ui * ui
		cubes[strconv.FormatUint(ui, 10)] = struct{}{}
	}
	return cubes
}

func generatePermutations(array []rune, n int, permutations *[]string) {
	if n == 1 {
		*permutations = append(*permutations, string(array))
	} else {
		for i := 0; i < n; i++ {
			generatePermutations(array, n-1, permutations)
			if n%2 == 0 {
				array[0], array[n-1] = array[n-1], array[0]
			} else {
				array[i], array[n-1] = array[n-1], array[i]
			}
		}
	}
}

func countPermutations(cubes map[string]struct{}, permutations []string) (count int) {
	for _, value := range permutations {
		_, found := cubes[value]
		if found {
			count++
		}
	}
	return
}

func cleanUpPermutations(cubes map[string]struct{}, permutations []string) {
	for _, value := range permutations {
		delete(cubes, value)
	}
}

func seekPermutations(cubes map[string]struct{}) {
	for cube := range cubes {
		permutations := make([]string, 0)
		digits := []rune(cube)
		generatePermutations(digits, len(digits), &permutations)
		count := countPermutations(cubes, permutations)
		if count != 5 {
			cleanUpPermutations(cubes, permutations)
		} else {
			p(permutations)
			break
		}
	}
}

func main() {
	p("Problem 62")
	cubes := genCubes()
	p(len(cubes))
	seekPermutations(cubes)
	p(len(cubes))
}
