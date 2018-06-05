// Problem 15
// Starting in the top left corner of a 2×2 grid, and only being able to move to the
// right and down, there are exactly 6 routes to the bottom right corner.
//
// ━━┓ ━┓   ━┓  ╻    ╻     ╻
//   ┃  ┗━┓  ┃  ┗━━┓ ┗━┓   ┃
//   ╹    ╹  ┗━    ╹   ┗━━ ┗━━╸
//
// How many such routes are there through a 20×20 grid?
package main

import (
	"fmt"
)

const GridSize = 20

// https://en.wikipedia.org/wiki/Hamming_weight
func popcount(x uint64) int {
	const (
		m1  = 0x5555555555555555
		m2  = 0x3333333333333333
		m4  = 0x0f0f0f0f0f0f0f0f
		h01 = 0x0101010101010101
	)
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return int((x * h01) >> 56)
}

func main() {
	fmt.Println("Problem 15")
	// The amount of edges is equals 2x the amount of vertices to create a path from
	// source to target
	// I know already that every "balanced path" reaches it's destination from before:
	// https://stackoverflow.com/questions/3522417/binary-numbers-with-the-same-quantity-of-0s-and-1s
	limit := uint64(1 << (GridSize * 2))
	var i, count uint64 = 0, 0
	fmt.Println(limit)

	for i < limit {
		// TODO: use concurrence to improve performance
		if popcount(i) == GridSize {
			count++
		}
		i++
	}
	fmt.Println(count)
}
