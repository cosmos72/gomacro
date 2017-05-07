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
 * type.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package type2

import (
	"go/types"
	"reflect"
)

// IsMethod reports whether a function type's contains a receiver, i.e. is a method.
// It panics if the type's Kind is not Func.
func (t *timpl) IsMethod() bool {
	if t.Kind() != reflect.Func {
		errorf("IsMethod of non-func type %v", t)
	}
	gtype := t.gtype.(*types.Signature)
	return gtype.Recv() != nil
}

// IsVariadic reports whether a function type's final input parameter
// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
// implicit actual type []T.
// IsVariadic panics if the type's Kind is not Func.
func (t *timpl) IsVariadic() bool {
	if t.Kind() != reflect.Func {
		errorf("In of non-func type %v", t)
	}
	return t.rtype.IsVariadic()
}

// In returns the type of a function type's i'th input parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumIn()).
func (t *timpl) In(i int) Type {
	if t.Kind() != reflect.Func {
		errorf("In of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	va := gtype.Params().At(i)
	if gtype.Recv() != nil {
		i++ // skip the receiver in reflect.Type
	}
	return maketype(va.Type(), t.rtype.In(i))
}

// NumIn returns a function type's input parameter count.
// It panics if the type's Kind is not Func.
func (t *timpl) NumIn() int {
	if t.Kind() != reflect.Func {
		errorf("NumIn of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	return gtype.Params().Len()
}

// NumOut returns a function type's output parameter count.
// It panics if the type's Kind is not Func.
func (t *timpl) NumOut() int {
	if t.Kind() != reflect.Func {
		errorf("NumOut of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	return gtype.Results().Len()
}

// Out returns the type of a function type's i'th output parameter.
// It panics if the type's Kind is not Func.
// It panics if i is not in the range [0, NumOut()).
func (t *timpl) Out(i int) Type {
	if t.Kind() != reflect.Func {
		errorf("Out of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	va := gtype.Results().At(i)
	return maketype(va.Type(), t.rtype.Out(i))
}

// Recv returns the type of a method type's receiver parameter.
// It panics if the type's Kind is not Func.
// It returns Type{} if t has no receiver.
func (t *timpl) Recv() Type {
	if t.Kind() != reflect.Func {
		errorf("Recv of non-func type %v", t)
	}
	gtype := t.underlying().(*types.Signature)
	va := gtype.Recv()
	if va == nil {
		return Type{}
	}
	return maketype(va.Type(), t.rtype.In(0))
}

func FuncOf(in []Type, out []Type, variadic bool) Type {
	return MethodOf(Type{}, in, out, variadic)
}

func MethodOf(recv Type, in []Type, out []Type, variadic bool) Type {
	gin := toGoTuple(in)
	gout := toGoTuple(out)
	rin := toReflectTypes(in)
	rout := toReflectTypes(out)
	var grecv *types.Var
	if recv.timpl != nil {
		rin = append([]reflect.Type{recv.rtype}, rin...)
		grecv = toGoParam(recv)
	}
	return maketype(
		types.NewSignature(grecv, gin, gout, variadic),
		reflect.FuncOf(rin, rout, variadic),
	)
}
