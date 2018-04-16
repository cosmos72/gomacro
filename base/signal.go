/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2018 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * signal.go
 *
 *  Created on: Apr 14, 2018
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"os"
	"os/signal"
	"sync/atomic"
	"unsafe"
)

// =======================================================================

func StartSignalHandler(handler func(os.Signal)) chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go signalHandlerGoroutine(c, handler)
	return c
}

func StopSignalHandler(c chan os.Signal) {
	close(c)
}

func signalHandlerGoroutine(c chan os.Signal, handler func(os.Signal)) {
	for {
		sig, ok := <-c
		if !ok {
			break
		}
		if handler != nil {
			handler(sig)
		}
	}
}

// =======================================================================

type Signal uint16

const (
	SigDefer Signal = 1 << iota // request to install a defer function
	SigReturn
	SigInterrupt // user pressed Ctrl+C, process received SIGINT, or similar

	SigNone = Signal(0) // no signal
	SigAll  = ^SigNone  // mask of all possible signals
)

func (sig Signal) String() string {
	var s string
	switch sig {
	case SigNone:
		s = "// signal: none"
	case SigReturn:
		s = "// signal: return"
	case SigDefer:
		s = "// signal: defer"
	case SigInterrupt:
		s = "// signal: interrupt"
	default:
		s = "// signal: unknown"
	}
	return s
}

type Signals struct {
	Sync  Signal
	Async Signal
}

func (s *Signals) IsEmpty() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(s))) == 0
}

