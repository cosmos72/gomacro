package main

import (
	r "reflect"
)

var Nil = r.ValueOf(nil)

var none struct{}
var None = r.ValueOf(none)

var typeOfString = r.TypeOf("")
var typeOfInterface = r.TypeOf((*interface{})(nil)).Elem()
var zeroStrings = []string{}

const temporaryFunctionName = "gorepl_temporary_function"
