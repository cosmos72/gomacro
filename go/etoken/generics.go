// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etoken

type Generics int

const (
	GENERICS_NONE Generics = iota
	// enables C++ style generics
	GENERICS_V1_CXX
	// enables generics "contracts are interfaces"
	GENERICS_V2_CTI
)

// can be changed at runtime. useful to enable them by in gomacro,
// without affecting packages that depend on gomacro, as Gophernotes
var GENERICS = GENERICS_NONE

func (g Generics) V1_CXX() bool {
	return g == GENERICS_V1_CXX
}

func (g Generics) V2_CTI() bool {
	return g == GENERICS_V2_CTI
}
