// n! means n × (n − 1) × ... × 3 × 2 × 1
// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
// Find the sum of the digits in the number 100!

package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func factorial(n *big.Int) *big.Int {
	total := big.NewInt(1)
	for n.Cmp(big.NewInt(0)) > 0 {
		total = total.Mul(total, n)
		n = n.Sub(n, big.NewInt(1))
	}
	return total
}

func sum(n *big.Int) *big.Int {
	total := big.NewInt(0)
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	modulus := big.NewInt(0)
	for n.Cmp(zero) > 0 {
		n.DivMod(n, ten, modulus)
		total.Add(total, modulus)
	}
	return total
}

func main() {
	fmt.Println("Problem 20")
	arg, _ := strconv.Atoi(os.Args[1])
	n := big.NewInt(int64(arg))
	factorialResult := factorial(n)
	fmt.Println("factorial = ", factorialResult)
	sumResult := sum(factorialResult)
	fmt.Println("sum of digits = ", sumResult)
}
