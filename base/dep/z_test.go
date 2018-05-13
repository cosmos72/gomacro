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
 *  Created on: May 03, 2018
 *      Author: Massimiliano Ghilardi
 */

package dep

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/cosmos72/gomacro/parser"
	"github.com/cosmos72/gomacro/token"
)

func TestRemoveItem(t *testing.T) {
	list := []string{"Env", "Stmt"}
	out := remove_item_inplace("Stmt", list)
	expect := []string{"Env"}
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("expected %v, actual %v", expect, out)
	}
}

func TestSortUnique1(t *testing.T) {
	in := []string{"c", "a", "c", "b", "a", "b", "x"}
	expect := []string{"a", "b", "c", "x"}
	_testSortUnique(t, in, expect)
}

func TestSortUnique2(t *testing.T) {
	in := []string{"Debugger", "Env", "IrGlobals", "Stmt", "Stmt", "poolCapacity"}
	expect := []string{"Debugger", "Env", "IrGlobals", "Stmt", "poolCapacity"}
	_testSortUnique(t, in, expect)
}

func _testSortUnique(t *testing.T, in []string, expect []string) {
	out := sort_unique_inplace(in)
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("expected %v, actual %v", expect, out)
	}
}

func TestSorter(t *testing.T) {
	tests := []struct {
		Name string
		Path string
	}{
		{"api", "api.go"},
		{"z_test_data_1", "z_test_data_1.txt"},
		{"z_test_data_2", "z_test_data_2.txt"},
		{"z_test_data_3", "z_test_data_3.txt"},
		{"fast_global", "../../fast/global.go"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_testSorter(t, test.Path)
		})
	}
}

func _testSorter(t *testing.T, filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("read file %q failed: %v", filename, err)
		return
	}

	var p parser.Parser
	fset := token.NewFileSet()
	p.Init(fset, filename, 0, bytes)

	nodes, err := p.Parse()
	if err != nil {
		t.Errorf("parse file %q failed: %v", filename, err)
		return
	}
	s := NewSorter()
	s.LoadNodes(nodes)

	for {
		sorted := s.Some()
		if len(sorted) == 0 {
			break
		}
		fmt.Print("---- sorted decls ----\n")
		sorted.Print()
	}
}
