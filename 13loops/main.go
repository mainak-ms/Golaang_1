package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println(days[i])
	}
	for index, day := range days {
		fmt.Printf("The index is %v and vale is %v\n ", index, day)
	} // this is a kind of for each loop
	anyValue := 1
	for anyValue < 10 {
		if anyValue == 2 {
			goto code
		}
		if anyValue == 5 {
			anyValue++
			continue
		}
		fmt.Println("The value is:", anyValue)
		anyValue++
	}
code:
	fmt.Println("Learn code")

	myruneslice := []rune("Learn code")
	fmt.Println(myruneslice)

}
