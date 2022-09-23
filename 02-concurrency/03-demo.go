package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(20)
}

func main() {
	//var wg sync.WaitGroup
	fmt.Println("Hit ENTER to continue...")
	fmt.Scanln()
	/*
		var count = flag.Int("count", 0, "number of goroutines")
		flag.Parse()


			wg := &sync.WaitGroup{}
			for i := 1; i <= *count; i++ {
				wg.Add(1)    //initializing the counter to 1
				go fn(i, wg) //scheduling the execution to the go scheduler
			}
			f2()
			wg.Wait() // waiting for the counter to become 0
	*/
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		f2()
		wg2.Done()
	}()
	wg2.Wait()
	fmt.Println("Done")

}

func fn(idx int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	fmt.Printf("fn started - [%d]\n", idx)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("fn completed - [%d]\n", idx)

}

func f2() {
	fmt.Printf("f2 started\n")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("f2 completed\n")
}
