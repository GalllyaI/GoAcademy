package main

import (
	"fmt"
	"goAcademy/week3/lecture9/task1/carddraw"
	"goAcademy/week3/lecture9/task1/cardgame"
	"log"
)

func main() {
	var d *cardgame.Deck
	deck := d.NewDeck()

	result, err := carddraw.DrawAllCards(deck)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}

	var c cardgame.Card
	deck.Cards = append(deck.Cards, *c.NewCard(20, 20))

	result, err = carddraw.DrawAllCards(deck)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}

}
