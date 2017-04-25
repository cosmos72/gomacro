/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * declaration.go
 *
 *  Created on Apr 01, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	"go/token"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

// Decl compiles a constant, variable, function or type declaration - or an import
func (c *Comp) Decl(node ast.Decl) {
	if node != nil {
		c.Pos = node.Pos()
	}
	switch node := node.(type) {
	case *ast.GenDecl:
		c.GenDecl(node)
	case *ast.FuncDecl:
		c.FuncDecl(node)
	default:
		c.Errorf("Compile: unsupported declaration, expecting <*ast.GenDecl> or <*ast.FuncDecl>, found: %v <%v>", node, r.TypeOf(node))
	}
}

// GenDecl compiles a constant, variable or type declaration - or an import
func (c *Comp) GenDecl(node *ast.GenDecl) {
	switch node.Tok {
	case token.IMPORT:
		for _, decl := range node.Specs {
			c.Import(decl)
		}
	case token.CONST:
		var defaultType ast.Expr
		var defaultExprs []ast.Expr
		top := c.Top()
		top.addIota()
		defer top.removeIota()
		for _, decl := range node.Specs {
			c.DeclConsts(decl, defaultType, defaultExprs)
			if valueSpec, ok := decl.(*ast.ValueSpec); ok && valueSpec.Values != nil {
				defaultType = valueSpec.Type
				defaultExprs = valueSpec.Values
			}
			top.incrementIota()
		}
	case token.TYPE:
		for _, decl := range node.Specs {
			c.DeclType(decl)
		}
	case token.VAR:
		for _, decl := range node.Specs {
			c.DeclVars(decl)
		}
	default:
		c.Errorf("Compile: unsupported declaration kind, expecting token.IMPORT, token.CONST, token.TYPE or token.VAR, found %v: %v %<v>",
			node.Tok, node, r.TypeOf(node))
	}
}

// DeclConsts compiles a set of constant declarations
func (c *Comp) DeclConsts(node ast.Spec, defaultType ast.Expr, defaultExprs []ast.Expr) {
	c.Pos = node.Pos()
	switch node := node.(type) {
	case *ast.ValueSpec:
		if node.Type != nil || node.Values != nil {
			defaultType = node.Type
			defaultExprs = node.Values
		}
		names, t, inits := c.prepareDeclConstsOrVars(tostrings(node.Names), defaultType, defaultExprs)
		c.DeclConsts0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported constant declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, r.TypeOf(node))
	}
}

// DeclVars compiles a set of variable declarations i.e. "var x1, x2... [type] = expr1, expr2..."
func (c *Comp) DeclVars(node ast.Spec) {
	c.Pos = node.Pos()
	switch node := node.(type) {
	case *ast.ValueSpec:
		names, t, inits := c.prepareDeclConstsOrVars(tostrings(node.Names), node.Type, node.Values)
		c.DeclVars0(names, t, inits)
	default:
		c.Errorf("Compile: unsupported variable declaration: expecting <*ast.ValueSpec>, found: %v <%v>", node, r.TypeOf(node))
	}
}

// DeclVarsShort compiles a set of variable short declarations i.e. "x1, x2... := expr1, expr2..."
func (c *Comp) DeclVarsShort(lhs []ast.Expr, rhs []ast.Expr) {
	n := len(lhs)
	names := make([]string, n)
	for i := range lhs {
		if ident, ok := lhs[i].(*ast.Ident); ok {
			names[i] = ident.Name
		} else {
			c.Errorf("non-name %v on left side of :=", lhs[i])
		}
	}
	_, t, inits := c.prepareDeclConstsOrVars(names, nil, rhs)
	c.DeclVars0(names, t, inits)
}

func tostrings(idents []*ast.Ident) []string {
	n := len(idents)
	names := make([]string, n)
	for i, ident := range idents {
		names[i] = ident.Name
	}
	return names
}

func (c *Comp) prepareDeclConstsOrVars(names []string, typ ast.Expr, exprs []ast.Expr) (names_out []string, t r.Type, inits []*Expr) {
	n := len(names)
	if typ != nil {
		t = c.Type(typ)
	}
	if exprs != nil {
		inits = c.ExprsMultipleValues(exprs, n)
	}
	return names, t, inits
}

func (c *Comp) DeclConsts0(names []string, t r.Type, inits []*Expr) {
	n := len(names)
	if inits == nil {
		c.Errorf("constants without initialization: %v", names)
	} else if len(inits) != n {
		c.Errorf("cannot declare %d constants with %d initializers: %v", n, len(inits), names)
	}
	for i, name := range names {
		init := inits[i]
		if !init.Const() {
			c.Errorf("const initializer for %q is not a constant", name)
		}
		c.DeclConst0(name, t, init.Value)
	}
}

