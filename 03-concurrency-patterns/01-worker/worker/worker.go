package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	/*  */
	workQueue chan Work
	wg        sync.WaitGroup
}

func New(workerCount int) *Worker {
	worker := &Worker{
		workQueue: make(chan Work),
	}
	worker.wg.Add(workerCount)
	for idx := 0; idx < workerCount; idx++ {
		go func(id int) {
			for w := range worker.workQueue {
				fmt.Println("Worker Id :", id)
				w.Task()
			}
			worker.wg.Done()
		}(idx)
	}
	return worker
}

func (w *Worker) Run(work Work) {
	w.workQueue <- work
}

func (w *Worker) Shutdown() {
	close(w.workQueue)
	w.wg.Wait()
	fmt.Println("Worker shutdown completed")
}
