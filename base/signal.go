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

type Signal uintptr

const (
	SigNone  Signal = 0
	SigDefer Signal = 1 << iota // request to install a defer function
	SigReturn
	SigInterrupt // user pressed Ctrl+C, process received SIGINT, or similar
)

func (sig Signal) String() string {
	var s string
	switch sig {
	case SigNone:
		s = "// no signal"
	case SigReturn:
		s = "// return signal"
	case SigDefer:
		s = "// defer signal"
	case SigInterrupt:
		s = "// interrupted"
	default:
		s = "// unknown signal"
	}
	return s
}

type Signals struct {
	mask Signal
}

func (s *Signals) Get() Signal {
	return Signal(atomic.LoadUintptr((*uintptr)(&s.mask)))
}

func (s *Signals) IsEmpty() bool {
	return s.Get() == 0
}

func (s *Signals) Set(bit Signal) {
	for {
		curr := s.Get()
		toset := curr | bit
		if curr == toset {
			break
		}
		if atomic.CompareAndSwapUintptr((*uintptr)(&s.mask), uintptr(curr), uintptr(toset)) {
			break
		}
	}
}

// return true if bit was set, false if it was already unset
func (s *Signals) Clear(bit Signal) bool {
	bit = ^bit
	for {
		curr := s.Get()
		toset := curr & bit
		if curr == toset {
			return false
		}
		if atomic.CompareAndSwapUintptr((*uintptr)(&s.mask), uintptr(curr), uintptr(toset)) {
			return true
		}
	}
}
