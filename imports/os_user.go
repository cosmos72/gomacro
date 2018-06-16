// this file was generated by gomacro command: import _b "os/user"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"os/user"
	. "reflect"
)

// reflection: allow interpreted code to import "os/user"
func init() {
	Packages["os/user"] = Package{
		Binds: map[string]Value{
			"Current":       ValueOf(user.Current),
			"Lookup":        ValueOf(user.Lookup),
			"LookupGroup":   ValueOf(user.LookupGroup),
			"LookupGroupId": ValueOf(user.LookupGroupId),
			"LookupId":      ValueOf(user.LookupId),
		}, Types: map[string]Type{
			"Group":               TypeOf((*user.Group)(nil)).Elem(),
			"UnknownGroupError":   TypeOf((*user.UnknownGroupError)(nil)).Elem(),
			"UnknownGroupIdError": TypeOf((*user.UnknownGroupIdError)(nil)).Elem(),
			"UnknownUserError":    TypeOf((*user.UnknownUserError)(nil)).Elem(),
			"UnknownUserIdError":  TypeOf((*user.UnknownUserIdError)(nil)).Elem(),
			"User":                TypeOf((*user.User)(nil)).Elem(),
		},
	}
}
