// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"io"
	"os"
)

func fib(a, b int) (int, int) { return b, a + b }

func main() {
	x, y := 1, 2
	println(x, y)
	x, y = fib(x, y)
	println(x, y)
	x, y = fib(fib(x, y))
	println(x, y)
	x, y = fib(fib(fib(x, y)))
	println(x, y)
	// m := [...]int{0x7ffffff: 3}
	// fmt.Println(m)
	// p := Pair{A: 1, B: true}
	// Pair{1, 2} = Pair{}
	// var f os.file
	// _ = bytes.Buffer{nil, 0}
}

func main1() {
	var x io.ReadWriteCloser = os.Stdin
	f := io.ReadWriteCloser.Close
	f(x)
	fmt.Printf("%T\n", f)
}
