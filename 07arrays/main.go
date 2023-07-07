package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Array")
	var familyList [4]string
	familyList[0] = "Trisha"
	familyList[1] = "Rita"
	familyList[2] = "Jyotish"
	familyList[3] = "Mainak"
	fmt.Println("The family members are \n ", familyList)
	fmt.Println("The length of the array is:", len(familyList))
	var foodStatus = [4]string{"Veg", "Veg", "Veg", "Non-Veg"}
	fmt.Println(foodStatus)
}
