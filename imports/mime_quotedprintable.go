// this file was generated by gomacro command: import _b "mime/quotedprintable"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	quotedprintable "mime/quotedprintable"
)

// reflection: allow interpreted code to import "mime/quotedprintable"
func init() {
	Packages["mime/quotedprintable"] = Package{
		Name: "quotedprintable",
		Binds: map[string]Value{
			"NewReader":	ValueOf(quotedprintable.NewReader),
			"NewWriter":	ValueOf(quotedprintable.NewWriter),
		}, Types: map[string]Type{
			"Reader":	TypeOf((*quotedprintable.Reader)(nil)).Elem(),
			"Writer":	TypeOf((*quotedprintable.Writer)(nil)).Elem(),
		}, 
	}
}
