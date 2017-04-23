// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"reflect"
)

/*
func compareClosures() {
	for i := 0; i <= 3; i++ {
		i := i
		f := func() int { return i }
		fmt.Printf("%v %x\n", f, **(**uintptr)(unsafe.Pointer(&f)))
	}
}
*/

type IFoo interface {
	foo()
}
type IBar interface {
	bar()
}

type Foo struct {
}
type Bar struct {
}

func (Foo) foo() {
}
func (Bar) bar() {
}

var TypeOfIBar = reflect.TypeOf((*IBar)(nil)).Elem()

func main() {
	fmt.Printf("%v", nil)
}
