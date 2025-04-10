package main

import "fmt"

func main() {
	//fmt.Println(factorial(5))
	//fmt.Println(factorial(10))

	//fmt.Println("sum of 0:", sumOfDigits(0))
	//fmt.Println("sum of 1:", sumOfDigits(1))
	fmt.Println("Sum of digits:", sumOfDigits(15))
	//fmt.Println("sum of 10:", sumOfDigits(10))
}

func factorial(n int) int {
	// Base Case => Factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive Case: factorial  of n is n * factorial(n-1)
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)

	/*
		Step1:
			n % 10 = 15 % 10 => 5
			n / 10 = 15 / 10 => 1
			now sumOfDigits(1) called
		Step2:
			1 < 10 = ✅ return n i.e. 1
			15 % 10 = 5 + 1 == 6
			return 6 to caller in main()
	*/
}
