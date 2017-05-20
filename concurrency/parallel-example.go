package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// self-executing anonymous function
	go func() {
		defer waitGrp.Done()

		// this will put ENTIRE PROGRAM to sleep for 5 seconds (if not using goroutine)
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	// self-executing anonymous function
	go func() {
		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()
}
