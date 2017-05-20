// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"go/token"
	"go/types"
	r "reflect"

	"golang.org/x/tools/go/types/typeutil"
)

type EmbedBasicType struct {
	int
}

func main() {
	var e EmbedBasicType
	fmt.Println(e.int)
	v := r.ValueOf(e)
	vi := v.Field(0)
	vi.Interface()
}

func add(a, b int) int {
	return a + b
}

func curry(f func(a, b int) int, a int) func(b int) int {
	return func(b int) int {
		return f(a, b)
	}
}

func reflectCurry(f r.Value, a r.Value) func([]r.Value) []r.Value {
	return func(args []r.Value) []r.Value {
		args = append([]r.Value{a}, args...)
		return f.Call(args)
	}
}

func main0() {
	add1 := curry(add, 1)
	fmt.Printf("add1 = %T\n", add1)
	fmt.Printf("add1(5) = %v\n", add1(5))

	add1refl := reflectCurry(r.ValueOf(add), r.ValueOf(1))
	fmt.Printf("add1refl = %T\n", add1refl)
	fmt.Printf("add1refl(5) = %v\n", add1refl([]r.Value{r.ValueOf(5)})[0])
}

func main1() {
	type Cons struct {
		A, B int
	}
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
