// this file was generated by gomacro command: import _b "testing/iotest"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	iotest "testing/iotest"
)

// reflection: allow interpreted code to import "testing/iotest"
func init() {
	Packages["testing/iotest"] = Package{
		Name: "iotest",
		Binds: map[string]Value{
			"DataErrReader":	ValueOf(iotest.DataErrReader),
			"ErrReader":	ValueOf(iotest.ErrReader),
			"ErrTimeout":	ValueOf(&iotest.ErrTimeout).Elem(),
			"HalfReader":	ValueOf(iotest.HalfReader),
			"NewReadLogger":	ValueOf(iotest.NewReadLogger),
			"NewWriteLogger":	ValueOf(iotest.NewWriteLogger),
			"OneByteReader":	ValueOf(iotest.OneByteReader),
			"TestReader":	ValueOf(iotest.TestReader),
			"TimeoutReader":	ValueOf(iotest.TimeoutReader),
			"TruncateWriter":	ValueOf(iotest.TruncateWriter),
		}, 
	}
}
