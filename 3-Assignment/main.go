package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := ""
	for _, e := range s {
		if e%2 == 0 {
			output = "is even"
		} else {
			output = "is odd"
		}
		fmt.Println(e, output)
	}
}
