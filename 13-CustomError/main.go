package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation completed successfully.")
}

type customError struct {
	code    int
	message string
	er      error
}

// Error returns the error message. Implementing Error() method of error interface.
func (ce *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", ce.code, ce.message, ce.er)
}

// Function that returns the custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "something went wrong!",
// 	}
// }

func doSomething() error {
	if err := doSomethingElse(); err != nil {
		return &customError{
			code:    500,
			message: "something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}
