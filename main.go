package main

func main() {
	cards := newDeck()
	cards = append(cards, "custom added card Joker")

	cards.deckPrinter()
}
