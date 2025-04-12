package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// learningStringsInGo()
	learningRuneInGo()
}

func learningStringsInGo() {
	message := "Hello, Go!"
	// In raw sequences, escape characters will not work.
	rawMessage := `Welcome\tGo`
	message1 := "Hello, \rGo!" // Go!lo
	name := "abrar"

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawMessage:", len(rawMessage))

	fmt.Println(message[0]) // Prints ASCII value of H
	fmt.Println(name[0])    // Prints ASCII value of a

	fmt.Println("--------Comparision---------")
	str1 := "Apple"  // A have an ASCII value 65
	str2 := "banana" // b have an ASCII value 98

	// So comparision will be (64 > 98) => false
	fmt.Println(str1 > str2) // Lexicographic Comparision or Lexical Comparision

	fmt.Println("----Iteration over message----")
	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		// fmt.Printf("%x\n", char) // %x hexadecimal value
		fmt.Printf("%v\n", char) // %v Actual Rune value
	}

	fmt.Println("Rune Count:", utf8.RuneCountInString("abrar"))
}

func learningRuneInGo() {
	words := []rune{'أ', 'ب', 'ر', 'ا', 'ر'}
	// fmt.Printf("%T\n", r)
	// fmt.Printf("%v\n", string(r))
	for _, word := range words {
		fmt.Println(string(word))
	}
}
