package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"strings"
)

type deck []string

const sep = ","

// newDeck creates a new, standard 52-card deck.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

// toString converts a deck to a string.
func (d deck) toString() string {
	return strings.Join([]string(d), sep)
}

// saveToFile saves a deck to a local file.
func (d deck) saveToFile(filename string) error {
	dBytes := []byte(d.toString())
	allAccessPerm := fs.FileMode(0666)
	return os.WriteFile(filename, dBytes, allAccessPerm)
}

// newDeckFromFile creates a new deck from a local file, exits the program if attempt fails.
func newDeckFromFile(filename string) deck {
	dBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dString := string(dBytes) // i.e. "Ace of Clubs,Two of Clubs,Three of Clubs,..."
	dSlice := strings.Split(dString, sep)
	return deck(dSlice)
}

// shuffle modifies the deck in place by shuffling the cards.
func (d deck) shuffle() {
	for i := range d {
		j := rand.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

// shuffleWithFY modifies the deck in place my shuffling the cards using Fisher–Yates algorithm to ensure unifrom distribution
func (d deck) shuffleWithFY() {
	for i := range d {
		// Generates a random index j such that i <= j < len(d)
		j := rand.Intn(len(d)-i) + i
		d[i], d[j] = d[j], d[i]
	}
}

// shuffleWithShuffle modifies the deck in place by shuffing the cards using the standard shuffle function.
func (d deck) shuffleWithShuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
