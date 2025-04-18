package main

import (
	"os"
	"testing"
)

const deckSize = 52
const firstCard = "Ace of Clubs"
const lastCard = "King of Diamonds"
const testFilename = "_decktesting"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	dLen := len(d)
	if dLen != deckSize {
		t.Errorf("Expected new deck size %v, got %v", deckSize, dLen)
	}

	dFirst := d[0]
	if dFirst != firstCard {
		t.Errorf("Expected first card of new deck to be %v, got %v", firstCard, dFirst)
	}

	dLast := d[dLen-1]
	if dLast != lastCard {
		t.Errorf("Expected last card of new deck to be %v, got %v", lastCard, dLast)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	err := os.Remove(testFilename)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("Attempt to cleanup previous test artifact %v failed: %v", testFilename, err)
	}

	d := newDeck()
	d.saveToFile(testFilename)

	newD := newDeckFromFile(testFilename)
	dLen := len(d)
	newDLen := len(newD)
	if newDLen != dLen {
		t.Errorf("Expected new deck size to match og deck size of %v, got %v", dLen, newDLen)
	}

	dFirst := d[0]
	newDFirst := newD[0]
	if newDFirst != dFirst {
		t.Errorf("Expected first card of new deck match og deck first card of %v, got %v", dFirst, newDFirst)
	}

	dLast := d[dLen-1]
	newDLast := newD[newDLen-1]
	if newDLast != dLast {
		t.Errorf("Expected last card of new deck match og deck last card of %v, got %v", dLast, newDLast)
	}

	err = os.Remove(testFilename)
	if err != nil {
		t.Fatalf("Attempt to cleanup current test artifact %v failed: %v", testFilename, err)
	}
}
