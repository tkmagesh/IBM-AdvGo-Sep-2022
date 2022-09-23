package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	quotient, remainder, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong!")
		fmt.Println(err)
		return
	}
	fmt.Printf("quotient = %d and remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := DivideByZeroError
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
