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
 * ast.go
 *
 *  Created on Feb 24, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
)

type (
	Ast interface {
		Op() token.Token
		Size() int
		// Get(i int) Ast
		// Set(child Ast, i int)
	}
	AstWithNode interface {
		Ast
		Node() ast.Node
	}
	AstWithResize interface {
		Ast
		Resize(n int)
	}

	ExprSlice  struct{ P *[]ast.Expr }
	FieldSlice struct{ P *[]*ast.Field }
	DeclSlice  struct{ P *[]ast.Decl }
	IdentSlice struct{ P *[]*ast.Ident }
	StmtSlice  struct{ P *[]ast.Stmt }
	SpecSlice  struct{ P *[]ast.Spec }

	ArrayType      struct{ *ast.ArrayType }
	AssignStmt     struct{ *ast.AssignStmt }
	BadDecl        struct{ *ast.BadDecl }
	BadExpr        struct{ *ast.BadExpr }
	BadStmt        struct{ *ast.BadStmt }
	BasicLit       struct{ *ast.BasicLit }
	BinaryExpr     struct{ *ast.BinaryExpr }
	BlockStmt      struct{ *ast.BlockStmt }
	BranchStmt     struct{ *ast.BranchStmt }
	CallExpr       struct{ *ast.CallExpr }
	CaseClause     struct{ *ast.CaseClause }
	ChanType       struct{ *ast.ChanType }
	CommClause     struct{ *ast.CommClause }
	CompositeLit   struct{ *ast.CompositeLit }
	DeclStmt       struct{ *ast.DeclStmt }
	DeferStmt      struct{ *ast.DeferStmt }
	Ellipsis       struct{ *ast.Ellipsis }
	EmptyStmt      struct{ *ast.EmptyStmt }
	ExprStmt       struct{ *ast.ExprStmt }
	Field          struct{ *ast.Field }
	FieldList      struct{ *ast.FieldList }
	File           struct{ *ast.File }
	ForStmt        struct{ *ast.ForStmt }
	FuncDecl       struct{ *ast.FuncDecl }
	FuncLit        struct{ *ast.FuncLit }
	FuncType       struct{ *ast.FuncType }
	GenDecl        struct{ *ast.GenDecl }
	GoStmt         struct{ *ast.GoStmt }
	Ident          struct{ *ast.Ident }
	IfStmt         struct{ *ast.IfStmt }
	ImportSpec     struct{ *ast.ImportSpec }
	IncDecStmt     struct{ *ast.IncDecStmt }
	IndexExpr      struct{ *ast.IndexExpr }
	InterfaceType  struct{ *ast.InterfaceType }
	KeyValueExpr   struct{ *ast.KeyValueExpr }
	LabeledStmt    struct{ *ast.LabeledStmt }
	MapType        struct{ *ast.MapType }
	Package        struct{ *ast.Package }
	ParenExpr      struct{ *ast.ParenExpr }
	RangeStmt      struct{ *ast.RangeStmt }
	ReturnStmt     struct{ *ast.ReturnStmt }
	SelectStmt     struct{ *ast.SelectStmt }
	SelectorExpr   struct{ *ast.SelectorExpr }
	SendStmt       struct{ *ast.SendStmt }
	SliceExpr      struct{ *ast.SliceExpr }
	StarExpr       struct{ *ast.StarExpr }
	StructType     struct{ *ast.StructType }
	SwitchStmt     struct{ *ast.SwitchStmt }
	TypeAssertExpr struct{ *ast.TypeAssertExpr }
	TypeSpec       struct{ *ast.TypeSpec }
	TypeSwitchStmt struct{ *ast.TypeSwitchStmt }
	UnaryExpr      struct{ *ast.UnaryExpr }
	ValueSpec      struct{ *ast.ValueSpec }
)

