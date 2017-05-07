package main

import "fmt"

func main() {
	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course) // Docker Deep Dive

	// pass pointer to course variables location in memory
	changeCourse(&course)

	fmt.Println("\nHi", name, "you're currently watching", course) // First Look: Native Docker Clustering
}

// askterisk tells Go that `course` is a POINTER to a string variable
func changeCourse(course *string) string {
	// asterisk tells Go that we're assigning to location in memory that the course pointer is referencing
	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your cousrse to", *course)
	return *course
}
