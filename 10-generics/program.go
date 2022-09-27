package main

import "fmt"

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func sumInt(nos []int) int {
	var result int
	for _, no := range nos {
		result += no
	}
	return result
}

func sumFloat(nos []float32) float32 {
	var result float32
	for _, no := range nos {
		result += no
	}
	return result
}

func sum[T Numbers](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}

func filter[T Numbers](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(sum(nos))
	fmt.Println(sum([]float32{3.5, 1.2, 4.9, 2.6, 5.7}))
	evenNos := filter(nos, func(no int) bool {
		return no%2 == 0
	})
	fmt.Println(evenNos)
}
