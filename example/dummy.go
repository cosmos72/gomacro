// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var ch1, ch2 chan int
	select {
	case i := <-ch1:
		println(i)
		break
	case ch2 <- 0:
		break
	}
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
