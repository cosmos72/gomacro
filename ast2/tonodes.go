/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
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
 * tonodes.go
 *
 *  Created on Mar 05, 2018
 *      Author Massimiliano Ghilardi
 */

package ast2

import (
	"go/ast"
	"go/token"
)

// ToNode recursively traverses Ast and extracts all the contained ast.Node:s
func ToNodes(x Ast) []ast.Node {
	return ToNodesAppend(nil, x)
}

func ToNodesAppend(dst []ast.Node, x Ast) []ast.Node {
	switch x := x.(type) {
	case nil:
		return dst
	case File:
		dst = collectImports(dst, x.X.Imports)
		// treat as AstWithSlice to traverse File contents
		break
	case AstWithNode:
		if x != nil {
			dst = append(dst, x.Node())
		}
		return dst
	case NodeSlice:
		// faster than generic AstWithSlice
		return append(dst, x.X...)
	case AstWithSlice:
		break
	default:
		y := x.Interface()
		errorf("cannot convert to []ast.Node: %v // %T", y, y)
		return nil
	}
	form, ok := x.(AstWithSlice)
	if ok && form != nil {
		n := form.Size()
		for i := 0; i < n; i++ {
			dst = ToNodesAppend(dst, form.Get(i))
		}
	}
	return dst
}

func collectImports(dst []ast.Node, imports []*ast.ImportSpec) []ast.Node {
	if n := len(imports); n != 0 {
		specs := make([]ast.Spec, n)
		for i, imp := range imports {
			specs[i] = imp
		}
		dst = append(dst, &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: specs,
		})
	}
	return dst
}
