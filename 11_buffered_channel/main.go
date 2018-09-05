package main

import (
	"fmt"
	"sync"
)

type task struct {
	name    string
	command string
}

const taskWorkerNum = 2

var taskQueue chan task
var wg sync.WaitGroup

func init() {
	taskQueue = make(chan task, 5)
}

func main() {
	wg.Add(taskWorkerNum)

	for i := 0; i < taskWorkerNum; i++ {
		go work(fmt.Sprintf("worker %d", i+1))
	}

	taskQueue <- task{
		name:    "work 1",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 2",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 3",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 4",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 5",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 6",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 7",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 8",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 9",
		command: "ls ~",
	}
	taskQueue <- task{
		name:    "work 10",
		command: "ls ~",
	}

	// close(taskQueue) // if omit this line, finally meet the deadlock
	// -> blocking -> non-blocking by select statement at work func

	wg.Wait()
}

func work(label string) {
	defer wg.Done()

L:
	for {
		select {
		case task, ok := <-taskQueue:
			if !ok {
				// close(taskQueue)
				continue
			}
			fmt.Println(task, label)
		default:
			break L
		}
	}
}
