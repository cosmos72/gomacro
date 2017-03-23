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

func (cmd *Cmd) Main(args []string) (err error) {
	if cmd.Env == nil {
		cmd.Init()
	}
	env := cmd.Env

	if len(args) == 0 {
		env.ReplStdin()
		return nil
	}
	quiet := false
	verbose := false

	env.Options &^= OptShowPrompt | OptShowAfterEval

	for len(args) > 0 {
		switch args[0] {
		case "-h":
			return cmd.Usage()
		case "-e":
			env.Options |= OptShowAfterEval
			env.Options = applyOptions(env.Options, quiet, verbose)

			buf := bytes.NewBufferString(strings.Join(args[1:], " "))
			buf.WriteByte('\n') // because ReadMultiLine() needs a final '\n'
			return cmd.EvalReader(buf)
		case "-q":
			quiet = true
			verbose = false
			args = args[1:]
			continue
		case "-v":
			verbose = true
			quiet = false
			args = args[1:]
			continue
		default:
			env.Options = applyOptions(env.Options, quiet, verbose)
			return cmd.EvalFilesAndDirs(args...)
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
	fmt.Print(`usage: gomacro [-q] [-v] files-and-dirs
       gomacro [-q] [-v] -e expression
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
