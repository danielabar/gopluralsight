package main

import "fmt"

var (
	name   string  //Name of subscriber
	course string  //Name of current course
	module float64 //Current place in course
)

func main() {
	fmt.Println("Name is set to", name)
	fmt.Println("Course is set to", course)
	fmt.Println("Module is set to", module)
}
