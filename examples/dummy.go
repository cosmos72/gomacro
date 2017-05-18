// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
)

func add3(*bool, x, y int) int {
	return  x + y
}

func main() {
	fmt.Printf("%T\n", add3)
	fmt.Printf("%v\n", add3(nil, 2, 3))
}
