package main

import "fmt"

var cards deck

func main() {
	cards = newDeck()
	hand, remainingDeck, _ := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()
	fmt.Println("----")

	fmt.Println("Remaining deck:")
	remainingDeck.print()
	fmt.Println("----")
}
