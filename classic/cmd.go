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
	Func func(ir *Interp, arg string, opt CmdOpt) (string, CmdOpt)
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
	'e': Cmd{"env", (*Interp).cmdEnv},
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
func (ir *Interp) Cmd(src string) (string, CmdOpt) {
	g := ir.Env.Globals
	var opt CmdOpt

	src = strings.TrimSpace(src)
	n := len(src)
	if n > 0 && src[0] == g.ReplCmdChar {
		prefix, arg := Split2(src[1:], ' ') // skip g.ReplCmdChar
		cmd, found := cmds.Lookup(prefix)
		if found {
			src, opt = cmd.Func(ir, arg, opt)
		} else {
			// ":<something>"
			// temporarily disable collection of declarations and statements,
			// and temporarily disable macroexpandonly (i.e. re-enable eval)
			opt |= CmdOptForceEval
			src = " " + src[1:] // slower than src = src[1:], but gives accurate column positions in error messages
		}
	}
	// :package and package are the same command
	if g.Options&OptMacroExpandOnly == 0 && (src == "package" || strings.HasPrefix(src, "package ")) {
		_, arg := Split2(src, ' ')
		src, opt = ir.cmdPackage(arg, opt)
	}
	return src, opt
}

func (ir *Interp) cmdEnv(arg string, opt CmdOpt) (string, CmdOpt) {
	ir.Env.ShowPackage(arg)
	return "", opt
}

func (ir *Interp) cmdHelp(arg string, opt CmdOpt) (string, CmdOpt) {
	g := ir.Env.ThreadGlobals.Globals
	g.ShowHelp()
	return "", opt
}

func (ir *Interp) cmdInspect(arg string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	if len(arg) == 0 {
		fmt.Fprint(env.Stdout, "// inspect: missing argument\n")
	} else {
		env.Inspect(arg)
	}
	return "", opt
}

func (ir *Interp) cmdOptions(arg string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	g := env.Globals

	if len(arg) != 0 {
		g.Options ^= ParseOptions(arg)
	} else {
		fmt.Fprintf(env.Stdout, "// current options: %v\n", g.Options)
		fmt.Fprintf(env.Stdout, "// unset   options: %v\n", ^g.Options)
	}
	return "", opt
}

// change package. path can be empty or a package path WITH quotes
// 'package NAME' where NAME is without quotes has no effect.
func (ir *Interp) cmdPackage(path string, cmdopt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	g := env.Globals
	path = strings.TrimSpace(path)
	n := len(path)
	if n == 0 {
		g.Fprintf(g.Stdout, "// current package: %s %q\n", env.Name, env.Path)
	} else if n > 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
		n -= 2
		ir.ChangePackage(path)
	} else if g.Options&OptShowPrompt != 0 {
		g.Debugf(`package %s has no effect. To switch to a different package, use package "PACKAGE/FULL/PATH" - note the quotes`, path)
	}
	return "", cmdopt
}

func (ir *Interp) cmdQuit(_ string, opt CmdOpt) (string, CmdOpt) {
	return "", opt | CmdOptQuit
}

// remove package 'path' from the list of known packages
func (ir *Interp) cmdUnload(path string, opt CmdOpt) (string, CmdOpt) {
	if n := len(path); n >= 2 && path[0] == '"' && path[n-1] == '"' {
		path = path[1 : n-1]
	}
	if len(path) != 0 {
		ir.Env.Globals.UnloadPackage(path)
	}
	return "", opt
}

func (ir *Interp) cmdWrite(filepath string, opt CmdOpt) (string, CmdOpt) {
	env := ir.Env
	if len(filepath) == 0 {
		env.WriteDeclsToStream(env.Stdout)
	} else {
		env.WriteDeclsToFile(filepath)
	}
	return "", opt
}
