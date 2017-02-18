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
 * string.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	// "errors"
	// "fmt"
	"strconv"
)

const (
	eNormal = iota
	eBackslash
	eHex
	eOctal
	eUni4
	eUni8
)

func unescapeChar(str string) rune {
	// fmt.Printf("debug unescapeChar(): parsing CHAR %#v", str)
	rs := []rune(str)
	n := len(rs)
	if n >= 2 && rs[0] == '\'' && rs[n-1] == '\'' {
		rs = rs[1 : n-1]
	}
	ret, _, _, err := strconv.UnquoteChar(string(rs), '\'')
	if err != nil {
		Error(err)
	}
	return ret
}

func unescapeString(str string) string {
	ret, err := strconv.Unquote(str)
	if err != nil {
		Error(err)
	}
	return ret
}
