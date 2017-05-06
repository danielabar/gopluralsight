package main

import "fmt"

func main() {
	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course) // Docker Deep Dive
	changeCourse(course)
	fmt.Println("\nHi", name, "you're currently watching", course) // Still Docker Deep Dive
}

// changeCourse gets a COPY of the course variable
func changeCourse(course string) string {
	course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your cousrse to", course)
	return course
}
