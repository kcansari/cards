package main

import "testing"

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
