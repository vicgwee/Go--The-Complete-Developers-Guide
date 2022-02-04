package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#0000ff",
	}

	/* Equivalent ways of making an empty map
	var colors map[string]string
	colors := make(map[string]string)
	*/

	// Add and delete individual key-value pairs
	colors["white"] = "#ffffff"
	delete(colors, "white")

	// Iterate over key-value pairs

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, "has a hex value of", v)
	}
}
