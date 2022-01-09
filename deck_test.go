package main

import (
	"os"
	"testing"
)

const fileName = "_decktesting"
const numbersOfCard = 16

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != numbersOfCard {
		t.Errorf("Expected deck length of %v, but got %v", numbersOfCard, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(fileName)

	d := newDeck()

	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != numbersOfCard {
		t.Errorf("Expected %v cards in deck, but got %v", numbersOfCard, len(loadedDeck))
	}

	os.Remove(fileName)
}
