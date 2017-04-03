// constants.go
package constants

import (
	r "reflect"
)

type none struct{}

var (
	Nil = r.Value{}

	None = r.ValueOf(none{}) // used to indicate "no value"

	True  = r.ValueOf(true)
	False = r.ValueOf(false)
)
