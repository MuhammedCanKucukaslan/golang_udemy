package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the file name as a command line argument! If the executable is named \"main\" and the file is named \"filename.extension\" then write")
		fmt.Print("\nmain filename.extension\n\n")

		os.Exit(21) // 400%126 -1
	}
	filename := os.Args[1]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(myWriter{}, file)

}

type myWriter struct{}

func (myWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
