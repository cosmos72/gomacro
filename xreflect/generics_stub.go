// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

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
 * generics_stub.go
 *
 *  Created on May 12, 2019
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"github.com/cosmos72/gomacro/go/mtoken"
)

func (v *Universe) addBasicTypesMethodsCTI() {
	if !mtoken.GENERICS_V2_CTI {
		return
	}
}
