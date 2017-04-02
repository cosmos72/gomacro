// constants.go
package constants

import (
	r "reflect"
)

var (
	Nil = r.Value{}

	none struct{}
	None = r.ValueOf(none) // used to indicate "no value"

	True  = r.ValueOf(true)
	False = r.ValueOf(false)
)
