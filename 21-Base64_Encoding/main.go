package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding")

	//Encode Base64
	encode := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encode)

	// Decode from Base64
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("error: while decoding the data")
		return
	}
	fmt.Println(string(decode))

	// URL Safe : Avoid '/' and '+'
	data1 := []byte("He~lo, Base64 Encoding")
	encode1 := base64.StdEncoding.EncodeToString(data1)
	fmt.Println(encode1)
	safeUncode := base64.URLEncoding.EncodeToString(data1)
	fmt.Println(safeUncode)
}
