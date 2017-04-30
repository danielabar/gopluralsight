package main

import (
	"fmt"
	"reflect"
)

var (
	// inferredName, inferredCourse, inferredModule = "Nigel", "Docker Deep Dive", 3.2
	inferredName   = "Nigel"
	inferredCourse = "Docker Deep Dive"
	inferredModule = 3.2
)

func main() {
	fmt.Println("Name is", inferredName, "and is of type", reflect.TypeOf(inferredName))
	fmt.Println("Module is", inferredModule, "and is of type", reflect.TypeOf(inferredModule))

	a := 10.0000000000
	b := 3
	fmt.Println("\nA is of type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))
}
