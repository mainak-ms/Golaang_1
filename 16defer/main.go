package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
} // Defer works in last in first out process
// Defer keyword made the execution late
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
