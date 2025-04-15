package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	Length float64
	Width  float64
}

// Method with Value Receiver
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Method with Pointer Receiver
func (r *Rectangle) Scale(factor float64) {
	r.Length *= factor
	r.Width *= factor
}

func main() {
	rect := Rectangle{
		Length: 10,
		Width:  9,
	}
	area := rect.Area()
	fmt.Println("Area of Rectangle:", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of Rectangle:", area)

	num1 := MyInt(-5)
	num2 := MyInt(5)
	fmt.Println(num1.IsPositive())
	fmt.Println(num2.IsPositive())
	fmt.Println(num1.WelcomeMessage())

	shape := Shape{
		Rectangle: Rectangle{
			Length: 10,
			Width:  9,
		},
	}
	fmt.Println(shape.Area())
}

type MyInt int

// Method on a user defined type
// mi is a instance of MyInt and IsPositive() is associated with MyInt type
func (mi MyInt) IsPositive() bool {
	return mi > 0
}

func (MyInt) WelcomeMessage() string {
	return "Welcome to Golang"
}
