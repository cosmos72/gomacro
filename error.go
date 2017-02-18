package main

import (
	"fmt"
	r "reflect"
)

type RuntimeError struct {
	Format string
	Args   []interface{}
}

func (err RuntimeError) Error() string {
	return fmt.Sprintf(err.Format, err.Args...)
}

func Error(err error) (r.Value, []r.Value) {
	panic(err)
}

func Errorf(format string, args ...interface{}) (r.Value, []r.Value) {
	panic(RuntimeError{format, args})
}

func Warnf(format string, args ...interface{}) {
	fmt.Fprint(Stderr, "warning: ")
	fmt.Fprintf(Stderr, format, args...)
}
