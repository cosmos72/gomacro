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

// func pair(a, b int) (int, int) { return a, b }
// var a, b, c = pair(1, 2), 3, 4

func main() {
	env := NewEnv(nil)
	env.Parsermode = parser.Trace

	args := os.Args
	// args = []string{"gomacro", "macro foo() { }"}
	if len(args) > 1 {
		str := strings.Join(args[1:], " ")
		env.ParseEvalPrint(str)
	} else {
		// testScanner(env)
		env.Repl()
	}
}
