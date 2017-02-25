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
 * ast_slice.go
 *
 *  Created on Feb 25, 2017
 *      Author Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
)

// Ast wrappers for variable-length fragments of ast.Nodes - they are not full-blown ast.Nodes

func (x ExprSlice) Op() token.Token  { return token.COMMA }     // FIXME
func (x FieldSlice) Op() token.Token { return token.SEMICOLON } // FIXME
func (x DeclSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x IdentSlice) Op() token.Token { return token.COMMA }     // FIXME
func (x SpecSlice) Op() token.Token  { return token.SEMICOLON } // FIXME
func (x StmtSlice) Op() token.Token  { return token.SEMICOLON } // FIXME

func (x ExprSlice) Size() int  { return len(*x.P) }
func (x FieldSlice) Size() int { return len(*x.P) }
func (x DeclSlice) Size() int  { return len(*x.P) }
func (x IdentSlice) Size() int { return len(*x.P) }
func (x SpecSlice) Size() int  { return len(*x.P) }
func (x StmtSlice) Size() int  { return len(*x.P) }

func (x ExprSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x FieldSlice) Get(i int) Ast { return ToAst((*x.P)[i]) }
func (x DeclSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x IdentSlice) Get(i int) Ast { return ToAst((*x.P)[i]) }
func (x SpecSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }
func (x StmtSlice) Get(i int) Ast  { return ToAst((*x.P)[i]) }

func (x ExprSlice) Set(i int, child Ast)  { (*x.P)[i] = ToExpr(child) }
func (x FieldSlice) Set(i int, child Ast) { (*x.P)[i] = ToField(child) }
func (x DeclSlice) Set(i int, child Ast)  { (*x.P)[i] = ToDecl(child) }
func (x IdentSlice) Set(i int, child Ast) { (*x.P)[i] = ToIdent(child) }
func (x SpecSlice) Set(i int, child Ast)  { (*x.P)[i] = ToSpec(child) }
func (x StmtSlice) Set(i int, child Ast)  { (*x.P)[i] = ToStmt(child) }

func (x ExprSlice) Resize(n int)  { s := make([]ast.Expr, n); copy(s, *x.P); *x.P = s }
func (x FieldSlice) Resize(n int) { s := make([]*ast.Field, n); copy(s, *x.P); *x.P = s }
func (x DeclSlice) Resize(n int)  { s := make([]ast.Decl, n); copy(s, *x.P); *x.P = s }
func (x IdentSlice) Resize(n int) { s := make([]*ast.Ident, n); copy(s, *x.P); *x.P = s }
func (x SpecSlice) Resize(n int)  { s := make([]ast.Spec, n); copy(s, *x.P); *x.P = s }
func (x StmtSlice) Resize(n int)  { s := make([]ast.Stmt, n); copy(s, *x.P); *x.P = s }

// variable-length ast.Nodes

func (x BlockStmt) Node() ast.Node  { return x }
func (x FieldList) Node() ast.Node  { return x }
func (x File) Node() ast.Node       { return x }
func (x GenDecl) Node() ast.Node    { return x }
func (x ReturnStmt) Node() ast.Node { return x }

func (x BlockStmt) Op() token.Token  { return token.LBRACE }
func (x FieldList) Op() token.Token  { return token.ELLIPSIS }
func (x File) Op() token.Token       { return token.EOF }
func (x GenDecl) Op() token.Token    { return x.Tok }
func (x ReturnStmt) Op() token.Token { return token.RETURN }

func (x BlockStmt) Size() int  { return len(x.List) }
func (x FieldList) Size() int  { return len(x.List) }
func (x File) Size() int       { return len(x.Decls) }
func (x GenDecl) Size() int    { return len(x.Specs) }
func (x ReturnStmt) Size() int { return len(x.Results) }

func (x BlockStmt) Get(i int) Ast  { return ToAst(x.List[i]) }
func (x FieldList) Get(i int) Ast  { return ToAst(x.List[i]) }
func (x File) Get(i int) Ast       { return ToAst(x.Decls[i]) }
func (x GenDecl) Get(i int) Ast    { return ToAst(x.Specs[i]) }
func (x ReturnStmt) Get(i int) Ast { return ToAst(x.Results[i]) }

func (x BlockStmt) Set(i int, child Ast)  { x.List[i] = ToStmt(child) }
func (x FieldList) Set(i int, child Ast)  { x.List[i] = ToField(child) }
func (x File) Set(i int, child Ast)       { x.Decls[i] = ToDecl(child) }
func (x GenDecl) Set(i int, child Ast)    { x.Specs[i] = ToSpec(child) }
func (x ReturnStmt) Set(i int, child Ast) { x.Results[i] = ToExpr(child) }

func (x BlockStmt) Resize(n int)  { s := make([]ast.Stmt, n); copy(s, x.List); x.List = s }
func (x FieldList) Resize(n int)  { s := make([]*ast.Field, n); copy(s, x.List); x.List = s }
func (x File) Resize(n int)       { s := make([]ast.Decl, n); copy(s, x.Decls); x.Decls = s }
func (x GenDecl) Resize(n int)    { s := make([]ast.Spec, n); copy(s, x.Specs); x.Specs = s }
func (x ReturnStmt) Resize(n int) { s := make([]ast.Expr, n); copy(s, x.Results); x.Results = s }
