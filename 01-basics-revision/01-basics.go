/* package declaration */
package main

/* import dependent packages */
import "fmt"

/* package level type / varaible declarations */
var msg string = "Hello World!"

/* package init function */
func init() {
	fmt.Println("Package initialized")
}

/* main function */
func main() {
	//fmt.Println("Hello World!")
	sayHello()

	/*
		var x int
		x = 100
	*/

	//var x int = 100

	var x = 100

	//x := 100
	y := 200
	result := x + y
	fmt.Println(result)
}

/* other functions */
func sayHello() {
	fmt.Println(msg)
}
