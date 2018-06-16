// this file was generated by gomacro command: import _b "image/jpeg"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"image/jpeg"
	. "reflect"
)

// reflection: allow interpreted code to import "image/jpeg"
func init() {
	Packages["image/jpeg"] = Package{
		Binds: map[string]Value{
			"Decode":         ValueOf(jpeg.Decode),
			"DecodeConfig":   ValueOf(jpeg.DecodeConfig),
			"DefaultQuality": ValueOf(jpeg.DefaultQuality),
			"Encode":         ValueOf(jpeg.Encode),
		}, Types: map[string]Type{
			"FormatError":      TypeOf((*jpeg.FormatError)(nil)).Elem(),
			"Options":          TypeOf((*jpeg.Options)(nil)).Elem(),
			"Reader":           TypeOf((*jpeg.Reader)(nil)).Elem(),
			"UnsupportedError": TypeOf((*jpeg.UnsupportedError)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"Reader": TypeOf((*P_image_jpeg_Reader)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"DefaultQuality": "int:75",
		},
	}
}

// --------------- proxy for image/jpeg.Reader ---------------
type P_image_jpeg_Reader struct {
	Object    interface{}
	Read_     func(_proxy_obj_ interface{}, p []byte) (n int, err error)
	ReadByte_ func(interface{}) (byte, error)
}

func (P *P_image_jpeg_Reader) Read(p []byte) (n int, err error) {
	return P.Read_(P.Object, p)
}
func (P *P_image_jpeg_Reader) ReadByte() (byte, error) {
	return P.ReadByte_(P.Object)
}
