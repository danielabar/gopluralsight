package main

import "fmt"

func main() {
	// create a channel using go's built-in make functions-primer
	// first argument is `chan` keyword indicating we want to build a channel
	// second argument is type of data that can flow through channel
	// third argument (optional) is channel capacity, set to 1 for this demo
	// because the same main thread is both sending and receiving a message,
	// if channel has no capacity to store a message then this code will deaadlock.
	ch := make(chan string, 1)

	// put a message in the channel using receive operator `<-`
	// with a basic channel, sending a message causes thread to block until something receives it
	ch <- "Hello"

	// get message to be printed from channel
	// use left facing arrow to drain channel `<-` (receive operator), followed by channel reference
	// goroutine will wait until message is available in channel
	fmt.Println(<-ch)
}
