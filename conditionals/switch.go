package main

import "fmt"

func main() {
	topic := "docker"

	switch topic {
	case "linux":
		fmt.Println("Here are some recommended Linux courses...")
	case "docker":
		fmt.Println("Here are some recommended Docker courses...")
	case "windows":
		fmt.Println("Here are some recommended Windows courses...")
	default:
		fmt.Println("Sorry we couldn't find a match, " +
			" why not try out Top 100 list!")
	}
}
