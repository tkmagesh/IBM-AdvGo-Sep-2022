package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateNos()
	for data := range ch {
		fmt.Println(data)
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
