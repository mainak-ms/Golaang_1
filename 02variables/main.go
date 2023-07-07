package main

import "fmt"

const LoginToken = "RTR" /* Public as the first letter of LoginToken
is capital letter or in upper case. So, we can access it from
anywhere*/

func main() {
	var username string = "trisha"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint = 255 // if it is >=256 it will give error
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default value
	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("Variable is of type: %T \n", anotherValue)

	//implicit type
	var website = "webtechsolz.com"
	fmt.Println(website)
	// no var type
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
