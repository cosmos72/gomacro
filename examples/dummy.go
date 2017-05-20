// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
)

type Cons struct {
	A, B int
}

func main() {
	p := Cons{1, 2}
	var i interface{} = p
	q := i.(Cons)
	q.A = 8
	fmt.Printf("%#v\n", p)
}

func main2() {
	interface1 := types.NewInterface(nil, nil).Complete()
	fmt.Printf("%v\n", interface1)

	pkg := types.NewPackage("main", "main")
	ret := types.NewVar(token.NoPos, pkg, "", types.Typ[types.Int])
	sig := types.NewSignature(nil, types.NewTuple(), types.NewTuple(ret), false)
	cap := types.NewFunc(token.NoPos, pkg, "Cap", sig)
	len := types.NewFunc(token.NoPos, pkg, "Len", sig)
	interface2 := types.NewInterface([]*types.Func{cap, len}, nil).Complete()
	fmt.Printf("%v\n", interface2)

	var m typeutil.Map
	m.Set(interface1, interface1)
	m.Set(interface2, interface2)
	fmt.Printf("%v\n", m.At(interface1))
	fmt.Printf("%v\n", m.At(interface2))
	fmt.Printf("%v\n", m.Len())
}
