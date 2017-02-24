package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"

	mp "github.com/cosmos72/gomacro/macroparser"
)

// Walk traverses the whole AST tree using pre-order traversal,
// and replaces each node with the result of tranformer(node).
// If tranformer(node) returns nil, Walk stops the traversal and immediately returns node (not nil)
// Warning: it modifies the AST tree in place!
func (ir *Interpreter) Walk(transformer func(ast.Node) ast.Node, node ast.Node) ast.Node {
	if node == nil {
		return nil
	}
	//ir.Debugf("Walk() %v", node)
	t := transformer
	node2 := t(node)
	if node2 == nil {
		// stop traversal, but return original node instead of nil
		return node
	}
	//if node2 != node {
	//	ir.Debugf("Walk() transformed %v to %v", node, node2)
	//}
	node = node2
	switch node := node.(type) {
	case *ast.ArrayType:
		ir.walkExpr(t, &node.Len)
		ir.walkExpr(t, &node.Elt)
	case *ast.AssignStmt:
		ir.walkExprs(t, node.Lhs)
		ir.walkExprs(t, node.Rhs)
	case *ast.BadDecl, *ast.BadExpr, *ast.BadStmt, *ast.BasicLit:
		// nothing to do
	case *ast.BinaryExpr:
		ir.walkExpr(t, &node.X)
		ir.walkExpr(t, &node.Y)
	case *ast.BlockStmt:
		ir.walkStmts(t, node.List)
	case *ast.BranchStmt:
		ir.walkIdent(t, &node.Label)
	case *ast.CallExpr:
		ir.walkExpr(t, &node.Fun)
		ir.walkExprs(t, node.Args)
	case *ast.CaseClause:
		ir.walkExprs(t, node.List)
		ir.walkStmts(t, node.Body)
	case *ast.ChanType:
		ir.walkExpr(t, &node.Value)
	case *ast.CommClause:
		ir.walkStmt(t, &node.Comm)
		ir.walkStmts(t, node.Body)
	case *ast.CompositeLit:
		ir.walkExpr(t, &node.Type)
		ir.walkExprs(t, node.Elts)
	case *ast.DeclStmt:
		ir.walkDecl(t, &node.Decl)
	case *ast.DeferStmt:
		ir.walkExpr(t, &node.Call.Fun)
		ir.walkExprs(t, node.Call.Args)
	case *ast.Ellipsis:
		ir.walkExpr(t, &node.Elt)
	case *ast.EmptyStmt:
		// nothing to do
	case *ast.ExprStmt:
		ir.walkExpr(t, &node.X)
	case *ast.Field:
		ir.walkIdents(t, node.Names)
		ir.walkExpr(t, &node.Type)
	case *ast.FieldList:
		ir.walkFields(t, node.List)
	case *ast.File:
		ir.walkIdent(t, &node.Name)
		ir.walkImports(t, node.Imports)
		ir.walkDecls(t, node.Decls)
	case *ast.ForStmt:
		ir.walkStmt(t, &node.Init)
		ir.walkExpr(t, &node.Cond)
		ir.walkStmt(t, &node.Post)
		ir.walkBlock(t, &node.Body)
	case *ast.FuncDecl:
		ir.walkFieldList(t, &node.Recv)
		ir.walkIdent(t, &node.Name)
		ir.walkFuncType(t, &node.Type)
		ir.walkBlock(t, &node.Body)
	case *ast.FuncLit:
		ir.walkFuncType(t, &node.Type)
		ir.walkBlock(t, &node.Body)
	case *ast.FuncType:
		ir.walkFieldList(t, &node.Params)
		ir.walkFieldList(t, &node.Results)
	case *ast.GenDecl:
		ir.walkSpecs(t, node.Specs)
	case *ast.GoStmt:
		ir.walkExpr(t, &node.Call.Fun)
		ir.walkExprs(t, node.Call.Args)
	case *ast.Ident:
		/* nothing to do */
	case *ast.IfStmt:
		ir.walkStmt(t, &node.Init)
		ir.walkExpr(t, &node.Cond)
		ir.walkBlock(t, &node.Body)
		ir.walkStmt(t, &node.Else)
	case *ast.ImportSpec:
		ir.walkIdent(t, &node.Name)
	case *ast.IncDecStmt:
		ir.walkExpr(t, &node.X)
	case *ast.IndexExpr:
		ir.walkExpr(t, &node.X)
		ir.walkExpr(t, &node.Index)
	case *ast.InterfaceType:
		ir.walkFieldList(t, &node.Methods)
	case *ast.KeyValueExpr:
		ir.walkExpr(t, &node.Key)
		ir.walkExpr(t, &node.Value)
	case *ast.LabeledStmt:
		ir.walkIdent(t, &node.Label)
		ir.walkStmt(t, &node.Stmt)
	case *ast.MapType:
		ir.walkExpr(t, &node.Key)
		ir.walkExpr(t, &node.Value)
	case *ast.Package:
		ir.WalkFiles(t, node.Files)
	case *ast.ParenExpr:
		ir.walkExpr(t, &node.X)
	case *ast.RangeStmt:
		ir.walkExpr(t, &node.Key)
		ir.walkExpr(t, &node.Value)
		ir.walkExpr(t, &node.X)
		ir.walkBlock(t, &node.Body)
	case *ast.ReturnStmt:
		ir.walkExprs(t, node.Results)
	case *ast.SelectStmt:
		ir.walkBlock(t, &node.Body)
	case *ast.SelectorExpr:
		ir.walkExpr(t, &node.X)
		ir.walkIdent(t, &node.Sel)
	case *ast.SendStmt:
		ir.walkExpr(t, &node.Chan)
		ir.walkExpr(t, &node.Value)
	case *ast.SliceExpr:
		ir.walkExpr(t, &node.X)
		ir.walkExpr(t, &node.Low)
		ir.walkExpr(t, &node.High)
		if node.Slice3 {
			ir.walkExpr(t, &node.Max)
		}
	case *ast.StarExpr:
		ir.walkExpr(t, &node.X)
	case *ast.StructType:
		ir.walkFieldList(t, &node.Fields)
	case *ast.SwitchStmt:
		ir.walkStmt(t, &node.Init)
		ir.walkExpr(t, &node.Tag)
		ir.walkBlock(t, &node.Body)
	case *ast.TypeAssertExpr:
		ir.walkExpr(t, &node.X)
		ir.walkExpr(t, &node.Type)
	case *ast.TypeSpec:
		ir.walkIdent(t, &node.Name)
		ir.walkExpr(t, &node.Type)
	case *ast.TypeSwitchStmt:
		ir.walkStmt(t, &node.Init)
		ir.walkStmt(t, &node.Assign)
		ir.walkBlock(t, &node.Body)
	case *ast.UnaryExpr:
		ir.walkExpr(t, &node.X)
	case *ast.ValueSpec:
		ir.walkIdents(t, node.Names)
		ir.walkExpr(t, &node.Type)
		ir.walkExprs(t, node.Values)
	}
	return node
}

