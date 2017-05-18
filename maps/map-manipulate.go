package main

import "fmt"

func main() {

	// declare a map using the composite literal form
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	// reference entry at key "C"
	fmt.Println(testMap["C"])

	// modify the value at key "A"
	testMap["A"] = 100

	// print out the entire map to verify the change
	fmt.Println(testMap)

	// insert a new value in map
	testMap["F"] = 6

	// verify change
	fmt.Println(testMap)

	// delete newly added key
	delete(testMap, "F")

	// verify
	fmt.Println(testMap)
}
