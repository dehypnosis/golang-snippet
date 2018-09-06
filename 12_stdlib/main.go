package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// embeded lock
var hits struct {
	sync.Mutex
	n int
}

type BB struct {
	CC string
	DD string
}

// nested struct
type config struct {
	AA string
	BB BB
}

func main() {
	hits.Lock()
	hits.n++
	hits.Unlock()

	// any type variable
	var obj interface{}

	// anonymous struct
	obj = []struct {
		a int
		b string
	}{
		{
			a: 10,
			b: "ten",
		},
		{
			a: 20,
			b: "twelve",
		},
	}
	describe(obj)

	switch obj.(type) {
	case rune:
		// ...
	default:
		print("v.(type) expression.. on switch statement")
	}

	c := config{
		AA: "a",
		BB: BB{
			CC: "c",
			DD: "d",
		},
	}
	fmt.Printf("%+v", c)

	// omit field on struct initilization
	d := config{"a", BB{"c", "d"}}
	fmt.Printf("%+v", d)

	fmt.Println("hello")
	println("what the") // stderr

	// set nil to channel
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				quit = nil
				fmt.Println(quit, "quited")
				// default:
				// 	println("hey")
			}
		}
	}()
	// time.Sleep(3 * time.Second)
	quit <- struct{}{}
	// time.Sleep(1 * time.Second)

	// use iota
	const (
		x = 10 << (iota + 1)
		y
		z
		w
		v
	)
	println(x, y, z, w, v)

	// defer/recover-panic flow control
	go func() {
		defer func() {
			println("I'm panicking...")
			// if condition can have pre-statement, only last statement is evalueated
			// use RECOVER to control panicked flow
			if x, y, r := 3, x+4, recover(); r != nil && y > x {
				println("stop panicking...!")
				if rs, ok := r.(string); ok { // need type assertion when convert interface{}
					println(rs)
				}
			}
		}()
		// panic("PANICKED!")

		// how about fatal?
		// log.Fatal("what about this?") //.. this ignore defer's and call os.Exit

		// how about Panic?
		// log.Panic("Panicked again..?!") //.. call panic and log.Print together
	}()

	// rune is int32 for representation of 4byte character (Unicode codepoint)
	character := rune('h')
	fmt.Printf("%+v, %T", character, character)

	// builtin print, println func: out to stderr
	time.Sleep(1 * time.Second)

	// Math
	print(math.IsNaN(math.NaN()), math.Inf(1), math.Inf(-1))

	// String
	print(`
	What
	the
hey
	`)

	// same type of args
	add := func(a, b int) int {
		return a + b
	}

	// dynamic slice of map
	m := []map[interface{}]interface{}{
		map[interface{}]interface{}{},
	} // m = [{}]
	m[0][0] = 1
	m[0]["key"] = "whatthe"
	fmt.Printf("%#v", m[0])

	print(add(1, 2))
}

func describe(i interface{}) {

	// printf detailed struct and type
	fmt.Printf("%+v (%T)", i, i)
}
