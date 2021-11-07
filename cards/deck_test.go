package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	exp_l := 52
	cards := newDeck()
	cards_l := len(cards)
	if len(cards) != exp_l {
		t.Errorf("Expected get length %v we got %v", exp_l, cards_l)
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades we got %v", cards[0])
	}
	if cards[cards_l-1] != "King of Clubs" {
		t.Errorf("Expected Ace of Spades we got %v", cards[cards_l-1])
	}
}

func TestSavingReading(t *testing.T) {
	filename := "_cards.txt"
	os.Remove(filename)
	cards := newDeck()
	cards.saveToFile(filename)

	hand := newDeckFromFile(filename)
	for i, _ := range cards {
		if cards[i] != hand[i] {
			t.Errorf("Expected %v we got %v", cards[i], hand[i])
		}
	}
}
