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
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"io"
	"io/ioutil"
	r "reflect"
	"strconv"
)

func (p *Parser) ReadFromSource(src interface{}) string {
	switch src := src.(type) {
	case string:
		return src
	case []byte:
		return string(src)
	default:
		if reader, ok := src.(io.Reader); ok {
			bytes, err := ioutil.ReadAll(reader)
			if err != nil {
				Errore(err)
			}
			return string(bytes)
		}
		p.Errorf("unsupported source, cannot read from: %v <%v>", src, r.TypeOf(src))
	}
	return ""
}

func unescapeChar(str string) rune {
	// fmt.Printf("debug unescapeChar(): parsing CHAR %#v", str)
	rs := []rune(str)
	n := len(rs)
	if n >= 2 && rs[0] == '\'' && rs[n-1] == '\'' {
		rs = rs[1 : n-1]
	}
	ret, _, _, err := strconv.UnquoteChar(string(rs), '\'')
	if err != nil {
		Errore(err)
	}
	return ret
}

func unescapeString(str string) string {
	ret, err := strconv.Unquote(str)
	if err != nil {
		Errore(err)
	}
	return ret
}

func skipCommentsAndSpaces(str string) int {
	n := len(str)
	const (
		Normal = iota
		Slash
		LineComment
		MultiLineComment
		MultiLineCommentStar
	)
	mode := Normal
	for i := 0; i < n; i++ {
		ch := str[i]
		switch mode {
		case Normal:
			if ch == '/' {
				mode = Slash
			} else if ch > ' ' {
				return i
			}
		case Slash:
			if ch == '/' {
				mode = LineComment
			} else if ch == '*' {
				mode = MultiLineComment
			} else {
				return i - 1
			}
		case LineComment:
			if ch == '\n' {
				mode = Normal
			}
		case MultiLineComment:
			if ch == '*' {
				mode = MultiLineCommentStar
			}
		case MultiLineCommentStar:
			if ch == '/' {
				mode = Normal
			} else {
				mode = MultiLineComment
			}
		}
	}
	return n
}

func extractFirstIdentifier(str string) string {
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') ||
			ch == '_' || ch >= 128 ||
			(i != 0 && (ch >= '0' && ch <= '9')) {
		} else {
			return str[:i]
		}
	}
	return str
}
