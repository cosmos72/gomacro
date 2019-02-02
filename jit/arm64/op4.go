/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * op4.go
 *
 *  Created on Jan 27, 2019
 *      Author Massimiliano Ghilardi
 */

package arch

import (
	"fmt"
)

// ============================================================================
// quaternary operation
type Op4 uint8

const (
/*
	LEA4 Op4 = 0x8D
*/
)

func (op Op4) String() string {
	return fmt.Sprintf("Op4(%d)", int(op))
}

// ============================================================================
func (asm *Asm) Op4(op Op4, a Arg, b Arg, c Arg, dst Arg) *Asm {
	errorf("Op4 not implemented: %v %v %v %v", op, a, b, c, dst)
	return asm
}
