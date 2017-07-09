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

	// close channel after we're finished putting words into it,
	// this prevents program from deadlocking
	close(ch)

	// run forever until break out
	// this loop is not dependent on phrase that serves as input data
	for {
		// `if` uses go's ability to initialize a variable at the same time its used in a condition
		// receive `<-` operator returns two variables when pulling a message from a channel:
		// 1. message itself
		// 2. boolean true if channel is open, false if channel is closed
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			//  break out of loop if channel is closed
			break
		}
	}
}