// DeclVars0 compiles a set of variable declarations
func (c *Comp) DeclVars0(names []string, t r.Type, inits []*Expr) {
	n := len(names)
	ni := len(inits)
	if ni == 0 {
		for i := 0; i < n; i++ {
			c.DeclVar0(names[i], t, nil)
		}
	} else if ni == n {
		for i := 0; i < n; i++ {
			c.DeclVar0(names[i], t, inits[i])
		}
	} else if ni == 1 && n > 1 {
		c.DeclMultiVar0(names, t, inits[0])
	} else {
		c.Errorf("cannot declare %d variables from %d expressions: %v", n, ni, names)
	}
}

// DeclConst0 compiles a constant declaration
func (c *Comp) DeclConst0(name string, t r.Type, value I) {
	if !isLiteral(value) {
		c.Errorf("const initializer for %q is not a constant: %v <%T>", name, value, value)
		return
	}
	lit := litValue(value)
	if t != nil {
		value = lit.ConstTo(t)
	}
	bind := c.AddBind(name, ConstBind, t)
	bind.Value = value // c.Binds[] is a map[string]*Bind => changes to *Bind propagate to the map
}

// AddFuncBind reserves space for a subsequent function declaration
func (c *Comp) AddFuncBind(name string, t r.Type) *Bind {
	bind := c.AddBind(name, FuncBind, t)
	if bind.Desc.Class() != FuncBind {
		c.Errorf("internal error! Comp.AddBind(name=%q, class=FuncBind, type=%v) returned class=%v, expecting FuncBind",
			name, t, bind.Desc.Class())
	}
	return bind
}

// AddBind reserves space for a subsequent constant, function or variable declaration
func (c *Comp) AddBind(name string, class BindClass, t r.Type) *Bind {
	if class == IntBind || class == VarBind {
		if IsCategory(t.Kind(), r.Bool, r.Int, r.Uint, r.Float64) || t.Kind() == r.Complex64 {
			class = IntBind
		} else {
			class = VarBind
		}
	}
	var index = NoIndex
	if name == "_" {
		// never store bindings for "_" in c.Binds
		desc := MakeBindDescriptor(class, index)
		bind := &Bind{Lit: Lit{Type: t}, Desc: desc, Name: name}
		return bind
	}
	if c.Binds == nil {
		c.Binds = make(map[string]*Bind)
	}
	if len(name) == 0 {
		// unnamed function result
	} else if bind := c.Binds[name]; bind != nil {
		c.Warnf("redefined identifier: %v", name)
		oldclass := bind.Desc.Class()
		if (oldclass == IntBind) == (class == IntBind) {
			// both are IntBind, or neither is.
			// we can reuse the bind index
			index = bind.Desc.Index()
		}
	}
	// allocate a slot either in Binds or in IntBinds
	switch class {
	case ConstBind:
		index = NoIndex
	default: // case FuncBind, VarBind:
		if index == NoIndex {
			if c.BindNum == NoIndex {
				c.BindNum++
			}
			index = c.BindNum
			c.BindNum++
		}
	case IntBind:
		if index == NoIndex {
			if c.IntBindNum == NoIndex {
				c.IntBindNum++
			}
			index = c.IntBindNum
			c.IntBindNum++
		}
	}
	desc := MakeBindDescriptor(class, index)
	bind := &Bind{Lit: Lit{Type: t}, Desc: desc, Name: name}
	if len(name) != 0 {
		// skip unnamed function results
		c.Binds[name] = bind
	}
	return bind
}

// DeclVar0 compiles a variable declaration. For caller's convenience, returns allocated Bind
func (c *Comp) DeclVar0(name string, t r.Type, init *Expr) *Bind {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare : %v", name)
		}
		t = init.DefaultType()
		if t == nil {
			c.Errorf("cannot declare variable as untyped nil: %v", name)
		}
		n := init.NumOut()
		if n == 0 {
			c.Errorf("initializer returns no values, cannot declare variable: %v", name)
		} else if n > 1 {
			c.Errorf("initializer returns %d values, using only the first one to declare variable: %v", name)
		}
	}
	bind := c.AddBind(name, VarBind, t)
	desc := bind.Desc
	switch desc.Class() {
	default:
		c.Errorf("internal error! Comp.AddBind(name=%q, class=VarBind, type=%v) returned class=%v, expecting VarBind or IntBind",
			name, t, desc.Class())
		return bind
	case IntBind:
		// no difference between declaration and assignment for these classes
		if init == nil {
			// no initializer... use the zero-value of t
			init = exprValue(r.Zero(t).Interface())
		}
		va := bind.AsVar(0, PlaceSettable)
		c.SetVar(va, token.ASSIGN, init)
	case VarBind:
		index := desc.Index()
		if index == NoIndex {
			// assigning a constant or expression to _
			// only keep the expression side effects
			c.Code.Append(init.AsStmt())
		}
		// declaring a variable in Env.Binds[], we must create a settable and addressable reflect.Value
		if init == nil {
			// no initializer... use the zero-value of t
			c.Code.Append(func(env *Env) (Stmt, *Env) {
				env.Binds[index] = r.New(t).Elem()
				env.IP++
				return env.Code[env.IP], env
			})
			return bind
		}
		if init.Const() {
			init.ConstTo(t) // convert untyped constants, check typed constants
		}
		fun := init.AsX1() // AsX1() panics if init.NumOut() == 0, warns if init.NumOut() > 1
		tfun := init.Out(0)
		if tfun != t && !tfun.AssignableTo(t) {
			c.Errorf("cannot assign <%v> to <%v> in variable declaration: %v <%v>", tfun, t, name, t)
			return bind
		}
		var ret func(env *Env) (Stmt, *Env)
		// optimization: no need to wrap multiple-valued function into a single-value function
		if f, ok := init.Fun.(func(*Env) (r.Value, []r.Value)); ok {
			ret = func(env *Env) (Stmt, *Env) {
				ret, _ := f(env)
				place := r.New(t).Elem()
				place.Set(ret.Convert(t))
				env.Binds[index] = place
				env.IP++
				return env.Code[env.IP], env
			}
		} else {
			ret = func(env *Env) (Stmt, *Env) {
				ret := fun(env)
				place := r.New(t).Elem()
				place.Set(ret.Convert(t))
				env.Binds[index] = place
				env.IP++
				return env.Code[env.IP], env
			}
		}
		c.Code.Append(ret)
	}
	return bind
}

