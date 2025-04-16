package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	Length, Width float64
}

type circle struct {
	Radius float64
}

func (r rect) area() float64 {
	return r.Length * r.Width
}

func (c circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// The only thing is different between rect and circle is, circle have one additional method diameter() that rect won't have.
func (c circle) diameter() float64 {
	return 2 * c.Radius
}

func measure(g geometry) {
	fmt.Println("Values:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}

func main() {

	r := rect{
		Length: 3,
		Width:  4,
	}
	fmt.Printf("Type of rect %T\n", r)

	c := circle{
		Radius: 5,
	}
	fmt.Printf("Type of circle %T\n", c)

	// We can print like this as well, but creating a method make it more readable, clearer and follows DRY principle
	// fmt.Println("Printing Rectangle details")
	// fmt.Println(r.area())
	// fmt.Println(r.perimeter())
	// fmt.Println("Printing Circle details")
	// fmt.Println(c.area())
	// fmt.Println(c.perimeter())

	// Calling measure()
	fmt.Println("Printing Rectangle details")
	measure(r)
	fmt.Println("Printing Circle details")
	measure(c)

	fmt.Println("*****Printing interface{}*****")
	myPrinter(1, "Abrar", true, 10.25, map[string]int{"Rob": 32}, []int{1, 2, 3, 4, 5})

	printType(23)
	printType("Golang")
	printType(3.258)
	printType(false)
}

func printType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("Type:string")
	case int:
		fmt.Println("Type:int")
	case float64:
		fmt.Println("Type:float64")
	default:
		fmt.Println("Type:unknown")
	}
}

func myPrinter(i ...interface{}) {
	for _, value := range i {
		fmt.Printf("Type: %T | Value: %[1]v\n", value)
	}
}
