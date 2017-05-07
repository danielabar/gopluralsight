package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	myEnv := os.Environ()
	fmt.Println("Type of env", reflect.TypeOf(myEnv))

	// nicer formatted print
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER")
	fmt.Println("Hello", name)
}
