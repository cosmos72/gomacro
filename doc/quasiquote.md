Quasiquote
==========

implementing quasiquote, unquote and unquote_splice in Go
--------------------------------------------------------

One of the main motivations behind the creation of Go interpreter `gomacro`
was to add Lisp-like macros to Go.

This includes implementing Common Lisp `quote`, `quasiquote` and, more crucially,
`unquote` and `unquote_splice` i.e. Common Lisp macro characters `'` `` ` `` `,` and `,@`

Since Go language is not homoiconic, i.e. (source) code and (program) data
are not represented identically, this is a challenge.

### Parser ###

The first (moderate) difficulty is adding support for `'` `` ` `` `,` and `,@` to Go parser.
It was solved by forking Go standard packages https://golang.org/pkg/go/scanner/
and https://golang.org/pkg/go/parser/ and patching them.

Actually, the characters `'` `` ` `` and `,` are already reserved in Go,
so the author decided to replace them as follows:
* quote          `'`  becomes `~'`
* quasiquote `` ` ``  becomes `~"` not ``~` `` because the latter messes up syntax hilighting in Go-aware editors and IDEs (starts a multiline raw string)
* unquote        `,`  becomes `~,`
* unquote_splice `,@` becomes `~,@`

and the prefix `~` is configurabile when manually instantiating the modified parser.

Go parser produces as output an abstract syntax tree (AST) represented as a tree of `ast.Node`,
from the standard package https://golang.org/pkg/go/ast/

Defining new node types is deliberately impossible (`ast.Node` is an interface with unexported methods),
luckily the existing types are flexible enough to accommodate the new syntax.

The chosen representation is somewhat ugly but effective:
newly created constants `token.QUOTE`, `token.QUASIQUOTE`, `token.UNQUOTE` and `token.UNQUOTE_SPLICE`
are used as unary operators of a fictitious closure containing the quoted code. Examples:
* `'x` must be written `~'x` and is parsed as if written `~' func() { x }`
* `` `{x = y}`` must be written `~"{x = y}` and is parsed as if written `~" func() { x = y }`
* `,{1 + 2}`  must be written `~,{1 + 2}` and is parsed as if written `~, func() { 1 + 2 }`
* `,@{f(a, b)}` must be written `~,@{f(a, b)}` and is parsed as if written `~,@ func() { f(a, b) }`

The fictitious closures are necessary because `ast.UnaryExpr` only allows an expression
as its operand - not arbitrary statements or declarations.
In Go, the only expression that can contain arbitrary statements and declarations is a closure (in Go terms, a "function literal")


