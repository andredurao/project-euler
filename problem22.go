// Using names.txt, a 46K text file containing over five-thousand first names,
// begin by sorting it into alphabetical order. Then working out the alphabetical value
// for each name, multiply this value by its alphabetical position in the list to obtain
// a name score.

// For example, when the list is sorted into alphabetical order, COLIN,
// which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
// So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func loadNames(filename string) []string {
	fileContent, _ := ioutil.ReadFile(filename)
	content := string(bytes.Replace(fileContent, []byte(`"`), []byte(``), -1))
	list := strings.Split(content, ",")
	sort.Sort(sort.StringSlice(list))
	return list
}

func main() {
	var total uint64 = 0
	fmt.Println("Problem 22")
	namesList := loadNames("p022_names.txt")

	for i, v := range namesList {
		index := i + 1
		total += uint64(index * charScore(v))
	}
	fmt.Println(total)
}

func charScore(value string) int {
	score := 0
	for _, v := range value {
		score += int(v - 'A' + 1)
	}
	return score
}
