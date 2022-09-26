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
	childCtxWithValue := context.WithValue(ctx, "k1", "v1")
	childCtx, cancel := context.WithCancel(childCtxWithValue)
	defer cancel()
	wg.Add(1)
	go fn1(childCtx, wg)
	go func() {
		fmt.Println("Hit ENTER to stop")
		fmt.Scanln()
		cancel()
	}()
	wg.Wait()
	fmt.Println("main completed")
}

func fn1(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("[@f1] Value from context [k1] =", ctx.Value("k1"))
	defer wg.Done()
	wg.Add(1)
	ctx2WithValue := context.WithValue(ctx, "k2", "v2")
	//ctx2, cancel := context.WithCancel(ctx)
	//defer cancel()
	go fn2(ctx2WithValue, wg)
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
	fmt.Println("[@f2] Value from context [k1] =", ctx.Value("k1"))
	fmt.Println("[@f2] Value from context [k2] =", ctx.Value("k2"))
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
