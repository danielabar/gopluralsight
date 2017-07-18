package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// mutex to guard goroutine to ensure only one thread can execute it at one time
	mutex := new(sync.Mutex)

	//  use up to 4 cpu's
	runtime.GOMAXPROCS(4)

	// increment two counters
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// simple goroutine responsible for printing out counter variables and their sums
			// demonstrates danger of shared memory because goroutines run on their own schedule
			// separate from the main outer loops.
			// Now to fix it with mutex
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
				// note this locking/unlocking (i.e. synchronizatin) slows the app down
			}()
		}
	}

	// prevent main function from exiting too early by making it wait for a keystroke
	fmt.Scanln()
}
