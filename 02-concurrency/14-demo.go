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
		timeoutCh := time.After(5 * time.Second) // built in alternative for our own timeout implementation
		i := 1
	LOOP:
		for {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-timeoutCh:
				fmt.Println("timeout occurred")
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
