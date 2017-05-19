// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"reflect"
)

func f1() interface{} {
	type Pair struct {
		A, B int
	}
	return Pair{}
}

func f2() interface{} {
	type Pair struct {
		A, B int
	}
	return Pair{}
}

func main() {
	p1, p2 := f1(), f2()
	t1, t2 := reflect.TypeOf(p1), reflect.TypeOf(p2)
	fmt.Printf("%v %T\n", p1, p1)
	fmt.Printf("%v %T\n", p2, p2)
	fmt.Printf("%t %t\n", p1 == p2, t1 == t2)
	fmt.Printf("%p %#v\n", t1, t1)
	fmt.Printf("%p %#v\n", t2, t2)
}
