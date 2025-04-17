package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

// Method to push element at the end of the stack
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

// Method to remove last element from the Stack
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

// Method to check if list is empty or not
func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

// Method to print elements
func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty...")
		return
	}
	fmt.Print("Stack Elements:")
	for _, element := range s.elements {
		fmt.Print(element, " ")
	}
	fmt.Println()
}

func main() {
	// x, y := 1, 2
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	// x1, y1 := "Go", "JavaScript"
	// x1, y1 = swap(x1, y1)
	// fmt.Println(x1, y1)

	// For int
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println(intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.isEmpty())
	intStack.printAll()

	fmt.Println("************************")
	// For string
	stringStack := Stack[string]{}
	stringStack.push("Go")
	stringStack.push("NodeJs")
	stringStack.push("Rust")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.isEmpty())
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.isEmpty())
	stringStack.printAll()
}
