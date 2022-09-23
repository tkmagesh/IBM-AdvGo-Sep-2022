package main

import "fmt"

func main() {

	/* anonymous function */
	func() {
		fmt.Println("anonymous function invoked")
	}()

	/* anonymous functions with arguments & return values */
	result := func(x, y int) int {
		return x + y
	}(100, 200)

	fmt.Println(result)

	/* assign functions to variables */
	/*
		var divide func(int, int) (int, int)
		divide = func(x, y int) (quotient, remainder int) {
			quotient, remainder = x/y, x%y
			return
		}
	*/
	divide := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	fmt.Println(divide(100, 7))

	//functions are arguments
	exec(fn)
	exec(func() {
		fmt.Println("anonymous function invoked")
	})
	logOperation(add, 100, 200)

	//functions as return values
	loggedDivide := getLoggedOperation(func(x, y int) {
		fmt.Println("Divide Result =", x/y)
	})
	loggedDivide(100, 7)

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))

	//closures
	fmt.Println("Closures")
	/*
		fmt.Println(increment())
		fmt.Println(increment())
		counter = 100
		fmt.Println(increment())
		fmt.Println(increment())
	*/
	increment := getIncrement()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func fn() {
	fmt.Println("fn invoked")
}

func exec(f func()) {
	fmt.Println("invocation started")
	f()
	fmt.Println("invocation completed")
}

func add(x, y int) {
	fmt.Println("Result =", x+y)
}

func logOperation(operation func(int, int), x, y int) {
	fmt.Println("operation started")
	operation(x, y)
	fmt.Println("operation completed")
}

func getLoggedOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("operation started")
		operation(x, y)
		fmt.Println("operation completed")
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

/*
var counter int

func increment() int {
	counter++
	return counter
}
*/

func getIncrement() func() int {
	var counter int
	increment := func() int {
		counter++
		return counter
	}
	return increment
}
