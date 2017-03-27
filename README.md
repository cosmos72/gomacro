## gomacro - A Go interpreter with Lisp-like macros

Gomacro is "yet another" Go interpreter.

Started as an experiment to add Lisp-like macros to Go,
it is a fairly complete Go interpreter, implemented in pure Go,
built on top of the go/ast and reflect packages.

Gomacro can be used as:
* a standalone executable with interactive Go REPL:  
  just run `gomacro` from your command line or, better, `rlwrap gomacro`
  (rlwrap is a wrapper that adds history and line editing to terminal-based programs - available on many platforms)
  Available options:  
    -e EXPRESSIONS: evaluate the expressions  
    -s: silent. suppress startup message and prompt (default when executing a file)  
    -v: verbose. show the result of expressions and statements (default for interactive REPL and -e EXPRESSIONS)

* a library that adds Eval() and scripting capabilities
  to your Go programs - provided you comply with its GPL license

* a way to execute Go source code on-the-fly without a Go compiler:  
  you can either run `gomacro FILENAME.go` (works on every supported platform)

  or you can insert a line `#!/usr/bin/env gomacro` at the beginning of a Go source file,
  then mark the file as executable with `chmod +x FILENAME.go` and finally execute it
  with `./FILENAME.go` (works only on Unix-like systems: Linux, *BSD, Mac OS X ...)

* a Go code generation tool:  
  run `gomacro -w FILENAMES` to parse and execute one or more files.
  For each filename on the command line, gomacro will parse and execute it,
  then create a corresponding FILENAME.go with the parsed and macroexpanded
  imports, declarations and statements.

  To parse and execute all *.gomacro files in a directory, run `gomacro -w DIRECTORY`

## Current Status

Fairly complete.

The intepreter supports:
* multiline input
* line comments starting with #! in addition to //
* constant, variable, type and function definitions (including variadic functions)
* primitive types: booleans, integers, floats, complex numbers
* composite types: arrays, channels, maps, pointers, slices, strings, structs
* composite literals
* the empty interface, i.e. interface{} - other interfaces not implemented yet
* channel send and receive
* function and method calls, including multiple return values
* if, for, for-range, break, continue, return (unimplemented: goto)
* select, switch, type switch, fallthrough
* defer, panic and recover
* imports: Go standard packages "just work", importing other packages requires Go 1.8+ and Linux
* switching to a different package
* macro definitions, for example `macro foo(a, b, c interface{}) interface{} { return b }`
* macro calls, for example `foo x; y; z`
* macroexpansion: code walker, MacroExpand and MacroExpand1
* quote and quasiquote. they take any number of arguments in curly braces, for example:
  `quote { x; y; z }`
* unquote and unquote_splice
* nesting macros, quotes and unquotes

Several things are still missing:
* the keyword "go"
* interfaces definition
* methods definition
* labeled statements, goto
* named return values
* ellipsis in array literal
* history/readline (rlwrap does the job in most cases)

## Why it was created

First of all, to experiment with Go :)

Second, to simplify Go code generation tools (keep reading for the gory details)

---

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
