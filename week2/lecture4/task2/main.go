package main

import (
	"fmt"
	"strconv"
)

func checkLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false

}

func isInputValid(month int, year int) bool {

	valMonth := month > 0 && month < 13

	valYear := false

	if len(strconv.Itoa(year)) == 4 {
		valYear = true
	}

	if valMonth && valYear {
		return true
	} else {
		return false
	}
}

func daysInMonth(month int, year int) (int, bool) {

	if !isInputValid(month, year) {
		return 0, false
	}

	switch month {
	case 2:
		leap := checkLeapYear(year)

		if leap {
			return 29, true
		} else {
			return 28, true
		}

	case 1, 3, 5, 7, 8, 10, 12:

		return 31, true

	case 4, 6, 9, 11:

		return 30, true

	default:
		return 0, false
	}
}

func main() {
	fmt.Println(daysInMonth(2, 2000))
	fmt.Println(daysInMonth(2, 1900))

	fmt.Println(daysInMonth(11, 1900))
	fmt.Println(daysInMonth(5, 2022))

	fmt.Println(daysInMonth(8, 20000))
	fmt.Println(daysInMonth(8, 200))
	fmt.Println(daysInMonth(88, 2000))
	fmt.Println(daysInMonth(88, 20000))
}
