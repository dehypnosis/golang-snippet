package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter       int32
	atomicCounter int32
	wg            sync.WaitGroup
)

func main() {
	wg.Add(2)

	go inc3("A")
	go inc3("B")

	wg.Wait()
	fmt.Printf("counter = %d, atomicCounter = %d\n", counter, atomicCounter)
}

func inc3(label string) {
	defer wg.Done()
	defer fmt.Println(label + " inc3 done")

	for i := 0; i < 100; i++ {
		v := counter
		time.Sleep(1)
		runtime.Gosched()
		counter = v + 1
		atomic.AddInt32(&atomicCounter, 1)
	}
}

// can see the data race by `go run -race ...` command
