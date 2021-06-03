package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// a slice of strings
type deck []string

// creates and return a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// returns a hand and the remaining cards in the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function, prints cards, any varibale of type "deck" now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// takes in a deck and returns that deck of cards as a string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// takes in a fileName and writes data to a file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// takes in a fileName and returns the contents of that file
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}
