package main

import "fmt"

func main() {

	// every field within struct must be unique
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// var DockerDeepDirve courseMeta
	// DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
}
