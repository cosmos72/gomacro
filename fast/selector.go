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

	"github.com/cosmos72/gomacro/base"
)

type rMethod struct {
	r.Method
	FieldIndexes []int
}

type rField r.StructField

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
	}
	c.Errorf("type %v has no field or method %q: %v", t, name, node)
	return nil
}

// lookup fields and methods at the same time... it's and error if both exist at the same depth
func (c *Comp) LookupFieldOrMethod(t r.Type, name string) (rField, bool, rMethod, bool) {
	field, fieldn := c.lookupField(t, name)
	mtd, mtdn := c.lookupMethod(t, name)
	fielddepth := len(field.Index)
	mtddepth := len(mtd.FieldIndexes) + 1
	if fieldn != 0 && mtdn != 0 {
		if fielddepth < mtddepth {
			// prefer the field
			mtdn = 0
		} else if fielddepth > mtddepth {
			// prefer the method
			fieldn = 0
		} else {
			c.Warnf("type %v has both %d field(s) and %d method(s) %q at the same depth=%d. this should not happen... using the method(s)",
				t, fieldn, mtdn, name, fielddepth)
			fieldn = 0
		}
	}
	if fieldn > 1 {
		c.Errorf("type %v has %d fields named %q, all at depth %d", t, fieldn, name, fielddepth)
	} else if mtdn > 1 {
		c.Errorf("type %v has %d methods named %q, all at depth %d", t, mtdn, name, mtddepth)
	}
	return field, fieldn == 1, mtd, mtdn == 1
}

// lookupField performs a breadth-first search for struct field with given name
func (c *Comp) lookupField(t r.Type, name string) (field rField, numfound int) {
	return c.lookupField0(t, 0, nil, name)
}

func (c *Comp) lookupField0(t r.Type, offset uintptr, index []int, name string) (field rField, numfound int) {
	var recurse []r.StructField
	if t.Kind() == r.Ptr {
		t = t.Elem()
		offset = 0
	}
	n := t.NumField()
	private := isPrivateFieldName(name)
	for i := 0; i < n; i++ {
		fieldi := t.Field(i)
		// check for exported/private (unexported) field
		if fieldi.Name == name {
			field = rField(fieldi)
			numfound++
			continue
		}
		// check for GensymPrivate fields - they emulate private fields
		if private && base.IsGensymPrivate(fieldi.Name) {
			fname := fieldi.Name[len(base.StrGensymPrivate):]
			if fname == name {
				field = rField(fieldi)
				numfound++
				continue
			}
		}
		// check for embedded fields
		if fieldi.Anonymous {
			if fieldi.Type.Name() == name {
				field = rField(fieldi)
				numfound++
				continue
			} else {
				recurse = append(recurse, fieldi)
			}
		}
		// check for GensymEmbedded fields - they emulate embedded fields
		if base.IsGensymEmbedded(fieldi.Name) {
			fname := fieldi.Name[len(base.StrGensymEmbedded):]
			if fname == name || c.NamedTypes[fieldi.Type].Name == name {
				// c.Debugf("lookupField: found GensymEmbedded field: %v", fieldi)
				field = rField(fieldi)
				field.Anonymous = true
				numfound++
				continue
			} else {
				recurse = append(recurse, fieldi)
			}
		}
	}
	if numfound != 0 {
		field.Offset += offset
		field.Index = append(index, field.Index...)
		return
	}
	// breadth-first recursion
	depth := len(index)
	index = append(index, 0)
	for _, recursei := range recurse {
		index[depth] = recursei.Index[len(recursei.Index)-1]
		field, n = c.lookupField0(recursei.Type, offset+recursei.Offset, index[:], name) // pass a copy of slice
		numfound += n
	}
	return
}

func isPrivateFieldName(name string) bool {
	if len(name) == 0 {
		return false
	}
	ch := name[0]
	return ch >= 'a' && ch <= 'z' || ch == '_'
}

func (c *Comp) lookupMethod(t r.Type, name string) (mtd rMethod, count int) {
	m, ok := t.MethodByName(name)
	if ok {
		mtd = rMethod{Method: m, FieldIndexes: nil}
		count++
	}
	return mtd, count

	// TODO check for methods declared on embedded field
	/*
		for i := 0; i < t.NumField(); i++ {
			if field := t.Field(i); field.Anonymous {
				if emtd, eok := methodByNameDepth(field.Type, name); eok {
				}
			}
		}
	*/
}

func (c *Comp) compileField(e *Expr, field rField) *Expr {
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

func (c *Comp) compileMethod(e *Expr, mtd rMethod) *Expr {
	// slow, but simple: return a closure with the receiver already bound
	index := mtd.Index
	t := r.Zero(e.Type).Method(index).Type()
	objfun := e.AsX1()
	fun := func(env *Env) r.Value {
		obj := objfun(env)
		return obj.Method(index)
	}
	return exprFun(t, fun)
}

// SelectorPlace compiles a.b returning a settable and addressable Place
func (c *Comp) SelectorPlace(node *ast.SelectorExpr, opt PlaceOption) *Place {
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
		field, fieldn := c.lookupField(t, name)
		if fieldn == 0 {
			break
		} else if fieldn > 1 {
			c.Errorf("type %v has %d fields named %q, all at depth %d", t, fieldn, name, len(field.Index))
			return nil
		}
		if t.Kind() == r.Struct {
			c.checkSettableField(node)
		}
		return c.compileFieldPlace(e, field)
	}
	c.Errorf("type %v has no field %q: %v", t, name, node)
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

func (c *Comp) compileFieldPlace(e *Expr, field rField) *Place {
	// c.Debugf("compileFieldPlace: field=%#v", field)
	objfun := e.AsX1()
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
