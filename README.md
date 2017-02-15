## gomacro - A Go interpreter with eval and Lisp-like macros

Problem: "go generate" and many other Go tools automatically create
Go source code from some kind of description - usually an interface
specifications as WSDL, XSD, JSON...

Typically such specification is **NOT** written in Go, and typically
external tools are needed to convert it to Go source code.

Also, Go is currently lacking generics (read: C++-like templates)
because of the rationale "we do not yet know how to do them right,
and once you do them wrong everybody is stuck with them"

The purpose of Lisp-like macros is to execute arbitrary code
while compiling, **in particular** to generate source code.

This makes them very well suited (although arguably a bit low level)
for both purposes: code generation and C++-like templates, which
are a special case of code generation.

Building a Go interpreter that supports Lisp-like macros,
allows to embed all these code-generation activities
into regular Go source code, without the need for external tools
(except for the intepreter itself).

As a free bonus, we get support for Eval()

## current status

ALPHA.

The intepreter works for simple variable definitions
and unary/binary expressions using booleans, integers,
floats and complex numbers.

A lot of things are still missing:
* support for strings
* fields access (a.b)
* function calls
* function definitions
* struct and interface definitions
* macros

