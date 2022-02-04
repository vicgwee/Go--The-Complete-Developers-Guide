package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://linkedin.com",
		"http://instagram.com",
		"http://tiktok.com",
	}

	c := make(chan string)

	for _, u := range urls {
		go checkURL(u, c)
	}

	/*
		// Simple go routine to iterate as data is received in channel c
		// Since u is a value type, it is passed into checkURL before it changes
		for u := range c {
			go checkURL(u, c)
		}
	*/

	/*
		To add a delay between requests, we cannot
		1. Place it in the for loop because it will block the main routine instead of the child routine
		2. Place it in checkURL since this prevents it from being reused.

		Solution: Use Function Literal to run more than 1 line of code in a goroutine
		Note the () after the closing } in order to call the func.
	*/
	/*
		for u := range c {
			go func() {
				time.Sleep(time.Second)
				checkURL(u, c)
			}()
		}
	*/

	/*
		Above code has unexpected behaviour since u changes before we run checkURL.
		Need to pass u (by value) into the function literal
		So that it is copied before it is changed by incoming data
	*/

	for u := range c {
		go func(url string) {
			time.Sleep(time.Second)
			checkURL(url, c)
		}(u)
	}

}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
	} else {
		fmt.Println(url, "is up!")
	}
	c <- url
}
