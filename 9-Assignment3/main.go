package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: number of arguments provided: ", len(os.Args)-1)
		os.Exit(1)
	}
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
