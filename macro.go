package main

import (
	"go/ast"
	r "reflect"
)

func (env *Env) evalQuote(nodes []ast.Expr) (r.Value, []r.Value) {
	switch n := len(nodes); n {
	case 0:
		return None, nil
	case 1:
		return r.ValueOf(&nodes[0]).Elem(), nil
	default:
		stmts := make([]ast.Stmt, n)
		for i := 0; i < n; i++ {
			stmts[i] = &ast.ExprStmt{X: nodes[i]}
		}

		var block ast.Node = &ast.BlockStmt{
			Lbrace: nodes[0].Pos(),
			List:   stmts,
			Rbrace: nodes[n-1].End(),
		}

		return r.ValueOf(&block).Elem(), nil
	}
}
