package main

import "fmt"

func main() {
	cards := deck{"A", "B", "C", "D"}
	cards = append(cards, "E")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
