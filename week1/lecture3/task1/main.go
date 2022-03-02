package main

import (
	"errors"
	"fmt"
	"log"
)

func CompareCards(cardAVal int, cardASuit int, cardBVal int, cardBSuit int) (int, error) {

	var result int

	if cardAVal < 2 || cardAVal > 13 || cardBVal < 0 || cardBVal > 13 || cardASuit < 0 || cardASuit > 4 || cardBSuit < 0 || cardBSuit > 4 {

		return result, errors.New("Invalid input")
	}

	if cardAVal > cardBVal {
		result = 1
	}

	if cardAVal < cardBVal {
		result = -1
	}

	if cardAVal == cardBVal {
		if cardASuit > cardBSuit {
			result = 1
		}

		if cardASuit < cardBSuit {
			result = -1
		}

		if cardASuit == cardBSuit {
			result = 0
		}
	}

	return result, nil
}

func main() {

	const (
		clubs = iota
		diamonds
		heart
		spade
	)

	res, err := CompareCards(4, clubs, 5, clubs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
