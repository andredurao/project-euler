// Coded triangle numbers
// Problem 42
// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
// so the first ten triangle numbers are:
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
// By converting each letter in a word to a number corresponding to its
// alphabetical position and adding these values we form a word value.
// For example, the word value for SKY is 19 + 11 + 25 = 55 = t(10).
// If the word value is a triangle number then that word is a triangle word.
// Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
// containing nearly two-thousand common English words,
// how many are triangle words?
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var p = fmt.Println

func loadWords(filename string) []string {
	fileContent, _ := ioutil.ReadFile(filename)
	content := string(bytes.Replace(fileContent, []byte(`"`), []byte(``), -1))
	list := strings.Split(content, ",")
	sort.Sort(sort.StringSlice(list))
	return list
}

func roots(a int, b int, c int) (solutions []float64) {
	delta := (b * b) - 4*(a*c)
	x1 := ((float64(b) * -1.0) + math.Sqrt(float64(delta))) / float64(2*a)
	x2 := ((float64(b) * -1.0) - math.Sqrt(float64(delta))) / float64(2*a)
	solutions = append(solutions, x1)
	solutions = append(solutions, x2)

	return
}

func isTriangleNumber(n int) bool {
	solutions := roots(1, 1, (2 * n * -1))
	for _, value := range solutions {
		if value > 0 && isIntegral(value) {
			return true
		}
	}
	return false
}

func isIntegral(value float64) bool {
	return value == float64(int(value))
}

func charactersSum(word string) (sum int) {
	for _, char := range word {
		sum += int(char) - 'A' + 1
	}
	return
}

func main() {
	p("Problem 42")

	total := 0
	words := loadWords("p042_words.txt")
	for _, word := range words {
		sum := charactersSum(word)
		if isTriangleNumber(sum) {
			total++
		}
	}
	p(total)
}
