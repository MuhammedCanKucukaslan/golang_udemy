package main // main.go

import "fmt"

func main() {
	fmt.Print("Hello, World!")
	print("\nHello,\tWorld!")
}

func print(str string) {
	fmt.Print(str)
}
