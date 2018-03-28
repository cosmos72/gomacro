/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * nummethod.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import "go/types"

// NumMethod returns the *total* number of methods for interface or named type t,
// including wrapper methods for embedded fields or embedded interfaces.
// Note: it has slightly different semantics from go/types.(*Named).NumMethods(),
//       since the latter returns 0 for named interfaces, and callers need to manually invoke
//       goNamedType.Underlying().NumMethods() to retrieve the number of methods
//       of a named interface
func (t *xtype) NumMethod() int {
	return goTypeNumMethod(t.gtype)
}

// recursively count total number of methods for type t,
// including wrapper methods for embedded fields or embedded interfaces
func goTypeNumMethod(gt types.Type) int {
	count := 0
again:
	switch t := gt.(type) {
	case *types.Named:
		count += t.NumMethods()
		u := t.Underlying()
		if u != gt {
			gt = u
			goto again
		}
	case *types.Interface:
		count += t.NumMethods()
	case *types.Struct:
		n := t.NumFields()
		for i := 0; i < n; i++ {
			if f := t.Field(i); f.Anonymous() {
				count += goTypeNumMethod(f.Type())
			}
		}
	}
	return count
}