func (x ArrayType) Node() ast.Node      { return x }
func (x AssignStmt) Node() ast.Node     { return x }
func (x BadDecl) Node() ast.Node        { return x }
func (x BadExpr) Node() ast.Node        { return x }
func (x BadStmt) Node() ast.Node        { return x }
func (x BasicLit) Node() ast.Node       { return x }
func (x BinaryExpr) Node() ast.Node     { return x }
func (x BranchStmt) Node() ast.Node     { return x }
func (x CallExpr) Node() ast.Node       { return x }
func (x CaseClause) Node() ast.Node     { return x }
func (x ChanType) Node() ast.Node       { return x }
func (x CommClause) Node() ast.Node     { return x }
func (x CompositeLit) Node() ast.Node   { return x }
func (x DeclStmt) Node() ast.Node       { return x }
func (x DeferStmt) Node() ast.Node      { return x }
func (x Ellipsis) Node() ast.Node       { return x }
func (x EmptyStmt) Node() ast.Node      { return x }
func (x ExprStmt) Node() ast.Node       { return x }
func (x Field) Node() ast.Node          { return x }
func (x ForStmt) Node() ast.Node        { return x }
func (x FuncDecl) Node() ast.Node       { return x }
func (x FuncLit) Node() ast.Node        { return x }
func (x FuncType) Node() ast.Node       { return x }
func (x GoStmt) Node() ast.Node         { return x }
func (x Ident) Node() ast.Node          { return x }
func (x IfStmt) Node() ast.Node         { return x }
func (x ImportSpec) Node() ast.Node     { return x }
func (x IncDecStmt) Node() ast.Node     { return x }
func (x IndexExpr) Node() ast.Node      { return x }
func (x InterfaceType) Node() ast.Node  { return x }
func (x KeyValueExpr) Node() ast.Node   { return x }
func (x LabeledStmt) Node() ast.Node    { return x }
func (x MapType) Node() ast.Node        { return x }
func (x Package) Node() ast.Node        { return x }
func (x ParenExpr) Node() ast.Node      { return x }
func (x RangeStmt) Node() ast.Node      { return x }
func (x SelectStmt) Node() ast.Node     { return x }
func (x SelectorExpr) Node() ast.Node   { return x }
func (x SendStmt) Node() ast.Node       { return x }
func (x SliceExpr) Node() ast.Node      { return x }
func (x StarExpr) Node() ast.Node       { return x }
func (x StructType) Node() ast.Node     { return x }
func (x SwitchStmt) Node() ast.Node     { return x }
func (x TypeAssertExpr) Node() ast.Node { return x }
func (x TypeSpec) Node() ast.Node       { return x }
func (x TypeSwitchStmt) Node() ast.Node { return x }
func (x UnaryExpr) Node() ast.Node      { return x }
func (x ValueSpec) Node() ast.Node      { return x }

