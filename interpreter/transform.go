/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * assignment.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"

	mt "github.com/cosmos72/gomacro/token"
)

// Transform traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of tranformer(node).
// If tranformer(node) returns nil, Transform stops the traversal and immediately returns node (not nil)
// Warning: it modifies the AST tree in place!
func (ir *Interpreter) Transform(transformer func(in ast.Node) (out ast.Node, traverseOut bool), node ast.Node) ast.Node {
	if node == nil {
		return nil
	}
	//ir.Debugf("Transform() %v", node)
	t := transformer
	out, traverseOut := t(node)
	if !traverseOut {
		return out
	}
	//if out != node {
	//	ir.Debugf("Transform() transformed %v to %v", node, out)
	//}
	switch node := out.(type) {
	case *ast.ArrayType:
		ir.transExpr(t, &node.Len)
		ir.transExpr(t, &node.Elt)
	case *ast.AssignStmt:
		ir.transExprs(t, node.Lhs)
		ir.transExprs(t, node.Rhs)
	case *ast.BadDecl, *ast.BadExpr, *ast.BadStmt, *ast.BasicLit:
		// nothing to do
	case *ast.BinaryExpr:
		ir.transExpr(t, &node.X)
		ir.transExpr(t, &node.Y)
	case *ast.BlockStmt:
		ir.transStmts(t, node.List)
	case *ast.BranchStmt:
		ir.transIdent(t, &node.Label)
	case *ast.CallExpr:
		ir.transExpr(t, &node.Fun)
		ir.transExprs(t, node.Args)
	case *ast.CaseClause:
		ir.transExprs(t, node.List)
		ir.transStmts(t, node.Body)
	case *ast.ChanType:
		ir.transExpr(t, &node.Value)
	case *ast.CommClause:
		ir.transStmt(t, &node.Comm)
		ir.transStmts(t, node.Body)
	case *ast.CompositeLit:
		ir.transExpr(t, &node.Type)
		ir.transExprs(t, node.Elts)
	case *ast.DeclStmt:
		ir.transDecl(t, &node.Decl)
	case *ast.DeferStmt:
		ir.transExpr(t, &node.Call.Fun)
		ir.transExprs(t, node.Call.Args)
	case *ast.Ellipsis:
		ir.transExpr(t, &node.Elt)
	case *ast.EmptyStmt:
		// nothing to do
	case *ast.ExprStmt:
		ir.transExpr(t, &node.X)
	case *ast.Field:
		ir.transIdents(t, node.Names)
		ir.transExpr(t, &node.Type)
	case *ast.FieldList:
		ir.transFields(t, node.List)
	case *ast.File:
		ir.transIdent(t, &node.Name)
		ir.transImports(t, node.Imports)
		ir.transDecls(t, node.Decls)
	case *ast.ForStmt:
		ir.transStmt(t, &node.Init)
		ir.transExpr(t, &node.Cond)
		ir.transStmt(t, &node.Post)
		ir.transBlock(t, &node.Body)
	case *ast.FuncDecl:
		ir.transFieldList(t, &node.Recv)
		ir.transIdent(t, &node.Name)
		ir.transFuncType(t, &node.Type)
		ir.transBlock(t, &node.Body)
	case *ast.FuncLit:
		ir.transFuncType(t, &node.Type)
		ir.transBlock(t, &node.Body)
	case *ast.FuncType:
		ir.transFieldList(t, &node.Params)
		ir.transFieldList(t, &node.Results)
	case *ast.GenDecl:
		ir.transSpecs(t, node.Specs)
	case *ast.GoStmt:
		ir.transExpr(t, &node.Call.Fun)
		ir.transExprs(t, node.Call.Args)
	case *ast.Ident:
		/* nothing to do */
	case *ast.IfStmt:
		ir.transStmt(t, &node.Init)
		ir.transExpr(t, &node.Cond)
		ir.transBlock(t, &node.Body)
		ir.transStmt(t, &node.Else)
	case *ast.ImportSpec:
		ir.transIdent(t, &node.Name)
	case *ast.IncDecStmt:
		ir.transExpr(t, &node.X)
	case *ast.IndexExpr:
		ir.transExpr(t, &node.X)
		ir.transExpr(t, &node.Index)
	case *ast.InterfaceType:
		ir.transFieldList(t, &node.Methods)
	case *ast.KeyValueExpr:
		ir.transExpr(t, &node.Key)
		ir.transExpr(t, &node.Value)
	case *ast.LabeledStmt:
		ir.transIdent(t, &node.Label)
		ir.transStmt(t, &node.Stmt)
	case *ast.MapType:
		ir.transExpr(t, &node.Key)
		ir.transExpr(t, &node.Value)
	case *ast.Package:
		ir.TransformFiles(t, node.Files)
	case *ast.ParenExpr:
		ir.transExpr(t, &node.X)
	case *ast.RangeStmt:
		ir.transExpr(t, &node.Key)
		ir.transExpr(t, &node.Value)
		ir.transExpr(t, &node.X)
		ir.transBlock(t, &node.Body)
	case *ast.ReturnStmt:
		ir.transExprs(t, node.Results)
	case *ast.SelectStmt:
		ir.transBlock(t, &node.Body)
	case *ast.SelectorExpr:
		ir.transExpr(t, &node.X)
		ir.transIdent(t, &node.Sel)
	case *ast.SendStmt:
		ir.transExpr(t, &node.Chan)
		ir.transExpr(t, &node.Value)
	case *ast.SliceExpr:
		ir.transExpr(t, &node.X)
		ir.transExpr(t, &node.Low)
		ir.transExpr(t, &node.High)
		if node.Slice3 {
			ir.transExpr(t, &node.Max)
		}
	case *ast.StarExpr:
		ir.transExpr(t, &node.X)
	case *ast.StructType:
		ir.transFieldList(t, &node.Fields)
	case *ast.SwitchStmt:
		ir.transStmt(t, &node.Init)
		ir.transExpr(t, &node.Tag)
		ir.transBlock(t, &node.Body)
	case *ast.TypeAssertExpr:
		ir.transExpr(t, &node.X)
		ir.transExpr(t, &node.Type)
	case *ast.TypeSpec:
		ir.transIdent(t, &node.Name)
		ir.transExpr(t, &node.Type)
	case *ast.TypeSwitchStmt:
		ir.transStmt(t, &node.Init)
		ir.transStmt(t, &node.Assign)
		ir.transBlock(t, &node.Body)
	case *ast.UnaryExpr:
		ir.transExpr(t, &node.X)
	case *ast.ValueSpec:
		ir.transIdents(t, node.Names)
		ir.transExpr(t, &node.Type)
		ir.transExprs(t, node.Values)
	}
	return out
}

