package main

import "fmt"

func main() {
	cards := newDeck() // create deck

	cards.toString()
	cards.saveToFile("my_cards")

	deckFromFile := newDeckFromFile("my_cards")

	fmt.Println(deckFromFile)
}
