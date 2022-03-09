package main

import "fmt"

type Card struct {
	value int
	suit  int
}

type CardComparator func(cOne Card, cTwo Card) int

func CompareCards(cardA Card, cardB Card) int {

	var result int

	if cardA.value > cardB.value {
		result = 1
	}

	if cardA.value < cardB.value {
		result = -1
	}

	if cardA.value == cardB.value {
		if cardA.suit > cardB.suit {
			result = 1
		}

		if cardA.suit < cardB.suit {
			result = -1
		}

		if cardA.suit == cardB.suit {
			result = 0
		}
	}

	return result
}

func maxCard(cards []Card, comparatorFunc CardComparator) Card {

	max := cards[0]

	for ind := range cards {
		if comparatorFunc(max, cards[ind]) == -1 {
			max = cards[ind]
		}
	}
	return max
}

// var anfunc  = func(cards []Card, comparatorFunc CardComparator) Card {

// }

func main() {

	cards := []Card{{7, 2}, {7, 2}, {3, 2}, {13, 2}, {13, 4}, {10, 4}, {14, 4}}

	mxCard := maxCard(cards, CompareCards)

	fmt.Println(mxCard)
	fmt.Println("---------")

	func() {
		fmt.Println(maxCard(cards, CompareCards))
	}()

}
