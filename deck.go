package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new data type with name "deck"
// which is the extension of existing slice functionality (in java we extends to different class)
type deck []string

//{d deck} is a receiver - Any variable of type deck now gets access to the "print" method
//cards is the reference to the actual copy of deck type variable
// func (cards deck) print() { d / cards it is similar to this in javascript - we usually by covention use single / two word -- receiver argument
//d - refers to value of type deck

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Go has support to retrun multiple from one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert the deck type to byte slice for I/O operations (which we are converting letter to ascii characters for machine readable format)
// we are going to use type conversion - string to byte by using []byte{"Pass the Values"}
// deck to string to byte
// ["red","yellow","blue"] to ["red,yellow,blue"]

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// error is a type in Go
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// bs - byte slice
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// option1 - log the error and return a callback to newDeck()
		// option2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //Ace of spades, One of spades
	return deck(s)
}
