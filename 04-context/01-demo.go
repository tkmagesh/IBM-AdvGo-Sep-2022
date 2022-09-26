package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	childCtx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go fn(childCtx, wg)
	go func() {
		fmt.Println("Hit ENTER to stop")
		fmt.Scanln()
		cancel()
	}()
	wg.Wait()
	fmt.Println("main completed")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn Done!")
			break LOOP
		default:
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
