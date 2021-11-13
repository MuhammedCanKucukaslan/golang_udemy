package main

import "fmt"

type bot interface {
	getGreeting() string
	// getBotVersion() float64
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	printGreeting(englishBot{})
	printGreeting(spanishBot{})

}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// ...
	return "Hi there"
}
func (eb spanishBot) getGreeting() string {
	// ...
	return "Hola!"
}
