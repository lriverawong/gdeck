package main

import (
	"os"
	"testing"
)

// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Check first card in deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck to start with <Ace of Spades> but got %v", d[0])
	}

	// Check last card in deck
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected deck to start with <King of Clubs> but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	deck := newDeck()

	// Remove any previous failed test runs
	os.Remove(testFileName)

	// Write test file
	deck.saveToFile(testFileName)

	// Read test file
	loadedDeck := newDeckFromFile(testFileName)

	// Testing
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	// Remove test file created
	os.Remove(testFileName)
}
