// this file was generated by gomacro command: import _b "crypto/sha1"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"crypto/sha1"
	. "reflect"
)

// reflection: allow interpreted code to import "crypto/sha1"
func init() {
	Packages["crypto/sha1"] = Package{
		Binds: map[string]Value{
			"BlockSize": ValueOf(sha1.BlockSize),
			"New":       ValueOf(sha1.New),
			"Size":      ValueOf(sha1.Size),
			"Sum":       ValueOf(sha1.Sum),
		}, Untypeds: map[string]string{
			"BlockSize": "int:64",
			"Size":      "int:20",
		},
	}
}