func (ir *Interpreter) walkDecls(t func(ast.Node) ast.Node, list []ast.Decl) {
	for i := range list {
		ir.walkDecl(t, &list[i])
	}
}

func (ir *Interpreter) walkExprs(t func(ast.Node) ast.Node, list []ast.Expr) {
	for i := range list {
		ir.walkExpr(t, &list[i])
	}
}

func (ir *Interpreter) walkFields(t func(ast.Node) ast.Node, list []*ast.Field) {
	for i := range list {
		ir.walkField(t, &list[i])
	}
}

func (ir *Interpreter) WalkFiles(t func(ast.Node) ast.Node, files map[string]*ast.File) {
	for name, file := range files {
		file2 := ir.ToFile(ir.Walk(t, file))
		if file2 != nil && file2 != file {
			files[name] = file2
		}
	}
}

func (ir *Interpreter) walkImports(t func(ast.Node) ast.Node, list []*ast.ImportSpec) {
	for i := range list {
		ir.walkImport(t, &list[i])
	}
}

func (ir *Interpreter) walkIdents(t func(ast.Node) ast.Node, list []*ast.Ident) {
	for i := range list {
		ir.walkIdent(t, &list[i])
	}
}

func (ir *Interpreter) walkSpecs(t func(ast.Node) ast.Node, list []ast.Spec) {
	for i := range list {
		ir.walkSpec(t, &list[i])
	}
}

func (ir *Interpreter) walkStmts(t func(ast.Node) ast.Node, list []ast.Stmt) {
	for i := range list {
		ir.walkStmt(t, &list[i])
	}
}

func (ir *Interpreter) walkDecl(t func(ast.Node) ast.Node, ptr *ast.Decl) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToDecl(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkBlock(t func(ast.Node) ast.Node, ptr **ast.BlockStmt) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToBlockStmt(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkExpr(t func(ast.Node) ast.Node, ptr *ast.Expr) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToExpr(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkField(t func(ast.Node) ast.Node, ptr **ast.Field) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToField(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkFuncType(t func(ast.Node) ast.Node, ptr **ast.FuncType) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToFuncType(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkFieldList(t func(ast.Node) ast.Node, ptr **ast.FieldList) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToFieldList(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkImport(t func(ast.Node) ast.Node, ptr **ast.ImportSpec) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToImportSpec(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkIdent(t func(ast.Node) ast.Node, ptr **ast.Ident) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToIdent(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkSpec(t func(ast.Node) ast.Node, ptr *ast.Spec) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToSpec(ir.Walk(t, *ptr)); node != nil && node != *ptr {
			*ptr = node
		}
	}
}

func (ir *Interpreter) walkStmt(t func(ast.Node) ast.Node, ptr *ast.Stmt) {
	if ptr != nil && *ptr != nil {
		if node := ir.ToStmt(ir.Walk(t, *ptr)); node != nil && node != *ptr {
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
	return &ast.UnaryExpr{Op: mp.MACRO, X: fun}
}
