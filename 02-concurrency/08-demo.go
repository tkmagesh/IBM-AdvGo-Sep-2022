package main

import (
	"fmt"
	"sync"
	"time"
)

//communicate by sharing memory

func main() {
	fmt.Println("main started")
	//using WaitGroup to synchronize goroutines
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("goroutine completed")
		wg.Done()
	}()
	wg.Wait()

	//using channels to synchronize goroutines
	/*
		done := make(chan struct{})
		go func() {
			time.Sleep(4 * time.Second)
			fmt.Println("goroutine completed")
			done <- struct{}{}
		}()
		<-done
	*/
	fmt.Println("main completed")
}
