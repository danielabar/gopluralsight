package main

import "fmt"

func main() {
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 7, 5, 2)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, f := range finishes {
		if f < best {
			best = f
		}
	}
	return best
}
