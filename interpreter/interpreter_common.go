package interpreter

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

type InterpreterCommon struct {
	InterpreterBase
	AllMethods map[r.Type]Methods // methods implemented by interpreted code
}

func NewInterpreterCommon() *InterpreterCommon {
	return &InterpreterCommon{
		InterpreterBase: MakeInterpreterBase(),
		AllMethods:      make(map[r.Type]Methods),
	}
}
