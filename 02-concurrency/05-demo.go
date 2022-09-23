package main

import (
	"fmt"
	"sync"
	"time"
)

//communicate by sharing memory

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //creating the channel
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch //receive operation
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result //send operation
	wg.Done()
}
