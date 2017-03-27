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
 * main.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	mp "github.com/cosmos72/gomacro/parser"
)

type Cmd struct {
	*Env
}

func (cmd *Cmd) Init() {
	cmd.Env = New()

	cmd.ParserMode = mp.Trace & 0
	cmd.Options = OptTrapPanic | OptShowPrompt | OptShowAfterEval // | OptShowAfterMacroExpansion // | OptDebugMacroExpand // |  OptDebugQuasiquote  // | OptShowEvalDuration // | OptShowAfterParse
}

type tristate int

const (
	tNochange tristate = iota
	tClear
	tSet
)

type UserOptions struct {
	verbose, collectDecls, collectStatements tristate
}

func (cmd *Cmd) Main(args []string) (err error) {
	if cmd.Env == nil {
		cmd.Init()
	}
	env := cmd.Env

	if len(args) == 0 {
		env.ReplStdin()
		return nil
	}
	var opts UserOptions

	for len(args) > 0 {
		switch args[0] {
		case "-i":
			env.Options |= OptShowPrompt | OptShowAfterEval
			env.Options = applyOptions(env.Options, opts)
			env.ReplStdin()
			args = args[1:]
		case "-e":
			env.Options |= OptShowAfterEval
			env.Options = applyOptions(env.Options, opts)

			buf := bytes.NewBufferString(strings.Join(args[1:], " "))
			buf.WriteByte('\n') // because ReadMultiLine() needs a final '\n'
			return cmd.EvalReader(buf)
		case "-h":
			return cmd.Usage()
		case "-o":
			args = args[1:]
			if len(args) > 0 {
				for _, str := range strings.Split(args[0], ",") {
					switch str {
					case "decl":
						opts.collectDecls = tSet
					case "^decl":
						opts.collectDecls = tClear
					case "stmt":
						opts.collectStatements = tSet
					case "^stmt":
						opts.collectStatements = tClear
					case "verbose":
						opts.verbose = tSet
					case "^verbose":
						opts.verbose = tClear
					}
				}
				args = args[1:]
			}
		case "-q":
			opts.verbose = tClear
			args = args[1:]
		case "-v":
			opts.verbose = tSet
			args = args[1:]
		default:
			env.Options &^= OptShowPrompt | OptShowAfterEval
			env.Options = applyOptions(env.Options, opts)
			return cmd.EvalFilesAndDirs(args...)
		}
	}
	env.Options |= OptShowPrompt | OptShowAfterEval
	env.Options = applyOptions(env.Options, opts)
	env.ReplStdin()
	return nil
}

func applyOptions(opts Options, user UserOptions) Options {
	switch user.verbose {
	case tClear:
		opts &^= OptShowAfterEval
	case tSet:
		opts |= OptShowAfterEval
	}
	switch user.collectDecls {
	case tClear:
		opts &^= OptCollectDeclarations
	case tSet:
		opts |= OptCollectDeclarations
	}
	switch user.collectStatements {
	case tClear:
		opts &^= OptCollectStatements
	case tSet:
		opts |= OptCollectStatements
	}
	return opts
}

func (cmd *Cmd) Usage() error {
	fmt.Print(`usage: gomacro [OPTIONS] files-and-dirs
       gomacro [OPTIONS] -e expressions

       Recognized options:
       -h      show this help and exit
       -i      interactive. immediately start a REPL, even in the middle
	                        of executing files and dirs
       -v      verbose. show startup message, prompt, and expressions result
       -q      quiet. do NOT show startup message, prompt, and expressions result
       -o LIST

        LIST is a comma-separated list of one or more:
         decl      collect declarations
         ^decl     do NOT collect declarations
         stmt      collect statements
         ^stmt     do NOT collect statements
         verbose   same as -v
         ^verbose  same as -q
       collected declarations and statements can be written to standard output
       or to a file with the REPL command :write
`)
	return nil
}

func (cmd *Cmd) EvalFilesAndDirs(filesAndDirs ...string) error {
	for _, fileOrDir := range filesAndDirs {
		err := cmd.EvalFileOrDir(fileOrDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cmd *Cmd) EvalFileOrDir(fileOrDir string) error {
	info, err := os.Stat(fileOrDir)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return cmd.EvalDir(fileOrDir)
	} else {
		return cmd.EvalFile(fileOrDir)
	}
}

func (cmd *Cmd) EvalDir(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		l := len(filename)
		if l > 3 && filename[l-3:] == ".go" || l > 8 && filename[l-8:] == ".gomacro" {
			err := cmd.EvalFile(filename)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cmd *Cmd) EvalFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()
	err = cmd.EvalReader(f)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *Cmd) EvalReader(src io.Reader) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = errors.New(fmt.Sprint(rec))
			}
		}
	}()
	in := bufio.NewReader(src)

	for cmd.Env.ReadParseEvalPrint(in) {
	}
	return nil
}
