/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * selector.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

// SelectorExpr compiles foo.bar, i.e. read access to methods and struct fields
func (c *Comp) SelectorExpr(node *ast.SelectorExpr) *Expr {
	e := c.Expr1(node.X)
	t := e.Type
	name := node.Sel.Name
	switch t.Kind() {
	case r.Ptr:
		t = t.Elem()
		if t.Kind() != r.Struct {
			break
		}
		fun := e.AsX1()
		e = exprFun(t, func(env *Env) r.Value {
			return fun(env).Elem()
		})
		fallthrough
	case r.Struct:
		field, fieldok, mtd, mtdok := c.LookupFieldOrMethod(t, name)
		if fieldok {
			return c.compileField(e, field)
		} else if mtdok {
			return c.compileMethod(e, mtd)
		}
	default:
		// interfaces and named types can have methods, but no fields
		if t.NumMethod() != 0 {
			mtd, mtdn := c.LookupMethod(t, name)
			switch mtdn {
			case 0:
			case 1:
				return c.compileMethod(e, mtd)
			default:
				c.Errorf("type %s has %d methods %q, expression is ambiguous: %v", t, mtdn, name, node)
			}
		}
	}
	c.Errorf("type %s has no field or method %q: %v", t, name, node)
	return nil
}

// lookup fields and methods at the same time... it's and error if both exist at the same depth
func (c *Comp) LookupFieldOrMethod(t xr.Type, name string) (xr.StructField, bool, xr.Method, bool) {
	field, fieldn := c.LookupField(t, name)
	mtd, mtdn := c.LookupMethod(t, name)
	fielddepth := len(field.Index)
	mtddepth := len(mtd.FieldIndex) + 1
	if fieldn != 0 && mtdn != 0 {
		if fielddepth < mtddepth {
			// prefer the field
			mtdn = 0
		} else if fielddepth > mtddepth {
			// prefer the method
			fieldn = 0
		} else {
			c.Errorf("type %v has %d field(s) and %d method(s) named %q at depth %d",
				t, fieldn, mtdn, name, fielddepth)
		}
	}
	if fieldn > 1 {
		c.Errorf("type %v has %d fields named %q at depth %d", t, fieldn, name, fielddepth)
	} else if mtdn > 1 {
		c.Errorf("type %v has %d methods named %q at depth %d", t, mtdn, name, mtddepth)
	}
	return field, fieldn == 1, mtd, mtdn == 1
}

// LookupField performs a breadth-first search for struct field with given name
func (c *Comp) LookupField(t xr.Type, name string) (field xr.StructField, numfound int) {
	return t.FieldByName(name, c.FileComp().Packagename)
}

// LookupMethod performs a breadth-first search for method with given name
func (c *Comp) LookupMethod(t xr.Type, name string) (mtd xr.Method, numfound int) {
	return t.MethodByName(name, c.FileComp().Packagename)
}

