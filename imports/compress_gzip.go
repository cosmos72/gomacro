// this file was generated by gomacro command: import _b "compress/gzip"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"compress/gzip"
	. "reflect"
)

// reflection: allow interpreted code to import "compress/gzip"
func init() {
	Packages["compress/gzip"] = Package{
		Binds: map[string]Value{
			"BestCompression":    ValueOf(gzip.BestCompression),
			"BestSpeed":          ValueOf(gzip.BestSpeed),
			"DefaultCompression": ValueOf(gzip.DefaultCompression),
			"ErrChecksum":        ValueOf(&gzip.ErrChecksum).Elem(),
			"ErrHeader":          ValueOf(&gzip.ErrHeader).Elem(),
			"HuffmanOnly":        ValueOf(gzip.HuffmanOnly),
			"NewReader":          ValueOf(gzip.NewReader),
			"NewWriter":          ValueOf(gzip.NewWriter),
			"NewWriterLevel":     ValueOf(gzip.NewWriterLevel),
			"NoCompression":      ValueOf(gzip.NoCompression),
		}, Types: map[string]Type{
			"Header": TypeOf((*gzip.Header)(nil)).Elem(),
			"Reader": TypeOf((*gzip.Reader)(nil)).Elem(),
			"Writer": TypeOf((*gzip.Writer)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"BestCompression":    "int:9",
			"BestSpeed":          "int:1",
			"DefaultCompression": "int:-1",
			"HuffmanOnly":        "int:-2",
			"NoCompression":      "int:0",
		},
	}
}
