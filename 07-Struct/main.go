package main

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell // Anonymous Embeded struct
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	City    string
	State   string
	Country string
}

func main() {
	p1 := Person{
		firstName: "Royal",
		lastName:  "Rumble",
		age:       31,
		address: Address{
			City:    "Kolhapur",
			State:   "Maharashtra",
			Country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "MyHome",
			cell: "MyCell",
		},
	}
	p1.address.Country = "Palestine🍉"
	p2 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       27,
	}

	// Anonymous struct
	p3 := struct {
		username string
		age      int
	}{
		username: "Gajodhar",
		age:      24,
	}
	p4 := struct {
		username string
		age      int
	}{
		username: "Gajodhar",
		age:      24,
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.address.Country)
	fmt.Println(p1.address.State)
	// fmt.Println(p1.City) => We can't do this coz it's a named struct
	fmt.Println(p1.cell) // We can do this bcoz it's a anonymous struct and it promote itself to called its fields by outer struct

	fmt.Println(p2.firstName)
	fmt.Println(p3.username)

	fmt.Println("----------------")
	result := p1.admin()
	fmt.Println(result)

	fmt.Println("-----Calling Pointer Function------")
	p1.incrementAgeByOne()
	fmt.Println(p1.age)

	// Comparision of two structs can be possible iff
	// 1️⃣both the instances of struct belongs to same struct
	// 2️⃣both struct have same fields as well as same fields value
	fmt.Println("Are p3 and p4 are equals:", p3 == p4)

}

// Struct Receiver
func (p Person) admin() string {
	return fmt.Sprintf("%s %s is %d years of young man.", p.firstName, p.lastName, p.age)
}

// Pointer receiver
func (p *Person) incrementAgeByOne() {
	p.age++
}
