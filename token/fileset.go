// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"go/token"
)

// -----------------------------------------------------------------------------
// FileSet

// A FileSet represents a set of source files.
// This is a wrapper for go/token.FileSet that adds a starting line offset to each file in the set
//
type FileSet struct {
	token.FileSet
	lines map[*token.File]int // starting line of each file in the set
}

// NewFileSet creates a new file set.
func NewFileSet() *FileSet {
	return &FileSet{
		FileSet: *token.NewFileSet(),
		lines:   make(map[*token.File]int),
	}
}

// AddFile adds a new file with a given filename, base offset, and file size
func (s *FileSet) AddFile(filename string, base, size, line int) *token.File {
	f := s.FileSet.AddFile(filename, base, size)
	s.lines[f] = line
	return f
}

// PositionFor converts a Pos p in the fileset into a Position value.
// If adjusted is set, the position may be adjusted by position-altering
// //line comments; otherwise those comments are ignored.
// p must be a Pos value in s or NoPos.
//
func (s *FileSet) PositionFor(p token.Pos, adjusted bool) (pos token.Position) {
	if f := s.File(p); f != nil {
		pos = f.PositionFor(p, adjusted)
		pos.Line += s.lines[f]
	}
	return
}

// Position converts a Pos p in the fileset into a Position value.
// Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).
//
func (s *FileSet) Position(p token.Pos) (pos token.Position) {
	return s.PositionFor(p, true)
}
