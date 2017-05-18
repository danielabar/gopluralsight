package main

import "fmt"

func main() {
	// declare a map using the composite literal form
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for key, value := range testMap {
		fmt.Printf("Key is: %v Value is : %v\n", key, value)
	}
}
