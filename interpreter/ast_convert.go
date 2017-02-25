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
 * ast_convert.go
 *
 *  Created on: Feb 25, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"

	mt "github.com/cosmos72/gomacro/token"
)

// ToAst2 returns either n0 (if i == 0) or n1, converted to Ast
func ToAst1(i int, node ast.Node) Ast {
	if i == 0 {
		return ToAst(node)
	} else {
		return BadIndex(i, 2)
	}
}

// ToAst2 returns either n0 (if i == 0) or n1, converted to Ast
func ToAst2(i int, n0 ast.Node, n1 ast.Node) Ast {
	var n ast.Node
	if i == 0 {
		n = n0
	} else if i == 1 {
		n = n1
	} else {
		return BadIndex(i, 2)
	}
	return ToAst(n)
}

func ToAst3(i int, n0 ast.Node, n1 ast.Node, n2 *ast.BlockStmt) Ast {
	var n ast.Node
	switch i {
	case 0:
		n = n0
	case 1:
		n1 = n1
	case 2:
		return BlockStmt{n2}
	default:
		return BadIndex(i, 3)
	}
	return ToAst(n)
}

// ToAst converts an ast.Node to Ast, providing uniform access to the node contents
//
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
		Errorf("unsupported node type <%v>", r.TypeOf(node))
	}
	return x
}

// ToNode converts Ast to ast.Node, or panics on failure
// (it fails if the argument is not AstWithNode)
func ToNode(x Ast) ast.Node {
	return x.(AstWithNode).Node()
}

func ToBasicLit(x Ast) *ast.BasicLit {
	switch node := ToNode(x).(type) {
	case *ast.BasicLit:
		return node
	case nil:
		break
	default:
		Errorf("cannot convert %v <%v> to *ast.BasicLit", node, r.TypeOf(node))
	}
	return nil
}

func ToBlockStmt(x Ast) *ast.BlockStmt {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.BlockStmt:
		return node
	default:
		stmt := ToStmt(x)
		return &ast.BlockStmt{Lbrace: stmt.Pos(), List: []ast.Stmt{stmt}, Rbrace: stmt.End()}
	}
	return nil
}

func ToCallExpr(x Ast) *ast.CallExpr {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.CallExpr:
		return node
	default:
		Errorf("cannot convert %v <%v> to *ast.CallExpr", node, r.TypeOf(node))
	}
	return nil
}

func ToDecl(x Ast) ast.Decl {
	switch node := ToNode(x).(type) {
	case nil:
	case ast.Decl:
		return node
	default:
		Errorf("cannot convert %v <%v> to ast.Decl", node, r.TypeOf(node))
	}
	return nil
}

func ToExpr(x Ast) ast.Expr {
	switch node := ToNode(x).(type) {
	case nil:
	case ast.Expr:
		return node
	case *ast.BlockStmt:
		return BlockStmtToExpr(node)
	case *ast.EmptyStmt:
		return &ast.Ident{NamePos: node.Semicolon, Name: "nil"}
	case *ast.ExprStmt:
		return node.X
	case ast.Stmt:
		list := []ast.Stmt{node}
		block := &ast.BlockStmt{List: list}
		return BlockStmtToExpr(block)
	default:
		Errorf("unimplemented conversion from %v <%v> to ast.Expr", node, r.TypeOf(node))
	}
	return nil
}

func ToExprSlice(x Ast) []ast.Expr {
	return *x.(ExprSlice).P
}

func ToField(x Ast) *ast.Field {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.Field:
		return node
	default:
		Errorf("cannot convert %v <%v> to *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func ToFile(x Ast) *ast.File {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.File:
		return node
	default:
		Errorf("cannot convert %v <%v> to *ast.File", node, r.TypeOf(node))
	}
	return nil
}

func ToFieldList(x Ast) *ast.FieldList {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.FieldList:
		return node
	case *ast.Field:
		return &ast.FieldList{Opening: node.Pos(), List: []*ast.Field{node}, Closing: node.End()}
	default:
		Errorf("cannot convert %v <%v> to *ast.Field", node, r.TypeOf(node))
	}
	return nil
}

func ToFuncType(x Ast) *ast.FuncType {
	switch node := ToNode(x).(type) {
	case nil:
	case *ast.FuncType:
		return node
	default:
		Errorf("cannot convert %v <%v> to *ast.FuncType", node, r.TypeOf(node))
	}
	return nil
}

func ToImportSpec(x Ast) *ast.ImportSpec {
	switch node := ToNode(x).(type) {
	case *ast.ImportSpec:
		return node
	case nil:
		break
	default:
		Errorf("cannot convert %v <%v> to *ast.ImportSpec", node, r.TypeOf(node))
	}
	return nil
}

func ToIdent(x Ast) *ast.Ident {
	switch node := ToNode(x).(type) {
	case *ast.Ident:
		return node
	case nil:
		break
	default:
		Errorf("cannot convert %v <%v> to *ast.Ident", node, r.TypeOf(node))
	}
	return nil
}

func ToIdentSlice(x Ast) []*ast.Ident {
	return *x.(IdentSlice).P
}

func ToSpec(x Ast) ast.Spec {
	switch node := ToNode(x).(type) {
	case ast.Spec:
		return node
	case nil:
		break
	default:
		Errorf("cannot convert %v <%v> to ast.Spec", node, r.TypeOf(node))
	}
	return nil
}

func ToStmt(x Ast) ast.Stmt {
	switch node := ToNode(x).(type) {
	case ast.Stmt:
		return node
	case ast.Decl:
		return &ast.DeclStmt{Decl: node}
	case ast.Expr:
		return &ast.ExprStmt{X: node}
	case nil:
		break
	default:
		Errorf("unimplemented conversion from %v <%v> to ast.Stmt", node, r.TypeOf(node))
	}
	return nil
}

func ToStmtSlice(x Ast) []ast.Stmt {
	return *x.(StmtSlice).P
}

func BlockStmtToExpr(node *ast.BlockStmt) ast.Expr {
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
