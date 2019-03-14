// Problem 15
// Starting in the top left corner of a 2×2 grid, and only being able to move to the
// right and down, there are exactly 6 routes to the bottom right corner.
//
// ━━┓ ━┓   ━┓  ╻    ╻     ╻
//   ┃  ┗━┓  ┃  ┗━━┓ ┗━┓   ┃
//   ╹    ╹  ┗━    ╹   ┗━━ ┗━━╸
//
// How many such routes are there through a 20×20 grid?
// A nice explanation: http://code.jasonbhill.com/python/project-euler-problem-15/
package main

import (
	"fmt"
)

const GridSize = 20

func main() {
	fmt.Println("Problem 15")

	// initialize grid
	grid := [GridSize + 1][GridSize + 1]int64{}

	//initialize borders
	for i := 0; i <= GridSize; i++ {
		grid[0][i] = 1
		grid[i][0] = 1
	}

	// Iterate through the grid elements
	for i := 1; i <= GridSize; i++ {
		for j := 1; j <= GridSize; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	//The Result is the target position
	fmt.Println(grid[GridSize][GridSize])
}
