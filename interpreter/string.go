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

package interpreter

import (
	"bytes"
	"io"
	r "reflect"
	"strconv"
)

func Read(src interface{}) []byte {
	switch s := src.(type) {
	case []byte:
		if s != nil {
			return s
		}
	case string:
		return []byte(s)
	case *bytes.Buffer:
		// is io.Reader, but src is already available in []byte form
		if s != nil {
			return s.Bytes()
		}
	case io.Reader:
		if s != nil {
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, s); err != nil {
				Error(err)
			}
			return buf.Bytes()
		}
	}
	Errorf("unsupported source, cannot read from: %v <%v>", src, r.TypeOf(src))
	return nil
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

func findFirstToken(src []byte) int {
	n := len(src)
	const (
		Normal = iota
		Slash
		LineComment
		MultiLineComment
		MultiLineCommentStar
	)
	mode := Normal
	for i := 0; i < n; i++ {
		ch := src[i]
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

func extractFirstIdentifier(src []byte) []byte {
	n := len(src)
	for i := 0; i < n; i++ {
		ch := src[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') ||
			ch == '_' || ch >= 128 ||
			(i != 0 && (ch >= '0' && ch <= '9')) {
		} else {
			return src[:i]
		}
	}
	return src
}
