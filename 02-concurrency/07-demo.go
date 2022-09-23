package main

import (
	"fmt"
	"time"
)

//communicate by sharing memory

func main() {
	ch := make(chan int) //creating the channel
	go add(100, 200, ch)
	result := <-ch //receive operation
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result //send operation
}
