package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 200 {
		t.Errorf("Expected deck lenght of 16, but go %v", len(d))
	}

}
