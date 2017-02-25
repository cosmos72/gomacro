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
