package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"
	fmt.Println("My language Map is \n", language)
	fmt.Println("JS shorts for:", language["JS"])
	delete(language, "RB")
	fmt.Println("Updated Map", language)

	//loop in maps of Golang
	for key, value := range language {
		fmt.Printf("For key %v and the value is %v\n", key, value)
	}
	sal := map[string]float64{
		"Blake":  6000.00,
		"Parker": 120000.50,
		"Dakota": 93000.00,
	}
	fmt.Println(sal)
	fmt.Println(sal)
	fmt.Println(sal)
}
