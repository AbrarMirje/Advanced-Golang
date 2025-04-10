package main

import "fmt"

func main() {
	// adder() calls only once and every time sequnce getting printed, so that value of i getting stored as a closure.
	//sequence := adder()
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println("-------------")
	//sequence2 := adder()
	//fmt.Println(sequence2())
	//fmt.Println(sequence2())

	decrement := func() func(int) int {
		countdown := 100
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	fmt.Println(decrement(1))
	fmt.Println(decrement(2))
	fmt.Println(decrement(3))
	fmt.Println(decrement(4))
}

//func adder() func() int {
//	i := 0
//	fmt.Println("Previous value of i:", i)
//	return func() int {
//		i++
//		fmt.Println("added 1 to i", i)
//		return i
//	}
//}