// DeclBindRuntimeValue compiles a variable, function or constant declaration with a reflect.Value passed at runtime
func (c *Comp) DeclBindRuntimeValue(bind *Bind) func(*Env, r.Value) {
	desc := bind.Desc
	index := desc.Index()
	if index == NoIndex {
		return nil
	}
	t := bind.Type
	switch desc.Class() {
	default:
		c.Errorf("cannot declare a %s with a value passed at runtime: %v <%v>", desc.Class(), bind.Name, t)
		return nil
	case FuncBind:
		// declaring a function in Env.Binds[], the reflect.Value must not be addressable or settable
		return func(env *Env, v r.Value) {
			env.Binds[index] = v.Convert(t)
		}
	case VarBind:
		// declaring a variable in Env.Binds[], we must create a settable and addressable reflect.Value
		return func(env *Env, v r.Value) {
			place := r.New(t).Elem()
			place.Set(v.Convert(t))
			env.Binds[index] = place
		}
	case IntBind:
		// no difference between declaration and assignment for IntBind
		return c.SetVarValue(bind.AsVar(0, PlaceSettable))
	}
}

// DeclMultiVar0 compiles multiple variable declarations from a single multi-valued expression
func (c *Comp) DeclMultiVar0(names []string, t r.Type, init *Expr) {
	if t == nil {
		if init == nil {
			c.Errorf("no value and no type, cannot declare variables: %v", names)
		}
	}
	n := len(names)
	if n == 1 {
		c.DeclVar0(names[0], t, init)
		return
	}
	ni := init.NumOut()
	if ni < n {
		c.Errorf("cannot declare %d variables from expression returning %d values: %v", n, ni, names)
	} else if ni > n {
		c.Warnf("declaring %d variables from expression returning %d values: %v", n, ni, names)
	}
	decls := make([]func(*Env, r.Value), n)
	for i, name := range names {
		ti := init.Out(i)
		if t != nil && t != ti {
			if !ti.AssignableTo(t) {
				c.Errorf("cannot assign <%v> to <%v> in variable declaration: %v", ti, t, names)
				return
			} else {
				ti = t // declared variable has type t, not the i-th type returned by multi-valued expression
			}
		}
		bind := c.AddBind(name, VarBind, ti)
		decls[i] = c.DeclBindRuntimeValue(bind)
	}
	fun := init.AsXV(0)
	c.Code.Append(func(env *Env) (Stmt, *Env) {
		// call the multi-valued function. we know ni > 1, so just use the []r.Value
		_, rets := fun(env)

		// declare and assign the variables one by one. we know n <= ni
		for i, decl := range decls {
			if decl != nil {
				decl(env, rets[i])
			}
		}
		env.IP++
		return env.Code[env.IP], env
	})
}

// DeclFunc0 compiles a function declaration. For caller's convenience, returns allocated Bind
func (c *Comp) DeclFunc0(name string, fun I) *Bind {
	funv := r.ValueOf(fun)
	t := funv.Type()
	if t.Kind() != r.Func {
		c.Errorf("DeclFunc0(%s): expecting a function, received %v <%v>", name, fun, t)
	}
	bind := c.AddFuncBind(name, t)
	index := bind.Desc.Index()
	ret := func(env *Env) (Stmt, *Env) {
		env.Binds[index] = funv
		env.IP++
		return env.Code[env.IP], env
	}
	c.Code.Append(ret)
	return bind
}

// DeclBuiltinFunc compiles a builtin function declaration. For caller's convenience, returns allocated Bind
func (c *Comp) DeclBuiltinFunc0(name string, builtin BuiltinFunc) *Bind {
	t := TypeOfBuiltinFunc
	bind := c.AddBind(name, ConstBind, t) // not a regular function... its type is not accurate
	bind.Value = builtin                  // c.Binds[] is a map[string]*Bind => changes to *Bind propagate to the map
	return bind
}
