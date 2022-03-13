package carddraw

import "goAcademy/week3/lecture8/task1/cardgame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	var cards []cardgame.Card

	for {
		card := dealer.Deal()
		if card != nil {
			cards = append(cards, *card)
		} else {
			break
		}
	}
	return cards
}
