package blackjack

import (
	"strings"
)

type Deck struct {
	Cards map[string]Card
}
type Card struct {
	Value int
}

// Make deck (non-suited) with key:value.
func MakeDeck() Deck {
	d := Deck{}
	d.Cards = make(map[string]Card)
	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for i, card := range types {
		switch card {
		case "Ace":
			i = 11
		case "Jack", "Queen", "King":
			i = 10
		default:
			i += 2
		}
		d.Cards[card] = Card{Value: i}
	}
	return d
}

// ParseCard returns the integer value of a card according to rules of BlackJack
func ParseCard(card string) int {
	d := MakeDeck()
	// Force card title to lowercase
	return d.Cards[strings.Title(strings.ToLower(card))].Value
}

// Return BlackJack if player cards total to 21, else false.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack && dealerScore < 10:
		return "W"
	case isBlackjack && dealerScore >= 10:
		return "S"
	default:
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17 || handScore >= 12 && handScore <= 16 && dealerScore <= 6:
		return "S"
	default:
		return "H"
	}
}
