package main

// var deckSize int

func main() {
	// var cards string = "Ace of spades";
	// We should specific data types while assigning the varibale just like java
	// Go will figure out what type of varibale it is based on the
	// use := very first initialize the variable, after that u can use = to change the values

	// var deckSize int
	// deckSize = 20

	// cards := "Ace of spades"
	// cards = "Five of Spades"

	// newCard := newCard()

	// fmt.Println(cards)
	// fmt.Println(deckSize)
	// fmt.Println(newCard())
	// fmt.Println(printState())

	//Array going to be fixed length of list whereas slice can grow or shrink
	//Both array and slice should be defined as same types / not mix dissimilar types

	// sliceCards := []string{newCard(), "Ace"}
	// sliceCards = append(sliceCards, "King")

	// fmt.Println(sliceCards)

	//range is the keyword we use whenever we want to loop through the slice
	//we are re-initializing ethe varibales verytime we for loop through the value thats
	//why we are using :=, we are dsicarding previous index and previous card had been declared

	// for i, card := range sliceCards {
	// 	fmt.Println(i, card)
	// }

	//Go is not a object oriented programming language, so no point of using classes

	// cards := deck{newCard(), "Ace"}
	// cards = append(cards, "King")

	// cards.print()

	// cards := newDeck()

	// cards.print()

	//After calling in the deal and passing in cards, does the list of string that cards variable point at change ? meaning Are we slicing the original
	// vlaue, Cards will be same before and after calling the deal function. we created two new refernces that point at subsections of cards slice, we never
	// modify the slice that cards is pointing at
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//convert string to byte slice - Type conversion

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	readCards := newDeckFromFile("my_cards")

	readCards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
