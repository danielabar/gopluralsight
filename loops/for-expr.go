package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Self destructing... BOOM!")
			break
		}
		fmt.Println("Initiating self destruct sequence in", timer, "seconds")
		time.Sleep(1 * time.Second)
	}
}
