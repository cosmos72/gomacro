// this file was generated by gomacro command: import _b "encoding/json"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"encoding/json"
	. "reflect"
)

// reflection: allow interpreted code to import "encoding/json"
func init() {
	Packages["encoding/json"] = Package{
		Binds: map[string]Value{
			"Compact":       ValueOf(json.Compact),
			"HTMLEscape":    ValueOf(json.HTMLEscape),
			"Indent":        ValueOf(json.Indent),
			"Marshal":       ValueOf(json.Marshal),
			"MarshalIndent": ValueOf(json.MarshalIndent),
			"NewDecoder":    ValueOf(json.NewDecoder),
			"NewEncoder":    ValueOf(json.NewEncoder),
			"Unmarshal":     ValueOf(json.Unmarshal),
			"Valid":         ValueOf(json.Valid),
		}, Types: map[string]Type{
			"Decoder":               TypeOf((*json.Decoder)(nil)).Elem(),
			"Delim":                 TypeOf((*json.Delim)(nil)).Elem(),
			"Encoder":               TypeOf((*json.Encoder)(nil)).Elem(),
			"InvalidUTF8Error":      TypeOf((*json.InvalidUTF8Error)(nil)).Elem(),
			"InvalidUnmarshalError": TypeOf((*json.InvalidUnmarshalError)(nil)).Elem(),
			"Marshaler":             TypeOf((*json.Marshaler)(nil)).Elem(),
			"MarshalerError":        TypeOf((*json.MarshalerError)(nil)).Elem(),
			"Number":                TypeOf((*json.Number)(nil)).Elem(),
			"RawMessage":            TypeOf((*json.RawMessage)(nil)).Elem(),
			"SyntaxError":           TypeOf((*json.SyntaxError)(nil)).Elem(),
			"Token":                 TypeOf((*json.Token)(nil)).Elem(),
			"UnmarshalFieldError":   TypeOf((*json.UnmarshalFieldError)(nil)).Elem(),
			"UnmarshalTypeError":    TypeOf((*json.UnmarshalTypeError)(nil)).Elem(),
			"Unmarshaler":           TypeOf((*json.Unmarshaler)(nil)).Elem(),
			"UnsupportedTypeError":  TypeOf((*json.UnsupportedTypeError)(nil)).Elem(),
			"UnsupportedValueError": TypeOf((*json.UnsupportedValueError)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Marshaler":   TypeOf((*P_encoding_json_Marshaler)(nil)).Elem(),
			"Unmarshaler": TypeOf((*P_encoding_json_Unmarshaler)(nil)).Elem(),
		},
	}
}

// --------------- proxy for encoding/json.Marshaler ---------------
type P_encoding_json_Marshaler struct {
	Object       interface{}
	MarshalJSON_ func(interface{}) ([]byte, error)
}

func (P *P_encoding_json_Marshaler) MarshalJSON() ([]byte, error) {
	return P.MarshalJSON_(P.Object)
}

// --------------- proxy for encoding/json.Unmarshaler ---------------
type P_encoding_json_Unmarshaler struct {
	Object         interface{}
	UnmarshalJSON_ func(interface{}, []byte) error
}

func (P *P_encoding_json_Unmarshaler) UnmarshalJSON(unnamed0 []byte) error {
	return P.UnmarshalJSON_(P.Object, unnamed0)
}
