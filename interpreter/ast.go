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
		Interface() interface{}
		Op() token.Token
		Size() int
		Get(i int) Ast
		Set(i int, child Ast)
	}
	AstWithNode interface {
		Ast
		Node() ast.Node
	}
	AstWithSlice interface {
		Ast
		Slice(lo, hi int)
		Append(child Ast)
	}

	NodeSlice  struct{ p []ast.Node }
	ExprSlice  struct{ p []ast.Expr }
	FieldSlice struct{ p []*ast.Field }
	DeclSlice  struct{ p []ast.Decl }
	IdentSlice struct{ p []*ast.Ident }
	StmtSlice  struct{ p []ast.Stmt }
	SpecSlice  struct{ p []ast.Spec }

	ArrayType      struct{ p *ast.ArrayType }
	AssignStmt     struct{ p *ast.AssignStmt }
	BadDecl        struct{ p *ast.BadDecl }
	BadExpr        struct{ p *ast.BadExpr }
	BadStmt        struct{ p *ast.BadStmt }
	BasicLit       struct{ p *ast.BasicLit }
	BinaryExpr     struct{ p *ast.BinaryExpr }
	BlockStmt      struct{ p *ast.BlockStmt }
	BranchStmt     struct{ p *ast.BranchStmt }
	CallExpr       struct{ p *ast.CallExpr }
	CaseClause     struct{ p *ast.CaseClause }
	ChanType       struct{ p *ast.ChanType }
	CommClause     struct{ p *ast.CommClause }
	CompositeLit   struct{ p *ast.CompositeLit }
	DeclStmt       struct{ p *ast.DeclStmt }
	DeferStmt      struct{ p *ast.DeferStmt }
	Ellipsis       struct{ p *ast.Ellipsis }
	EmptyStmt      struct{ p *ast.EmptyStmt }
	ExprStmt       struct{ p *ast.ExprStmt }
	Field          struct{ p *ast.Field }
	FieldList      struct{ p *ast.FieldList }
	File           struct{ p *ast.File }
	ForStmt        struct{ p *ast.ForStmt }
	FuncDecl       struct{ p *ast.FuncDecl }
	FuncLit        struct{ p *ast.FuncLit }
	FuncType       struct{ p *ast.FuncType }
	GenDecl        struct{ p *ast.GenDecl }
	GoStmt         struct{ p *ast.GoStmt }
	Ident          struct{ p *ast.Ident }
	IfStmt         struct{ p *ast.IfStmt }
	ImportSpec     struct{ p *ast.ImportSpec }
	IncDecStmt     struct{ p *ast.IncDecStmt }
	IndexExpr      struct{ p *ast.IndexExpr }
	InterfaceType  struct{ p *ast.InterfaceType }
	KeyValueExpr   struct{ p *ast.KeyValueExpr }
	LabeledStmt    struct{ p *ast.LabeledStmt }
	MapType        struct{ p *ast.MapType }
	Package        struct{ p *ast.Package }
	ParenExpr      struct{ p *ast.ParenExpr }
	RangeStmt      struct{ p *ast.RangeStmt }
	ReturnStmt     struct{ p *ast.ReturnStmt }
	SelectStmt     struct{ p *ast.SelectStmt }
	SelectorExpr   struct{ p *ast.SelectorExpr }
	SendStmt       struct{ p *ast.SendStmt }
	SliceExpr      struct{ p *ast.SliceExpr }
	StarExpr       struct{ p *ast.StarExpr }
	StructType     struct{ p *ast.StructType }
	SwitchStmt     struct{ p *ast.SwitchStmt }
	TypeAssertExpr struct{ p *ast.TypeAssertExpr }
	TypeSpec       struct{ p *ast.TypeSpec }
	TypeSwitchStmt struct{ p *ast.TypeSwitchStmt }
	UnaryExpr      struct{ p *ast.UnaryExpr }
	ValueSpec      struct{ p *ast.ValueSpec }
)
