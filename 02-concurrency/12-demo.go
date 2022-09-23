package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(400 * time.Millisecond)
		fmt.Println(<-ch3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case ch3 <- 300:
			fmt.Println("sent 300 to the channel")
		case data2 := <-ch2:
			fmt.Println(data2)
		}
	}
}
