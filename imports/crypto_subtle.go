// this file was generated by gomacro command: import _b "crypto/subtle"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"crypto/subtle"
	. "reflect"
)

// reflection: allow interpreted code to import "crypto/subtle"
func init() {
	Packages["crypto/subtle"] = Package{
		Binds: map[string]Value{
			"ConstantTimeByteEq":   ValueOf(subtle.ConstantTimeByteEq),
			"ConstantTimeCompare":  ValueOf(subtle.ConstantTimeCompare),
			"ConstantTimeCopy":     ValueOf(subtle.ConstantTimeCopy),
			"ConstantTimeEq":       ValueOf(subtle.ConstantTimeEq),
			"ConstantTimeLessOrEq": ValueOf(subtle.ConstantTimeLessOrEq),
			"ConstantTimeSelect":   ValueOf(subtle.ConstantTimeSelect),
		},
	}
}
