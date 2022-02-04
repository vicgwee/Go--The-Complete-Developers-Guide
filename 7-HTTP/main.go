package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	url := "http://google.com"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		// Manually reading and writing
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// Using io.Copy
	//io.Copy(os.Stdout, resp.Body)

	// Custom Writer implementation
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println("Log: First 10 chars:", string(p[:10]))
	return 10, nil
}
