package main

import "fmt"

func main() {
	leageTitles := make(map[string]int)
	leageTitles["Sunderland"] = 6
	leageTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeage titltes: %v\nRecent head to heads: %v\n",
		leageTitles, recentHead2Head)
}
