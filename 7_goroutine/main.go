package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	wg.Add(numCPU + 2)

	fmt.Println("Go routine!")

	go func() {
		defer wg.Done()

		for c := 0; c < 3; c++ {
			for ch := 'a'; ch <= 'z'; ch++ {
				fmt.Printf("%c ", ch)
			}
		}
	}()

	go func() {
		// this will be performed before the function returns
		defer wg.Done()

		for c := 0; c < 3; c++ {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%c ", ch)
			}
		}
	}()

	for i := 0; i < numCPU; i++ {
		go printPrime(fmt.Sprintf("%c", 'A'+i))
	}

	// wait until wait group counts zero
	wg.Wait()
}

func printPrime(prefix string) {
	defer wg.Done()

outerLoop:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			// fmt.Println(outer, inner)
			if outer%inner == 0 {
				continue outerLoop
			}
		}
		fmt.Printf("%s:%d ", prefix, outer)
	}
	fmt.Println("Finished: ", prefix)

}
