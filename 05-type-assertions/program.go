package main

import "fmt"

func main() {
	//var x interface{}
	var x any
	//x = 100
	//x = "This is a string"
	//x = true
	//x = 123.456
	//x = struct{}{}
	//x = []int{10, 20, 30, 40}
	x = []string{}
	fmt.Println(x)
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("Not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64")
	case bool:
		fmt.Println("x is a bool with value :", val)
	case struct{}:
		fmt.Println("x is an empty struct")
	case []int:
		fmt.Println("x is a list of int, list size = ", len(val))
	default:
		fmt.Println("Unknown type")
	}

}
