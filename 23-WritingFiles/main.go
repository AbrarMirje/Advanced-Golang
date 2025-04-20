package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("error: while creating file.", err)
		return
	}
	defer file.Close()

	// Write data to file
	msg := "We are going to write into the file."
	// Write()
	n, err := file.Write([]byte(msg))
	if err != nil {
		fmt.Println("error: while writing into file", err)
		return
	}
	fmt.Printf("File created successfully and there are %d char in the file.\n", n)

	// WriteString()
	file, err = os.Create("writestring.txt")
	if err != nil {
		fmt.Println("error: while writing a file.", err)
	}
	n, err = file.WriteString("We are now using WriteString() function to write string into the file directly.")
	if err != nil {
		fmt.Println("error: while writing into file", err)
		return
	}
	fmt.Printf("WriteString file created successfully with %d bytes", n)
}
