// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.11

package etoken

// AddLineColumnInfo adds alternative file, line, and column number
// information for a given file offset. The offset must be larger
// than the offset for the previously added alternative line info
// and smaller than the file size; otherwise the information is
// ignored.
//
// AddLineColumnInfo is typically used to register alternative position
// information for line directives such as //line filename:line:column.
//
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int) {
	// our github.com/cosmos72/gomacro/go/scanner calls this method,
	// but go/token implements it only for go >= 1.11
	//
	// approximate it on older go versions
	f.File.AddLineInfo(offset, filename, line)
}
