package main

import (
	"fmt"
	"goAcademy/week3/lecture8/task1/carddraw"
	"goAcademy/week3/lecture8/task1/cardgame"
)

func main() {
	var d *cardgame.Deck
	deck := d.NewDeck()

	result := carddraw.DrawAllCards(deck)
	fmt.Println(result)
}
