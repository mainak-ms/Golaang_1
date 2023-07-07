package main

import (
	"fmt"
)

// Declaring a map of map
var players = make(map[string]map[string]bool)

// Global variable
func main() {
	fmt.Println("welcome to Golang Functions")
	greet()
	sum := adder(5, 4)
	fmt.Println(sum)
	proSum, proString := proAdder(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(proSum)
	fmt.Println(proString)
	addPlayer("Trisha", "Halder")
	addPlayer("Rita", "Saha")
	addPlayer("Jyotish", "Saha")
	addPlayer("Mainak", "Saha")
	fmt.Println(players)
	fmt.Println(hasPlayer("Trisha", "Halder"))
	fmt.Println(hasPlayer("Tisa", "Halder"))

}

func greet() {
	fmt.Println("Hello I am a Golang Function")
}
func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

// we can pass multiple values with variadic function (values ...type)
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Summation done"
}

func addPlayer(fName, lName string) {
	var n = players[fName]
	if n == nil {
		n = make(map[string]bool)
		players[fName] = n
	}
	n[lName] = true
}
func hasPlayer(fName, lName string) bool {
	return players[fName][lName]
}
