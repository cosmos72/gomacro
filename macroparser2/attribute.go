/*
 * macroparser2 - A Go source code parser with support for Lisp-like macros
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
 * attribute.go
 *
 *  Created on: Feb 20, 2017
 *      Author: Massimiliano Ghilardi
 */

package macroparser2

import (
	"fmt"
	_ "go/ast"
	"go/token"
)

func (ident Ident) Bad() bool {
	return ident.string == nil || len(*ident.string) == 0
}

func (slice Slice) Bad() bool {
	return len(slice) == 0 || slice[0] == BAD_s
}

func (node Ident) String() string {
	if node.Bad() {
		return "_"
	}
	return *node.string
}

func (node Ident) set(p *MacroParser, attrs Attrs) Ident {
	if node.string == nil {
		p.error(p.pos, "cannot set attributes of nil *string")
	} else {
		p.identAttrs[node.string] = attrs
	}
	return node
}

func (node Slice) set(p *MacroParser, attrs Attrs) Slice {
	if len(node) == 0 {
		p.error(p.pos, fmt.Sprintf("cannot set attributes of zero-length %T", node))
	} else {
		p.sliceAttrs[&node[0]] = attrs
	}
	return node
}

func (node Ident) setAttr(p *MacroParser, key Attr, value interface{}) Ident {
	if node.string == nil {
		p.error(p.pos, "cannot set attribute of nil *string")
	} else {
		var attrs Attrs
		var ok bool
		n0 := node.string
		if attrs, ok = p.identAttrs[n0]; !ok {
			attrs = make(Attrs)
			p.identAttrs[n0] = attrs
		}
		attrs[key] = value
	}
	return node
}

func (node Slice) setAttr(p *MacroParser, key Attr, value interface{}) Slice {
	if len(node) == 0 {
		p.error(p.pos, fmt.Sprintf("cannot set attribute of zero-length %T", node))
	} else {
		var attrs Attrs
		var ok bool
		n0 := &node[0]
		if attrs, ok = p.sliceAttrs[n0]; !ok {
			attrs = make(Attrs)
			p.sliceAttrs[n0] = attrs
		}
		attrs[key] = value
	}
	return node
}

func (node Ident) Attr(p *MacroParser, key Attr) interface{} {
	if node.string == nil {
		p.error(p.pos, "cannot get attribute of nil *string")
	} else {
		n0 := node.string
		if attrs, ok := p.identAttrs[n0]; ok {
			if attr, ok := attrs[key]; ok {
				return attr
			}
		}
	}
	return nil
}

func (node Slice) Attr(p *MacroParser, key Attr) interface{} {
	if len(node) == 0 {
		p.error(p.pos, fmt.Sprintf("cannot get attribute of zero-length %T", node))
	} else {
		n0 := &node[0]
		if attrs, ok := p.sliceAttrs[n0]; ok {
			if attr, ok := attrs[key]; ok {
				return attr
			}
		}
	}
	return nil
}

func (node Ident) setObj(p *MacroParser, obj interface{}) Ident {
	return node.setAttr(p, OBJ, obj)
}

func (node Ident) Obj(p *MacroParser) interface{} {
	return node.Attr(p, OBJ)
}

func (node Ident) Pos(p *MacroParser) token.Pos {
	pos := node.Attr(p, POS)
	if pos != nil {
		if pos, ok := pos.(token.Pos); ok {
			return pos
		}
	}
	return token.NoPos
}

func (node Slice) Pos(p *MacroParser) token.Pos {
	pos := node.Attr(p, POS)
	if pos != nil {
		if pos, ok := pos.(token.Pos); ok {
			return pos
		}
	}
	return token.NoPos
}

func (node Ident) End(p *MacroParser) token.Pos {
	pos := node.Pos(p)
	if pos == token.NoPos || node.Bad() {
		return pos
	}
	return token.Pos(int(pos) + len(*node.string))
}

func (node Slice) End(p *MacroParser) token.Pos {
	end := node.Attr(p, END)
	if end != nil {
		if end, ok := end.(token.Pos); ok {
			if end != token.NoPos {
				return end
			}
		}
	}
	if node.Bad() {
		return token.NoPos
	}
	last := node[len(node)-1]
	switch last := last.(type) {
	case Node:
		return last.End(p)
	default:
		return token.NoPos
	}
}

/*
func (p *MacroParser) parseAny() Any {
	var node Any

	if p.tok == token.COMMENT {
		// advance to the next non-comment token
		p.next()
	}
	switch p.tok {
	case token.PACKAGE:
		node = p.parseFile()
	case token.IMPORT:
		node = p.parseGenDecl(token.IMPORT, p.parseImportSpec)
	case token.CONST, token.FUNC, token.TYPE, token.VAR, MACRO:
		node = p.parseDecl(syncDecl)
	default:
		node = p.parseStmt()
		if expr, ok := node.(*ast.ExprStmt); ok {
			// unwrap expressions
			node = expr.X
		}
	}
	return node
}
*/
