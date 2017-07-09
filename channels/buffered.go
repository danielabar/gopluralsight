package main

import (
	"fmt"
	"strings"
)

// Even though there is only one thread in this program,
// can still do asynchronous processing because channels are always asynchronous
func main() {
	phrase := "These are the times that try men's souls.\n"

	// divide phrase into individual strings using `strings` package `Split` function
	words := strings.Split(phrase, " ")

	// create channel for strings with capacity for all the words
	ch := make(chan string, len(words))

	// go over slide of words pushing into channel
	for _, word := range words {
		ch <- word
	}

	// drain channel
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}
