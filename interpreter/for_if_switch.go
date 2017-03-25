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
 * for_if_switch.go
 *
 *  Created on: Feb 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	"go/token"
	r "reflect"
)

type eBreak struct {
	label string
}

type eContinue struct {
	label string
}

type eReturn struct {
	results []r.Value
}

func (_ eBreak) Error() string {
	return "break outside for or switch"
}

func (_ eContinue) Error() string {
	return "continue outside for"
}

func (_ eReturn) Error() string {
	return "return outside function"
}

func (env *Env) evalBranch(node *ast.BranchStmt) (r.Value, []r.Value) {
	var label string
	if node.Label != nil {
		label = node.Label.Name
	}
	switch node.Tok {
	case token.BREAK:
		panic(eBreak{label})
	case token.CONTINUE:
		panic(eContinue{label})
	case token.GOTO:
		return env.errorf("unimplemented: goto")
	case token.FALLTHROUGH:
		return env.errorf("invalid fallthrough: not the last statement in a case")
	default:
		return env.errorf("unimplemented branch: %v <%v>", node, r.TypeOf(node))
	}
}

func (env *Env) evalReturn(node *ast.ReturnStmt) (r.Value, []r.Value) {
	var rets []r.Value
	if len(node.Results) == 1 {
		// return foo() returns *all* the values returned by foo, not just the first one
		rets = packValues(env.evalExpr(node.Results[0]))
	} else {
		rets = env.evalExprs(node.Results)
	}
	panic(eReturn{rets})
}

