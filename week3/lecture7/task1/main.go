package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	value int
	suit  int
}

type Deck struct {
	cards []card
}

func NewDeck() *Deck {

	var cards []card

	for i := 2; i <= 3; i++ {
		for j := 0; j < 4; j++ {
			cards = append(cards, card{i, j})
		}
	}

	return &Deck{cards: cards}
}

func (d *Deck) Shuffle() *Deck {

	if len(d.cards) == 0 {
		fmt.Println("Deck is empty")
		return nil
	}

	for i := 0; i < len(d.cards); i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(d.cards) - 1
		j := rand.Intn(max-min+1) + min
		temp := d.cards[i]
		d.cards[i] = d.cards[j]
		d.cards[j] = temp
	}

	return d
}

func (d *Deck) Deal() card {

	dealCard := card{}
	length := len(d.cards)

	switch length {

	case 0:
		fmt.Println("Deck is empty")
		return dealCard

	case 1:
		dealCard = d.cards[0]
		d.cards = []card{}
		return dealCard

	case 2:
		dealCard = d.cards[0]
		d.cards = []card{d.cards[1]}
		return dealCard

	default:
		dealCard = d.cards[0]
		d.cards = append([]card{d.cards[1]}, d.cards[2:]...)
		return dealCard
	}
}

func main() {

	deck := NewDeck()

	fmt.Println(deck.cards)
	deck.Shuffle()
	fmt.Println(deck.cards)
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Shuffle()
	fmt.Println(deck.cards)
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Deal()
	deck.Shuffle()
}
