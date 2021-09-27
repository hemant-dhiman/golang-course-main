package main

import "fmt"

// custom type declaration
type deck []string

// new deck generater func
func newDeck() deck {
	cards := deck{}

	cardName := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValue := []string{"ace", "two", "three", "four", "five"}

	for _, cardN := range cardName {
		for _, cardV := range cardValue {
			cards = append(cards, cardN+" of "+cardV)
		}
	}

	return cards
}

// print method for custom type object (deck)
func (d deck) deckPrinter() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
