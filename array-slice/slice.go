package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4]) //5

	// modify slice
	mySlice[1] = 0
	fmt.Println(mySlice) //[1 0 3 4 5 6 7 8 9 10]

	// make a smaller slice
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice) //[3 4 5]
}
