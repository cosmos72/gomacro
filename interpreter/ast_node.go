/*
 * gomacro - A Go intepreter with Lisp-like macros
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
 * ast_node.go
 *
 *  Created on Feb 25, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
)

func (x ArrayType) Interface() interface{}      { return x.p }
func (x AssignStmt) Interface() interface{}     { return x.p }
func (x BadDecl) Interface() interface{}        { return x.p }
func (x BadExpr) Interface() interface{}        { return x.p }
func (x BadStmt) Interface() interface{}        { return x.p }
func (x BasicLit) Interface() interface{}       { return x.p }
func (x BinaryExpr) Interface() interface{}     { return x.p }
func (x BranchStmt) Interface() interface{}     { return x.p }
func (x CallExpr) Interface() interface{}       { return x.p }
func (x CaseClause) Interface() interface{}     { return x.p }
func (x ChanType) Interface() interface{}       { return x.p }
func (x CommClause) Interface() interface{}     { return x.p }
func (x CompositeLit) Interface() interface{}   { return x.p }
func (x DeclStmt) Interface() interface{}       { return x.p }
func (x DeferStmt) Interface() interface{}      { return x.p }
func (x Ellipsis) Interface() interface{}       { return x.p }
func (x EmptyStmt) Interface() interface{}      { return x.p }
func (x ExprStmt) Interface() interface{}       { return x.p }
func (x Field) Interface() interface{}          { return x.p }
func (x ForStmt) Interface() interface{}        { return x.p }
func (x FuncDecl) Interface() interface{}       { return x.p }
func (x FuncLit) Interface() interface{}        { return x.p }
func (x FuncType) Interface() interface{}       { return x.p }
func (x GoStmt) Interface() interface{}         { return x.p }
func (x Ident) Interface() interface{}          { return x.p }
func (x IfStmt) Interface() interface{}         { return x.p }
func (x ImportSpec) Interface() interface{}     { return x.p }
func (x IncDecStmt) Interface() interface{}     { return x.p }
func (x IndexExpr) Interface() interface{}      { return x.p }
func (x InterfaceType) Interface() interface{}  { return x.p }
func (x KeyValueExpr) Interface() interface{}   { return x.p }
func (x LabeledStmt) Interface() interface{}    { return x.p }
func (x MapType) Interface() interface{}        { return x.p }
func (x Package) Interface() interface{}        { return x.p }
func (x ParenExpr) Interface() interface{}      { return x.p }
func (x RangeStmt) Interface() interface{}      { return x.p }
func (x SelectStmt) Interface() interface{}     { return x.p }
func (x SelectorExpr) Interface() interface{}   { return x.p }
func (x SendStmt) Interface() interface{}       { return x.p }
func (x SliceExpr) Interface() interface{}      { return x.p }
func (x StarExpr) Interface() interface{}       { return x.p }
func (x StructType) Interface() interface{}     { return x.p }
func (x SwitchStmt) Interface() interface{}     { return x.p }
func (x TypeAssertExpr) Interface() interface{} { return x.p }
func (x TypeSpec) Interface() interface{}       { return x.p }
func (x TypeSwitchStmt) Interface() interface{} { return x.p }
func (x UnaryExpr) Interface() interface{}      { return x.p }
func (x ValueSpec) Interface() interface{}      { return x.p }

func (x ArrayType) Node() ast.Node      { return x.p }
func (x AssignStmt) Node() ast.Node     { return x.p }
func (x BadDecl) Node() ast.Node        { return x.p }
func (x BadExpr) Node() ast.Node        { return x.p }
func (x BadStmt) Node() ast.Node        { return x.p }
func (x BasicLit) Node() ast.Node       { return x.p }
func (x BinaryExpr) Node() ast.Node     { return x.p }
func (x BranchStmt) Node() ast.Node     { return x.p }
func (x CallExpr) Node() ast.Node       { return x.p }
func (x CaseClause) Node() ast.Node     { return x.p }
func (x ChanType) Node() ast.Node       { return x.p }
func (x CommClause) Node() ast.Node     { return x.p }
func (x CompositeLit) Node() ast.Node   { return x.p }
func (x DeclStmt) Node() ast.Node       { return x.p }
func (x DeferStmt) Node() ast.Node      { return x.p }
func (x Ellipsis) Node() ast.Node       { return x.p }
func (x EmptyStmt) Node() ast.Node      { return x.p }
func (x ExprStmt) Node() ast.Node       { return x.p }
func (x Field) Node() ast.Node          { return x.p }
func (x ForStmt) Node() ast.Node        { return x.p }
func (x FuncDecl) Node() ast.Node       { return x.p }
func (x FuncLit) Node() ast.Node        { return x.p }
func (x FuncType) Node() ast.Node       { return x.p }
func (x GoStmt) Node() ast.Node         { return x.p }
func (x Ident) Node() ast.Node          { return x.p }
func (x IfStmt) Node() ast.Node         { return x.p }
func (x ImportSpec) Node() ast.Node     { return x.p }
func (x IncDecStmt) Node() ast.Node     { return x.p }
func (x IndexExpr) Node() ast.Node      { return x.p }
func (x InterfaceType) Node() ast.Node  { return x.p }
func (x KeyValueExpr) Node() ast.Node   { return x.p }
func (x LabeledStmt) Node() ast.Node    { return x.p }
func (x MapType) Node() ast.Node        { return x.p }
func (x Package) Node() ast.Node        { return x.p }
func (x ParenExpr) Node() ast.Node      { return x.p }
func (x RangeStmt) Node() ast.Node      { return x.p }
func (x SelectStmt) Node() ast.Node     { return x.p }
func (x SelectorExpr) Node() ast.Node   { return x.p }
func (x SendStmt) Node() ast.Node       { return x.p }
func (x SliceExpr) Node() ast.Node      { return x.p }
func (x StarExpr) Node() ast.Node       { return x.p }
func (x StructType) Node() ast.Node     { return x.p }
func (x SwitchStmt) Node() ast.Node     { return x.p }
func (x TypeAssertExpr) Node() ast.Node { return x.p }
func (x TypeSpec) Node() ast.Node       { return x.p }
func (x TypeSwitchStmt) Node() ast.Node { return x.p }
func (x UnaryExpr) Node() ast.Node      { return x.p }
func (x ValueSpec) Node() ast.Node      { return x.p }

func (x ArrayType) Op() token.Token  { return token.LBRACK }
func (x AssignStmt) Op() token.Token { return x.Op() }
func (x BadDecl) Op() token.Token    { return token.ILLEGAL }
func (x BadExpr) Op() token.Token    { return token.ILLEGAL }
func (x BadStmt) Op() token.Token    { return token.ILLEGAL }
func (x BasicLit) Op() token.Token   { return x.p.Kind }
func (x BinaryExpr) Op() token.Token { return x.p.Op }
func (x BranchStmt) Op() token.Token { return x.p.Tok }
func (x CallExpr) Op() token.Token   { return token.RPAREN }
func (x CaseClause) Op() token.Token {
	if len(x.p.List) != 0 {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}
func (x ChanType) Op() token.Token { return token.CHAN }
func (x CommClause) Op() token.Token {
	if x.p.Comm != nil {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}

func (x CompositeLit) Op() token.Token   { return token.RBRACE }
func (x DeclStmt) Op() token.Token       { return x.p.Decl.(*ast.GenDecl).Tok }
func (x DeferStmt) Op() token.Token      { return token.DEFER }
func (x Ellipsis) Op() token.Token       { return token.ELLIPSIS }
func (x EmptyStmt) Op() token.Token      { return token.SEMICOLON }
func (x ExprStmt) Op() token.Token       { return token.ELSE } // FIXME
func (x Field) Op() token.Token          { return token.PERIOD }
func (x ForStmt) Op() token.Token        { return token.FOR }
func (x FuncDecl) Op() token.Token       { return token.FUNC }
func (x FuncLit) Op() token.Token        { return token.FUNC }
func (x FuncType) Op() token.Token       { return token.FUNC }
func (x GoStmt) Op() token.Token         { return token.GO }
func (x Ident) Op() token.Token          { return token.IDENT }
func (x IfStmt) Op() token.Token         { return token.IF }
func (x ImportSpec) Op() token.Token     { return token.IMPORT }
func (x IncDecStmt) Op() token.Token     { return x.p.Tok }
func (x IndexExpr) Op() token.Token      { return token.RBRACK }
func (x InterfaceType) Op() token.Token  { return token.INTERFACE }
func (x KeyValueExpr) Op() token.Token   { return token.COLON } // FIXME
func (x LabeledStmt) Op() token.Token    { return token.COLON } // FIXME
func (x MapType) Op() token.Token        { return token.MAP }
func (x Package) Op() token.Token        { return token.PACKAGE }
func (x ParenExpr) Op() token.Token      { return token.RPAREN }
func (x RangeStmt) Op() token.Token      { return token.RANGE }
func (x SelectStmt) Op() token.Token     { return token.SELECT }
func (x SelectorExpr) Op() token.Token   { return token.CASE }
func (x SendStmt) Op() token.Token       { return token.CHAN }   // FIXME
func (x SliceExpr) Op() token.Token      { return token.RBRACK } // FIXME
func (x StarExpr) Op() token.Token       { return token.MUL }
func (x StructType) Op() token.Token     { return token.STRUCT }
func (x SwitchStmt) Op() token.Token     { return token.SWITCH }
func (x TypeAssertExpr) Op() token.Token { return token.TYPE } // FIXME
func (x TypeSpec) Op() token.Token       { return token.TYPE }
func (x TypeSwitchStmt) Op() token.Token { return token.SWITCH } // FIXME
func (x UnaryExpr) Op() token.Token      { return x.p.Op }
func (x ValueSpec) Op() token.Token      { return token.VAR } // can be VAR or CONST

func (x ArrayType) Size() int      { return 2 }
func (x AssignStmt) Size() int     { return 2 }
func (x BadDecl) Size() int        { return 0 }
func (x BadExpr) Size() int        { return 0 }
func (x BadStmt) Size() int        { return 0 }
func (x BasicLit) Size() int       { return 0 }
func (x BinaryExpr) Size() int     { return 2 }
func (x BranchStmt) Size() int     { return 1 }
func (x CallExpr) Size() int       { return 2 }
func (x CaseClause) Size() int     { return 2 }
func (x ChanType) Size() int       { return 1 }
func (x CommClause) Size() int     { return 2 }
func (x CompositeLit) Size() int   { return 2 }
func (x DeclStmt) Size() int       { return 1 }
func (x DeferStmt) Size() int      { return 1 }
func (x Ellipsis) Size() int       { return 2 }
func (x EmptyStmt) Size() int      { return 0 }
func (x ExprStmt) Size() int       { return 1 }
func (x Field) Size() int          { return 2 }
func (x ForStmt) Size() int        { return 4 }
func (x FuncDecl) Size() int       { return 4 }
func (x FuncLit) Size() int        { return 2 }
func (x FuncType) Size() int       { return 2 }
func (x GoStmt) Size() int         { return 1 }
func (x Ident) Size() int          { return 0 }
func (x IfStmt) Size() int         { return 4 }
func (x ImportSpec) Size() int     { return 2 }
func (x IncDecStmt) Size() int     { return 1 }
func (x IndexExpr) Size() int      { return 2 }
func (x InterfaceType) Size() int  { return 1 }
func (x KeyValueExpr) Size() int   { return 2 }
func (x LabeledStmt) Size() int    { return 2 }
func (x MapType) Size() int        { return 2 }
func (x Package) Size() int        { return 2 }
func (x ParenExpr) Size() int      { return 1 }
func (x RangeStmt) Size() int      { return 4 }
func (x SelectStmt) Size() int     { return 1 }
func (x SelectorExpr) Size() int   { return 2 }
func (x SendStmt) Size() int       { return 2 }
func (x SliceExpr) Size() int      { return 4 }
func (x StarExpr) Size() int       { return 1 }
func (x StructType) Size() int     { return 1 }
func (x SwitchStmt) Size() int     { return 3 }
func (x TypeAssertExpr) Size() int { return 2 }
func (x TypeSpec) Size() int       { return 2 }
func (x TypeSwitchStmt) Size() int { return 3 }
func (x UnaryExpr) Size() int      { return 1 }
func (x ValueSpec) Size() int      { return 3 }

func (x ArrayType) Get(i int) Ast { return ToAst2(i, x.p.Len, x.p.Elt) }
func (x AssignStmt) Get(i int) Ast {
	var addr *[]ast.Expr
	if i == 0 {
		addr = &x.p.Lhs
	} else if i == 1 {
		addr = &x.p.Rhs
	} else {
		return BadIndex(i, 2)
	}
	return ExprSlice{addr}
}
func (x BadDecl) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BadExpr) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BadStmt) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BasicLit) Get(i int) Ast   { return BadIndex(i, 0) }
func (x BinaryExpr) Get(i int) Ast { return ToAst2(i, x.p.X, x.p.Y) }
func (x BranchStmt) Get(i int) Ast { return Ident{x.p.Label} }
func (x CallExpr) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.p.Fun)
	} else if i == 1 {
		return ExprSlice{&x.p.Args}
	} else {
		return BadIndex(i, 2)
	}
}
func (x CaseClause) Get(i int) Ast {
	if i == 0 {
		return ExprSlice{&x.p.List}
	} else if i == 1 {
		return StmtSlice{&x.p.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x ChanType) Get(i int) Ast { return ToAst1(i, x.p.Value) }
func (x CommClause) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.p.Comm)
	} else if i == 1 {
		return StmtSlice{&x.p.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x CompositeLit) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.p.Type)
	} else if i == 1 {
		return ExprSlice{&x.p.Elts}
	} else {
		return BadIndex(i, 2)
	}
}
func (x DeclStmt) Get(i int) Ast  { return ToAst1(i, x.p.Decl) }
func (x DeferStmt) Get(i int) Ast { return CallExpr{x.p.Call} }
func (x Ellipsis) Get(i int) Ast  { return ToAst1(i, x.p.Elt) }
func (x EmptyStmt) Get(i int) Ast { return BadIndex(i, 0) }
func (x ExprStmt) Get(i int) Ast  { return ToAst1(i, x.p.X) }
func (x Field) Get(i int) Ast {
	if i == 0 {
		return IdentSlice{&x.p.Names}
	} else if i == 1 {
		return ToAst(x.p.Type)
	} else {
		return BadIndex(i, 2)
	}
}
func (x ForStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.Init
	case 1:
		node = x.p.Cond
	case 2:
		node = x.p.Post
	case 3:
		return BlockStmt{x.p.Body}
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x FuncDecl) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.Recv
	case 1:
		return Ident{x.p.Name}
	case 2:
		node = x.p.Type
	case 3:
		return BlockStmt{x.p.Body}
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x FuncLit) Get(i int) Ast {
	if i == 0 {
		return FuncType{x.p.Type}
	} else if i == 1 {
		return BlockStmt{x.p.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x FuncType) Get(i int) Ast {
	var ret *ast.FieldList
	if i == 0 {
		ret = x.p.Params
	} else if i == 1 {
		ret = x.p.Results
	} else {
		return BadIndex(i, 2)
	}
	return FieldList{ret}
}
func (x GoStmt) Get(i int) Ast { return CallExpr{x.p.Call} }
func (x Ident) Get(i int) Ast  { return BadIndex(i, 0) }
func (x IfStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.Init
	case 1:
		node = x.p.Cond
	case 2:
		return BlockStmt{x.p.Body}
	case 3:
		node = x.p.Else
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}

func (x ImportSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.p.Name}
	} else if i == 1 {
		return BasicLit{x.p.Path}
	} else {
		return BadIndex(i, 2)
	}
}
func (x IncDecStmt) Get(i int) Ast    { return ToAst1(i, x.p.X) }
func (x IndexExpr) Get(i int) Ast     { return ToAst2(i, x.p.X, x.p.Index) }
func (x InterfaceType) Get(i int) Ast { return FieldList{x.p.Methods} }
func (x KeyValueExpr) Get(i int) Ast  { return ToAst2(i, x.p.Key, x.p.Value) }
func (x LabeledStmt) Get(i int) Ast {
	if i == 0 {
		return Ident{x.p.Label}
	} else if i == 1 {
		return ToAst(x.p.Stmt)
	} else {
		return BadIndex(i, 2)
	}
}
func (x MapType) Get(i int) Ast   { return ToAst2(i, x.p.Key, x.p.Value) }
func (x Package) Get(i int) Ast   { return nil } // TODO
func (x ParenExpr) Get(i int) Ast { return ToAst1(i, x.p.X) }
func (x RangeStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.Key
	case 1:
		node = x.p.Value
	case 2:
		node = x.p.X
	case 3:
		return BlockStmt{x.p.Body}
	default:
		BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x SelectStmt) Get(i int) Ast   { return BlockStmt{x.p.Body} }
func (x SelectorExpr) Get(i int) Ast { return ToAst2(i, x.p.X, x.p.Sel) }
func (x SendStmt) Get(i int) Ast     { return ToAst2(i, x.p.Chan, x.p.Value) }
func (x SliceExpr) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.X
	case 1:
		node = x.p.Low
	case 2:
		node = x.p.High
	case 3:
		node = x.p.Max
	default:
		BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x StarExpr) Get(i int) Ast       { return ToAst1(i, x.p.X) }
func (x StructType) Get(i int) Ast     { return FieldList{x.p.Fields} }
func (x SwitchStmt) Get(i int) Ast     { return ToAst3(i, x.p.Init, x.p.Tag, x.p.Body) }
func (x TypeAssertExpr) Get(i int) Ast { return ToAst2(i, x.p.X, x.p.Type) }
func (x TypeSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.p.Name}
	} else if i == 1 {
		return ToAst(x.p.Type)
	} else {
		return BadIndex(i, 2)
	}
}
func (x TypeSwitchStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.p.Init
	case 1:
		node = x.p.Assign
	case 2:
		return BlockStmt{x.p.Body}
	default:
		BadIndex(i, 3)
	}
	return ToAst(node)
}
func (x UnaryExpr) Get(i int) Ast { return ToAst1(i, x.p.X) }
func (x ValueSpec) Get(i int) Ast {
	switch i {
	case 0:
		return IdentSlice{&x.p.Names}
	case 1:
		return ToAst(x.p.Type)
	case 2:
		return ExprSlice{&x.p.Values}
	default:
		return BadIndex(i, 3)
	}
}

func (x ArrayType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.Len = expr
	} else if i == 1 {
		x.p.Elt = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x AssignStmt) Set(i int, child Ast) {
	exprs := ToExprSlice(child)
	if i == 0 {
		x.p.Lhs = exprs
	} else if i == 1 {
		x.p.Rhs = exprs
	} else {
		BadIndex(i, 2)
	}
}
func (x BadDecl) Set(i int, child Ast)  { BadIndex(i, 0) }
func (x BadExpr) Set(i int, child Ast)  { BadIndex(i, 0) }
func (x BadStmt) Set(i int, child Ast)  { BadIndex(i, 0) }
func (x BasicLit) Set(i int, child Ast) { BadIndex(i, 0) }
func (x BinaryExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.X = expr
	} else if i == 1 {
		x.p.Y = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x BranchStmt) Set(i int, child Ast) { x.p.Label = ToIdent(child) }
func (x CallExpr) Set(i int, child Ast) {
	if i == 0 {
		x.p.Fun = ToExpr(child)
	} else if i == 1 {
		x.p.Args = ToExprSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x CaseClause) Set(i int, child Ast) {
	if i == 0 {
		x.p.List = ToExprSlice(child)
	} else if i == 1 {
		x.p.Body = ToStmtSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x ChanType) Set(i int, child Ast) { x.p.Value = ToExpr(child) }
func (x CommClause) Set(i int, child Ast) {
	if i == 0 {
		x.p.Comm = ToStmt(child)
	} else if i == 1 {
		x.p.Body = ToStmtSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x CompositeLit) Set(i int, child Ast) {
	if i == 0 {
		x.p.Type = ToExpr(child)
	} else if i == 1 {
		x.p.Elts = ToExprSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x DeclStmt) Set(i int, child Ast)  { x.p.Decl = ToDecl(child) }
func (x DeferStmt) Set(i int, child Ast) { x.p.Call = ToCallExpr(child) }
func (x Ellipsis) Set(i int, child Ast)  { x.p.Elt = ToExpr(child) }
func (x EmptyStmt) Set(i int, child Ast) { BadIndex(i, 0) }
func (x ExprStmt) Set(i int, child Ast)  { x.p.X = ToExpr(child) }
func (x Field) Set(i int, child Ast) {
	if i == 0 {
		x.p.Names = ToIdentSlice(child)
	} else if i == 1 {
		x.p.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x ForStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Init = ToStmt(child)
	case 1:
		x.p.Cond = ToExpr(child)
	case 2:
		x.p.Post = ToStmt(child)
	case 3:
		x.p.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x FuncDecl) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Recv = ToFieldList(child)
	case 1:
		x.p.Name = ToIdent(child)
	case 2:
		x.p.Type = ToFuncType(child)
	case 3:
		x.p.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x FuncLit) Set(i int, child Ast) {
	if i == 0 {
		x.p.Type = ToFuncType(child)
	} else if i == 1 {
		x.p.Body = ToBlockStmt(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x FuncType) Set(i int, child Ast) {
	list := ToFieldList(child)
	if i == 0 {
		x.p.Params = list
	} else if i == 1 {
		x.p.Results = list
	} else {
		BadIndex(i, 2)
	}
}
func (x GoStmt) Set(i int, child Ast) { x.p.Call = ToCallExpr(child) }
func (x Ident) Set(i int, child Ast)  { BadIndex(i, 0) }
func (x IfStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Init = ToStmt(child)
	case 1:
		x.p.Cond = ToExpr(child)
	case 2:
		x.p.Body = ToBlockStmt(child)
	case 3:
		x.p.Else = ToStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x ImportSpec) Set(i int, child Ast) {
	if i == 0 {
		x.p.Name = ToIdent(child)
	} else if i == 1 {
		x.p.Path = ToBasicLit(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x IncDecStmt) Set(i int, child Ast) { x.p.X = ToExpr(child) }
func (x IndexExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.X = expr
	} else if i == 1 {
		x.p.Index = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x InterfaceType) Set(i int, child Ast) { x.p.Methods = ToFieldList(child) }
func (x KeyValueExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.Key = expr
	} else if i == 1 {
		x.p.Value = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x LabeledStmt) Set(i int, child Ast) {
	if i == 0 {
		x.p.Label = ToIdent(child)
	} else if i == 1 {
		x.p.Stmt = ToStmt(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x MapType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.Key = expr
	} else if i == 1 {
		x.p.Value = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x Package) Set(i int, child Ast)   {} // TODO
func (x ParenExpr) Set(i int, child Ast) { x.p.X = ToExpr(child) }
func (x RangeStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Key = ToExpr(child)
	case 1:
		x.p.Value = ToExpr(child)
	case 2:
		x.p.X = ToExpr(child)
	case 3:
		x.p.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x SelectStmt) Set(i int, child Ast) { x.p.Body = ToBlockStmt(child) }
func (x SelectorExpr) Set(i int, child Ast) {
	if i == 0 {
		x.p.X = ToExpr(child)
	} else if i == 1 {
		x.p.Sel = ToIdent(child)
	}
}
func (x SendStmt) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.p.Chan = expr
	} else if i == 1 {
		x.p.Chan = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x SliceExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	switch i {
	case 0:
		x.p.X = expr
	case 1:
		x.p.Low = expr
	case 2:
		x.p.High = expr
	case 3:
		x.p.Max = expr
		x.p.Slice3 = expr != nil
	default:
		BadIndex(i, 4)
	}
}
func (x StarExpr) Set(i int, child Ast)   { x.p.X = ToExpr(child) }
func (x StructType) Set(i int, child Ast) { x.p.Fields = ToFieldList(child) }
func (x SwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Init = ToStmt(child)
	case 1:
		x.p.Tag = ToExpr(child)
	case 2:
		x.p.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 3)
	}
}
func (x TypeAssertExpr) Set(i int, child Ast) {
	if i == 0 {
		x.p.X = ToExpr(child)
	} else if i == 1 {
		x.p.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x TypeSpec) Set(i int, child Ast) {
	if i == 0 {
		x.p.Name = ToIdent(child)
	} else if i == 1 {
		x.p.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x TypeSwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Init = ToStmt(child)
	case 1:
		x.p.Assign = ToStmt(child)
	case 2:
		x.p.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 3)
	}
}
func (x UnaryExpr) Set(i int, child Ast) { x.p.X = ToExpr(child) }
func (x ValueSpec) Set(i int, child Ast) {
	switch i {
	case 0:
		x.p.Names = ToIdentSlice(child)
	case 1:
		x.p.Type = ToExpr(child)
	case 2:
		x.p.Values = ToExprSlice(child)
	default:
		BadIndex(i, 3)
	}
}
