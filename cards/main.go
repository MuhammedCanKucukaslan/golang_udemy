package main

import (
	"fmt"
)

func main() {
	filename := "hand.txt"
	cards := newDeck()
	cards.print()
	fmt.Println("\n\n")
	cards.saveToFile(filename)

	cards.shuffle()
	cards.print()
	fmt.Println("\n\n")
	cards = newDeckFromFile(filename)
	cards.print()

}

func test(as, ax int, st string) {
	fmt.Println("Hello World")
}

func newCard() string {
	return "Five of Diamonds"
}
