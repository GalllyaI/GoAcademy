package main

import (
	"fmt"
	"math/rand"
	"time"
)

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}

func groupSlices(cities []string, prices []int) map[string][]int {

	result := make(map[string][]int)

	for ind, city := range cities {
		if result[city] == nil {
			result[city] = append(result[city], prices[ind])
		} else {
			result[city] = append(result[city], prices[ind])
		}

	}

	return result
}

func main() {

	cities, prices := citiesAndPrices()
	// cities := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London", "Berlin", "Moscow", "Chicago", "Tokyo", "London", "Berlin", "Moscow", "Chicago"}
	// prices := []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 11, 22, 33}
	output := groupSlices(cities, prices)

	fmt.Println(output)
}
