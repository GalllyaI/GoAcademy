package main

import (
	"errors"
	"fmt"
)

type Card struct {
	value int
	suit  int
}

func CompareCards(cardA Card, cardB Card) (int, error) {

	var result int

	if cardA.value < 2 || cardA.value > 13 || cardB.value < 0 || cardB.value > 13 || cardA.suit < 0 || cardA.suit > 4 || cardB.suit < 0 || cardB.suit > 4 {

		return result, errors.New("Invalid input")
	}

	if cardA.value > cardB.value {
		result = 1
	}

	if cardA.value < cardB.value {
		result = -1
	}

	if cardA.value == cardB.value {
		if cardA.suit > cardB.suit {
			result = 1
		}

		if cardA.suit < cardB.suit {
			result = -1
		}

		if cardA.suit == cardB.suit {
			result = 0
		}
	}

	return result, nil
}

func main() {

	cardOne := Card{7, 2}
	cardTwo := Card{7, 2}
	cardThree := Card{3, 2}
	cardFour := Card{13, 2}
	cardFive := Card{13, 4}
	cardNotValid1 := Card{15, 4}
	cardNotValid2 := Card{5, -4}

	fmt.Println(CompareCards(cardOne, cardTwo))
	fmt.Println(CompareCards(cardTwo, cardThree))
	fmt.Println(CompareCards(cardThree, cardFour))
	fmt.Println(CompareCards(cardFour, cardFive))
	fmt.Println(CompareCards(cardFour, cardNotValid1))
	fmt.Println(CompareCards(cardFour, cardNotValid2))

}
