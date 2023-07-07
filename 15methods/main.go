package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//No inheritance in golang
	choi := User{"Trisha", "trisha@go.dev", true, 28}
	fmt.Println(choi)
	fmt.Printf("choi details are: %+v\n", choi)
	fmt.Printf("Particular detail %v \n", choi.Email)
	choi.GetStatus()
	choi.NewMail()
	fmt.Printf("Particular detail %v \n", choi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}
func (u User) NewMail() {
	u.Email = "main@go.dev"
	fmt.Println("New Mail ID is", u.Email)
} /* this method will only send the copy of the new email. So, it won't
change the initial Email. This is why pointer is needed*/
