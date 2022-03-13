package cardgame

type Card struct {
	value int
	suit  int
}

type Deck struct {
	cards []Card
}

func (d *Deck) NewDeck() *Deck {

	var cards []Card

	for i := 2; i <= 14; i++ {
		for j := 0; j < 4; j++ {
			cards = append(cards, Card{i, j})
		}
	}

	return &Deck{cards: cards}
}

func (d *Deck) Deal() *Card {
	var dealCard *Card
	length := len(d.cards)

	switch length {

	case 0:
		return dealCard

	case 1:
		dealCard = &d.cards[0]
		d.cards = []Card{}
		return dealCard

	case 2:
		dealCard = &d.cards[0]
		d.cards = []Card{d.cards[1]}
		return dealCard

	default:
		dealCard = &d.cards[0]
		d.cards = append([]Card{d.cards[1]}, d.cards[2:]...)
		return dealCard
	}
}
