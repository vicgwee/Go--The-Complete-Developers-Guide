package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first element Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)

	d := newDeck()
	d.saveToFile(f)

	d2 := newDeckFromFile(f)

	if len(d2) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d2))
	}

	os.Remove(f)
}
