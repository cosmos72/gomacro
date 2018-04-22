/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
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
 * debugger.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"runtime/debug"

	"github.com/cosmos72/gomacro/base"
)

func (d *Debugger) Help() {
	g := d.globals
	g.Fprintf(g.Stdout, "%s", `// debugger commands:
env [NAME]    show available functions, variables and constants
              in current scope, or from imported package NAME
continue      resume normal execution
?             show this help
help          show this help
inspect EXPR  inspect expression interactively
next          execute a single statement, skipping functions
print   EXPR  print expression, statement or declaration
// abbreviations are allowed if unambiguous.
`)
	/*
		not implemented yet:

		backtrace [N] show function stack frames
		finish        run until the end of current function
		step          execute a single statement, entering functions
	*/
}

func (d *Debugger) Show(breakpoint bool) {
	// d.env is the Env being debugged.
	// to execute code at debugger prompt, use d.interp
	env := d.env
	pos := env.DebugPos
	g := d.globals
	ip := env.IP

	var label string
	if breakpoint {
		label = "breakpoint"
	} else {
		label = "stopped"
	}
	if ip < len(pos) && g.Fileset != nil {
		source, pos := g.Fileset.Source(pos[ip])
		g.Fprintf(g.Stdout, "// %s at %s IP=%d. type ? for debugger help\n", label, pos, ip)
		if len(source) != 0 {
			g.Fprintf(g.Stdout, "%s\n", source)
			d.showCaret(source, pos.Column)
		}
	} else {
		g.Fprintf(g.Stdout, "// %s at IP=%d. type ? for debugger help\n", label, ip)
	}
}

var spaces = []byte("                                                                      ")

func (d *Debugger) showCaret(source string, col int) {
	col--
	n := len(source)
	if col >= 0 && col < n && n >= 5 {
		out := d.globals.Stdout
		chunk := len(spaces)
		for col >= chunk {
			out.Write(spaces)
			col -= chunk
		}
		out.Write(spaces[:col])
		out.Write([]byte("^^^\n"))
	}
}

func (d *Debugger) Repl() DebugOp {
	g := d.globals
	var opts base.ReadOptions
	if g.Options&base.OptShowPrompt != 0 {
		opts |= base.ReadOptShowPrompt
	}

	op := DebugRepl
	for op == DebugRepl {
		src, firstToken := g.ReadMultiline(opts, "debug> ")
		empty := len(src) == 0
		if firstToken < 0 && empty {
			// EOF
			op = DebugContinue
			break
		}
		if empty {
			// keyboard enter repeats last command
			src = d.lastcmd
		}
		op = d.Eval(src)
	}
	return op
}

func (d *Debugger) Eval(src string) DebugOp {
	return d.Cmd(src)
}

func (d *Debugger) Print(src string) {
	g := d.globals
	trap := g.Options&base.OptTrapPanic != 0

	// do NOT debug expression evaluated at debugger prompt!
	sig := &d.env.ThreadGlobals.Signals
	sigdebug := sig.Debug
	sig.Debug = base.SigNone

	defer func() {
		sig.Debug = sigdebug
		if trap {
			rec := recover()
			if g.Options&base.OptPanicStackTrace != 0 {
				g.Fprintf(g.Stderr, "%v\n%s", rec, debug.Stack())
			} else {
				g.Fprintf(g.Stderr, "%v\n", rec)
			}
		}
	}()

	ir := d.interp
	expr := ir.Compile(src)
	values := base.PackValues(ir.RunExpr(expr))
	types := base.PackTypes(expr.Type, expr.Types)

	g.Print(values, types)

	trap = false // no panic happened
}
