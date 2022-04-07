package main

import (
	"fmt"
	"goAcademy/week6/lecture17/task1/cardgame"
)

func main() {

	c0 := cardgame.Card{Value: 2, Suit: 4}
	c1 := cardgame.Card{Value: 2, Suit: 4}
	c2 := cardgame.Card{Value: 13, Suit: 4}
	c3NotValid := cardgame.Card{Value: 15, Suit: 4}
	c4 := cardgame.Card{Value: 6, Suit: 4}
	c5 := cardgame.Card{Value: 9, Suit: 2}
	c6 := cardgame.Card{Value: 11, Suit: 1}
	c7 := cardgame.Card{Value: 4, Suit: 2}
	c8 := cardgame.Card{Value: 12, Suit: 4}
	c9 := cardgame.Card{Value: 13, Suit: 4}
	c10 := cardgame.Card{Value: 3, Suit: 4}
	c11 := cardgame.Card{Value: 10, Suit: 2}
	c12 := cardgame.Card{Value: 5, Suit: 3}
	c13 := cardgame.Card{Value: 5, Suit: 3}
	c14 := cardgame.Card{Value: 7, Suit: 1}
	c15NotValid := cardgame.Card{Value: 5, Suit: -4}
	c16 := cardgame.Card{Value: 3, Suit: 1}

	cards := []cardgame.Card{c0, c1, c2, c3NotValid, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15NotValid, c16}

	max := cardgame.MaxCard(cards, 7, -14)

	fmt.Println(max)
}
