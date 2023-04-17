package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected a deck of length of 20, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected the first card to be ace of spades, but we got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {

		t.Errorf("expected the first card to be four of clubs, but we got %v", d[len(d)-1])
	}
}

func TestSaveToDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")
	deck := newDeck()

	deck.saveToFIle("_decktesting")
	loadedDeck := newDeckFromFIle("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
