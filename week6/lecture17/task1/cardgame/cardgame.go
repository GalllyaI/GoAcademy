package cardgame

import (
	"fmt"
)

type Card struct {
	Value int
	Suit  int
}

func CompareCards(cardA Card, cardB Card) (int, error) {

	var result int

	if cardA.Value < 2 || cardA.Value > 13 || cardB.Value < 0 || cardB.Value > 13 || cardA.Suit < 0 || cardA.Suit > 4 || cardB.Suit < 0 || cardB.Suit > 4 {
		err := fmt.Errorf("invalid input")
		return result, err
	}

	if cardA.Value > cardB.Value {
		result = 1
	}

	if cardA.Value < cardB.Value {
		result = -1
	}

	if cardA.Value == cardB.Value {
		if cardA.Suit > cardB.Suit {
			result = 1
		}

		if cardA.Suit < cardB.Suit {
			result = -1
		}

		if cardA.Suit == cardB.Suit {
			result = 0
		}
	}

	return result, nil
}

func MaxCard(cards []Card, idx1 int, idx2 int) Card {

	if idx1 < 0 || idx1 > len(cards) || idx2 < 0 || idx2 > len(cards) {
		fmt.Println("invalid index")
		return Card{}
	}

	cardMax := Card{Value: 2, Suit: 1}

	for i := idx1; i <= idx2; i++ {

		result, err := CompareCards(cardMax, cards[i])

		if err != nil {
			fmt.Println(err.Error())
			return Card{}

		}

		if result == -1 || result == 0 {
			cardMax = cards[i]
		}

	}
	return cardMax
}
