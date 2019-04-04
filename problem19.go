// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Problem 19")
	fmt.Println(seekUsingAPI())
	fmt.Println(seekManually())
}

// monday = 0 ... sunday = 6
func isSunday(day int) bool {
	return day%7 == 6
}

func seekManually() int {
	day, count := 0, 0
	for year := 1900; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			if isSunday(day) && year > 1900 {
				count++
			}
			day += daysInMonth(year, month)
		}
	}
	return count
}

func daysInMonth(year int, month int) int {
	var daysMap = map[int]int{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30,
		10: 31, 11: 30, 12: 31}
	if month == 2 && ((year%4 == 0 && year%100 != 0) || (year%400 == 0)) {
		return 29
	} else {
		return daysMap[month]
	}
}

func seekUsingAPI() int {
	count := 0
	for year := 1901; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			// Using golang time API
			date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				count++
			}
		}
	}
	return count
}
