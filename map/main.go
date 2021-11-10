package main

import (
	"fmt"
)

func main() {

	var m map[string]string

	m = map[string]string{
		"red":  "#ff0000",
		"blue": "#0000ff",
	}

	printMap(m)

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "white")

	fmt.Println(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, ":", hex)
	}
	/*
		for color:= range m {
			fmt.Println(color, ":", m[color])
		}
	*/
}
