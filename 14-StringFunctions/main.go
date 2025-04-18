package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	// strings.Split() splits the string by ,
	fruits := "Mango, Banana, Apple"
	sepString := strings.Split(fruits, ",")
	fmt.Println(&fruits)
	for _, val := range sepString {
		fmt.Println(&val)
	}

	countries := []string{"India", "Palestine", "Saudi Arabia"}
	// strings.Join() is used to add seperator to the give data structure
	joined := strings.Join(countries, ", ")
	fmt.Println(countries)
	fmt.Println(joined)

	isContain := strings.Contains(fruits, "Mango")
	fmt.Println(isContain)

	phrase := "We are working on Go, because Go is one of the most powerfull and popular and easy language"

	// phrasedString := strings.Replace(phrase, "Go", "OG", 1)
	phrasedString := strings.ReplaceAll(phrase, "Go", "OG")
	fmt.Println(phrasedString)

	tripString := " Hello Golang  "
	fmt.Println(tripString)
	trimmedSpace := strings.TrimSpace(tripString)
	fmt.Println(trimmedSpace)

	fmt.Println(strings.ToLower(trimmedSpace))
	fmt.Println(strings.ToUpper(trimmedSpace))

	// Repeat a string for number of times
	fmt.Println(strings.Repeat("Abrar ", 4))

	// Count number of occurences
	fmt.Println(strings.Count("Abrar", "B"))

	fmt.Println(strings.HasPrefix("Abrar", "A"))
	fmt.Println(strings.HasSuffix("Abrar", "r"))

	str5 := "Hello 123 let's Go 1.24"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	runeCountString := "Hello السلام عليكم"
	fmt.Println(utf8.RuneCountInString(runeCountString))
}
