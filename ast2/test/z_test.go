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
 * z_test.go
 *
 *  Created on: May 05, 2018
 *      Author: Massimiliano Ghilardi
 */

package test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/cosmos72/gomacro/ast2"
	"github.com/cosmos72/gomacro/base"
	"github.com/cosmos72/gomacro/parser"
	"github.com/cosmos72/gomacro/token"
)

func TestReader(t *testing.T) {
	tests := []struct {
		Name string
		Path string
	}{
		{"z_test_data_2", "z_test_data_2.txt"},
		{"fast_global", "../../fast/global.go"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_testReader(t, test.Path)
		})
	}
}

func _testReader(t *testing.T, filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("read file %q failed: %v", filename, err)
		return
	}

	fset := token.NewFileSet()
	st := base.Stringer{Fileset: fset}

	var p parser.Parser
	p.Init(fset, filename, 0, bytes)

	nodes, err := p.Parse()
	if err != nil {
		t.Errorf("parse file %q failed: %v", filename, err)
		return
	}
	r := AstReader(NodeSlice{nodes})

	for r.Next(); !r.Empty(); r.Next() {
		st.Fprintf(os.Stdout, "%v\n", r.Node)
	}
}
