package classic

import (
	r "reflect"

	"github.com/cosmos72/gomacro/base"
)

type ThreadGlobals struct {
	base.Globals
	AllMethods map[r.Type]Methods // methods implemented by interpreted code
	CompEnv    interface{}        // *fast.CompEnv // temporary...
}

func NewThreadGlobals() *ThreadGlobals {
	tg := &ThreadGlobals{
		AllMethods: make(map[r.Type]Methods),
	}
	tg.Globals.Init()
	return tg
}