func (c *Comp) compileField(e *Expr, field xr.StructField) *Expr {
	objfun := e.AsX1()
	t := field.Type
	var fun I
	index := field.Index
	// c.Debugf("compileField: field=%#v", field)
	if len(index) == 1 {
		index0 := index[0]
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				return obj.Field(index0).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				return int(obj.Field(index0).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				return int8(obj.Field(index0).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				return int16(obj.Field(index0).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				return int32(obj.Field(index0).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				return obj.Field(index0).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				return uint(obj.Field(index0).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				return uint8(obj.Field(index0).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				return uint16(obj.Field(index0).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				return uint32(obj.Field(index0).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				return obj.Field(index0).Uint()
			}
		case r.Uintptr:
			fun = func(env *Env) uintptr {
				obj := objfun(env)
				return uintptr(obj.Field(index0).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				return float32(obj.Field(index0).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				return obj.Field(index0).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				return complex64(obj.Field(index0).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				return obj.Field(index0).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj.Field(index0).String()
			}
		default:
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				return obj.Field(index0)
			}
		}
	} else {
		switch t.Kind() {
		case r.Bool:
			fun = func(env *Env) bool {
				obj := objfun(env)
				return obj.FieldByIndex(index).Bool()
			}
		case r.Int:
			fun = func(env *Env) int {
				obj := objfun(env)
				return int(obj.FieldByIndex(index).Int())
			}
		case r.Int8:
			fun = func(env *Env) int8 {
				obj := objfun(env)
				return int8(obj.FieldByIndex(index).Int())
			}
		case r.Int16:
			fun = func(env *Env) int16 {
				obj := objfun(env)
				return int16(obj.FieldByIndex(index).Int())
			}
		case r.Int32:
			fun = func(env *Env) int32 {
				obj := objfun(env)
				return int32(obj.FieldByIndex(index).Int())
			}
		case r.Int64:
			fun = func(env *Env) int64 {
				obj := objfun(env)
				return obj.FieldByIndex(index).Int()
			}
		case r.Uint:
			fun = func(env *Env) uint {
				obj := objfun(env)
				return uint(obj.FieldByIndex(index).Uint())
			}
		case r.Uint8:
			fun = func(env *Env) uint8 {
				obj := objfun(env)
				return uint8(obj.FieldByIndex(index).Uint())
			}
		case r.Uint16:
			fun = func(env *Env) uint16 {
				obj := objfun(env)
				return uint16(obj.FieldByIndex(index).Uint())
			}
		case r.Uint32:
			fun = func(env *Env) uint32 {
				obj := objfun(env)
				return uint32(obj.FieldByIndex(index).Uint())
			}
		case r.Uint64:
			fun = func(env *Env) uint64 {
				obj := objfun(env)
				return obj.FieldByIndex(index).Uint()
			}
		case r.Uintptr:

			fun = func(env *Env) uintptr {
				obj := objfun(env)
				return uintptr(obj.FieldByIndex(index).Uint())
			}
		case r.Float32:
			fun = func(env *Env) float32 {
				obj := objfun(env)
				return float32(obj.FieldByIndex(index).Float())
			}
		case r.Float64:
			fun = func(env *Env) float64 {
				obj := objfun(env)
				return obj.FieldByIndex(index).Float()
			}
		case r.Complex64:
			fun = func(env *Env) complex64 {
				obj := objfun(env)
				return complex64(obj.FieldByIndex(index).Complex())
			}
		case r.Complex128:
			fun = func(env *Env) complex128 {
				obj := objfun(env)
				return obj.FieldByIndex(index).Complex()
			}
		case r.String:
			fun = func(env *Env) string {
				obj := objfun(env)
				return obj.FieldByIndex(index).String()
			}
		default:
			fun = func(env *Env) r.Value {
				obj := objfun(env)
				return obj.FieldByIndex(index)
			}
		}
	}
	return exprFun(t, fun)
}

func (c *Comp) removeReceiver(t xr.Type) (trecv, tfunc xr.Type) {
	nin, nout := t.NumIn(), t.NumOut()
	params := make([]xr.Type, nin)
	results := make([]xr.Type, nout)
	for i := 0; i < nin; i++ {
		params[i] = t.In(i)
	}
	for i := 0; i < nout; i++ {
		results[i] = t.Out(i)
	}
	if trecv = t.Recv(); trecv == nil {
		trecv = params[0]
		params = params[1:]
	}
	return trecv, c.Universe.FuncOf(params, results, t.IsVariadic())
}

// compileMethod compiles a method call.
// relatively slow, but simple: return a closure with the receiver already bound
func (c *Comp) compileMethod(e *Expr, mtd xr.Method) *Expr {
	fieldindex := mtd.FieldIndex
	t := e.Type

	// descend embedded fields
	for i, index := range fieldindex {
		if i > 0 && t.Kind() == r.Ptr && t.Elem().Kind() == r.Struct {
			// embedded field is a pointer, dereference it.
			t = t.Elem()
		}
		t = t.Field(index).Type
	}
	rtype := t.ReflectType()
	index := mtd.Index
	tmethod := t.Method(index)
	tfunc := tmethod.Type
	trecv, tclosure := c.removeReceiver(tfunc)
	rtclosure := tclosure.ReflectType()

	recvPointer := trecv.Kind() == r.Ptr // method with pointer receiver?
	objPointer := t.Kind() == r.Ptr      // field is pointer?
	addressof := recvPointer && !objPointer
	deref := !recvPointer && objPointer

	if addressof && xr.PtrTo(t).AssignableTo(trecv) {
		// ok
	} else if deref && t.Elem().AssignableTo(trecv) {
		// ok
	} else if !t.AssignableTo(trecv) {
		c.Errorf("cannot use <%v> as <%v> in receiver of method <%v>", t, trecv, tfunc)
	}

	objfun := e.AsX1()
	var ret func(env *Env) r.Value

	if t.NumMethod() == rtype.NumMethod() && t.Named() && SameName(t.Name(), t.PkgPath(), rtype.Name(), rtype.PkgPath()) {
		// closures for methods declared by compiled code are accessible
		// simply with reflect.Value.Method(index). Easy.
		switch len(fieldindex) {
		case 0:
			ret = func(env *Env) r.Value {
				obj := objfun(env)
				return obj.Method(index)
			}
		case 1:
			fieldindex := fieldindex[0]
			ret = func(env *Env) r.Value {
				obj := objfun(env).Field(fieldindex)
				return obj.Method(index)
			}
		default:
			ret = func(env *Env) r.Value {
				obj := objfun(env).FieldByIndex(fieldindex)
				return obj.Method(index)
			}
		}
	} else if tmethod.Func == Nil {
		c.Errorf("method declared but not yet implemented: %v.%v", t, tmethod.Name)
	} else {
		// method declared by interpreted code, manually build the closure.
		//
		// It's not possible to call r.MakeFunc() only once at compile-time,
		// because the closure passed to it needs access to a variable holding the receiver.
		// such variable would be allocated only once at compile-time,
		// not once per goroutine!
		fun := tmethod.Func
		nin := tclosure.NumIn() + 1

		switch len(fieldindex) {
		case 0:
			ret = func(env *Env) r.Value {
				obj := objfun(env)
				if addressof {
					obj = obj.Addr()
				} else if deref {
					obj = obj.Elem()
				}
				return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
					fullargs := make([]r.Value, nin)
					fullargs[0] = obj
					copy(fullargs[1:], args)
					// Debugf("invoking <%v> with args %v", fun.Type(), fullargs)
					return fun.Call(fullargs)
				})
			}
		case 1:
			fieldindex := fieldindex[0]
			ret = func(env *Env) r.Value {
				obj := objfun(env).Field(fieldindex)
				if addressof {
					obj = obj.Addr()
				} else if deref {
					obj = obj.Elem()
				}
				return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
					fullargs := make([]r.Value, nin)
					fullargs[0] = obj
					copy(fullargs[1:], args)
					// Debugf("invoking <%v> with args %v", fun.Type(), fullargs)
					return fun.Call(fullargs)
				})
			}
		default:
			ret = func(env *Env) r.Value {
				obj := objfun(env).FieldByIndex(fieldindex)
				if addressof {
					obj = obj.Addr()
				} else if deref {
					obj = obj.Elem()
				}
				return r.MakeFunc(rtclosure, func(args []r.Value) []r.Value {
					fullargs := make([]r.Value, nin)
					fullargs[0] = obj
					copy(fullargs[1:], args)
					// Debugf("invoking <%v> with args %v", fun.Type(), fullargs)
					return fun.Call(fullargs)
				})
			}
		}
	}
	return exprX1(tclosure, ret)
}

// SelectorPlace compiles a.b returning a settable and addressable Place
func (c *Comp) SelectorPlace(node *ast.SelectorExpr, opt PlaceOption) *Place {
	obje := c.Expr1(node.X)
	te := obje.Type
	name := node.Sel.Name
	switch te.Kind() {
	case r.Ptr:
		te = te.Elem()
		if te.Kind() != r.Struct {
			break
		}
		objfun := obje.AsX1()
		obje = exprFun(te, func(env *Env) r.Value {
			obj := objfun(env)
			// Debugf("SelectorPlace: obj = %v <%v> (expecting pointer to struct)", obj, obj.Type())
			return obj.Elem()
		})
		fallthrough
	case r.Struct:
		field, fieldn := c.LookupField(te, name)
		if fieldn == 0 {
			break
		} else if fieldn > 1 {
			c.Errorf("type %v has %d fields named %q, all at depth %d", te, fieldn, name, len(field.Index))
			return nil
		}
		if te.Kind() == r.Struct {
			c.checkSettableField(node)
		}
		return c.compileFieldPlace(obje, field)
	}
	c.Errorf("type %v has no field %q: %v", te, name, node)
	return nil
}

// checkSettableField check that a struct field is settable and addressable.
// by Go specs, this requires the struct itself to be settable and addressable.
func (c *Comp) checkSettableField(node *ast.SelectorExpr) {
	panicking := true
	defer func() {
		if panicking {
			recover()
			c.Pos = node.Pos()
			c.Errorf("cannot assign to %v", node)
		}
	}()
	c.placeOrAddress(node.X, PlaceAddress)
	panicking = false
}

func (c *Comp) compileFieldPlace(obje *Expr, field xr.StructField) *Place {
	// c.Debugf("compileFieldPlace: field=%#v", field)
	objfun := obje.AsX1()
	t := field.Type
	var fun, addr func(*Env) r.Value
	index := field.Index
	if len(index) == 1 {
		index0 := index[0]
		fun = func(env *Env) r.Value {
			obj := objfun(env)
			return obj.Field(index0)
		}
		addr = func(env *Env) r.Value {
			obj := objfun(env)
			return obj.Field(index0).Addr()
		}
	} else {
		fun = func(env *Env) r.Value {
			obj := objfun(env)
			return obj.FieldByIndex(index)
		}
		addr = func(env *Env) r.Value {
			obj := objfun(env)
			return obj.FieldByIndex(index).Addr()
		}
	}
	return &Place{Var: Var{Type: t, Name: field.Name}, Fun: fun, Addr: addr}
}
