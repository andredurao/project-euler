// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//What is the sum of the digits of the number 2^1000?
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func sum(v string) string {
	c, _ := strconv.Atoi(v)
	zero := new(big.Int)
	ten := new(big.Int).SetUint64(10)
	exp := new(big.Int)
	total := new(big.Int)
	modulus := new(big.Int)
	x := new(big.Int).SetUint64(2)
	y := new(big.Int).SetUint64(uint64(c))
	exp = exp.Exp(x, y, nil)
	for exp.Cmp(zero) == 1 {
		exp.DivMod(exp, ten, modulus)
		total.Add(total, modulus)
	}
	return total.String()
}

func main() {
	fmt.Println("Problem 16")
	result := sum(os.Args[1])
	fmt.Println(result)
}
