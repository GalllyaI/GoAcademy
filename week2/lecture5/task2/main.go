package main

import (
	"fmt"
	"task2/card"
	"task2/compareCards"
)

// type Card struct {
// 	Value int
// 	Suit  int
// }

func main() {
	c0 := card.Card{Value: 2, Suit: 4}
	c1 := card.Card{Value: 2, Suit: 4}
	c2 := card.Card{Value: 13, Suit: 4}
	c3 := card.Card{Value: 3, Suit: 3}
	c4 := card.Card{Value: 6, Suit: 4}
	c5 := card.Card{Value: 9, Suit: 2}
	c6 := card.Card{Value: 11, Suit: 1}
	c7 := card.Card{Value: 4, Suit: 2}
	c8 := card.Card{Value: 12, Suit: 4}
	c9 := card.Card{Value: 13, Suit: 4}
	c10 := card.Card{Value: 3, Suit: 4}
	c11 := card.Card{Value: 10, Suit: 2}
	c12 := card.Card{Value: 5, Suit: 3}
	c13 := card.Card{Value: 5, Suit: 3}
	c14 := card.Card{Value: 7, Suit: 1}
	c15 := card.Card{Value: 3, Suit: 1}

	cards := []card.Card{c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15}

	//please change the values for idx1 and idx2 for different inputs

	idx1 := 11
	idx2 := 15

	cardMax := card.Card{Value: 2, Suit: 1}

	for i := idx1; i < idx2; i++ {

		result := compareCards.Compare(cardMax, cards[i])

		if result == -1 || result == 0 {
			cardMax = cards[i]
		}

	}

	fmt.Println(cardMax)

}
