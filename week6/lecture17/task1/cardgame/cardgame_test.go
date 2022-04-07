package cardgame

import (
	"testing"
)

func TestCompare(t *testing.T) {
	card1 := Card{Value: 3, Suit: 2}
	card2 := Card{Value: 2, Suit: 2}
	c3NotValid := Card{Value: 15, Suit: 4}
	expected, errExp := 1, "invalid input"
	ans, err := CompareCards(card1, card2)

	if ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	ans, err = CompareCards(card2, card1)
	expected, errExp = -1, "invalid input"

	if ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	ans, err = CompareCards(card2, card2)
	expected, errExp = 0, "invalid input"

	if ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	if err != nil {
		t.Errorf("got %s, want nil, %s", err, errExp)
	}

	ans, err = CompareCards(card2, c3NotValid)
	expected, errExp = 0, "invalid input"

	if err == nil {
		t.Errorf("got %d and %s , want error %d %s", ans, err, expected, errExp)
	}

}

func TestMax(t *testing.T) {
	ind1 := 0
	ind2 := 3
	cards := []Card{{2, 3}, {13, 2}, {4, 3}, {5, 1}}
	ans := MaxCard(cards, ind1, ind2)
	expected := Card{13, 2}

	if ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	ind1 = -1
	expected = Card{}
	ans = MaxCard(cards, ind1, ind2)

	if ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

}
