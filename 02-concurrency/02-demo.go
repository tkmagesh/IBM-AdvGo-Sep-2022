package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //initializing the counter to 1
	go f1()   //scheduling the execution to the go scheduler
	f2()
	wg.Wait() // waiting for the counter to become 0
}

func f1() {
	fmt.Println("f1 invoked")
	wg.Done() // decrementing the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
