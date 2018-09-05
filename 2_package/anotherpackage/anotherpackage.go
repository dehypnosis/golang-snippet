package anotherpackage

import (
	"errors"
	"fmt"
)

// ErrXx blabla is blabla.
var ErrXx error = errors.New("unexpected EOF")

// SomeFunc blabla.
func SomeFunc() error {
	fmt.Printf(ErrXx.Error())
	return ErrXx
}
