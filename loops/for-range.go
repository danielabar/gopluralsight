package main

import "fmt"

func main() {
	// declare and define a list of courses,
	// technically, its a "slice": unordered numbered list of items
	// behind the scenes, its a reference to an array
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}

	for _, i := range coursesInProg {
		fmt.Println(i)
	}
}
