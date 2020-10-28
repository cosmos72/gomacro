// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gc
// +build wasm

package gls

// return the current goroutine ID.
//
// note that the returned value is DIFFERENT from most other goroutine libraries:
// this GoID() returns the address, converted to uintptr, of the runtime.g struct.
// NOT the runtime.g.goid field returned by most other libraries.
func GoID() uintptr {
	// stub implementation, WASM does not have threads yet.
	// Question: does GOARCH=wasm support goroutines?
	// If yes, than a stub implementation is not enough.
	return 0
}
