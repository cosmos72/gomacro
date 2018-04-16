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
 * cmd.go
 *
 *  Created on: Apr 11, 2018
 *      Author: Massimiliano Ghilardi
 */

package classic

import (
	"fmt"
	"strings"

	. "github.com/cosmos72/gomacro/base"
)

type Cmd struct {
	Name string
	Func func(ir *Interp, arg string, opt CmdOpt, in Readline) (string, CmdOpt)
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
	'c': Cmd{"classic", (*Interp).cmdClassic},
	'e': Cmd{"env", (*Interp).cmdEnv},
	'f': Cmd{"fast", (*Interp).cmdFast},
	'h': Cmd{"help", (*Interp).cmdHelp},
	'i': Cmd{"inspect", (*Interp).cmdInspect},
	'o': Cmd{"options", (*Interp).cmdOptions},
	'p': Cmd{"package", (*Interp).cmdPackage},
	'q': Cmd{"quit", (*Interp).cmdQuit},
	'u': Cmd{"unload", (*Interp).cmdUnload},
	'w': Cmd{"write", (*Interp).cmdWrite},
}

// execute one of the REPL commands starting with ':'
// return any remainder string to be evaluated, and the options to evaluate it
func (ir *Interp) Cmd(src string, in Readline) (string, CmdOpt) {
	g := ir.Env.Globals
	var opt CmdOpt
	if g.Options&OptFastInterpreter != 0 {
		opt |= CmdOptFast
	}

	src = strings.TrimSpace(src)
	n := len(src)
	if n > 0 && src[0] == g.ReplCmdChar {
		prefix, arg := split2(src[1:], ' ') // skip g.ReplCmdChar
		cmd, found := cmds.Lookup(prefix)
		if found {
			src, opt = cmd.Func(ir, arg, opt, in)
		} else {
			// ":<something>"
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			opt |= CmdOptForceEval
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		}
	}
	// :package and package are the same command
	if src == "package" || strings.HasPrefix(src, "package ") {
		_, arg := split2(src, ' ')
		src, opt = ir.cmdPackage(arg, opt, in)
	}
	return src, opt
}

// split 's' into a prefix and suffix separated by 'separator'.
// suffix is trimmed with strings.TrimSpace() before returning it
func split2(s string, separator rune) (string, string) {
	var prefix, suffix string
	if space := strings.IndexByte(s, ' '); space > 0 {
		prefix = s[:space]
		suffix = strings.TrimSpace(s[space+1:])
	} else {
		prefix = s
	}
	return prefix, suffix
}

func (ir *Interp) cmdClassic(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	g := ir.Env.Globals
	if len(arg) == 0 {
		if g.Options&OptFastInterpreter != 0 {
			g.Debugf("switched to classic interpreter")
		}
		g.Options &^= OptFastInterpreter
		arg = ""
	} else {
		// temporary override, do not change g.Options
		arg = strings.TrimSpace(arg)
	}
	opt &^= CmdOptFast
	return arg, opt
}

func (ir *Interp) cmdEnv(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	if opt&CmdOptFast != 0 {
		ir.fastShowPackage(arg)
	} else {
		ir.Env.ShowPackage(arg)
	}
	return "", opt
}

func (ir *Interp) cmdFast(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	g := ir.Env.Globals
	if len(arg) == 0 {
		if g.Options&OptFastInterpreter == 0 {
			g.Debugf("switched to fast interpreter")
		}
		g.Options |= OptFastInterpreter
		arg = ""
	} else {
		// temporary override, do not change g.Options
		arg = strings.TrimSpace(arg)
	}
	opt |= CmdOptFast
	return arg, opt
}

func (ir *Interp) cmdHelp(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	g := ir.Env.ThreadGlobals.Globals
	g.ShowHelp()
	return "", opt
}

func (ir *Interp) cmdInspect(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	env := ir.Env
	if len(arg) == 0 {
		fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
	} else {
		env.Inspect(in, arg, opt&CmdOptFast != 0)
	}
	return "", opt
}

func (ir *Interp) cmdOptions(arg string, opt CmdOpt, in Readline) (string, CmdOpt) {
	env := ir.Env
	g := env.Globals

	if len(arg) != 0 {
		g.Options ^= ParseOptions(arg)
		env.fastUpdateOptions() // to set fastInterp.Comp.CompGlobals.Universe.DebugDepth

		if g.Options&OptFastInterpreter != 0 {
			opt |= CmdOptFast
		} else {
			opt &^= CmdOptFast
		}
	} else {
		fmt.Fprintf(env.Stdout, "// current options: %v\n", g.Options)
		fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^g.Options)
	}
	return "", opt
}

// change package. path can be empty or a package path with or without quotes
func (ir *Interp) cmdPackage(path string, cmdopt CmdOpt, in Readline) (string, CmdOpt) {
	env := ir.Env
	path = strings.TrimSpace(path)
	n := len(path)
	if n >= 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
		n -= 2
	}
	if cmdopt&CmdOptFast != 0 {
		ir.fastCmdPackage(path)
	} else if n == 0 {
		fmt.Fprintf(env.Stdout, "// current package: %s %q\n", env.Name, env.Path)
	} else {
		ir.ChangePackage(path)
	}
	return "", cmdopt
}

func (ir *Interp) cmdQuit(_ string, opt CmdOpt, in Readline) (string, CmdOpt) {
	return "", opt | CmdOptQuit
}

// remove package 'path' from the list of known packages
func (ir *Interp) cmdUnload(path string, opt CmdOpt, in Readline) (string, CmdOpt) {
	if len(path) != 0 {
		if opt&CmdOptFast != 0 {
			ir.fastUnloadPackage(path)
		} else {
			ir.Env.Globals.UnloadPackage(path)
		}
	}
	return "", opt
}

func (ir *Interp) cmdWrite(filepath string, opt CmdOpt, in Readline) (string, CmdOpt) {
	env := ir.Env
	if len(filepath) == 0 {
		env.WriteDeclsToStream(env.Stdout)
	} else {
		env.WriteDeclsToFile(filepath)
	}
	return "", opt
}