func (x ArrayType) Op() token.Token  { return token.LBRACK }
func (x AssignStmt) Op() token.Token { return x.Op() }
func (x BadDecl) Op() token.Token    { return token.ILLEGAL }
func (x BadExpr) Op() token.Token    { return token.ILLEGAL }
func (x BadStmt) Op() token.Token    { return token.ILLEGAL }
func (x BasicLit) Op() token.Token   { return x.Kind }
func (x BinaryExpr) Op() token.Token { return x.BinaryExpr.Op }
func (x BranchStmt) Op() token.Token { return x.Tok }
func (x CallExpr) Op() token.Token   { return token.RPAREN }
func (x CaseClause) Op() token.Token {
	if len(x.List) != 0 {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}
func (x ChanType) Op() token.Token { return token.CHAN }
func (x CommClause) Op() token.Token {
	if x.Comm != nil {
		return token.CASE
	} else {
		return token.DEFAULT
	}
}

func (x CompositeLit) Op() token.Token   { return token.RBRACE }
func (x DeclStmt) Op() token.Token       { return x.Decl.(*GenDecl).Tok }
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
func (x IncDecStmt) Op() token.Token     { return x.Tok }
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
func (x UnaryExpr) Op() token.Token      { return x.UnaryExpr.Op }
func (x ValueSpec) Op() token.Token      { return token.VAR } // can be VAR or CONST

func (x ArrayType) Size() int      { return 2 }
func (x AssignStmt) Size() int     { return 2 }
func (x BadDecl) Size() int        { return 0 }
func (x BadExpr) Size() int        { return 0 }
func (x BadStmt) Size() int        { return 0 }
func (x BasicLit) Size() int       { return 1 }
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

func (x ArrayType) Get(i int) Ast { return ToAst2(i, x.Len, x.Elt) }
func (x AssignStmt) Get(i int) Ast {
	var addr *[]ast.Expr
	if i == 0 {
		addr = &x.Lhs
	} else if i == 1 {
		addr = &x.Rhs
	} else {
		return BadIndex(i, 2)
	}
	return ExprSlice{addr}
}
func (x BadDecl) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BadExpr) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BadStmt) Get(i int) Ast    { return BadIndex(i, 0) }
func (x BasicLit) Get(i int) Ast   { return BadIndex(i, 0) }
func (x BinaryExpr) Get(i int) Ast { return ToAst2(i, x.X, x.Y) }
func (x BranchStmt) Get(i int) Ast { return Ident{x.Label} }
func (x CallExpr) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Fun)
	} else if i == 1 {
		return ExprSlice{&x.Args}
	} else {
		return BadIndex(i, 2)
	}
}
func (x CaseClause) Get(i int) Ast {
	if i == 0 {
		return ExprSlice{&x.List}
	} else if i == 1 {
		return StmtSlice{&x.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x ChanType) Get(i int) Ast { return ToAst1(i, x.Value) }
func (x CommClause) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Comm)
	} else if i == 1 {
		return StmtSlice{&x.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x CompositeLit) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Type)
	} else if i == 1 {
		return ExprSlice{&x.Elts}
	} else {
		return BadIndex(i, 2)
	}
}
func (x DeclStmt) Get(i int) Ast  { return ToAst1(i, x.Decl) }
func (x DeferStmt) Get(i int) Ast { return CallExpr{x.Call} }
func (x Ellipsis) Get(i int) Ast  { return ToAst1(i, x.Elt) }
func (x EmptyStmt) Get(i int) Ast { return BadIndex(i, 0) }
func (x ExprStmt) Get(i int) Ast  { return ToAst1(i, x.X) }
func (x Field) Get(i int) Ast {
	if i == 0 {
		return IdentSlice{&x.Names}
	} else if i == 1 {
		return ToAst(x.Type)
	} else {
		return BadIndex(i, 2)
	}
}
func (x ForStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.Init
	case 1:
		node = x.Cond
	case 2:
		node = x.Post
	case 3:
		return BlockStmt{x.Body}
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x FuncDecl) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.Recv
	case 1:
		return Ident{x.Name}
	case 2:
		node = x.Type
	case 3:
		return BlockStmt{x.Body}
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x FuncLit) Get(i int) Ast {
	if i == 0 {
		return FuncType{x.Type}
	} else if i == 1 {
		return BlockStmt{x.Body}
	} else {
		return BadIndex(i, 2)
	}
}
func (x FuncType) Get(i int) Ast {
	var ret *ast.FieldList
	if i == 0 {
		ret = x.Params
	} else if i == 1 {
		ret = x.Results
	} else {
		return BadIndex(i, 2)
	}
	return FieldList{ret}
}
func (x GoStmt) Get(i int) Ast { return CallExpr{x.Call} }
func (x Ident) Get(i int) Ast  { return BadIndex(i, 0) }
func (x IfStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.Init
	case 1:
		node = x.Cond
	case 2:
		return BlockStmt{x.Body}
	case 3:
		node = x.Else
	default:
		return BadIndex(i, 4)
	}
	return ToAst(node)
}

func (x ImportSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Name}
	} else if i == 1 {
		return BasicLit{x.Path}
	} else {
		return BadIndex(i, 2)
	}
}
func (x IncDecStmt) Get(i int) Ast    { return ToAst1(i, x.X) }
func (x IndexExpr) Get(i int) Ast     { return ToAst2(i, x.X, x.Index) }
func (x InterfaceType) Get(i int) Ast { return FieldList{x.Methods} }
func (x KeyValueExpr) Get(i int) Ast  { return ToAst2(i, x.Key, x.Value) }
func (x LabeledStmt) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Label}
	} else if i == 1 {
		return ToAst(x.Stmt)
	} else {
		return BadIndex(i, 2)
	}
}
func (x MapType) Get(i int) Ast   { return ToAst2(i, x.Key, x.Value) }
func (x Package) Get(i int) Ast   { return nil } // TODO
func (x ParenExpr) Get(i int) Ast { return ToAst1(i, x.X) }
func (x RangeStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.Key
	case 1:
		node = x.Value
	case 2:
		node = x.X
	case 3:
		return BlockStmt{x.Body}
	default:
		BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x SelectStmt) Get(i int) Ast   { return BlockStmt{x.Body} }
