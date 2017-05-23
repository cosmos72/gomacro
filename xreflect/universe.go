/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * universe.go
 *
 *  Created on May 14, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/types"
	"reflect"
	"sync"

	"github.com/cosmos72/gomacro/typeutil"
)

type Types struct {
	gmap typeutil.Map
}

type Universe struct {
	Types
	ReflectTypes    map[reflect.Type]Type
	BasicTypes      []Type
	TypeOfInterface Type
	TypeOfError     Type
	TryResolve      func(name, pkgpath string) Type
	Packages        map[string]*Package
	Importer        types.ImporterFrom
	RebuildDepth    int
	mutex           sync.Mutex
	debugmutex      int
	ThreadSafe      bool
}

func lock(v *Universe) *Universe {
	if v.debugmutex != 0 {
		errorf(nil, "deadlocking universe %p", v)
	}
	v.mutex.Lock()
	v.debugmutex++
	return v
}

func un(v *Universe) {
	// debugf("unlocking universe %p", v)
	v.mutex.Unlock()
	v.debugmutex--
}

func (v *Universe) rebuild() bool {
	return v.RebuildDepth >= 0
}

func (v *Universe) cache(rt reflect.Type, t Type) Type {
	if v.ReflectTypes == nil {
		v.ReflectTypes = make(map[reflect.Type]Type)
	}
	v.ReflectTypes[rt] = t
	// debugf("added rtype to cache: %v -> %v (%v)", rt, t, t.ReflectType())
	return t
}

func (v *Universe) CachePackage(pkg *types.Package) {
	if pkg == nil {
		return
	}
	if v.Packages == nil {
		v.Packages = make(map[string]*Package)
	}
	v.Packages[pkg.Path()] = (*Package)(pkg)
}

func (v *Universe) namedTypeFromImport(rtype reflect.Type) Type {
	t := v.namedTypeFromPackageCache(rtype)
	// importer gives accurate view of wrapper methods for embedded fields... use if type has methods
	if t != nil || rtype.NumMethod() == 0 {
		return t
	}
	if v.Importer == nil {
		v.Importer = DefaultImporter()
	}
	pkgpath := rtype.PkgPath()
	pkg, err := v.Importer.Import(pkgpath)
	if err != nil || pkg == nil {
		return nil
	}
	v.CachePackage(pkg)

	return v.namedTypeFromPackage(rtype, pkg)
}

func (v *Universe) namedTypeFromPackageCache(rtype reflect.Type) Type {
	pkgpath := rtype.PkgPath()
	pkg := (*types.Package)(v.Packages[pkgpath])
	if pkg != nil {
		return v.namedTypeFromPackage(rtype, pkg)
	}
	return nil
}

func (v *Universe) namedTypeFromPackage(rtype reflect.Type, pkg *types.Package) Type {
	name := rtype.Name()
	if scope := pkg.Scope(); scope != nil && len(name) != 0 {
		if obj := scope.Lookup(name); obj != nil {
			if gtype := obj.Type(); gtype != nil {
				// debugf("imported named type %v for %v", gtype, rtype)
				// not v.MakeType, because we already hold the lock
				return v.maketype3(gtypeToKind(nil, gtype), gtype, rtype)
			}
		}
	}
	return nil
}
