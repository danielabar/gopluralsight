package main

import "fmt"

func main() {
	// declare and define a list of courses,
	// technically, its a "slice": unordered numbered list of items
	// behind the scenes, its a reference to an array
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}
	coursesCompleted := []string{"Docker Deep Dive", "Go Fundamentals", "Puppet Fundamentals"}

	for _, i := range coursesInProg {
		fmt.Println("Course in progress:", i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("Found the same value in both lists:", j)
			}
		}
	}
}
