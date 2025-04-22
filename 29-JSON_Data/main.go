package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	// steve := Person{
	// 	FirstName: "Steve",
	// 	Age:       32,
	// 	Email:     "steve@gmail.com",
	// 	Address: Address{
	// 		City:    "Kolhapur",
	// 		Country: "India",
	// 	},
	// }

	// jsonData, err := json.MarshalIndent(steve, "", "  ")
	// if err != nil {
	// 	fmt.Println("error: while marshalling")
	// 	return
	// }
	// fmt.Println(string(jsonData))

	// Unmarshal JSON Object
	// jsonData := `{
	// 	"first_name": "Steve",
	// 	"age": 32,
	// 	"email": "steve@gmail.com",
	// 	"address": {
	// 	  "city": "Kolhapur",
	// 	  "country": "India"
	// 	}
	//   }`

	// var person Person
	// if err := json.Unmarshal([]byte(jsonData), &person); err != nil {
	// 	fmt.Println("error: while unmarshalling")
	// 	return
	// }
	// fmt.Println(person)
	// person.Address.City = "Pune"
	// fmt.Println("City changes:", person.Address.City)
	// fmt.Println(person)

	// // Manage List (Slice/Arrays)
	// listOfCityCountry := []Address{
	// 	{City: "Mumbai", Country: "India"},
	// 	{City: "Makkah", Country: "UAE"},
	// 	{City: "Madinah", Country: "UAE"},
	// 	{City: "Gaza", Country: "Palestine"},
	// }

	// fmt.Println(listOfCityCountry)
	// listJsonData, err := json.Marshal(listOfCityCountry)
	// if err != nil {
	// 	fmt.Println("error: while marshaling list")
	// 	return
	// }
	// fmt.Println("List Data:", string(listJsonData))

	// Handling unknown JSON structures
	jsonData := `{
			"first_name": "Steve",
			"age": 32,
			"email": "steve@gmail.com",
			"address": {
			  "city": "Kolhapur",
			  "country": "India"
			}
		  }`

	var unmarshalJson map[string]any

	if err := json.Unmarshal([]byte(jsonData), &unmarshalJson); err != nil {
		fmt.Println("error: while handling unknown json")
		return
	}
	fmt.Println(unmarshalJson)
	fmt.Println(unmarshalJson["address"])

}
