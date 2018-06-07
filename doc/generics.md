Generics
========

implementing generics Go
------------------------

The branch generics-v1 is an experiment to add generics to gomacro.

This file contains observations, choices, difficulties and solutions found
during such work.

Due to the author's personal taste and his familiarity with C++, generics are
named 'templates' in this document.

### Parser ###

#### Declaring templates ####

Adding a new syntax to declare template types, function and methods is easy:
it's just a matter of inventing the syntax and choosing a representation in
terms of go/ast.Nodes

Current syntax is:
```
template [T1, T2...] type ...
template [T1, T2...] func ...
```

Template type declarations are represented with *ast.TypeSpec as usual,
with the difference that TypeSpec.Type now contains
`&ast.CompositeLit{Type: <type declaration>, Elts: [T1, T2 ...]}`

Template function and method declarations are represented with *ast.FuncDecl
as usual, with the difference that len(FuncDecl.Recv.List) > 1:
the first receiver is nil for functions and non-nil for methods,
further receivers are &ast.Field{Names: nil, Type: <i-th template arg>}

#### Using templates ####

The main idea is that template functions and methods will be used mostly
in the same ways non-template ones, i.e. `Func(args)` and `obj.Method(args)`
exploiting appropriate type inference (exact inference rules need to be defined).

In some cases, using template functions and methods will need to specify
the exact template arguments. Template types will need such explicit
qualification most of (or maybe all) the time.

For example, after a declaration
```
template [T1, T2] type Pair struct { First T1; Second T2 }
```
it is tempting to say that the syntax to specify the template arguments
(= to qualify the template name) is
```
Pair[int, string]
```
i.e. the template name is immediately followed by '[' and the comma-separated
list of template arguments.

Alas, such syntax is too ambiguous for current Go parser. Take for example the
code fragment
```
func Nop(Pair[int, int]) { }
```
By manual inspection, it's clear that `Pair` is a type name, not a parameter
name. But compare the fragment above with this:
```
func Nop(Pair []int) { }
```
where `Pair` is a parameter name with type `[]int`.

In both cases, the parser will encounter `Pair` followed by `[` and must
decide how to parse them without further look-ahead.

The current parser algorithm for this case assumes that `Pair` is an
identifier and that `[` starts a type expression to be parsed.

To avoid breaking lots of existing code, the current parser algorithm for
this case must be preserved. So we need a different, less ambiguous syntax to
qualify template names.

One of the suggestions in latest Ian Lance Taylor
[Type parameters (December 2013)](https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md)
proposal is "using double square brackets, as in `Vector[[int]]`, or perhaps
some other character(s)."

The authors' current decision - but it's trivial to change it - is to write
`Pair#[int, string]` and similarly `Vector#[int]`. The reason is twofold:

1. double square brackets look too "magical"
2. the hash character `#` is currently not used in Go syntax, and usually does not have
   strong connotations in programmers' minds. The alternatives are the other
   ASCII characters currently not used in Go syntax: `?` `@` `$` `~`
   * the question mark `?` is better suited for conditionals, as for example
     the C ternary operator `?:`
   * the at sign `@` already has several common meanings (at, email...).
   * the dollar sign `$` seems inappropriate, in the author's opinion, for
     this usage.
   * the tilde sign `~` is already used by gomacro for quasiquote and friends.

Implementation choice: `Pair#[int, string]` is represented as
```
&ast.IndexExpr{X: Pair, Index: &ast.CompositeLit{Elts: [T1, T2...]} }
```
The simpler `&ast.CompositeLit{Type: Pair, Elts: [T1, T2...]} }` would suffice
for the parser, but compiling it is much more ambiguous, since it could be
interpreted as the composite literal `Pair{T1, T2}`

#### Composite Literals ####

The parser had to be extended to recognize things like `Pair#[T1,T2] {}`
as a valid composite literal.

In practice, `isTypeName()` and `isLiteralType()` now return true for *ast.IndexExpr.

This solution should be better examined or tested to check whether the increased
syntax ambiguity is a problem, but an official implementation will surely create
new ast.Node types to hold template declarations and template uses, bypassing
this potential problem.

### Declaration ###

The declaration of template types and functions is straightforward.

For each template declaration found, the compiler now collects it in the same
`map[string]*Bind` used for const, var and func declarations.

Such declarations store in the *Bind their **source code** as an ast.Node, in order to
retrieve and compile it on-demand when the template type or function needs to be
instantiated.

This is easy for an interpreter, but more tricky for a compiler:
since a package A may use a template B.C declared in package B,
the compiler may need to instantiate B.C while compiling A.

There are at least two solutions:
1. for each compiled package, store in the compiled archive packagename.a
   the **source code** of each template, alongside with binary code.

   This may not play well with commercial, binary-only libraries since,
   with a little effort, the source code of templates could be extracted.

2. for each compiled package, require its source code to be available
   in order to instantiate its templates.

   This has the same problem as above, only in stronger form:
   the **full** source code of B must be available when compiling A.

Another question is: where to store the B.C instantiated while compiling A ?

For templates declared in the standard library and instantiated by non-root users,
$GOROOT may not be writeable, so it should probably be stored in
$GOPATH/pkg/$GOOS_$GOARCH/path/to/package, using a name like B.<somehash>.a

### Instantiation ###

Instantiantion is a regular compile, with some careful setup.

Since a template may access global symbols in the scope where it was declared,
it must be compiled in that **same** scope. Better yet, it can be compiled
in a new inner scope, that defines the template arguments to use for instantiation.

An example can help to understand the abstract sentence above:
suppose package B contains
```
package B

const N = 10

template[T] type Array [N]T
```

and is later used by package A as
```
package A

import "B"

var arr B.Array#[int]
```

the technique described abstractly above means: to compile `B.Array#[int]`,
pretend that package B contains the following (pseudo-code, it's not valid Go):
```
{ // open new scope

	type T = int // inject the concrete template argument

	// inject the template declaration literally - no replacements needed
	type Array [N]T // finds T immediately above, and N in outer scope
}
```

There are two small issues with this approach:

1. the name `Array` may conflict with the template parameters names - in this case `T`.
   Solutions: `Array` name may be replaced/augmented with a generated hash, or some
   inner compiler function could be used that compiles an **anonymous** type with
   underlying type `[N]T`.
   The best solution also depends on the desired behaviour of reflect.Type.Name():
   what name should it return for a type `Array#[int]` ?

2. compiling recursive template functions or types requires some additional trickery
