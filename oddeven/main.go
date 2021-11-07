package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var msg_odd string = " is odd"
	var msg_even string = " is even"

	for _, val := range arr {
		if val%2 == 0 {
			fmt.Println(val, msg_even)
		}
		if val%2 != 0 {
			fmt.Println(val, msg_odd)
		}
	}
}
