package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
