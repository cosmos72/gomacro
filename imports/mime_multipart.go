// this file was generated by gomacro command: import _b "mime/multipart"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"mime/multipart"
	. "reflect"
)

// reflection: allow interpreted code to import "mime/multipart"
func init() {
	Packages["mime/multipart"] = Package{
		Binds: map[string]Value{
			"ErrMessageTooLarge": ValueOf(&multipart.ErrMessageTooLarge).Elem(),
			"NewReader":          ValueOf(multipart.NewReader),
			"NewWriter":          ValueOf(multipart.NewWriter),
		}, Types: map[string]Type{
			"File":       TypeOf((*multipart.File)(nil)).Elem(),
			"FileHeader": TypeOf((*multipart.FileHeader)(nil)).Elem(),
			"Form":       TypeOf((*multipart.Form)(nil)).Elem(),
			"Part":       TypeOf((*multipart.Part)(nil)).Elem(),
			"Reader":     TypeOf((*multipart.Reader)(nil)).Elem(),
			"Writer":     TypeOf((*multipart.Writer)(nil)).Elem(),
		}, Proxies: map[string]Type{
			"File": TypeOf((*P_mime_multipart_File)(nil)).Elem(),
		},
	}
}

// --------------- proxy for mime/multipart.File ---------------
type P_mime_multipart_File struct {
	Object  interface{}
	Close_  func(interface{}) error
	Read_   func(_proxy_obj_ interface{}, p []byte) (n int, err error)
	ReadAt_ func(_proxy_obj_ interface{}, p []byte, off int64) (n int, err error)
	Seek_   func(_proxy_obj_ interface{}, offset int64, whence int) (int64, error)
}

func (P *P_mime_multipart_File) Close() error {
	return P.Close_(P.Object)
}
func (P *P_mime_multipart_File) Read(p []byte) (n int, err error) {
	return P.Read_(P.Object, p)
}
func (P *P_mime_multipart_File) ReadAt(p []byte, off int64) (n int, err error) {
	return P.ReadAt_(P.Object, p, off)
}
func (P *P_mime_multipart_File) Seek(offset int64, whence int) (int64, error) {
	return P.Seek_(P.Object, offset, whence)
}
