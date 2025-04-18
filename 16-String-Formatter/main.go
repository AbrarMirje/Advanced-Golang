package main

import "fmt"

func main() {
	num := 123
	/*
		%0 => Leading numbers should be zero (0)
		5d => There should me minimum 5 numbers
		\n => new line
	*/
	fmt.Printf("%05d\n", num)

	message := "Hello"
	/*
		|(pipe operator) => is not required, just to check white spaces.
		- sign => for trailing spaces
	*/
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

}
