package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) ([]string, error) {

	var dateList []time.Time
	var sortedDates []string

	for _, date := range dates {
		d, err := time.Parse(format, date)

		if err != nil {
			return nil, errors.New("invalid date")

		}
		dateList = append(dateList, d)
	}

	sort.Slice(dateList, func(i, j int) bool {
		return dateList[i].Before(dateList[j])
	})

	for _, d := range dateList {
		sortedDates = append(sortedDates, d.Format("Jan-02-2006"))
	}

	return sortedDates, nil
}

func main() {
	dates := []string{"Dec-03-2021", "Mar-18-2022", "Sep-14-2008"}
	format := "Jan-02-2006"
	dateList, _ := sortDates(format, dates...)

	fmt.Println(dateList)

}
