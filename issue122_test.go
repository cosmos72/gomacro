/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * issue122_test.go
 *
 *  Created on: Jun 25 2022
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	"reflect"

	"github.com/cosmos72/gomacro/imports"
)

type Base struct{}
type Wrapper struct{ *Base }

func (f *Base) IsBase(struct{}) bool {
	return true
}

func (b *Wrapper) IsWrapper() bool {
	return true
}

func init() {
	imports.Packages["test/issue122"] = imports.Package{
		Types: map[string]reflect.Type{
			"Base":    reflect.TypeOf((*Base)(nil)).Elem(),
			"Wrapper": reflect.TypeOf((*Wrapper)(nil)).Elem(),
		},
	}
}
