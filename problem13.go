// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func ScanFile() *big.Int {
	file, _ := os.Open("problem13numbers.txt")
	fileScanner := bufio.NewScanner(file)
	total := big.NewInt(0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		current := big.NewInt(0)
		current.SetString(line, 10)
		total.Add(total, current)
	}
	return total
}

func main() {
	fmt.Println("Problem 13")
	total := ScanFile()
	fmt.Println(total)
	strTotal := total.String()
	fmt.Println(strTotal[:10])
}
