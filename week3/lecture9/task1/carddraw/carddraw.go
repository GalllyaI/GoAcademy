package carddraw

import (
	"goAcademy/week3/lecture9/task1/cardgame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var cards []cardgame.Card
	for {
		card, err := dealer.Deal()

		if err != nil {
			if dealer.Done() {
				return cards, nil
			} else {
				return nil, err
			}
		}
		cards = append(cards, *card)
	}
}
