package main

import (
	"fmt"
	"math/rand"
	"time"
	"worker-demo/worker"
)

type MyWork struct {
	Id int
}

//implementation of the worker.Work interface
func (myWork *MyWork) Task() {
	fmt.Println("task started - ", myWork.Id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println("task completed - ", myWork.Id)
}

func main() {
	w := worker.New(5)
	for i := 1; i <= 20; i++ {
		w.Run(MyWork{Id: i})
	}
	fmt.Println("All tasks are assigned")
	w.Shutdown() //wait for all the tasks to complete
}
