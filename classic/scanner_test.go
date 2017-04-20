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
 * scanner.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func testScanner(env *Env) {
	src := []byte(" Quote { /*foo*/ x } y")

	var s scanner.Scanner
	var errs scanner.ErrorList

	pos0 := env.Fileset.Base()
	file := env.Fileset.AddFile("temp.go", pos0, len(src))

	s.Init(file, src, errs.Add, 0 /*scanner.ScanComments*/)

	for {
		pos, tok, str := s.Scan()
		fmt.Printf("%v\t%v\t%#v\n", pos, tok, str)
		if tok == token.EOF {
			break
		}
	}
}
