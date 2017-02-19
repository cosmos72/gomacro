package main

import (
	"os"
	r "reflect"
)

var Stdout = os.Stdout
var Stderr = os.Stderr

var Nil = r.ValueOf(nil)

var none struct{}
var None = r.ValueOf(none)

var typeOfString = r.TypeOf("")
var typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
var zeroStrings = []string{}
