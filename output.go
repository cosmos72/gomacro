package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"io"
	r "reflect"
	"sort"
)

type RuntimeError struct {
	*Parser
	Format string
	Args   []interface{}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf(err.Format, err.toPrintables(err.Args)...)
}

func Errore(err error) interface{} {
	panic(err)
}

func (p *Parser) Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{p, format, args})
}

func (p *Parser) PackErrorf(format string, args ...interface{}) []r.Value {
	panic(RuntimeError{p, format, args})
}

func (p *Parser) Warnf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(p.Stderr, "warning: %s\n", str)
}

func (p *Parser) Debugf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(p.Stdout, "debug: %s\n", str)
}

func (p *Parser) FprintValues(out io.Writer, values ...r.Value) {
	if len(values) == 0 {
		fmt.Fprint(out, "// no value\n")
		return
	}
	for _, v := range values {
		p.FprintValue(out, v)
	}
}

func (p *Parser) FprintValue(out io.Writer, v r.Value) {
	var vi interface{}
	var vt r.Type
	if v == None {
		fmt.Fprint(out, "// no value\n")
		return
	} else if v == Nil {
		vi = nil
		vt = nil
	} else {
		vi = v.Interface()
		vt = v.Type()
	}
	vi = p.toPrintable(vi)
	if vi == nil && vt == nil {
		fmt.Fprint(out, "<nil>\n")
		return
	}
	switch vi := vi.(type) {
	case uint, uint8, uint32, uint64, uintptr:
		fmt.Fprintf(out, "%d <%v>\n", vi, vt)
	default:
		if vt == typeOfString {
			fmt.Fprintf(out, "%#v <%v>\n", vi, vt)
		} else {
			fmt.Fprintf(out, "%v <%v>\n", vi, vt)
		}
	}
}

func (p *Parser) Sprintf(format string, values ...interface{}) string {
	values = p.toPrintables(values)
	return fmt.Sprintf(format, values...)
}

func (p *Parser) toString(separator string, values ...interface{}) string {
	if len(values) == 0 {
		return ""
	}
	values = p.toPrintables(values)
	var buf bytes.Buffer
	for i, value := range values {
		if i != 0 {
			buf.WriteString(separator)
		}
		fmt.Fprint(&buf, value)
	}
	return buf.String()
}

func (p *Parser) toPrintables(values []interface{}) []interface{} {
	for i, vi := range values {
		values[i] = p.toPrintable(vi)
	}
	return values
}

func (p *Parser) toPrintable(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if v, ok := value.(r.Value); ok {
		return p.valueToPrintable(v)
	}
	v := r.ValueOf(value)
	k := v.Kind()
	if k == r.Array || k == r.Slice {
		n := v.Len()
		values := make([]interface{}, n)
		for i := 0; i < n; i++ {
			values[i] = p.toPrintable(v.Index(i))
		}
		return values
	}
	if node, ok := value.(ast.Node); ok {
		return p.nodeToPrintable(node)
	}
	return value
}

func (p *Parser) valueToPrintable(value r.Value) interface{} {
	if value == None {
		return "/*no value*/"
	} else if value == Nil {
		return nil
	} else {
		return p.toPrintable(value.Interface())
	}
}

func (p *Parser) nodeToPrintable(node ast.Node) interface{} {
	var buf bytes.Buffer
	err := format.Node(&buf, p.Fileset, node)
	if err != nil {
		return err
	}
	return buf.String()
}

func (p *Parser) showHelp(out io.Writer) {
	fmt.Fprint(out, `// interpreter commands:
:env    show available functions, variables and constants
:help   show this help
:quit   quit the interpreter
`)
}

func (env *Env) showEnv(out io.Writer) {
	binds := env.binds
	keys := make([]string, len(binds))
	i := 0
	for k := range binds {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	spaces15 := "               "
	for _, k := range keys {
		n := len(k) & 15
		fmt.Fprintf(out, "%s%s = ", k, spaces15[n:])
		env.FprintValue(out, binds[k])
	}
}
