package main

import (
	"os"
	"testing"
)

// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d.Cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d.Cards))
	}

	// Check first card in deck
	firstCard := Card{Suit: "Spades", Value: "Ace"}
	if d.Cards[0] != firstCard {
		t.Errorf("Expected deck to start with <Ace of Spades> but got %v", d.Cards[0])
	}

	// Check last card in deck
	lastCard := Card{Suit: "Clubs", Value: "King"}
	if d.Cards[len(d.Cards)-1] != lastCard {
		t.Errorf("Expected deck to start with <King of Clubs> but got %v", d.Cards[len(d.Cards)-1])
	}

}

func TestStreamSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	deck := newDeck()

	// Remove any previous failed test runs
	os.Remove(testFileName)

	// Write test file
	deck.streamSaveToFile(testFileName)

	// Read test file
	loadedDeck := streamNewDeckFromFile(testFileName)

	// Testing
	if len(loadedDeck.Cards) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck.Cards))
	}

	// Remove test file created
	os.Remove(testFileName)
}
