package main

import "fmt"

func main() {

	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

}

func test(as, ax int, st string) {
	fmt.Println("Hello World")
}

func newCard() string {
	return "Five of Diamonds"
}
