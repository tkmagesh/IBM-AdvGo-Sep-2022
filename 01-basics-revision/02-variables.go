package main

import "fmt"

//var x int
//var x int = 100
//var x = 100
//x := 100 //=> NOT applicable at the package level

var myVar int // => unused variables are allowed at the package level

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		//type inteference
		var msg = "Hello World!"
	*/

	//the below syntax is application ONLY in a function (function scope) and NOT at the package level
	msg := "Hello World"
	fmt.Println(msg)

	//unused variables are NOT allowed at the function level
	var myVar int
	myVar = 100
	fmt.Println(myVar)

	/* multiple variables */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of 100 and 200 is"
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 and 200 is"
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 is"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of 100 and 200 is"
	result := x + y
	fmt.Println(str, result)

	//type conversion
	var n int = 100
	var f float32
	f = float32(n) //converting int to float32
	fmt.Println("f = ", f)

	const pi = 3.14
	//const pi float32 = 3.14 //unused constants are allowed
	//pi = 2 //=> assignment to a constant not allowed

	//iota
	/*
		const red = 1
		const green = 2
		const blue = 3
	*/

	/*
		const (
			red   = 1
			green = 2
			blue  = 3
		)
	*/

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	const (
		red = iota + 2
		green
		_
		blue
	)

	//fmt.Printf("Red = %d, Green = %d, Blue = %d\n", red, green, blue)
	fmt.Println("Red =", red, "Green =", green, "Blue =", blue)

	//Usage
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)

	//complex type
	var n1 complex64 = 4 + 2i
	fmt.Println(n1, real(n1), imag(n1))
	var n2 complex64 = 5 + 3i
	fmt.Println(n1 + n2)

	f1 := 87352.5678464
	fmt.Println(f1)
	fmt.Printf("%.2f\n", f1)
}
