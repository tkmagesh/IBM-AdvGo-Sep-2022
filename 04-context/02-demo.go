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
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	wg.Add(1)
	go fn(childCtx, wg)
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
