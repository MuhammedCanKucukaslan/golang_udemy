package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com") //"http://api.acikkuran.com/surah/113") //
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)
	/*
		bs := make([]byte, 99_999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	//fmt.Println("Wrote ", len(bs), " many bytes.")
	return len(bs), nil
}
