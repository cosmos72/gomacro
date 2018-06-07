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
