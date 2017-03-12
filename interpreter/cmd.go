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
	"sort"
	"strings"

	mp "github.com/cosmos72/gomacro/parser"
)

type Cmd struct {
	*Env
}

func (cmd *Cmd) Init() {
	cmd.Env = New()

	cmd.ParserMode = mp.Trace & 0                                 // | mp.TraceMacro
	cmd.Options = OptTrapPanic | OptShowPrompt | OptShowAfterEval // | OptShowAfterMacroExpansion // | OptDebugMacroExpand // |  OptDebugQuasiquote  // | OptShowEvalDuration // | OptShowAfterParse
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
	silent := false
	verbose := false

	env.Options &^= OptShowPrompt | OptShowAfterEval

	for len(args) > 0 {
		switch args[0] {
		case "-h":
			return cmd.Usage()
		case "-f":
			env.Options = applyOptions(env.Options, silent, verbose)
			return cmd.EvalFiles(args[1:]...)
		case "-d":
			env.Options = applyOptions(env.Options, silent, verbose)
			return cmd.EvalDirs(args[1:]...)
		case "-s":
			silent = true
			verbose = false
			args = args[1:]
			continue
		case "-v":
			verbose = true
			silent = false
			args = args[1:]
			continue
		case ".":
			return cmd.EvalDirs(".")
		default:
			env.Options |= OptShowAfterEval
			env.Options = applyOptions(env.Options, silent, verbose)

			buf := bytes.NewBufferString(strings.Join(args, " "))
			buf.WriteByte('\n') // because ReadMultiLine() needs a final '\n'
			return cmd.EvalReader(buf)
		}
		break
	}
	return nil
}

func applyOptions(opts Options, silent bool, verbose bool) Options {
	if silent {
		opts &^= OptShowAfterEval
	} else if verbose {
		opts |= OptShowAfterEval
	}
	return opts
}

func (cmd *Cmd) Usage() error {
	fmt.Print(`usage: gomacro [expression]
       gomacro -f [filenames...]
       gomacro -d [dirnames...]
       gomacro .
`)
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

func (cmd *Cmd) EvalFiles(filenames ...string) (err error) {
	for _, filename := range filenames {
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
	}
	return nil
}

func (cmd *Cmd) EvalDirs(dirnames ...string) error {
	for _, dirname := range dirnames {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			return err
		}
		filenames := make([]string, 0, len(files))
		for _, file := range files {
			filename := file.Name()
			l := len(filename)
			if l > 3 && filename[l-3:] == ".go" || l > 8 && filename[l-8:] == ".gomacro" {
				filenames = append(filenames, filename)
			}
		}
		sort.Strings(filenames)

		for _, filename := range filenames {
			err := cmd.EvalFiles(filename)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