func (x SelectorExpr) Get(i int) Ast { return ToAst2(i, x.X, x.Sel) }
func (x SendStmt) Get(i int) Ast     { return ToAst2(i, x.Chan, x.Value) }
func (x SliceExpr) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.X
	case 1:
		node = x.Low
	case 2:
		node = x.High
	case 3:
		node = x.Max
	default:
		BadIndex(i, 4)
	}
	return ToAst(node)
}
func (x StarExpr) Get(i int) Ast       { return ToAst1(i, x.X) }
func (x StructType) Get(i int) Ast     { return FieldList{x.Fields} }
func (x SwitchStmt) Get(i int) Ast     { return ToAst3(i, x.Init, x.Tag, x.Body) }
func (x TypeAssertExpr) Get(i int) Ast { return ToAst2(i, x.X, x.Type) }
func (x TypeSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Name}
	} else if i == 1 {
		return ToAst(x.Type)
	} else {
		return BadIndex(i, 2)
	}
}
func (x TypeSwitchStmt) Get(i int) Ast {
	var node ast.Node
	switch i {
	case 0:
		node = x.Init
	case 1:
		node = x.Assign
	case 2:
		return BlockStmt{x.Body}
	default:
		BadIndex(i, 3)
	}
	return ToAst(node)
}
func (x UnaryExpr) Get(i int) Ast { return ToAst1(i, x.X) }
func (x ValueSpec) Get(i int) Ast {
	switch i {
	case 0:
		return IdentSlice{&x.Names}
	case 1:
		return ToAst(x.Type)
	case 2:
		return ExprSlice{&x.Values}
	default:
		return BadIndex(i, 3)
	}
}

