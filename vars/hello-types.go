package main

import (
	"fmt"
	"reflect"
)

var (
	name   string  //Name of subscriber
	course string  //Name of current course
	module float64 //Current place in course
)

func main() {
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
}
