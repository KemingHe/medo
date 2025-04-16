package main

import "fmt"

var cards deck

func main() {
	cards = append(cards, newCard())
	cards = append(cards, "Six of Hearts")
	cards = append(cards, "Five of Spades")
	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Ace of Diamonds"
}
