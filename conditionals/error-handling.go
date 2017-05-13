package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Open returns two values but we don't care about the first,
	// so it's effectively ignored with underscore
	_, err := os.Open("/someuser/temp/test1.txt")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}
