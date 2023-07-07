package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Webtechsolz office")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our work between 1 to 5")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	// converting string to float using parse
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	/* We are trimming the user input string because the user input is
	   is taking the input as '4\n'*/
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
