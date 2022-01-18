package main

import (
	"fmt"
	"net/http"
	"time"
)

//  concurrency is not parallelism!
func main() {
	urls := []string{
		"https://www.udemy.com/",
		"https://google.com/",
		"https://amazon.com/",
		"https://go.dev/",
		"https://facebook.com/",
		//"https://nonexistingbutreally.com/",
	}

	c := make(chan string)

	for _, link := range urls {
		go check(link, c)
	}
	for l := range c {
		go func(url string) {
			time.Sleep(7 * time.Second)
			check(url, c)
		}(l)
	}
}

func check(url string, c chan string) {
	_, err := http.Get(url) // bottleneck!
	//time.Sleep(7 * time.Second)
	if err != nil {
		c <- url
		fmt.Println("Some error occured accessing the ", url)
		return
	}
	fmt.Println(url, " is working!")
	c <- url
}
