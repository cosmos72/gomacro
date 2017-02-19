package main

import (
	"bufio"
	"fmt"
	"os"
	r "reflect"
	"strings"
)

type Binds map[string]r.Value

type Env struct {
	*Parser
	binds      Binds
	Outer      *Env
	iotaOffset int
}

func NewEnv(outer *Env) *Env {
	env := Env{}
	env.binds = make(map[string]r.Value)
	env.iotaOffset = 1
	if outer == nil {
		env.Parser = NewParser()
		env.addBuiltins()
		env.addInterpretedBuiltins()
	} else {
		env.Parser = outer.Parser
		env.Outer = outer
	}
	return &env
}

func (env *Env) Repl() {
	in := bufio.NewReader(os.Stdin)
	env.Repl1(in)
}

func (env *Env) Repl1(in *bufio.Reader) {
	fmt.Fprint(env.Stdout, "// Welcome to gomacro. Type :help for help\n")

	for env.ReadEvalPrint(in) {
	}
}

func (env *Env) ReadEvalPrint(in *bufio.Reader) (ret bool) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Fprintln(env.Stderr, rec)
			ret = true
		}
	}()

	fmt.Fprint(env.Stdout, "go> ")

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Fprintln(env.Stderr, err)
		return false
	}

	pos := findFirstToken(line)
	line = strings.TrimSpace(line[pos:])

	if len(line) > 0 && line[0] == ':' {
		if line == ":quit" {
			return false
		} else if line == ":env" {
			env.showEnv(env.Stdout)
			return true
		} else if line == ":help" {
			env.showHelp(env.Stdout)
			return true
		}
	}
	ast := env.Parse(line)
	// env.FprintValue(Stdout, r.ValueOf(ast))
	value, values := env.Eval(ast)
	if len(values) != 0 {
		env.FprintValues(env.Stdout, values...)
	} else {
		env.FprintValues(env.Stdout, value)
	}
	return true
}
