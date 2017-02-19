package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

type Parser struct {
	Packagename string
	Filename    string
	Fileset     *token.FileSet
	Parsermode  parser.Mode
	Stdout      io.Writer
	Stderr      io.Writer
}

func NewParser() *Parser {
	p := Parser{}
	p.Packagename = "main"
	p.Filename = "main.go"
	p.Fileset = token.NewFileSet()
	// using both os.Stdout and os.Stderr can interleave impredictably
	// normal output and diagnostic messages - ugly in interactive use
	p.Stdout = os.Stdout
	p.Stderr = os.Stdout
	return &p
}

func (p *Parser) Parse(src interface{}) ast.Node {
	str := p.ReadFromSource(src)
	node, err := p.parseOrError(str)
	if err != nil {
		Errore(err)
		return nil
	}
	return node
}

func (p *Parser) parseOrError(str string) (ast.Node, error) {
	pos := findFirstToken(str)
	str = str[pos:]
	expr, err := parser.ParseExprFrom(p.Fileset, p.Filename, str, 0)
	if err == nil {
		if p.Parsermode == 0 {
			return expr, nil
		}
		return parser.ParseExprFrom(p.Fileset, p.Filename, str, p.Parsermode)
	}
	firstIdent := extractFirstIdentifier(str)
	switch firstIdent {
	case "package":
		/* nothing to do */
	case "const", "func", "import", "type", "var":
		str = fmt.Sprintf("package %s; %s", p.Packagename, str)
	default:
		str = fmt.Sprintf("package %s; func %s() { %s }", p.Packagename, temporaryFunctionName, str)
	}
	return parser.ParseFile(p.Fileset, p.Filename, str, p.Parsermode)
}
