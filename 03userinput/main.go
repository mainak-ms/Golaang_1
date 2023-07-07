package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "You are welcome Mainak!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your name:")

	/*comma ok syntax or comma err syntax to catch any error
	during user input*/
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Tea,", input)
}
