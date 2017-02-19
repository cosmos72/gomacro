package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"io"
	r "reflect"
)

type RuntimeError struct {
	Format string
	Args   []interface{}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf(err.Format, err.Args...)
}

func Errore(err error) interface{} {
	panic(err)
}

func (p *Parser) Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{format, args})
}

func (p *Parser) PackErrorf(format string, args ...interface{}) []r.Value {
	panic(RuntimeError{format, args})
}

func (p *Parser) Warnf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(Stderr, "warning: %s\n", str)
}

func (p *Parser) Debugf(format string, args ...interface{}) {
	str := p.Sprintf(format, args...)
	fmt.Fprintf(Stdout, "debug: %s\n", str)
}

func (p *Parser) FprintMultipleValues(out io.Writer, values ...r.Value) {
	if len(values) == 0 {
		fmt.Fprint(out, "// no values\n")
		return
	}
	var vi interface{}
	var vt r.Type
	for _, v := range values {
		if v == None {
			fmt.Fprint(out, "// no values\n")
			continue
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
			continue
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
