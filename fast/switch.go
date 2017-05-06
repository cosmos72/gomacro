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
 * switch.go
 *
 *  Created on May 06, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
	"sort"
)

func (c *Comp) Switch(node *ast.SwitchStmt, labels []string) {
	c.Errorf("unimplemented statement: %v <%v>", node, r.TypeOf(node))

	initLocals := false
	var initBinds [2]int
	c, initLocals = c.pushEnvIfLocalBinds(&initBinds, node.Init)
	if node.Init != nil {
		c.Stmt(node.Init)
	}
	var tag *Expr
	if node.Tag == nil {
		// "switch { }" without an expression means "switch true { }"
		tag = exprValue(true)
	} else {
		tag = c.Expr1(node.Tag)
	}
	var ibreak int
	sort.Strings(labels)
	c.Loop = &LoopInfo{
		Break:      &ibreak,
		ThisLabels: labels,
	}

	if node.Body != nil {
		c.switchCase(node.Body.List, tag)
	}

	c = c.popEnvIfLocalBinds(initLocals, &initBinds, node.Init)
}

func (c *Comp) switchCase(list []ast.Stmt, tag *Expr) {
	c.Errorf("unimplemented: switch { case }")
}
