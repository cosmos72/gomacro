// this file was generated by gomacro command: import _b "crypto/sha256"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	sha256 "crypto/sha256"
)

// reflection: allow interpreted code to import "crypto/sha256"
func init() {
	Packages["crypto/sha256"] = Package{
		Name: "sha256",
		Binds: map[string]Value{
			"BlockSize":	ValueOf(sha256.BlockSize),
			"New":	ValueOf(sha256.New),
			"New224":	ValueOf(sha256.New224),
			"Size":	ValueOf(sha256.Size),
			"Size224":	ValueOf(sha256.Size224),
			"Sum224":	ValueOf(sha256.Sum224),
			"Sum256":	ValueOf(sha256.Sum256),
		}, Untypeds: map[string]string{
			"BlockSize":	"int:64",
			"Size":	"int:32",
			"Size224":	"int:28",
		}, 
	}
}
