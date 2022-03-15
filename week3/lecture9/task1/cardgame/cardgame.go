package cardgame

import "errors"

type Card struct {
	value int
	suit  int
}

func (c *Card) NewCard(value int, suit int) *Card {
	return &Card{value, suit}
}

func (c *Card) ValidCheck() (*Card, error) {
	if !(c.value >= 2 && c.value < 15 && c.suit >= 0 && c.suit < 4) {
		return nil, errors.New("invalid card")
	}
	return c, nil
}

type Deck struct {
	Cards []Card
}

func (d *Deck) NewDeck() *Deck {

	var cards []Card

	for i := 2; i <= 3; i++ {
		for j := 0; j < 2; j++ {
			cards = append(cards, Card{i, j})
		}
	}

	return &Deck{Cards: cards}
}

func (d *Deck) Done() bool {

	return len(d.Cards) == 0
}

func (d *Deck) Deal() (*Card, error) {
	var dealCard *Card
	length := len(d.Cards)
	if length == 0 {
		return nil, errors.New("deck is empty")

	}

	dealCard = &d.Cards[0]
	card, err := dealCard.ValidCheck()

	if err != nil {
		return nil, err
	}

	// if err == nil {

	switch {

	case length == 1:
		d.Cards = []Card{}

	case length == 2:
		d.Cards = []Card{d.Cards[1]}

	case length > 2:
		d.Cards = append([]Card{d.Cards[1]}, d.Cards[2:]...)
		// }
	}

	return card, nil
}
