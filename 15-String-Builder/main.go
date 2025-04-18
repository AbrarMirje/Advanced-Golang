package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create string builder
	var builder strings.Builder

	// Write some string
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")

	// Convert builder to string
	result := builder.String()
	fmt.Println(result)

	// Using WriteRune to add a character
	builder.WriteRune(' ')
	len, _ := builder.WriteString("How are you")
	fmt.Println("Length:", len)

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("New Fresh String")
	result = builder.String()
	fmt.Println(result)
}

// Proper way to prevent from panic
//https://go.dev/play/p/Ht8o91tgIOR
