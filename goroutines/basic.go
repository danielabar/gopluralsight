package main

import (
	"runtime"
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10 ms ")

	// add parallelism, use 2 processors
	runtime.GOMAXPROCS(2)

	// define anonymous self-executing functions
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello: ", i)
			// all the "go" will get printed out first, then all the "hello"
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("go: ", i)
			time.Sleep(godur)
		}
	}()

	// create duration of 1 second
	dur, _ := time.ParseDuration("1s")
	// tell main thread to sleep to give the other concurrent goroutines a chance to execute
	time.Sleep(dur)
}
