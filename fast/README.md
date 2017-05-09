## gomacro - A Go interpreter with Lisp-like macros

The package `fast` contains a faster reimplementation of gomacro interpreter.

To learn about gomacro, download, compile and use it, please refer to the original implementation [README.md](../README.md)

If you want to help with the reimplementation, or you are simply curious, read on :)

## Current Status

BETA.

The fast interpreter supports:
* multiline input - shared with the classic interpreter
* line comments starting with #! in addition to // - shared with the classic interpreter
* basic types: booleans, integers, floats, complex numbers, strings (and iota)
* the empty interface, i.e. interface{} - other interfaces not implemented yet
* constant, variable and type declarations (including structs and untyped constants)
* unary and binary operators
* assignment, i.e. operators = += -= *= /= %= &= |= ^= &^= (unimplemented: <<= and >>=)
* composite types: arrays, channels, maps, pointers, slices, strings, structs (unimplemented: composite literals)
* reading and writing struct fields, including embedded fields
* slicing
* type assertions and type conversions
* function declarations, including variadic functions
* closures
* transparent invocation of compiled functions from interpreter, and vice-versa
* function calls, including calls to variadic functions, with or without ... after last call argument
* if, for, switch, break, continue, fallthrough, return (unimplemented: for range, type switch, goto, return foo() where foo() returns multiple values)
* go i.e. goroutines
* all builtins except defer and recover
* ~quote

Missing features - you are welcome to contribute:
* composite literals
* naked return, i.e. return without arguments in a function with named results
* return foo() where foo() returns multiple values
* bar(foo()) where foo() returns multiple values
* x,y = foo() where foo() returns multiple values
* shifted assignment, i.e. operators <<= and >>=
* for range (normal for is implemented)
* type switch
* method definition
* method call
* interface definition
* conversion from/to interpreted interfaces
* defer and recover
* macro definition
* ~quasiquote, ~unquote, ~unquote_splice

Current limitations:
* no distinction between named and unnamed types created by interpreted code.
  For the interpreter, `struct { A, B int }` and `type Pair struct { A, B int }`
  are exactly the same type. This has subtle consequences, including the risk
  that two different packages define the same type and overwrite each other's methods.

  The reason for such limitation is simple: the interpreter uses `reflect.StructOf()`
  to define new types, which can only create unnamed types.
  The interpreter then defines named types as aliases for the underlying unnamed types.

* cannot create recursive types, as for example `type List struct { First interface{}; Rest *List}`
  The reason is the same as above: the interpreter uses `reflect.StructOf()` to define new types,
  which cannot create recursive types

* operators << and >> on untyped constants do not follow the exact type deduction rules.
  The implemented behavior is:
  * an untyped constant shifted by a non-constant expression always returns an int
  * an untyped floating point constant shifted by a constant expression returns an untyped integer constant.
    the interpreter signals an error during the precompile phase
    if the left operand has a non-zero fractional or imaginary part,
    or it overflows both int64 and uint64.
  See [Go Language Specification](https://golang.org/ref/spec#Operators) for the correct behavior

Misc TODO notes:
* when importing a package, reuse compiled .so if exists already
* gomacro FILE: execute main() if (re)defined and package == "main"
* try to run Go compiler tests
