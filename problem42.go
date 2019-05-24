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

func main() {
	p("Problem 42")
	words := loadWords("p042_words.txt")
	p(words)
}
