// this file was generated by gomacro command: import "testing/iotest"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "testing/iotest"
	. "reflect"
)

func init() {
	Binds["testing/iotest"] = map[string]Value{
		"DataErrReader":	ValueOf(pkg.DataErrReader),
		"ErrTimeout":	ValueOf(&pkg.ErrTimeout).Elem(),
		"HalfReader":	ValueOf(pkg.HalfReader),
		"NewReadLogger":	ValueOf(pkg.NewReadLogger),
		"NewWriteLogger":	ValueOf(pkg.NewWriteLogger),
		"OneByteReader":	ValueOf(pkg.OneByteReader),
		"TimeoutReader":	ValueOf(pkg.TimeoutReader),
		"TruncateWriter":	ValueOf(pkg.TruncateWriter),
	}
	Types["testing/iotest"] = map[string]Type{
	}
}