func (x ArrayType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.Len = expr
	} else if i == 1 {
		x.Elt = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x AssignStmt) Set(i int, child Ast) {
	exprs := ToExprSlice(child)
	if i == 0 {
		x.Lhs = exprs
	} else if i == 1 {
		x.Rhs = exprs
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
		x.X = expr
	} else if i == 1 {
		x.Y = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x BranchStmt) Set(i int, child Ast) { x.Label = ToIdent(child) }
func (x CallExpr) Set(i int, child Ast) {
	if i == 0 {
		x.Fun = ToExpr(child)
	} else if i == 1 {
		x.Args = ToExprSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x CaseClause) Set(i int, child Ast) {
	if i == 0 {
		x.List = ToExprSlice(child)
	} else if i == 1 {
		x.Body = ToStmtSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x ChanType) Set(i int, child Ast) { x.Value = ToExpr(child) }
func (x CommClause) Set(i int, child Ast) {
	if i == 0 {
		x.Comm = ToStmt(child)
	} else if i == 1 {
		x.Body = ToStmtSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x CompositeLit) Set(i int, child Ast) {
	if i == 0 {
		x.Type = ToExpr(child)
	} else if i == 1 {
		x.Elts = ToExprSlice(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x DeclStmt) Set(i int, child Ast)  { x.Decl = ToDecl(child) }
func (x DeferStmt) Set(i int, child Ast) { x.Call = ToCallExpr(child) }
func (x Ellipsis) Set(i int, child Ast)  { x.Elt = ToExpr(child) }
func (x EmptyStmt) Set(i int, child Ast) { BadIndex(i, 0) }
func (x ExprStmt) Set(i int, child Ast)  { x.X = ToExpr(child) }
func (x Field) Set(i int, child Ast) {
	if i == 0 {
		x.Names = ToIdentSlice(child)
	} else if i == 1 {
		x.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x ForStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Init = ToStmt(child)
	case 1:
		x.Cond = ToExpr(child)
	case 2:
		x.Post = ToStmt(child)
	case 3:
		x.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x FuncDecl) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Recv = ToFieldList(child)
	case 1:
		x.Name = ToIdent(child)
	case 2:
		x.Type = ToFuncType(child)
	case 3:
		x.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x FuncLit) Set(i int, child Ast) {
	if i == 0 {
		x.Type = ToFuncType(child)
	} else if i == 1 {
		x.Body = ToBlockStmt(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x FuncType) Set(i int, child Ast) {
	list := ToFieldList(child)
	if i == 0 {
		x.Params = list
	} else if i == 1 {
		x.Results = list
	} else {
		BadIndex(i, 2)
	}
}
func (x GoStmt) Set(i int, child Ast) { x.Call = ToCallExpr(child) }
func (x Ident) Set(i int, child Ast)  { BadIndex(i, 0) }
func (x IfStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Init = ToStmt(child)
	case 1:
		x.Cond = ToExpr(child)
	case 2:
		x.Body = ToBlockStmt(child)
	case 3:
		x.Else = ToStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x ImportSpec) Set(i int, child Ast) {
	if i == 0 {
		x.Name = ToIdent(child)
	} else if i == 1 {
		x.Path = ToBasicLit(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x IncDecStmt) Set(i int, child Ast) { x.X = ToExpr(child) }
func (x IndexExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.X = expr
	} else if i == 1 {
		x.Index = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x InterfaceType) Set(i int, child Ast) { x.Methods = ToFieldList(child) }
func (x KeyValueExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.Key = expr
	} else if i == 1 {
		x.Value = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x LabeledStmt) Set(i int, child Ast) {
	if i == 0 {
		x.Label = ToIdent(child)
	} else if i == 1 {
		x.Stmt = ToStmt(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x MapType) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.Key = expr
	} else if i == 1 {
		x.Value = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x Package) Set(i int, child Ast)   {} // TODO
func (x ParenExpr) Set(i int, child Ast) { x.X = ToExpr(child) }
func (x RangeStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Key = ToExpr(child)
	case 1:
		x.Value = ToExpr(child)
	case 2:
		x.X = ToExpr(child)
	case 3:
		x.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 4)
	}
}
func (x SelectStmt) Set(i int, child Ast) { x.Body = ToBlockStmt(child) }
func (x SelectorExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X = ToExpr(child)
	} else if i == 1 {
		x.Sel = ToIdent(child)
	}
}
func (x SendStmt) Set(i int, child Ast) {
	expr := ToExpr(child)
	if i == 0 {
		x.Chan = expr
	} else if i == 1 {
		x.Chan = expr
	} else {
		BadIndex(i, 2)
	}
}
func (x SliceExpr) Set(i int, child Ast) {
	expr := ToExpr(child)
	switch i {
	case 0:
		x.X = expr
	case 1:
		x.Low = expr
	case 2:
		x.High = expr
	case 3:
		x.Max = expr
		x.Slice3 = expr != nil
	default:
		BadIndex(i, 4)
	}
}
func (x StarExpr) Set(i int, child Ast)   { x.X = ToExpr(child) }
func (x StructType) Set(i int, child Ast) { x.Fields = ToFieldList(child) }
func (x SwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Init = ToStmt(child)
	case 1:
		x.Tag = ToExpr(child)
	case 2:
		x.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 3)
	}
}
func (x TypeAssertExpr) Set(i int, child Ast) {
	if i == 0 {
		x.X = ToExpr(child)
	} else if i == 1 {
		x.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x TypeSpec) Set(i int, child Ast) {
	if i == 0 {
		x.Name = ToIdent(child)
	} else if i == 1 {
		x.Type = ToExpr(child)
	} else {
		BadIndex(i, 2)
	}
}
func (x TypeSwitchStmt) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Init = ToStmt(child)
	case 1:
		x.Assign = ToStmt(child)
	case 2:
		x.Body = ToBlockStmt(child)
	default:
		BadIndex(i, 3)
	}
}
func (x UnaryExpr) Set(i int, child Ast) { x.X = ToExpr(child) }
func (x ValueSpec) Set(i int, child Ast) {
	switch i {
	case 0:
		x.Names = ToIdentSlice(child)
	case 1:
		x.Type = ToExpr(child)
	case 2:
		x.Values = ToExprSlice(child)
	default:
		BadIndex(i, 3)
	}
}
