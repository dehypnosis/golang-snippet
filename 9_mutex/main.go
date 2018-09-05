package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int32
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg.Add(3)

	go work("A")
	go work("B")
	go work("C")

	time.Sleep(1000 * time.Millisecond)
	atomic.StoreInt32(&shutdown, 1)

	wg.Wait()
}

func work(label string) {
	defer wg.Done()
	defer fmt.Println("Done.....", label)

	for {
		fmt.Println("WIP.....", label)
		time.Sleep(250 * time.Millisecond)

		// if shutdown == 1 {
		// 	return
		// }

		// prevent deadlock from racing by atomic functions
		// if atomic.LoadInt32(&shutdown) == 1 {
		// 	return
		// }

		// or can use mutex
		mutex.Lock()
		if shutdown == 1 {
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}
