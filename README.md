## gomacro - A Go interpreter with Lisp-like macros

Problem: "go generate" and many other Go tools automatically create
Go source code from some kind of description - usually an interface
specifications as WSDL, XSD, JSON...

Typically such specification is **NOT** written in Go, and typically
a variety of external tools are needed to convert it to Go source code.

Also, Go is currently lacking generics (read: C++-like templates)
because of the rationale "we do not yet know how to do them right,
and once you do them wrong everybody is stuck with them"

The purpose of Lisp-like macros is to execute arbitrary code
while compiling, **in particular** to generate source code.

This makes them very well suited (although arguably a bit low level)
for both purposes: code generation and C++-like templates, which
are a special case of code generation - for a demonstration of how
to implement C++-like templates on top of Lisp-like macros,
see for example the project https://github.com/cosmos72/cl-parametric-types
from the same author.

Building a Go interpreter that supports Lisp-like macros,
allows to embed all these code-generation activities
into regular Go source code, without the need for external tools
(except for the intepreter itself).

As a free bonus, we get support for Eval()

## Current Status

ALPHA.

The intepreter supports:
* type inference and Go multiple values
* variable and function definitions
* function calls (no imports yet, so builtins only)
* strings, booleans, integers, floats, complex numbers and interface{}

A lot of things are still missing:
* support to create arrays, slices, maps, channels, pointers
* fields access (a.b)
* method calls (i.e. functions with receivers)
* struct and interface definitions
* packages and imports
* macros

