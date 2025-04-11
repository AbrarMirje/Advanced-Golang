package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// fmt.Println("----------------")
	// modification(ptr)
	// fmt.Println(ptr)
	// fmt.Println(a)
	fmt.Println("----------------")
	modificationWithoutPtr(*ptr)
	fmt.Println(ptr)
	fmt.Println(a)
}

func modification(ptr *int) {
	*ptr++
}

func modificationWithoutPtr(ptr int) {
	ptr++
}
