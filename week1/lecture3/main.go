package main

import (
	"fmt"
	cards "lecture3/compareCards"
	"log"
)

func main() {

	const (
		clubs = iota
		diamonds
		heart
		spade
	)

	res, err := cards.CompareCards(4, clubs, 5, clubs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
