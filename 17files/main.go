package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File in Golang")
	files, err := os.Create("./mygofile.txt")
	if err != nil {
		panic(err)
	}
	outFile, err := io.WriteString(files, "My created Golang file")
	if err != nil {
		panic(err)
	}
	fmt.Println("output file is:", outFile)
	defer files.Close()
	readFile("./mygofile.txt")

}
func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err) // panic keyword just stops the execution if error
	}
	fmt.Println("The raw data inside the file is\n", databyte)
	fmt.Println("The data inside the file is:\n", string(databyte))
}
