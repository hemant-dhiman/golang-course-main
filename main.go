package main

import "fmt"

func main() {
	cards := newDeck()
	cards = append(cards, "custom added card Joker")

	fmt.Println("Complete Deck")
	cards.deckPrinter()

	fmt.Println() // new line

	cardDeal, remainingCards := deal(cards, 7)

	fmt.Println("\nhand of cards")
	cardDeal.deckPrinter()

	fmt.Println("\nremainig of cards")
	remainingCards.deckPrinter()

	// fmt.Println("actual deck")
	// cards.deckPrinter()

	fmt.Println("deck to string conversion!")
	fmt.Println(cards.deckToString())

}
