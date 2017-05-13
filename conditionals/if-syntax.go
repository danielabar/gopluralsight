package main

import "fmt"

func main() {
	// Variables to store course rankings
	firstRank := "39"   //Docker Deep Dive
	secondRank := "614" //Docker Clustering

	if firstRank < secondRank {
		fmt.Println("First course is doing better than second course")
	} else if firstRank > secondRank {
		fmt.Println("Oh dear... your first course must be doing abysmally!")
	} else {
		fmt.Println("Both courses are either the same or something weird is going on")
	}
}
