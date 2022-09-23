package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ch := generateNos(done)
	go func() {
		fmt.Println("Hit ENTER to stop....")
		fmt.Scanln()
		done <- true
	}()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func generateNos(done chan bool) chan int {
	ch := make(chan int)
	go func() {
		i := 1
	LOOP:
		for {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-done:
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
