package main

import (
	"fmt"
	"reflect"
)

func main() {
	module := 3.2
	ptr := &module
	fmt.Println("Ptr is", ptr, "and is of type", reflect.TypeOf(ptr))
	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)
}
