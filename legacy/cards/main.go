package main

// import "fmt"

var cards deck

func main() {
	// const filename = "current_deck"
	cards = newDeck()
	// hand, remainingDeck, _ := deal(cards, 5)

	// fmt.Println("Hand:")
	// hand.print()
	// fmt.Println("----")

	// fmt.Println("Remaining deck:")
	// remainingDeck.print()
	// fmt.Println("----")

	// fmt.Println("Deck as a singlestring:")
	// fmt.Println(cards.toString())
	// cards.saveToFile(filename)
	// newCards := newDeckFromFile(filename)
	// newCards.print()
	cards.shuffle()
	cards.print()
}
