package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("expected 20, but got %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("Expected First card of \"ace of spades\" but got: %v", d[0])
	}

	if d[len(d)-1] != "five of clubs" {
		t.Errorf("Expected First card of \"five of clubs\" but got: %v", d[0])
	}
}
