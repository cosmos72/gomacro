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
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"go/parser"
	"os"
	"strings"
)

func main() {
	env := NewEnv(nil)
	env.ParserMode = parser.Trace & 0
	env.Options = OptShowEvalDuration & 0

	args := os.Args
	// args = []string{"gomacro", "macro foo(a, b, c interface{}) interface{} { b }\nMacroExpand1(quote{foo x; y; z})"}
	if len(args) > 1 {
		str := strings.Join(args[1:], " ")
		env.ParseEvalPrint(str)
	} else {
		// testScanner(env)
		env.Repl()
	}
}
