package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got: %v.", len(d))
	}
	if d[0] != firstCard {
		t.Errorf("Expected %v as first card, got: %v", firstCard, d[0])
	}
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected %v as last card, got: %v", lastCard, d[len(d) - 1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := deckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got: %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}