func (env *Env) evalIf(node *ast.IfStmt) (r.Value, []r.Value) {
	if node.Init != nil {
		env = NewEnv(env, "if {}")
		_, _ = env.evalStatement(node.Init)
	}
	cond, _ := env.Eval(node.Cond)
	if cond.Kind() != r.Bool {
		cf := cond.Interface()
		return env.errorf("if: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
	}
	if cond.Bool() {
		return env.evalBlock(node.Body)
	} else if node.Else != nil {
		return env.evalStatement(node.Else)
	} else {
		return Nil, nil
	}
}

func (env *Env) evalFor(node *ast.ForStmt) (r.Value, []r.Value) {
	// Debugf("evalFor() init = %#v, cond = %#v, post = %#v, body = %#v", node.Init, node.Cond, node.Post, node.Body)

	if node.Init != nil {
		env = NewEnv(env, "for {}")
		env.evalStatement(node.Init)
	}
	for {
		if node.Cond != nil {
			cond := env.evalExpr1(node.Cond)
			if cond.Kind() != r.Bool {
				cf := cond.Interface()
				return env.errorf("for: invalid condition type <%T> %#v, expecting <bool>", cf, cf)
			}
			if !cond.Bool() {
				break
			}
		}
		if !env.evalForBodyOnce(node.Body) {
			break
		}
		if node.Post != nil {
			env.evalStatement(node.Post)
		}
	}
	return None, nil
}

func (env *Env) evalForRange(node *ast.RangeStmt) (r.Value, []r.Value) {
	// Debugf("evalForRange() init = %#v, cond = %#v, post = %#v, body = %#v", node.Init, node.Cond, node.Post, node.Body)

	container := env.evalExpr1(node.X)
	if container == Nil || container == None {
		return env.errorf("invalid for range: cannot iterate on nil: %v evaluated to %v", node.X, container)
	}

	switch container.Kind() {
	case r.Chan:
		return env.evalForRangeChannel(container, node)
	case r.Map:
		return env.evalForRangeMap(container, node)
	case r.Slice, r.Array:
		return env.evalForRangeSlice(container, node)
	case r.String:
		// Golang specs https://golang.org/ref/spec#RangeClause
		// "For a string value, the "range" clause iterates over the Unicode code points in the string"
		return env.evalForRangeString(container.String(), node)
	case r.Ptr:
		if container.Elem().Kind() == r.Array {
			return env.evalForRangeSlice(container.Elem(), node)
		}
	}
	return env.errorf("invalid for range: expecting array, channel, map, slice, string, or pointer to array, found: %v <%v>",
		container, typeOf(container))
}

func (env *Env) evalForRangeMap(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range map {}")
		t := obj.Type()
		k := env.defineForIterVar(knode, t.Key())
		v := env.defineForIterVar(vnode, t.Elem())

		for _, key := range obj.MapKeys() {
			if k != Nil {
				k.Set(key)
			}
			if v != Nil {
				v.Set(obj.MapIndex(key))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for _, key := range obj.MapKeys() {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, key)
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, obj.MapIndex(key))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeChannel(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	if node.Value != nil {
		return env.errorf("range expression is a channel: expecting at most one iteration variable, found two: %v %v", node.Key, node.Value)
	}

	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range channel {}")
		k := env.defineForIterVar(knode, obj.Type().Elem())

		for {
			recv, ok := obj.Recv()
			if !ok {
				break
			}
			if k != Nil {
				k.Set(recv)
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for {
			recv, ok := obj.Recv()
			if !ok {
				break
			}
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, recv)
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeString(str string, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range string {}")
		k := env.defineForIterVar(knode, typeOfInt)
		v := env.defineForIterVar(vnode, typeOfRune)

		for i, rune := range str {
			if k != Nil {
				k.Set(r.ValueOf(i))
			}
			if v != Nil {
				v.Set(r.ValueOf(rune))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		for i, rune := range str {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, r.ValueOf(i))
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, r.ValueOf(rune))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForRangeSlice(obj r.Value, node *ast.RangeStmt) (r.Value, []r.Value) {
	knode := nilIfIdentUnderscore(node.Key)
	vnode := nilIfIdentUnderscore(node.Value)
	tok := node.Tok
	switch tok {
	case token.DEFINE:
		env = NewEnv(env, "range slice/array {}")
		k := env.defineForIterVar(knode, typeOfInt)
		v := env.defineForIterVar(vnode, obj.Type().Elem())

		n := obj.Len()
		for i := 0; i < n; i++ {
			if k != Nil {
				k.Set(r.ValueOf(i))
			}
			if v != Nil {
				v.Set(obj.Index(i))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	case token.ASSIGN:
		n := obj.Len()
		for i := 0; i < n; i++ {
			// Golang specs https://golang.org/ref/spec#RangeClause
			// "Function calls on the left are evaluated once per iteration"
			//
			// we actually evaluate once per iteration the full expressions on the left
			if knode != nil {
				kplace := env.evalPlace(knode)
				env.assignPlace(kplace, tok, r.ValueOf(i))
			}
			if vnode != nil {
				vplace := env.evalPlace(vnode)
				env.assignPlace(vplace, tok, obj.Index(i))
			}
			if !env.evalForBodyOnce(node.Body) {
				break
			}
		}
	}
	return None, nil
}

func (env *Env) evalForBodyOnce(node *ast.BlockStmt) (cont bool) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case eBreak:
				cont = false
			case eContinue:
				cont = true
			default:
				panic(rec)
			}
		}
	}()
	env.evalBlock(node)
	return true
}

func (env *Env) defineForIterVar(node ast.Expr, t r.Type) r.Value {
	if node == nil || t == nil {
		return Nil
	}
	name := node.(*ast.Ident).Name
	env.defineVar(name, t, r.Zero(t))
	return env.Binds[name]
}

func nilIfIdentUnderscore(node ast.Expr) ast.Expr {
	if ident, ok := node.(*ast.Ident); ok {
		if ident.Name == "_" {
			return nil
		}
	}
	return node
}

func (env *Env) evalSwitch(node *ast.SwitchStmt) (ret r.Value, rets []r.Value) {
	if node.Init != nil {
		// the scope of variables defined in the init statement of a switch
		// is the switch itself
		env = NewEnv(env, "switch")
		env.evalStatement(node.Init)
	}
	var tag r.Value
	if node.Tag == nil {
		tag = valueOfTrue
	} else {
		tag = env.evalExpr1(node.Tag)
	}
	if node.Body == nil || len(node.Body.List) == 0 {
		return None, nil
	}
	isFallthrough := false
	cases := node.Body.List
	n := len(cases)
	default_i := n
	for i := 0; i < n; i++ {
		case_ := cases[i].(*ast.CaseClause)
		if !isFallthrough && case_.List == nil {
			// default will be executed later, if no case matches
			default_i = i
		} else if isFallthrough || env.caseMatchesTag(tag, case_.List) {
			ret, rets, isFallthrough = env.evalCaseBody(i == default_i, case_)
			if !isFallthrough {
				return ret, rets
			}
		}
	}
	// even "default:" can end with fallthrough...
	for i := default_i; i < n; i++ {
		case_ := cases[i].(*ast.CaseClause)
		ret, rets, isFallthrough = env.evalCaseBody(i == default_i, case_)
		if !isFallthrough {
			return ret, rets
		}
	}
	return None, nil
}

func (env *Env) evalCaseBody(isDefault bool, case_ *ast.CaseClause) (ret r.Value, rets []r.Value, isFallthrough bool) {
	if case_ == nil || len(case_.Body) == 0 {
		return None, nil, false
	}
	body := case_.Body
	n := len(body)
	// implement fallthrough
	if last, ok := body[n-1].(*ast.BranchStmt); ok {
		if last.Tok == token.FALLTHROUGH {
			isFallthrough = true
			body = body[0 : n-1]
		}
	}

	// each case body has its own environment
	label := "case:"
	if isDefault {
		label = "default:"
	}
	panicking := true
	defer func() {
		if panicking {
			switch pan := recover().(type) {
			case eBreak:
				isFallthrough = false
			default:
				panic(pan)
			}
		}
	}()
	env = NewEnv(env, label)
	ret, rets = env.evalStatements(body)
	panicking = false
	return
}

func (env *Env) caseMatchesTag(tag r.Value, list []ast.Expr) bool {
	var i interface{}
	var t r.Type = nil
	if tag != None && tag != Nil {
		i = tag.Interface()
		t = tag.Type()
	}
	for _, expr := range list {
		v := env.evalExpr1(expr)
		if t == nil {
			if v == Nil || v == None {
				return true
			}
		} else {
			v = env.valueToType(v, t)
			// https://golang.org/pkg/reflect
			// "To compare two Values, compare the results of the Interface method"
			if v.Interface() == i {
				return true
			}
		}
	}
	return false
}
