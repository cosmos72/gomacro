package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"reflect"
	"strings"
)

const TemporaryFunctionName = "gorepl_temporary_function"

var Nil reflect.Value = reflect.ValueOf(nil)

type Binds map[string]reflect.Value

type Env struct {
	Binds
	Outer *Env
}

type Interpreter struct {
	*Env
	Packagename string
	Filename    string
	Fileset     *token.FileSet
	Parsermode  parser.Mode
	iotaOffset  int
}

func New() *Interpreter {
	ir := Interpreter{}
	ir.PushEnv()
	ir.Packagename = "main"
	ir.Filename = "main.go"
	ir.Fileset = token.NewFileSet()
	ir.iotaOffset = 1
	return &ir
}

func (ir *Interpreter) PushEnv() {
	ir.Env = &Env{make(Binds), ir.Env}
}

func (ir *Interpreter) PopEnv() {
	if ir.Env != nil {
		ir.Env = ir.Env.Outer
	}
}

func (ir *Interpreter) Parse(src interface{}) (ast.Node, error) {
	expr, err := parser.ParseExprFrom(ir.Fileset, ir.Filename, src, 0)
	if err == nil {
		if ir.Parsermode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(ir.Fileset, ir.Filename, src, ir.Parsermode)
	}
	str, ok := src.(string)
	if ok {
		str = strings.TrimLeft(str, " \f\t\r\n\v")
		firstWord := str
		space := strings.IndexAny(str, " \f\t\r\n\v")
		if space >= 0 {
			firstWord = str[:space]
		}
		switch firstWord {
		case "package":
			/* nothing to do */
		case "func", "var", "const", "type":
			str = fmt.Sprintf("package %s; %s", ir.Packagename, str)
		default:
			str = fmt.Sprintf("package %s; func %s() { %s }", ir.Packagename, TemporaryFunctionName, str)
		}
		src = str
	}
	return parser.ParseFile(ir.Fileset, ir.Filename, src, ir.Parsermode)
}

func (ir *Interpreter) PrintAst(out io.Writer, prefix string, node ast.Node) {
	fmt.Fprint(out, prefix)
	printer.Fprint(out, ir.Fileset, node)
	fmt.Fprintln(out)
}

func (ir *Interpreter) Print(out io.Writer, value reflect.Value) {
	v := value.Interface()
	switch v := v.(type) {
	case uint64:
		fmt.Fprintf(out, "%d\n", v)
	default:
		fmt.Fprintf(out, "%#v\n", v)
	}
}

func (ir *Interpreter) Loop() {
	in := bufio.NewReader(os.Stdin)
	ir.Loop3(in, os.Stdout, os.Stderr)
}

func (ir *Interpreter) Loop1(in *bufio.Reader) {
	ir.Loop3(in, os.Stdout, os.Stderr)
}

func (ir *Interpreter) Loop2(in *bufio.Reader, out_and_err io.Writer) {
	ir.Loop3(in, out_and_err, out_and_err)
}

func (ir *Interpreter) Loop3(in *bufio.Reader, out io.Writer, eout io.Writer) {
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(eout, err)
			break
		}
		ast, err := ir.Parse(line)
		if err != nil {
			fmt.Fprintln(eout, err)
			continue
		}
		// ir.PrintAst(out, "parsed: ", ast)
		value, err := ir.Eval(ast)
		if err != nil {
			fmt.Fprintln(eout, err)
			continue
		}
		ir.Print(out, value)
	}
}

func main() {
	interpreter := New()
	// ir.Parsermode = parser.Trace
	interpreter.Loop()
}
