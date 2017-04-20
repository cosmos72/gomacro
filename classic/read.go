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

package classic

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
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
	ReadOptShowPrompt ReadOptions = 1 << iota
	ReadOptCollectAllComments
)

func ReadMultiline(in *bufio.Reader, opts ReadOptions, out io.Writer, prompt string) (src string, comments string, err error) {
	type Mode int
	const (
		mNormal Mode = iota
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
	tokens := false // true if some non-comment was found
	optPrompt := opts&ReadOptShowPrompt != 0
	optAllComments := opts&ReadOptCollectAllComments != 0

	if optPrompt {
		fmt.Fprint(out, prompt)
	}
	var line, buf, commentbuf []byte
	for {
		line, err = in.ReadBytes('\n')
		for i, ch := range line {
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
					continue
				case '#':
					mode = mHash // support #! line comments
					continue
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
					continue
				case '*':
					mode = mComment
					continue
				default:
					mode = mNormal
				}
			case mHash:
				switch ch {
				case '!':
					mode = mLineComment
					line[i-1] = '/'
					line[i] = '/'
					continue
				default:
					mode = mNormal
				}
			case mLineComment:
				switch ch {
				case '\n':
					mode = mNormal
				default:
					continue
				}
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
					continue
				}
			case mTilde:
				mode = mNormal
			}
			if !tokens && optAllComments {
				commentbuf = append(commentbuf, line[:i+1]...)
			}
			tokens = true
		}
		if !tokens && optAllComments {
			commentbuf = append(commentbuf, line...)
		}
		buf = append(buf, line...)
		if err != nil || paren <= 0 && mode == mNormal && (tokens || !optAllComments) {
			break
		}
		if optPrompt {
			printDots(out, 4+2*paren)
		}
	}
	if err != nil {
		if err == io.EOF && paren > 0 {
			err = errors.New("unexpected EOF")
		}
		return string(buf), string(commentbuf), err
	}
	return string(buf), string(commentbuf), nil
}

func invalidChar(ch byte, ctx string) (string, string, error) {
	return "\n", "", errors.New(fmt.Sprintf("unexpected character %q inside %s literal", ch, ctx))
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
