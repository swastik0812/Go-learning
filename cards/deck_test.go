package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length of 16, but got", len(d))
	}

	if d[0] != "Ace of Spade" {

		t.Error("Expected first card of Ace of Spade, but got", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {

		t.Error("Expected last card Four of Clubs, but got", d[len(d)-1])
	}

}

//  *testing.T this will specify the type of value

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected deck length of 16, but got", len(loadedDeck))
	}
	
	os.Remove("_deckTesting")

}
