package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in Golang")
	//No inheritance in golang
	choi := User{"Trisha", "trisha@go.dev", true, 28}
	fmt.Println(choi)
	fmt.Printf("choi details are: %+v\n", choi)
	fmt.Printf("Particular detail %v", choi.Email)
	fmt.Println("-----------------------------")
	p := &choi
	p.Age = 29
	p.Email = "trisha@webtechsolz"
	p.Name = "Trisha Mainak Saha"
	fmt.Println(*p)
	fmt.Printf("%x %x", &p.Name, &choi.Name)
	koyel := User{"Trisha Mainak Saha", "trisha@webtechsolz", true, 29}
	fmt.Println(compare1(choi, koyel))
	fmt.Println(koyel)
	fmt.Println(compare2(&choi, &koyel))
	fmt.Println(koyel)
	x1 := User2{"Choi", 29}
	x2 := User2{"koyel", 29}
	fmt.Println(x1)
	fmt.Println(x2)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func compare1(p1 User, p2 User) bool {
	if p1.Name == p2.Name {
		p2.Name = "Trisha Halder"
		return true
	}
	return false
} // This is the call by value which does not modify the basic struct
func compare2(p1 *User, p2 *User) bool {
	if p1.Name == p2.Name {
		p2.Name = "Trisha Halder"
		return true
	}
	return false
} // This is call by reference which can modify the basic struct
type User2 struct {
	string
	int
} // Anonymous struct