func (ir *Interpreter) transDecls(t func(ast.Node) (ast.Node, bool), list []ast.Decl) {
	for i := range list {
		ir.transDecl(t, &list[i])
	}
}

func (ir *Interpreter) transExprs(t func(ast.Node) (ast.Node, bool), list []ast.Expr) {
	for i := range list {
		ir.transExpr(t, &list[i])
	}
}

func (ir *Interpreter) transFields(t func(ast.Node) (ast.Node, bool), list []*ast.Field) {
	for i := range list {
		ir.transField(t, &list[i])
	}
}

func (ir *Interpreter) TransformFiles(t func(ast.Node) (ast.Node, bool), files map[string]*ast.File) {
	for name, file := range files {
		file2 := ir.ToFile(ir.Transform(t, file))
		if file2 != nil && file2 != file {
			files[name] = file2
		}
	}
}

func (ir *Interpreter) transImports(t func(ast.Node) (ast.Node, bool), list []*ast.ImportSpec) {
	for i := range list {
		ir.transImport(t, &list[i])
	}
}

func (ir *Interpreter) transIdents(t func(ast.Node) (ast.Node, bool), list []*ast.Ident) {
	for i := range list {
		ir.transIdent(t, &list[i])
	}
}

func (ir *Interpreter) transSpecs(t func(ast.Node) (ast.Node, bool), list []ast.Spec) {
	for i := range list {
		ir.transSpec(t, &list[i])
	}
}

func (ir *Interpreter) transStmts(t func(ast.Node) (ast.Node, bool), list []ast.Stmt) {
	for i := range list {
		ir.transStmt(t, &list[i])
	}
}

