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
 * interp.go
 *
 *  Created on: Jun 15, 2017
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"bufio"
	"fmt"
	"os"
	r "reflect"
	"runtime/debug"
	"strings"
	"time"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

type Interp struct {
	*Env
	currOpt CmdOpt
}

func New() *Interp {
	top := NewEnv(nil, "builtin")
	env := NewEnv(top, "main")
	return &Interp{Env: env}
}

func (ir *Interp) ChangePackage(path string) {
	ir.Env = ir.Env.ChangePackage(path)
}

var historyfile = Subdir(UserHomeDir(), ".gomacro_history")

func (ir *Interp) ReplStdin() {
	if ir.Options&OptShowPrompt != 0 {
		fmt.Fprint(ir.Stdout, `// GOMACRO, an interactive Go interpreter with macros <https://github.com/cosmos72/gomacro>
// Copyright (C) 2017-2018 Massimiliano Ghilardi
// License LGPL v3+: GNU Lesser GPL version 3 or later <https://gnu.org/licenses/lgpl>
// This is free software with ABSOLUTELY NO WARRANTY.
//
// Type :help for help
`)
	}
	tty, _ := MakeTtyReadline(historyfile)
	defer tty.Close(historyfile) // restore normal tty mode

	c := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(c)

	ir.Line = 0
	for ir.ReadParseEvalPrint(tty) {
		ir.Line = 0
	}
	os.Stdout.WriteString("\n")
}

func (ir *Interp) Repl(in *bufio.Reader) {
	r := MakeBufReadline(in, ir.Stdout)

	c := StartSignalHandler(ir.Interrupt)
	defer StopSignalHandler(c)

	for ir.ReadParseEvalPrint(r) {
	}
}

func (ir *Interp) ReadParseEvalPrint(in Readline) (callAgain bool) {
	str, firstToken := ir.Read(in)
	if firstToken < 0 {
		// skip comment-only lines and continue, but fail on EOF or other errors
		return len(str) != 0
	}
	return ir.ParseEvalPrint(str[firstToken:], in)
}

// return read string and position of first non-comment token.
// return "", -1 on EOF
func (ir *Interp) Read(in Readline) (string, int) {
	var opts ReadOptions
	if ir.Options&OptShowPrompt != 0 {
		opts |= ReadOptShowPrompt
	}
	str, firstToken := ir.Env.Globals.ReadMultiline(in, opts)
	if firstToken < 0 {
		ir.IncLine(str)
	} else if firstToken > 0 {
		ir.IncLine(str[0:firstToken])
	}
	return str, firstToken
}

func (ir *Interp) ParseEvalPrint(str string, in Readline) (callAgain bool) {
	var t1 time.Time
	trap := ir.Options&OptTrapPanic != 0
	duration := ir.Options&OptShowTime != 0
	if duration {
		t1 = time.Now()
	}
	defer func() {
		ir.IncLine(str)
		if trap {
			rec := recover()
			if ir.Options&OptPanicStackTrace != 0 {
				fmt.Fprintf(ir.Stderr, "%v\n%s", rec, debug.Stack())
			} else {
				fmt.Fprintf(ir.Stderr, "%v\n", rec)
			}
			callAgain = true
		}
		if duration {
			delta := time.Since(t1)
			ir.Debugf("eval time %v", delta)
		}
	}()
	callAgain = ir.parseEvalPrint(str, in)
	trap = false // no panic happened
	return callAgain
}

func (ir *Interp) parseEvalPrint(src string, in Readline) (callAgain bool) {
	if len(strings.TrimSpace(src)) == 0 {
		return true // no input. don't print anything
	}
	env := ir.Env
	g := env.Globals

	src, opt := ir.Cmd(src, in)

	callAgain = opt&CmdOptQuit == 0
	if len(src) == 0 || !callAgain {
		return callAgain
	}

	if opt&CmdOptForceEval != 0 {
		// temporarily disable collection of declarations and statements,
		// and temporarily re-enable eval (i.e. disable macroexpandonly)
		const todisable = OptMacroExpandOnly | OptCollectDeclarations | OptCollectStatements
		if g.Options&todisable != 0 {
			g.Options &^= todisable
			defer func() {
				g.Options |= todisable
			}()
		}
	}

	ir.currOpt = opt // store options where Interp.Interrupt() can find them

	// parse phase. no macroexpansion/collect yet
	form := env.ParseOnly(src)

	// macroexpand + collect + eval phase
	var values []r.Value
	var types []xr.Type
	if form != nil {
		if opt&CmdOptFast != 0 {
			values, types = env.fastEval(form)
		} else {
			values = env.classicEval(form)
		}
	}

	// print phase
	g.Print(values, types)
	return true
}

func (ir *Interp) Interrupt(sig os.Signal) {
	if ir.currOpt&CmdOptFast != 0 {
		ir.Env.fastInterrupt()
	} else {
		// TODO not implemented
	}
}
