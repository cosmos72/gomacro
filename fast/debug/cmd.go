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
 * global.go
 *
 *  Created on Apr 21, 2018
 *      Author Massimiliano Ghilardi
 */

package debug

import (
	"strings"

	"github.com/cosmos72/gomacro/base"
)

type Cmd struct {
	Name string
	Func func(d *Debugger, arg string) DebugOp
}

type Cmds map[byte]Cmd

func (cmd *Cmd) Match(prefix string) bool {
	return strings.HasPrefix(cmd.Name, prefix)
}

func (cmds Cmds) Lookup(prefix string) (Cmd, bool) {
	if len(prefix) != 0 {
		cmd, found := cmds[prefix[0]]
		if found && cmd.Match(prefix) {
			return cmd, true
		}
	}
	return Cmd{}, false
}

var cmds = Cmds{
	'b': Cmd{"backtrace", (*Debugger).cmdBacktrace},
	'c': Cmd{"continue", (*Debugger).cmdContinue},
	'e': Cmd{"env", (*Debugger).cmdEnv},
	'f': Cmd{"finish", (*Debugger).cmdFinish},
	'h': Cmd{"help", (*Debugger).cmdHelp},
	'?': Cmd{"?", (*Debugger).cmdHelp},
	'i': Cmd{"inspect", (*Debugger).cmdInspect},
	'l': Cmd{"list", (*Debugger).cmdList},
	'n': Cmd{"next", (*Debugger).cmdNext},
	'p': Cmd{"print", (*Debugger).cmdPrint},
	's': Cmd{"step", (*Debugger).cmdStep},
}

// execute one of the debugger commands
func (d *Debugger) Cmd(src string) DebugOp {
	op := DebugRepl
	src = strings.TrimSpace(src)
	n := len(src)
	if n > 0 {
		prefix, arg := base.Split2(src, ' ')
		cmd, found := cmds.Lookup(prefix)
		if found {
			d.lastcmd = src
			op = cmd.Func(d, arg)
		} else {
			d.Help()
		}
	}
	return op
}

func (d *Debugger) cmdBacktrace(arg string) DebugOp {
	return DebugRepl
}

func (d *Debugger) cmdContinue(arg string) DebugOp {
	return DebugContinue
}

func (d *Debugger) cmdEnv(arg string) DebugOp {
	d.interp.ShowPackage(arg)
	return DebugRepl
}

func (d *Debugger) cmdFinish(arg string) DebugOp {
	return DebugFinish
}

func (d *Debugger) cmdHelp(arg string) DebugOp {
	d.Help()
	return DebugRepl
}

func (d *Debugger) cmdInspect(arg string) DebugOp {
	if len(arg) == 0 {
		g := d.globals
		g.Fprintf(g.Stdout, "// inspect: missing argument\n")
	} else {
		d.interp.Inspect(arg)
	}
	return DebugRepl
}

func (d *Debugger) cmdList(arg string) DebugOp {
	d.Show(false)
	return DebugRepl
}

func (d *Debugger) cmdNext(arg string) DebugOp {
	return DebugNext
}

func (d *Debugger) cmdPrint(arg string) DebugOp {
	if len(arg) == 0 {
		g := d.globals
		g.Fprintf(g.Stdout, "// print: missing argument\n")
	} else {
		d.Eval(arg)
	}
	return DebugRepl
}

func (d *Debugger) cmdStep(arg string) DebugOp {
	return DebugStep
}
