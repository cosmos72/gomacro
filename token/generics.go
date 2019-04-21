// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

// enable C++-style generics?
const GENERICS_V1_CXX = false

// enable generics "constraints are interfaces" ?
const GENERICS_V2_CTI = true

// can only enable one style of generics
func init() {
	if GENERICS_V1_CXX && GENERICS_V2_CTI {
		panic("github.com/cosmos72/gomacro/token: cannot enable both GENERICS_V1_CXX and GENERICS_V2_CI. Please disable at least one of them.")
	}
}
