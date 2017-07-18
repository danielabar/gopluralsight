package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	//  use up to 4 cpu's
	runtime.GOMAXPROCS(4)

	// create log file
	f, _ := os.Create("./log.txt")
	// close log file explicitly because logging routine will open it whenever it needs to write
	f.Close()

	// create channel to hold messages to be logged
	// give it a large buffer to prevent app from i/o bound
	// (writing to disk is slower than calculations)
	logCh := make(chan string, 50)

	// goroutine to handle logging
	go func() {
		// infinite loop
		for {
			// wait for message to come in on log channel
			msg, ok := <-logCh
			// if channel is open
			if ok {
				// open log file in append mode
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				// prepare log entry
				logTime := time.Now().Format(time.RFC3339)
				// write log msg to file
				f.WriteString(logTime + " - " + msg)
				fmt.Println("wrote to log file")
				// close file so changes can be flushed to disk
				f.Close()
			} else {
				// channel is closed, break out of loop and terminate goroutine
				break
			}
		}
	}()

	// take advantage of blocking behaviour of channels that have full buffers
	// in this scenario, we don't care about the message
	mutex := make(chan bool, 1)

	// increment two counters
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// instead of locking a mutex, put a mesage in the channel
			// message content doesn't matter so just arbitrarily pick true
			mutex <- true
			go func() {
				// store the message
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				// send the message to log channel
				logCh <- msg
				// also print message to console
				fmt.Printf(msg)
				// drain the channel without storing the result in a variable
				<-mutex
			}()
		}
	}

	// prevent main function from exiting too early by making it wait for a keystroke
	fmt.Scanln()
}
