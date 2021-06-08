package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedSize := 52

	if len(d) != expectedSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedSize, len(d))
	}

	expectedStartingCard := "Ace of Spades"
	actualStartingCard := d[0]

	if actualStartingCard != expectedStartingCard {
		t.Errorf("Expected first card to be %v but found %v", expectedStartingCard, actualStartingCard)
	}

	expectedEndingCard := "King of Diamonds"
	actualEndingCard := d[len(d)-1]

	if actualEndingCard != expectedEndingCard {
		t.Errorf("Expected end card to be %v but found %v", expectedEndingCard, actualEndingCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	expectedSize := 52

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != expectedSize {
		t.Errorf("Expected deck length of %v found %v", expectedSize, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
