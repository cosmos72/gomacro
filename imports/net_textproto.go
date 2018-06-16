// this file was generated by gomacro command: import _b "net/textproto"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"net/textproto"
	. "reflect"
)

// reflection: allow interpreted code to import "net/textproto"
func init() {
	Packages["net/textproto"] = Package{
		Binds: map[string]Value{
			"CanonicalMIMEHeaderKey": ValueOf(textproto.CanonicalMIMEHeaderKey),
			"Dial":       ValueOf(textproto.Dial),
			"NewConn":    ValueOf(textproto.NewConn),
			"NewReader":  ValueOf(textproto.NewReader),
			"NewWriter":  ValueOf(textproto.NewWriter),
			"TrimBytes":  ValueOf(textproto.TrimBytes),
			"TrimString": ValueOf(textproto.TrimString),
		}, Types: map[string]Type{
			"Conn":          TypeOf((*textproto.Conn)(nil)).Elem(),
			"Error":         TypeOf((*textproto.Error)(nil)).Elem(),
			"MIMEHeader":    TypeOf((*textproto.MIMEHeader)(nil)).Elem(),
			"Pipeline":      TypeOf((*textproto.Pipeline)(nil)).Elem(),
			"ProtocolError": TypeOf((*textproto.ProtocolError)(nil)).Elem(),
			"Reader":        TypeOf((*textproto.Reader)(nil)).Elem(),
			"Writer":        TypeOf((*textproto.Writer)(nil)).Elem(),
		}, Wrappers: map[string][]string{
			"Conn": []string{"DotReader", "DotWriter", "EndRequest", "EndResponse", "Next", "PrintfLine", "ReadCodeLine", "ReadContinuedLine", "ReadContinuedLineBytes", "ReadDotBytes", "ReadDotLines", "ReadLine", "ReadLineBytes", "ReadMIMEHeader", "ReadResponse", "StartRequest", "StartResponse"},
		},
	}
}
