package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := getNewDeck()
	firstCard := d[0]
	lastCard := d[len(d)-1]

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", firstCard)
	}

	if lastCard != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	deck := getNewDeck()

	deck.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(testFile)
}