func (ir *Interpreter) transDecl(t func(ast.Node) (ast.Node, bool), ptr *ast.Decl) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToDecl(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transBlock(t func(ast.Node) (ast.Node, bool), ptr **ast.BlockStmt) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToBlockStmt(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transExpr(t func(ast.Node) (ast.Node, bool), ptr *ast.Expr) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToExpr(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transField(t func(ast.Node) (ast.Node, bool), ptr **ast.Field) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToField(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transFuncType(t func(ast.Node) (ast.Node, bool), ptr **ast.FuncType) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToFuncType(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transFieldList(t func(ast.Node) (ast.Node, bool), ptr **ast.FieldList) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToFieldList(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transImport(t func(ast.Node) (ast.Node, bool), ptr **ast.ImportSpec) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToImportSpec(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transIdent(t func(ast.Node) (ast.Node, bool), ptr **ast.Ident) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToIdent(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transSpec(t func(ast.Node) (ast.Node, bool), ptr *ast.Spec) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToSpec(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) transStmt(t func(ast.Node) (ast.Node, bool), ptr *ast.Stmt) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToStmt(ir.Transform(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) ToBlockStmt(node ast.Node) *ast.BlockStmt {
	switch node := node.(type) {
	case *ast.BlockStmt:
		return node
	case nil:
		break
	default:
		stmt := ir.ToStmt(node)
		return &ast.BlockStmt{Lbrace: stmt.Pos(), List: []ast.Stmt{stmt}, Rbrace: stmt.End()}
	}
	return nil
}

func (ir *Interpreter) ToDecl(node ast.Node) ast.Decl {
	switch node := node.(type) {
	case ast.Decl:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type ast.Decl", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToField(node ast.Node) *ast.Field {
	switch node := node.(type) {
	case *ast.Field:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToFile(node ast.Node) *ast.File {
	switch node := node.(type) {
	case *ast.File:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.File", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToFieldList(node ast.Node) *ast.FieldList {
	switch node := node.(type) {
	case *ast.FieldList:
		return node
	case *ast.Field:
		return &ast.FieldList{Opening: node.Pos(), List: []*ast.Field{node}, Closing: node.End()}
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToFuncType(node ast.Node) *ast.FuncType {
	switch node := node.(type) {
	case *ast.FuncType:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.FuncType", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToExpr(node ast.Node) ast.Expr {
	switch node := node.(type) {
	case ast.Expr:
		return node
	case *ast.BlockStmt:
		return ir.BlockStmtToExpr(node)
	case *ast.EmptyStmt:
		return &ast.Ident{NamePos: node.Semicolon, Name: "nil"}
	case *ast.ExprStmt:
		return node.X
	case ast.Stmt:
		list := []ast.Stmt{node}
		block := &ast.BlockStmt{List: list}
		return ir.BlockStmtToExpr(block)
	case nil:
		break
	default:
		ir.Errorf("unimplemented conversion from <%v> to ast.Expr", r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToImportSpec(node ast.Node) *ast.ImportSpec {
	switch node := node.(type) {
	case *ast.ImportSpec:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.ImportSpec", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToIdent(node ast.Node) *ast.Ident {
	switch node := node.(type) {
	case *ast.Ident:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type *ast.Ident", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToSpec(node ast.Node) ast.Spec {
	switch node := node.(type) {
	case ast.Spec:
		return node
	case nil:
		break
	default:
		ir.Errorf("cannot use %v <%v> as type ast.Spec", node, r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) ToStmt(node ast.Node) ast.Stmt {
	switch node := node.(type) {
	case ast.Stmt:
		return node
	case ast.Decl:
		return &ast.DeclStmt{Decl: node}
	case ast.Expr:
		return &ast.ExprStmt{X: node}
	case nil:
		break
	default:
		ir.Errorf("unimplemented conversion from <%v> to ast.Stmt", r.TypeOf(node))
	}
	return nil
}

func (ir *Interpreter) BlockStmtToExpr(node *ast.BlockStmt) ast.Expr {
	if node == nil {
		return nil
	}
	list := node.List
	switch len(list) {
	case 0:
		// convert {} to nil, because {} in expression context means "no useful value"
		return &ast.Ident{NamePos: node.Lbrace, Name: "nil"}
	case 1:
		// check if we are lucky...
		switch node := list[0].(type) {
		case *ast.ExprStmt:
			return node.X
		case *ast.EmptyStmt:
			// convert { ; } to nil, because { ; } in expression context means "no useful value"
			return &ast.Ident{NamePos: node.Semicolon, Name: "nil"}
		}
	}

	// due to go/ast strictly typed model, there is only one mechanism
	// to insert a statement inside an expression: use a closure.
	// so we return a unary expression: MACRO (func() { /*block*/ })
	typ := &ast.FuncType{Func: token.NoPos, Params: &ast.FieldList{}}
	fun := &ast.FuncLit{Type: typ, Body: node}
	return &ast.UnaryExpr{Op: mt.MACRO, X: fun}
}
