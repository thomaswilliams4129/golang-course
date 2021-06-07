package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedSize := 53

	if len(d) != expectedSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedSize, len(d))
	}

}
