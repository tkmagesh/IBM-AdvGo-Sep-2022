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
	childCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	wg.Add(1)
	go fn1(childCtx, wg)
	wg.Wait()
	fmt.Println("main completed")
}

func fn1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()
	go fn2(ctx2, wg)
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn1 Done!")
			break LOOP
		default:
			fmt.Print(".")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func fn2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn2 Done!")
			break LOOP
		default:
			fmt.Print("tick")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
