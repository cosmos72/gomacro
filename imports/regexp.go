// this file was generated by gomacro command: import _b "regexp"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	regexp "regexp"
)

// reflection: allow interpreted code to import "regexp"
func init() {
	Packages["regexp"] = Package{
		Name: "regexp",
		Binds: map[string]Value{
			"Compile":	ValueOf(regexp.Compile),
			"CompilePOSIX":	ValueOf(regexp.CompilePOSIX),
			"Match":	ValueOf(regexp.Match),
			"MatchReader":	ValueOf(regexp.MatchReader),
			"MatchString":	ValueOf(regexp.MatchString),
			"MustCompile":	ValueOf(regexp.MustCompile),
			"MustCompilePOSIX":	ValueOf(regexp.MustCompilePOSIX),
			"QuoteMeta":	ValueOf(regexp.QuoteMeta),
		}, Types: map[string]Type{
			"Regexp":	TypeOf((*regexp.Regexp)(nil)).Elem(),
		}, 
	}
}
