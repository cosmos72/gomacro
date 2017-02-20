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
 * macro.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/ast"
	"go/token"
	r "reflect"

	mp "github.com/cosmos72/gomacro/macroparser"
)

func (env *Env) evalQuote(op token.Token, node *ast.BlockStmt) (r.Value, []r.Value) {
	var ret ast.Node
	switch op {
	case mp.QUOTE, mp.QUASIQUOTE:
		if len(node.List) != 1 {
			return env.Errorf("invalid %s: contains %d statements, expecting exactly one: %v", mp.TokenString(op), node)
		}
		ret = node.List[0]
	case mp.UNQUOTE, mp.UNQUOTE_SPLICE:
		return env.Errorf("unimplemented %s: %v", mp.TokenString(op), node)
	}
	return r.ValueOf(&ret).Elem(), nil
}
