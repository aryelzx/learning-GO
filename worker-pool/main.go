package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func main() {

}

// Way to process the tasks
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)

	// fmt.Println("Task", t.ID, "processed")
}

type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}
