package main

import "fmt"

type deck []string

// newDeck creates a new, standard 52-card deck.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print prints each card in the deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal deals a hand of cards, returning the hand, the remaining deck, and an error if the hand size is invalid.
func deal(d deck, handSize int) (deck, deck, error) {
	if handSize <= 0 || handSize > len(d) {
		return nil, nil, fmt.Errorf("invalid hand size: %d (deck size: %d)", handSize, len(d))
	}
	return d[:handSize], d[handSize:], nil
}
