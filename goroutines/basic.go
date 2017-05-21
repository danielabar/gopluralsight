package main

func main() {
	// define an anonymous self-executing function that prints hello to console
	func() {
		println("Hello")
	}()
}
