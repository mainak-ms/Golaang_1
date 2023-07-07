package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in Golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Want to create a date with time
	createdDate := time.Date(2011, time.August, 19, 10, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
