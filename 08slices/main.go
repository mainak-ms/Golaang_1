package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Golang Slices")
	var fruitList = []string{"Appel", "Tomato", "Peach"}
	// In slice we don't need to specify the number inside[]
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	highScores := make([]int, 4) // Use of make() to create maps
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 897
	highScores = append(highScores, 555, 671, 321)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	//How to remove a slice based on index
	var course = []string{"Flutter", "Golang", "RestAPI", "Networking"}
	fmt.Println(course)
	var index = 3
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
