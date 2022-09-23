package main

import (
	"fmt"
	"time"
)

/*
func main() {
	ch := make(chan int)
	go generateNos(ch)
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func generateNos(ch chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
}
*/

func main() {
	ch := generateNos()
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("Done")
}

func generateNos() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
