/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * util.go
 *
 *  Created on: Nov 23, 2019
 *      Author: Massimiliano Ghilardi
 */

package util

import (
	"strings"
	"unicode"
)

// return the string after last '/' in path
func FileName(path string) string {
	return path[1+strings.LastIndexByte(path, '/'):]
}

func TailIdentifier(s string) string {
	if len(s) == 0 {
		return s
	}
	// work on unicode runes, not on bytes
	chars := []rune(s)
	var i, n = 0, len(chars)
	var digit bool
	for i = n - 1; i >= 0; i-- {
		ch := chars[i]
		if ch < 0x80 {
			if ch >= 'A' && ch <= 'Z' || ch == '_' || ch >= 'a' && ch <= 'z' {
				digit = false
			} else if ch >= '0' && ch <= '9' {
				digit = true
			} else {
				break
			}
		} else if unicode.IsLetter(ch) {
			digit = false
		} else if unicode.IsDigit(ch) {
			digit = true
		} else {
			break
		}
	}
	if digit {
		i++
	}
	return string(chars[i+1:])
}
