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
 * util.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"go/build"
	"os"
	r "reflect"
	"strings"

	xr "github.com/cosmos72/gomacro/xreflect"
)

func PackValues(val0 r.Value, values []r.Value) []r.Value {
	if len(values) == 0 && val0 != None {
		values = []r.Value{val0}
	}
	return values
}

func PackTypes(typ0 xr.Type, types []xr.Type) []xr.Type {
	if len(types) == 0 && typ0 != nil {
		types = []xr.Type{typ0}
	}
	return types
}

func UnpackValues(vals []r.Value) (r.Value, []r.Value) {
	val0 := None
	if len(vals) > 0 {
		val0 = vals[0]
	}
	return val0, vals
}

// ValueInterface() is a zero-value-safe version of reflect.Value.Interface()
func ValueInterface(v r.Value) interface{} {
	if !v.IsValid() || !v.CanInterface() || v == None {
		return nil
	}
	return v.Interface()
}

// ValueType() is a zero-value-safe version of reflect.Value.Type()
func ValueType(value r.Value) r.Type {
	if !value.IsValid() || value == None {
		return nil
	}
	return value.Type()
}

func IsNillableKind(k r.Kind) bool {
	switch k {
	case r.Invalid, // nil is nillable...
		r.Chan, r.Func, r.Interface, r.Map, r.Ptr, r.Slice:
		return true
	default:
		return false
	}
}

// split 's' into a prefix and suffix separated by 'separator'.
// suffix is trimmed with strings.TrimSpace() before returning it
func Split2(s string, separator rune) (string, string) {
	var prefix, suffix string
	if space := strings.IndexByte(s, ' '); space > 0 {
		prefix = s[:space]
		suffix = strings.TrimSpace(s[space+1:])
	} else {
		prefix = s
	}
	return prefix, suffix
}

// always use forward slashes. they work also on Windows...
func unixpath(path string) string {
	if os.PathSeparator != '/' && len(path) != 0 {
		path = strings.Replace(path, string(os.PathSeparator), "/", -1)
	}
	return path
}

// find user's home directory, see https://stackoverflow.com/questions/2552416/how-can-i-find-the-users-home-dir-in-a-cross-platform-manner-using-c
// without importing "os/user" - which requires cgo to work thus makes cross-compile difficult, see https://github.com/golang/go/issues/11797
func UserHomeDir() string {
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = os.Getenv("USERPROFILE")
		if len(home) == 0 {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		}
	}
	return unixpath(home)
}

func Subdir(dirs ...string) string {
	// should use string(os.PathSeparator) instead of "/', but:
	// 1) package names use '/', not os.PathSeparator
	// 2) it would complicate DirName()
	return strings.Join(dirs, "/")
}

var (
	GoSrcDir = Subdir(build.Default.GOPATH, "src")

	GomacroDir = Subdir(GoSrcDir, "github.com", "cosmos72", "gomacro") // vendored copies of gomacro may need to change this
)
