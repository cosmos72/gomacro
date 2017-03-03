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
* type inference
* variable and function definitions
* strings, booleans, integers, floats, complex numbers and interface{}
* function calls, including multiple return values
  (no dotted notation `foo.bar` yet, so no method calls or imported symbols)
* imports: Go standard packages "just work", importing other packages requires Go 1.8+ and Linux
  (but they cannot be used yet, dotted notation `foo.bar` still unimplemented)
* macro definitions, for example `macro foo(a, b, c interface{}) interface{} { return b }`
* macro calls, for example `foo x; y; z`
* macroexpansion code walker, MacroExpand and MacroExpand1
* quote and quasiquote. they take any number of arguments in curly braces, for example:
  `quote { x; y; z }`
* unquote and unquote_splice

A lot of things are still missing:
* support to create arrays, slices, maps, channels, pointers
* dotted notation i.e. `foo.bar` for fields access and imported symbols 
* method calls (i.e. functions with receivers)
* struct and interface definitions
* switching to a different package
* multiline input, history/readline
* nested macro calls and quoted macros:
  both `foo bar baz` and `quote{foo bar baz}` parse,
  but `foo {bar baz}` and `quote{foo {bar baz}}` don't parse yet
  because foo{bar /*...*/} is a Go composite literal
  The case `foo 1; bar baz; 2` instead parses, but
  all statements after `bar` are consumed by `bar`, not by `foo`

* support to quote or quasiquote macro calls, i.e. `quasiquote{some_macro ...}`
  it currently works only if some_macro is already defined (see also next point)

* support to embed a variable number of spliced arguments in quoted calls to macros, i.e.
    `quasiquote{some_macro unquote_splice{...} ...}`
  or even in quoted calls to Go reserved keywords, i.e.
	`quasiquote{for unquote_splice{...} ...}`
  see parser/macro.go: Env.parseQuote() for details
