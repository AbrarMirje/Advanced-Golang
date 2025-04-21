package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	// Keyword to filter lines
	keyword := "little"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "fewer")
			fmt.Printf("%d Filtered line:%v\n", lineNumber, line)
			fmt.Printf("%d Updated line:%v\n", lineNumber, updatedLine)
			lineNumber++
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("error: while scanning file:", err)
		return
	}
}
