package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers notes")
	var ptr *int
	fmt.Println("The value of the pointer ", ptr)
	myNumber := 9
	ptr = &myNumber
	fmt.Println("The address of the pointer ", ptr)
	fmt.Println("The value in pointer ", *ptr)
	*ptr = *ptr + 4
	fmt.Println("New value is ", myNumber)
	fmt.Println("The address is ", ptr)
	fmt.Println("New value is ", *ptr)
}
