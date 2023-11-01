// +build !go1.18

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2023 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * genimport.go
 *
 *  Created on May 26, 2017
 *      Author Massimiliano Ghilardi
 */

package genimport

var go_1_18_required = error! go toolchain >= 1.18 required
