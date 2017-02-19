package main

import (
	"go/ast"
	r "reflect"
)

func (env *Env) evalQuote(node *ast.BlockStmt) (r.Value, []r.Value) {
	var ret ast.Node
	switch len(node.List) {
	case 0:
		ret = nil
	case 1:
		ret = node.List[0]
	default:
		ret = node
	}
	return r.ValueOf(&ret).Elem(), nil
}
