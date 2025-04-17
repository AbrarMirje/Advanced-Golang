package main

import (
	"errors"
	"fmt"
)

// func square(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("square root of negative number")
// 	}
// 	return 1, nil
// }

func process(b []byte) error {
	if len(b) == 0 {
		return errors.New("length is zero")
	}
	return nil
}

func main() {
	// result, err := square(-10)
	// if err != nil {
	// 	fmt.Println("Error while calculating square")
	// 	return
	// }
	// fmt.Println(result)

	// sl := []byte{1, 2, 3}
	// if err := process(sl); err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	// fmt.Println("Data proccessed successfully!!")

	// if err1 := eprocess(); err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

type myError struct {
	message string
}

func (me *myError) Error() string {
	return fmt.Sprintf("Error:%s", me.message)
}

func eprocess() error {
	return &myError{
		message: "Custom Error Message",
	}
}

func readData() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}
