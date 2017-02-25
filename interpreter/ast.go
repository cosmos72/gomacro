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
 * assignment.go
 *
 *  Created on Feb 24, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	_ "github.com/cosmos72/gomacro/token"
)

type (
	Ast interface {
		Op() token.Token
		N() int
		Get(i int) Ast
		// Set(child Ast, i int)
	}
	Node interface {
		Node() ast.Node
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
func (x BlockStmt) Node() ast.Node      { return x }
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
func (x FieldList) Node() ast.Node      { return x }
func (x File) Node() ast.Node           { return x }
func (x ForStmt) Node() ast.Node        { return x }
func (x FuncDecl) Node() ast.Node       { return x }
func (x FuncLit) Node() ast.Node        { return x }
func (x FuncType) Node() ast.Node       { return x }
func (x GenDecl) Node() ast.Node        { return x }
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
func (x ReturnStmt) Node() ast.Node     { return x }
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
func (x BlockStmt) Op() token.Token  { return token.LBRACE }
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

func (x ExprSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x FieldSlice) Op() token.Token { return token.SEMICOLON } // FIXME
func (x DeclSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x IdentSlice) Op() token.Token { return token.COMMA }     // FIXME
func (x SpecSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x StmtSlice) Op() token.Token  { return token.SEMICOLON } // FIXME

func (x CompositeLit) Op() token.Token   { return token.RBRACE }
func (x DeclStmt) Op() token.Token       { return x.Decl.(*GenDecl).Tok }
func (x DeferStmt) Op() token.Token      { return token.DEFER }
func (x Ellipsis) Op() token.Token       { return token.ELLIPSIS }
func (x EmptyStmt) Op() token.Token      { return token.SEMICOLON }
func (x ExprStmt) Op() token.Token       { return token.ELSE } // FIXME
func (x Field) Op() token.Token          { return token.PERIOD }
func (x FieldList) Op() token.Token      { return token.ELLIPSIS }
func (x File) Op() token.Token           { return token.EOF }
func (x ForStmt) Op() token.Token        { return token.FOR }
func (x FuncDecl) Op() token.Token       { return token.FUNC }
func (x FuncLit) Op() token.Token        { return token.FUNC }
func (x FuncType) Op() token.Token       { return token.FUNC }
func (x GenDecl) Op() token.Token        { return x.Tok }
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
func (x ReturnStmt) Op() token.Token     { return token.RETURN }
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

func (x ExprSlice) N() int  { return len(*x.P) }
func (x FieldSlice) N() int { return len(*x.P) }
func (x DeclSlice) N() int  { return len(*x.P) }
func (x IdentSlice) N() int { return len(*x.P) }
func (x SpecSlice) N() int  { return len(*x.P) }
func (x StmtSlice) N() int  { return len(*x.P) }

func (x ArrayType) N() int      { return 2 }
func (x AssignStmt) N() int     { return 2 }
func (x BadDecl) N() int        { return 0 }
func (x BadExpr) N() int        { return 0 }
func (x BadStmt) N() int        { return 0 }
func (x BasicLit) N() int       { return 1 }
func (x BinaryExpr) N() int     { return 2 }
func (x BlockStmt) N() int      { return len(x.List) }
func (x BranchStmt) N() int     { return 1 }
func (x CallExpr) N() int       { return 2 }
func (x CaseClause) N() int     { return 2 }
func (x ChanType) N() int       { return 1 }
func (x CommClause) N() int     { return 2 }
func (x CompositeLit) N() int   { return 2 }
func (x DeclStmt) N() int       { return 1 }
func (x DeferStmt) N() int      { return 1 }
func (x Ellipsis) N() int       { return 2 }
func (x EmptyStmt) N() int      { return 0 }
func (x ExprStmt) N() int       { return 1 }
func (x Field) N() int          { return 2 }
func (x FieldList) N() int      { return len(x.List) }
func (x File) N() int           { return len(x.Decls) }
func (x ForStmt) N() int        { return 4 }
func (x FuncDecl) N() int       { return 4 }
func (x FuncLit) N() int        { return 2 }
func (x FuncType) N() int       { return 2 }
func (x GenDecl) N() int        { return len(x.Specs) }
func (x GoStmt) N() int         { return 1 }
func (x Ident) N() int          { return 0 }
func (x IfStmt) N() int         { return 4 }
func (x ImportSpec) N() int     { return 2 }
func (x IncDecStmt) N() int     { return 1 }
func (x IndexExpr) N() int      { return 2 }
func (x InterfaceType) N() int  { return 1 }
func (x KeyValueExpr) N() int   { return 2 }
func (x LabeledStmt) N() int    { return 2 }
func (x MapType) N() int        { return 2 }
func (x Package) N() int        { return 2 }
func (x ParenExpr) N() int      { return 1 }
func (x RangeStmt) N() int      { return 4 }
func (x ReturnStmt) N() int     { return len(x.Results) }
func (x SelectStmt) N() int     { return 1 }
func (x SelectorExpr) N() int   { return 2 }
func (x SendStmt) N() int       { return 2 }
func (x SliceExpr) N() int      { return 4 }
func (x StarExpr) N() int       { return 1 }
func (x StructType) N() int     { return 1 }
func (x SwitchStmt) N() int     { return 3 }
func (x TypeAssertExpr) N() int { return 2 }
func (x TypeSpec) N() int       { return 2 }
func (x TypeSwitchStmt) N() int { return 3 }
func (x UnaryExpr) N() int      { return 1 }
func (x ValueSpec) N() int      { return 3 }

func (x ExprSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x FieldSlice) Get(i int) Ast { return ToAst((*x.P)[i]) }
func (x DeclSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x IdentSlice) Get(i int) Ast { return ToAst((*x.P)[i]) }
func (x SpecSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x StmtSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }

func (x ArrayType) Get(i int) Ast { return ToAst2(i, x.Len, x.Elt) }
func (x AssignStmt) Get(i int) Ast {
	var addr *[]ast.Expr
	if i == 0 {
		addr = &x.Lhs
	} else {
		addr = &x.Rhs
	}
	return ExprSlice{addr}
}
func (x BadDecl) Get(i int) Ast    { return nil }
func (x BadExpr) Get(i int) Ast    { return nil }
func (x BadStmt) Get(i int) Ast    { return nil }
func (x BasicLit) Get(i int) Ast   { return nil }
func (x BinaryExpr) Get(i int) Ast { return ToAst2(i, x.X, x.Y) }
func (x BlockStmt) Get(i int) Ast  { return ToAst(x.List[i]) }
func (x BranchStmt) Get(i int) Ast { return Ident{x.Label} }
func (x CallExpr) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Fun)
	} else {
		return ExprSlice{&x.Args}
	}
}
func (x CaseClause) Get(i int) Ast {
	if i == 0 {
		return ExprSlice{&x.List}
	} else {
		return StmtSlice{&x.Body}
	}
}
func (x ChanType) Get(i int) Ast { return ToAst(x.Value) }
func (x CommClause) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Comm)
	} else {
		return StmtSlice{&x.Body}
	}
}
func (x CompositeLit) Get(i int) Ast {
	if i == 0 {
		return ToAst(x.Type)
	} else {
		return ExprSlice{&x.Elts}
	}
}
func (x DeclStmt) Get(i int) Ast  { return ToAst(x.Decl) }
func (x DeferStmt) Get(i int) Ast { return CallExpr{x.Call} }
func (x Ellipsis) Get(i int) Ast  { return ToAst(x.Elt) }
func (x EmptyStmt) Get(i int) Ast { return nil }
func (x ExprStmt) Get(i int) Ast  { return ToAst(x.X) }
func (x Field) Get(i int) Ast {
	if i == 0 {
		return IdentSlice{&x.Names}
	} else {
		return ToAst(x.Type)
	}
}
func (x FieldList) Get(i int) Ast { return ToAst(x.List[i]) }
func (x File) Get(i int) Ast      { return ToAst(x.Decls[i]) }
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
	}
	return ToAst(node)
}
func (x FuncLit) Get(i int) Ast {
	if i == 0 {
		return FuncType{x.Type}
	} else {
		return BlockStmt{x.Body}
	}
}
func (x FuncType) Get(i int) Ast {
	var ret *ast.FieldList
	if i == 0 {
		ret = x.Params
	} else {
		ret = x.Results
	}
	return FieldList{ret}
}
func (x GenDecl) Get(i int) Ast { return ToAst(x.Specs[i]) }
func (x GoStmt) Get(i int) Ast  { return CallExpr{x.Call} }
func (x Ident) Get(i int) Ast   { return nil }
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
	}
	return ToAst(node)
}
func (x ImportSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Name}
	} else {
		return BasicLit{x.Path}
	}
}
func (x IncDecStmt) Get(i int) Ast    { return ToAst(x.X) }
func (x IndexExpr) Get(i int) Ast     { return ToAst2(i, x.X, x.Index) }
func (x InterfaceType) Get(i int) Ast { return FieldList{x.Methods} }
func (x KeyValueExpr) Get(i int) Ast  { return ToAst2(i, x.Key, x.Value) }
func (x LabeledStmt) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Label}
	} else {
		return ToAst(x.Stmt)
	}
}
func (x MapType) Get(i int) Ast   { return ToAst2(i, x.Key, x.Value) }
func (x Package) Get(i int) Ast   { return nil } // TODO
func (x ParenExpr) Get(i int) Ast { return ToAst(x.X) }
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
	}
	return ToAst(node)
}
func (x ReturnStmt) Get(i int) Ast   { return ToAst(x.Results[i]) }
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
	}
	return ToAst(node)
}
func (x StarExpr) Get(i int) Ast       { return ToAst(x.X) }
func (x StructType) Get(i int) Ast     { return FieldList{x.Fields} }
func (x SwitchStmt) Get(i int) Ast     { return ToAst3(i, x.Init, x.Tag, x.Body) }
func (x TypeAssertExpr) Get(i int) Ast { return ToAst2(i, x.X, x.Type) }
func (x TypeSpec) Get(i int) Ast {
	if i == 0 {
		return Ident{x.Name}
	} else {
		return ToAst(x.Type)
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
	}
	return ToAst(node)
}
func (x UnaryExpr) Get(i int) Ast { return ToAst(x.X) }
func (x ValueSpec) Get(i int) Ast {
	switch i {
	case 0:
		return IdentSlice{&x.Names}
	case 1:
		return ToAst(x.Type)
	case 2:
		return ExprSlice{&x.Values}
	default:
		return nil
	}
}

