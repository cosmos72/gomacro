// this file was generated by gomacro command: import _b "net/http/httptest"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"net/http/httptest"
	. "reflect"
)

// reflection: allow interpreted code to import "net/http/httptest"
func init() {
	Packages["net/http/httptest"] = Package{
		Binds: map[string]Value{
			"DefaultRemoteAddr":  ValueOf(httptest.DefaultRemoteAddr),
			"NewRecorder":        ValueOf(httptest.NewRecorder),
			"NewRequest":         ValueOf(httptest.NewRequest),
			"NewServer":          ValueOf(httptest.NewServer),
			"NewTLSServer":       ValueOf(httptest.NewTLSServer),
			"NewUnstartedServer": ValueOf(httptest.NewUnstartedServer),
		}, Types: map[string]Type{
			"ResponseRecorder": TypeOf((*httptest.ResponseRecorder)(nil)).Elem(),
			"Server":           TypeOf((*httptest.Server)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"DefaultRemoteAddr": "string:1.2.3.4",
		},
	}
}
