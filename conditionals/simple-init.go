package main

import "fmt"

func main() {
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("First course is doing better than second course")
	} else if firstRank > secondRank {
		fmt.Println("Oh dear... your first course must be doing abysmally!")
	} else {
		fmt.Println("Both courses are either the same or something weird is going on")
	}
}