// ToAst2 returns n0 (if i == 0) or n1 converted to Ast
func ToAst2(i int, n0 ast.Node, n1 ast.Node) Ast {
	n := n0
	if i != 0 {
		n = n1
	}
	return ToAst(n)
}

func ToAst3(i int, n0 ast.Node, n1 ast.Node, n2 *ast.BlockStmt) Ast {
	switch i {
	case 0:
		/*nop*/
	case 1:
		n1 = n0
	case 2:
		return BlockStmt{n2}
	}
	return ToAst(n1)
}

func ToAst(node ast.Node) Ast {
	var x Ast
	switch node := node.(type) {
	case *ast.ArrayType:
		x = ArrayType{node}
	case *ast.AssignStmt:
		x = AssignStmt{node}
	case *ast.BadDecl:
		x = BadDecl{node}
	case *ast.BadExpr:
		x = BadExpr{node}
	case *ast.BadStmt:
		x = BadStmt{node}
	case *ast.BasicLit:
		x = BasicLit{node}
	case *ast.BinaryExpr:
		x = BinaryExpr{node}
	case *ast.BlockStmt:
		x = BlockStmt{node}
	case *ast.BranchStmt:
		x = BranchStmt{node}
	case *ast.CallExpr:
		x = CallExpr{node}
	case *ast.CaseClause:
		x = CaseClause{node}
	case *ast.ChanType:
		x = ChanType{node}
	case *ast.CommClause:
		x = CommClause{node}
	case *ast.CompositeLit:
		x = CompositeLit{node}
	case *ast.DeclStmt:
		x = DeclStmt{node}
	case *ast.DeferStmt:
		x = DeferStmt{node}
	case *ast.Ellipsis:
		x = Ellipsis{node}
	case *ast.EmptyStmt:
		x = EmptyStmt{node}
	case *ast.ExprStmt:
		x = ExprStmt{node}
	case *ast.Field:
		x = Field{node}
	case *ast.FieldList:
		x = FieldList{node}
	case *ast.File:
		x = File{node}
	case *ast.ForStmt:
		x = ForStmt{node}
	case *ast.FuncDecl:
		x = FuncDecl{node}
	case *ast.FuncLit:
		x = FuncLit{node}
	case *ast.FuncType:
		x = FuncType{node}
	case *ast.GenDecl:
		x = GenDecl{node}
	case *ast.GoStmt:
		x = GoStmt{node}
	case *ast.Ident:
		x = Ident{node}
	case *ast.IfStmt:
		x = IfStmt{node}
	case *ast.ImportSpec:
		x = ImportSpec{node}
	case *ast.IncDecStmt:
		x = IncDecStmt{node}
	case *ast.IndexExpr:
		x = IndexExpr{node}
	case *ast.InterfaceType:
		x = InterfaceType{node}
	case *ast.KeyValueExpr:
		x = KeyValueExpr{node}
	case *ast.LabeledStmt:
		x = LabeledStmt{node}
	case *ast.MapType:
		x = MapType{node}
	case *ast.Package:
		x = Package{node}
	case *ast.ParenExpr:
		x = ParenExpr{node}
	case *ast.RangeStmt:
		x = RangeStmt{node}
	case *ast.ReturnStmt:
		x = ReturnStmt{node}
	case *ast.SelectStmt:
		x = SelectStmt{node}
	case *ast.SelectorExpr:
		x = SelectorExpr{node}
	case *ast.SendStmt:
		x = SendStmt{node}
	case *ast.SliceExpr:
		x = SliceExpr{node}
	case *ast.StarExpr:
		x = StarExpr{node}
	case *ast.StructType:
		x = StructType{node}
	case *ast.SwitchStmt:
		x = SwitchStmt{node}
	case *ast.TypeAssertExpr:
		x = TypeAssertExpr{node}
	case *ast.TypeSpec:
		x = TypeSpec{node}
	case *ast.TypeSwitchStmt:
		x = TypeSwitchStmt{node}
	case *ast.UnaryExpr:
		x = UnaryExpr{node}
	case *ast.ValueSpec:
		x = ValueSpec{node}
	case nil:
		break
	default:
		panic(errors.New(fmt.Sprintf("unsupported node type <%v>", reflect.TypeOf(node))))
	}
	return x
}
