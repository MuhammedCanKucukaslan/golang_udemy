package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

// define deck type as a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// create a new deck
func newDeck() deck {
	var cards deck = deck{}

	var cardSuits []string = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardValues []string = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
			//cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func deal(hand deck, size int) (deck, deck) {
	return hand[:size], hand[size:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func newDeckFromFile(filename string) deck {
	var d deck = deck{}
	// using the function
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)

	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	str := string(byteSlice)
	arr := strings.Split(str, ", ")
	// instead of "for" we could have written deck(arr)
	for _, str := range arr {
		d = append(d, str)
	}
	return d
}

func (d deck) shuffle() {
	for index := range d {
		tmp, err := rand.Int(rand.Reader, big.NewInt(51))
		newPos := tmp.Int64()
		if err != nil {
			os.Exit(70)
		}
		card := d[index]
		d[index] = d[newPos]
		d[newPos] = card
	}
}
