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
	WriteDeclsAndStmtsToFile, OverwriteFiles bool
}

func (cmd *Cmd) Init() {
	cmd.Env = New()
	cmd.ParserMode = mp.Trace & 0
	cmd.Options = OptTrapPanic | OptShowPrompt | OptShowEval // | OptShowAfterMacroExpansion // | OptDebugMacroExpand // |  OptDebugQuasiquote  // | OptShowEvalDuration // | OptShowAfterParse
	cmd.WriteDeclsAndStmtsToFile = false
	cmd.OverwriteFiles = false
}

func (cmd *Cmd) Main(args []string) (err error) {
	if cmd.Env == nil {
		cmd.Init()
	}
	env := cmd.Env

	var set, clear Options
	repl := len(args) == 0
	cmd.WriteDeclsAndStmtsToFile = false
	cmd.OverwriteFiles = false

	for len(args) > 0 {
		switch args[0] {
		case "-e":
			if len(args) > 1 {
				buf := bytes.NewBufferString(args[1])
				buf.WriteByte('\n') // because ReadMultiLine() needs a final '\n'
				env.Options |= OptShowEval
				env.Options = (env.Options | set) &^ clear
				err := cmd.EvalReader(buf)
				if err != nil {
					return err
				}
				args = args[1:]
			}
		case "-f":
			cmd.OverwriteFiles = true
		case "-h":
			return cmd.Usage()
		case "-i":
			repl = true
		case "-o":
			if len(args) > 1 {
				for _, str := range strings.Split(args[1], ",") {
					switch str {
					case "decl":
						set |= OptCollectDeclarations
						clear &^= OptCollectDeclarations
					case "^decl":
						set &^= OptCollectDeclarations
						clear |= OptCollectDeclarations
					case "stmt":
						set |= OptCollectStatements
						clear &^= OptCollectStatements
					case "^stmt":
						set &^= OptCollectStatements
						clear |= OptCollectStatements
					case "verbose":
						set |= OptShowEval
						clear &^= OptShowEval
					case "^verbose":
						set &^= OptShowEval
						clear |= OptShowEval
					}
				}
				args = args[1:]
			}
		case "-s":
			set &^= OptShowEval
			clear |= OptShowEval
		case "-v":
			set |= OptShowEval
			clear &^= OptShowEval
		case "-w":
			cmd.WriteDeclsAndStmtsToFile = true
		default:
			if cmd.WriteDeclsAndStmtsToFile {
				env.Options |= OptCollectDeclarations | OptCollectStatements
			}
			env.Options &^= OptShowPrompt | OptShowEval
			env.Options = (env.Options | set) &^ clear
			cmd.EvalFileOrDir(args[0])

			env.Imports, env.Declarations, env.Statements = nil, nil, nil
		}
		args = args[1:]
	}
	if repl {
		env.Options |= OptShowPrompt | OptShowEval
		env.Options = (env.Options | set) &^ clear
		env.ReplStdin()
	}
	return nil
}

func (cmd *Cmd) Usage() error {
	fmt.Print(`usage: gomacro [OPTIONS] [files-and-dirs]

       Recognized options:
       -e EXPR evaluate expression
       -f      force option -w to overwrite existing files
       -h      show this help and exit
       -i      interactive. start a REPL after evaluating expression, files and dirs.
       -o LIST set/unset options, see below.
       -s      silent. do NOT show startup message, prompt, and expressions result
       -v      verbose. show startup message, prompt, and expressions result
       -w      write collected declarations and statements to *.go files

        LIST is a comma-separated list of one or more:
         decl      collect declarations
         ^decl     do NOT collect declarations
         stmt      collect statements
         ^stmt     do NOT collect statements
         verbose   same as -v
         ^verbose  same as -s
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
		if !file.IsDir() && endsWith(filename, ".gomacro") {
			filename = fmt.Sprintf("%s%c%s", dirname, os.PathSeparator, filename)
			err := cmd.EvalFile(filename)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cmd *Cmd) EvalFile(filename string) (err error) {
	env := cmd.Env
	env.Declarations = nil
	env.Statements = nil

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

	if cmd.WriteDeclsAndStmtsToFile {
		outname := filename
		if dot := strings.LastIndexByte(outname, '.'); dot >= 0 {
			// sanity check: dot must be in the file name, NOT in its path
			if slash := strings.LastIndexByte(outname, os.PathSeparator); slash < dot {
				outname = outname[0:dot]
			}
		}
		outname += ".go"
		if !cmd.OverwriteFiles {
			_, err := os.Stat(outname)
			if err == nil {
				env.warnf("file exists already, use -f to force overwriting: %v", outname)
				return nil
			}
		}
		env.writeDeclsToFile(outname)

		if env.Options&OptShowEval != 0 {
			fmt.Fprintf(env.Stdout, "// processed file: %v\t-> %v\n", filename, outname)
		}
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
