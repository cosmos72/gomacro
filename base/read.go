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
 * read.go
 *
 *  Created on: Mar 12, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	r "reflect"
)

func ReadBytes(src interface{}) []byte {
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

func ReadString(src interface{}) string {
	switch s := src.(type) {
	case []byte:
		if s != nil {
			return string(s)
		}
	case string:
		return s
	case *bytes.Buffer:
		// is io.Reader, but src is already available in string form
		if s != nil {
			return s.String()
		}
	case io.Reader:
		if s != nil {
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, s); err != nil {
				Error(err)
			}
			return buf.String()
		}
	}
	Errorf("unsupported source, cannot read from: %v <%v>", src, r.TypeOf(src))
	return ""
}

type ReadOptions int

const (
	ReadOptShowPrompt         ReadOptions = 1 << iota
	ReadOptCollectAllComments             // continue until non-comment is found. default is to return comments one by one
)

func ReadMultiline(in *bufio.Reader, opts ReadOptions, out io.Writer, prompt string) (src string, firstToken int, err error) {
	type Mode int
	const (
		mNormal Mode = iota
		mPlus
		mMinus
		mRune
		mString
		mRuneEscape
		mStringEscape
		mRawString
		mSlash
		mHash
		mLineComment
		mComment
		mCommentStar
		mTilde
	)
	mode := mNormal
	paren := 0
	optPrompt := opts&ReadOptShowPrompt != 0
	optAllComments := opts&ReadOptCollectAllComments != 0
	ignorenl := false
	firstToken = -1

	if optPrompt {
		fmt.Fprint(out, prompt)
	}
	// comments do not reset ignorenl
	resetnl := func(paren int, mode Mode) bool {
		return paren != 0 ||
			(mode != mNormal && mode != mSlash && mode != mHash &&
				mode != mLineComment && mode != mComment && mode != mCommentStar)
	}
	var line, buf []byte

	for {
		line, err = in.ReadBytes('\n')
		for i, ch := range line {
			Debugf("ReadMultiline: found %q, mode = %d, ignorenl = %t", ch, mode, ignorenl)
			switch mode {
			case mPlus, mMinus:
				if ch == '+' {
					if mode == mPlus {
						mode = mNormal
					} else {
						mode = mPlus
					}
					break
				} else if ch == '-' {
					if mode == mMinus {
						mode = mNormal
					} else {
						mode = mMinus
					}
					break
				}
				mode = mNormal
				ignorenl = true
				if ch <= ' ' {
					continue
				}
				fallthrough
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
					continue // no tokens yet
				case '#':
					mode = mHash // support #! line comments
					continue     // no tokens yet
				case '~':
					mode = mTilde
				case '!', '%', '&', '*', ',', '.', '<', '=', '>', '^', '|':
					ignorenl = paren == 0
				case '+':
					ignorenl = false
					if paren == 0 {
						mode = mPlus
					}
				case '-':
					ignorenl = false
					if paren == 0 {
						mode = mMinus
					}
				default:
					if ch <= ' ' {
						continue // not a token
					}
					ignorenl = false // found a token
				}
			case mRune:
				switch ch {
				case '\\':
					mode = mRuneEscape
				case '\'':
					mode = mNormal
				default:
					if ch < ' ' {
						return merge(buf, line[:i]), firstToken, invalidChar(ch, "rune")
					}
				}
			case mRuneEscape:
				if ch < ' ' {
					return merge(buf, line[:i]), firstToken, invalidChar(ch, "rune")
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
						return merge(buf, line[:i]), firstToken, invalidChar(ch, "string")
					}
				}
			case mStringEscape:
				if ch < ' ' {
					return merge(buf, line[:i]), firstToken, invalidChar(ch, "string")
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
					continue // no tokens
				case '*':
					mode = mComment
					continue // no tokens
				default:
					mode = mNormal
					if paren == 0 {
						ignorenl = true
					}
					if firstToken < 0 {
						firstToken = len(buf) + i - 1
					}
				}
			case mHash:
				switch ch {
				case '!':
					mode = mLineComment
					line[i-1] = '/'
					line[i] = '/'
					continue // no tokens
				default:
					mode = mNormal
					if firstToken < 0 {
						firstToken = len(buf) + i - 1
					}
				}
			case mLineComment:
				continue
			case mComment:
				switch ch {
				case '*':
					mode = mCommentStar
				}
				continue
			case mCommentStar:
				switch ch {
				case '/':
					mode = mNormal
				default:
					mode = mComment
				}
				continue
			case mTilde:
				mode = mNormal
			}
			// Debugf("ReadMultiline: mode = %d, ignorenl = %t, resetnl = %t", mode, ignorenl, resetnl(paren, mode))
			if resetnl(paren, mode) {
				ignorenl = false
				// Debugf("ReadMultiline: cleared ignorenl")
			}
			if firstToken < 0 {
				firstToken = len(buf) + i
				// Debugf("ReadMultiline: setting firstToken to %d, line up to it = %q", firstToken, line[:i+1])
			}
		}
		buf = append(buf, line...)
		if mode == mLineComment {
			mode = mNormal
		}
		if err != nil || paren <= 0 && !ignorenl && mode == mNormal && (firstToken >= 0 || !optAllComments) {
			break
		}
		// Debugf("ReadMultiline: continuing, mode = %d, paren == %d, ignorenl = %t", mode, paren, ignorenl)
		if mode == mPlus || mode == mMinus {
			mode = mNormal
		}
		if optPrompt {
			printDots(out, 4+2*paren)
		}
	}
	if err != nil {
		if err == io.EOF && paren > 0 {
			err = errors.New("unexpected EOF")
		}
		return string(buf), firstToken, err
	}
	// Debugf("ReadMultiline: read %d bytes, firstToken at %d", len(buf), firstToken)
	// if firstToken >= 0 {
	//     Debugf("ReadMultiline: comments: %q", buf[:firstToken])
	//     Debugf("ReadMultiline: tokens: %q", buf[firstToken:])
	// } else {
	//     Debugf("ReadMultiline: comments: %q", buf)
	// }
	return string(buf), firstToken, nil
}

func merge(buf, tail []byte) string {
	return string(append(buf, tail...))
}

func invalidChar(ch byte, ctx string) error {
	return errors.New(fmt.Sprintf("unexpected character %q inside %s literal", ch, ctx))
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
