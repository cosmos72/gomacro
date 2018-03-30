/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * output.go
 *
 *  Created on: Mar 30, 2018
 *      Author: Massimiliano Ghilardi
 */

package fast

import (
	"fmt"
	"io"
	r "reflect"
	"sort"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (b Builtin) String() string {
	return fmt.Sprintf("%p", b.Compile)
}

func (imp Import) String() string {
	return fmt.Sprintf("{%s %q, %d binds, %d types}", imp.Name, imp.Path, len(imp.Binds), len(imp.Types))
}

func (ir *Interp) ShowPackage(name string) {
	if len(name) != 0 {
		ir.ShowImportedPackage(name)
		return
	}
	// show current package
	for {
		c := ir.Comp
		env := ir.env
		ir.ShowAsPackage()
		for i := 0; i < c.UpCost && env != nil; i++ {
			env = env.Outer
		}
		c = c.Outer
		if env == nil || c == nil {
			return
		}
		ir = &Interp{c, env}
	}
}

func (ir *Interp) ShowAsPackage() {
	c := ir.Comp
	out := c.Stdout
	if binds := c.Binds; len(binds) > 0 {
		fmt.Fprintf(out, "// ----- %s binds -----\n", c.Path)

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			bind := binds[k]
			if bind == nil {
				continue
			}
			if bind.Const() {
				showValue(out, k, bind.ConstValue(), bind.Type)
				continue
			}
			expr := c.Symbol(bind.AsSymbol(0))
			showValue(out, k, ir.RunExpr1(expr), expr.Type)
		}
		fmt.Fprintln(out)
	}
	showTypes(out, c.Path, c.Types)
}

func (ir *Interp) ShowImportedPackage(name string) {
	var imp Import
	var ok bool
	if bind := ir.Comp.Binds[name]; bind != nil && bind.Const() && bind.Type != nil && bind.Type.ReflectType() == rtypeOfImport {
		imp, ok = bind.Value.(Import)
	}
	if !ok {
		ir.Comp.Warnf("not an imported package: %q", name)
		return
	}
	out := ir.Comp.Stdout
	if binds := imp.Binds; len(binds) > 0 {
		fmt.Fprintf(out, "// ----- %s binds -----\n", imp.Path)

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			showValue(out, k, binds[k], imp.BindTypes[k])
		}
		fmt.Fprintln(out)
	}
	showTypes(out, imp.Path, imp.Types)
}

func showTypes(out io.Writer, path string, types map[string]xr.Type) {
	if len(types) > 0 {
		fmt.Fprintf(out, "// ----- %s types -----\n", path)

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			t := types[k]
			if t != nil {
				showType(out, k, t)
			}
		}
		fmt.Fprintln(out)
	}
}

const spaces15 = "               "

func showValue(out io.Writer, name string, v r.Value, t xr.Type) {
	n := len(name) & 15
	if v == base.Nil || v == base.None {
		fmt.Fprintf(out, "%s%s = nil\t// %v\n", name, spaces15[n:], t)
	} else {
		fmt.Fprintf(out, "%s%s = %v\t// %v\n", name, spaces15[n:], v, t)
	}
}

func showType(out io.Writer, name string, t xr.Type) {
	n := len(name) & 15
	fmt.Fprintf(out, "%s%s %v <%v>\n", name, spaces15[n:], t.Kind(), t)
}
