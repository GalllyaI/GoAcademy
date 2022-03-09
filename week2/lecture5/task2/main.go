package main

import "fmt"

type Card struct {
	Value int
	Suit  int
}

func Compare(cardA Card, cardB Card) int {

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

func main() {
	c0 := Card{Value: 2, Suit: 4}
	c1 := Card{Value: 2, Suit: 4}
	c2 := Card{Value: 13, Suit: 4}
	c3 := Card{Value: 3, Suit: 3}
	c4 := Card{Value: 6, Suit: 4}
	c5 := Card{Value: 9, Suit: 2}
	c6 := Card{Value: 11, Suit: 1}
	c7 := Card{Value: 4, Suit: 2}
	c8 := Card{Value: 12, Suit: 4}
	c9 := Card{Value: 13, Suit: 4}
	c10 := Card{Value: 3, Suit: 4}
	c11 := Card{Value: 10, Suit: 2}
	c12 := Card{Value: 5, Suit: 3}
	c13 := Card{Value: 5, Suit: 3}
	c14 := Card{Value: 7, Suit: 1}
	c15 := Card{Value: 3, Suit: 1}

	cards := []Card{c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15}

	//please change the values for idx1 and idx2 for different inputs

	idx1 := 11
	idx2 := 15

	cardMax := Card{Value: 2, Suit: 1}

	for i := idx1; i < idx2; i++ {

		result := Compare(cardMax, cards[i])

		if result == -1 || result == 0 {
			cardMax = cards[i]
		}

	}

	fmt.Println(cardMax)

}
