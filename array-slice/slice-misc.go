package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice) // [1 2 3 4 5]

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice) // 1 2 3 4 5 10 20 30]
}
