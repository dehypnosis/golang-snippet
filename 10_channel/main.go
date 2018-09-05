package main

import (
	"fmt"
	"sync"
)

var ch chan int
var chbuf chan int
var wg sync.WaitGroup

func init() {
	ch = make(chan int)
	chbuf = make(chan int, 10)
}

func main() {
	wg.Add(1)

	go ping(ch)
	ch <- 1

	wg.Wait()
}

func ping(ch chan int) {
	v := <-ch
	fmt.Println("got", v)

	if v > 100 {
		wg.Done()
		close(ch)
		fmt.Println("pingpong done")
		return
	}

	fmt.Println("ping", v)
	go ping(ch)
	ch <- v + 1
}
