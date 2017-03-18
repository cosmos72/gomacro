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
 * readmultiline.go
 *
 *  Created on: Mar 12, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	r "reflect"
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
				error_(err)
			}
			return buf.Bytes()
		}
	}
	errorf("unsupported source, cannot read from: %v <%v>", src, r.TypeOf(src))
	return nil
}

func ReadMultiline(in *bufio.Reader, showPrompt bool, out io.Writer, prompt string) (string, error) {
	var buf []byte
	type Mode int
	const (
		mNormal Mode = iota
		mRune
		mString
		mRuneEscape
		mStringEscape
		mRawString
		mSlash
		mLineComment
		mComment
		mCommentStar
		mTilde
	)
	mode := mNormal
	paren := 0

	if showPrompt {
		fmt.Fprint(out, prompt)
	}
	for {
		line, err := in.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		for _, ch := range line {
			switch mode {
			case mNormal:
				switch ch {
				case '(', '[', '{':
					paren++
				case ')', ']', '}':
					paren--
				case '\'':
					mode = mRune
				case '"':
					mode = mString
				case '`':
					mode = mRawString
				case '/':
					mode = mSlash
				case '~':
					mode = mTilde
				}
			case mRune:
				switch ch {
				case '\\':
					mode = mRuneEscape
				case '\'':
					mode = mNormal
				default:
					if ch < ' ' {
						return invalidChar(ch, "rune")
					}
				}
			case mRuneEscape:
				if ch < ' ' {
					return invalidChar(ch, "rune")
				}
				mode = mRune
			case mString:
				switch ch {
				case '\\':
					mode = mStringEscape
				case '"':
					mode = mNormal
				default:
					if ch < ' ' {
						return invalidChar(ch, "string")
					}
				}
			case mStringEscape:
				if ch < ' ' {
					return invalidChar(ch, "string")
				}
				mode = mString
			case mRawString:
				switch ch {
				case '`':
					mode = mNormal
				}
			case mSlash:
				switch ch {
				case '/':
					mode = mLineComment
				case '*':
					mode = mComment
				default:
					mode = mNormal
				}
			case mLineComment:
				switch ch {
				case '\n':
					mode = mNormal
				}
			case mComment:
				switch ch {
				case '*':
					mode = mCommentStar
				}
			case mCommentStar:
				switch ch {
				case '/':
					mode = mNormal
				default:
					mode = mComment
				}
			case mTilde:
				mode = mNormal
			}
		}
		buf = append(buf, line...)
		if paren <= 0 && mode == mNormal {
			break
		}
		if showPrompt {
			printDots(out, 4+2*paren)
		}
	}
	return string(buf), nil
}

func invalidChar(ch byte, ctx string) (string, error) {
	return "", errors.New(fmt.Sprintf("unexpected character %q inside %s literal", ch, ctx))
}

func printDots(out io.Writer, count int) {
	const (
		dots  = ". . . . . . . . . . . . . . . . "
		ndots = len(dots)
	)
	for count >= ndots {
		fmt.Fprint(out, dots)
		count -= ndots
	}
	if count > 0 {
		fmt.Fprint(out, dots[0:count])
	}
}
