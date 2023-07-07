package main

import "fmt"

func main() {
	fmt.Println("If and Else in Golang")
	logincount := 22
	var result string
	if logincount < 10 {
		result = "Regular count"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)
	// Below if syntax used during web request
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}
