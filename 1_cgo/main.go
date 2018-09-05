package main

/*
#include "./clib/hello.h"
#cgo LDFLAGS: -L./clib -lhello
*/
import "C"
import "fmt"

func main() {
	C.helloFromC()
	fmt.Println("Hello World!")
}
