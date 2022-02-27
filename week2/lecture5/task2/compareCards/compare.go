package compareCards

import "task2/card"

// type Card struct {
// 	Value int
// 	Suit  int
// }

func Compare(cardA card.Card, cardB card.Card) int {

	var result int

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

	return result
}
