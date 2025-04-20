package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("error: while open the file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing the open file...")
		file.Close()
	}()
	fmt.Println("File opened successfully!")

	// Read the content of the opened file
	// data := make([]byte, 1024) // Buffer to read data into
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("error: while reading the file:", err)
	// 	return
	// }
	// fmt.Println("File content:", string(data))

	// Reading file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("error: while reading the file:", err)
		return
	}
}
