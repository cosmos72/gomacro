// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"unsafe"
)

func compareClosures() {
	for i := 0; i <= 3; i++ {
		i := i
		f := func() int { return i }
		fmt.Printf("%v %x\n", f, **(**uintptr)(unsafe.Pointer(&f)))
	}
}

func main() {
	var e1 *int
	var e2 *int = e1
	var i interface{} = e2
	println(e1 == nil)
	println(e2 == nil)
	println(i == nil)
}
