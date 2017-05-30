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
* constant, variable and type declarations (including untyped constants)
* Go 1.9 type aliases (experimental)
* unary and binary operators
* assignment, i.e. operators := = += -= *= /= %= &= |= ^= <<= >>= &^=
* composite types: arrays, channels, maps, pointers, slices, strings, structs (including composite literals)
* accessing struct fields, including embedded fields
* slicing
* type assertions and type conversions
* interface declarations (**only** declarations. interfaces cannot be implemented or used yet)
* importing and using interfaces declared (and implemented) by compiled code
* function declarations and calls, including variadic functions
* method declarations and calls, including wrapper methods for embedded fields
* closures
* Type.Method i.e. converting methods to functions (examples: time.Duration.Hours, fmt.Stringer.String)
* seamless invocation of compiled functions from interpreter, and vice-versa
* if, for, switch, break, continue, fallthrough, return (unimplemented: for-range, select, type switch, goto)
* all builtins except defer and recover
* go i.e. goroutines
* imports
  * Go standard packages "just work"
  * importing other packages requires the "plugin" package (available only on Linux with Go 1.8+)
* ~quote

Missing features - you are welcome to contribute:
* for-range (normal 'for' is implemented)
* select
* type switch
* switch is implemented but not (yet) thread-safe
* interfaces. They can be declared, but nothing more: there is no way to implement them or call their methods
  (interfaces declared in compiled code can be used, but not yet implemented by interpreted code)
* conversion from/to interpreted interfaces
* defer and recover
* macro definition
* ~quasiquote, ~unquote, ~unquote_splice

Current limitations:
* named types declared by interpreted code are approximated.
  Inside the interpreter they look and behave correctly, but if you pass them to compiled code,
  the type is actually unnamed.

  For example, if interpreted code declares `type Pair struct { A, B int }`,
  then passes a `Pair` to compiled code, it will be received as `struct { A, B int }`

  The reason for such limitation is simple: the function `reflect.NamedOf()` does not exist,
  so the interpreter uses `reflect.StructOf()` to define new types,
  which can only create unnamed types.

* recursive types declared by interpreted code are approximated.
  Inside the interpreter they look and behave correctly, but if you pass them to compiled code,
  the type is unnamed (as above) and self-references are actually interface{}.

  For example, if interpreted code declares `type List struct { First int; Rest *List }`
  then passes a `List` to compiled code, it will be received as `struct { First int; Rest *interface{} }`

  The reason is the same as above: the interpreter uses `reflect.StructOf()` to define new types,
  which cannot create recursive types.

  Interestingly, this means the interpreter also accepts the following declaration,
  which is rejected by Go compiler: `type List2 struct { First int; Rest List2 }`
  Note that `Rest` is a `List2` **not** a pointer to `List2`

* interfaces declared by interpreted code are emulated.
  Inside the interpreter they look and behave correctly, but if you pass them to compiled code,
  the type is actually a pointer to a struct containing a header and a lot of functions.

  The reason is: the function `reflect.InterfaceOf()` does not exist,
  so the interpreter has to emulate interfaces with `reflect.StructOf()` and a lot of bookkeeping

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
