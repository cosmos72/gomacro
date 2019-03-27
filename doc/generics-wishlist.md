Generics in Go
==============

a wishlist
----------

This file contains experiences and desiderata for a Go implementation of generics,
i.e. polymorphic types and functions.

It is a feature present in many other languages with varying names. A few examples:
* C++ [templates](https://www.geeksforgeeks.org/templates-cpp)
* Java [generics](https://en.wikipedia.org/wiki/Generics_in_Java)
* Haskell [generics](https://wiki.haskell.org/Generics)

The author has experience using generics in the three languages listed above,
which will also be used for comparison and reference in the rest of this document.

In addition, the author personally added generics to three programming languages:
* Go: the unofficial interpreter [gomacro](https://github.com/cosmos72/gomacro)
  contains a Go implementation of generics, modeled after C++ templates.
* Common Lisp: the library [cl-parametric-types](https://github.com/cosmos72/cl-parametric-types)
  contains a Common Lisp implementation of generics, again modeled after C++ templates.
* the [lfyre](https://sourceforge.net/projects/lfyre) programming language,
  created by the author, used to contain an implementation of generics.
  It now has a different maintainer.

# Goals

The reasons to implement generics in Go can be many, and sometimes contradicting.
The author's personal list of reasons, which can also be interpreted as goals
that Go generics are expected to achieve, are:

* reusable, flexible algorithms and types. Examples:

  a single `sort` function that can sort any slice of any ordered type.

  a single `cache` type that can cache key/value pairs of any type,
  provided that keys can be compared.

* type-safety:

  generic functions and types should be instantiable on arbitrary,
  concrete types - for example `sort#[int]` would only accept `[]int` slices
  and `cache#[uint64, []byte]` would only accept `uint64` keys and `[]byte` values.
  In particular, generic functions and types should not need to use `interface{}`,
  either internally or in they exported API, and should not need type assertions at runtime.

* high runtime speed, low runtime overhead:

  generic functions and types should be reified in order to maximize code execution speed
  and have low or zero data representation overhead in memory.

  Reified means that `sort#[int]` and `sort#[uint]` will be two different and unrelated functions,
  one only dealing with `int` slices and the other only dealing with `uint` slices,
  and that `cache#[uint64, []byte]` and `cache#[uint32, []byte]` will be two different
  and unrelated types, with (possibly) different layout in memory.

  While reified generics also have disadvantages (see for example
  https://gbracha.blogspot.com/2018/10/reified-generics-search-for-cure.html)
  the author has extensive experience with both reified generics (C++, Haskell)
  and non reified generics (Java), and he is convinced that reified generics
  are a better fit for Go - the reasons can be explained if needed.

  One obvious disadvantage of reified generics is that each instantiation
  of a generic function must be compiled separately, for example `sort#[int]`
  and `sort#[uint]`, increasing build time.

  Luckily, Go `import`s compiled packages instead of `#include`-ing their source code,
  which is expected to contain build time for two reasons:

  1. each generic function will be parsed only once. Instead C++ `#include` mechanism
  typically needs to parse again the same generic function each time it is included
  by a different source file.

  2. each instantiation of a generic function - say `sort#[int]` - will be compiled
  only once, provided that Go implements a cache of instantiated functions, similarly
  to how it implements a cache of compiled packages.\
  Instead C++ `#include` mechanism typically needs to compile again
  the same generic function - say `sort<int>` even if it's instantiated with the same types
  from a different source file - for example `a.cpp` and `b.cpp` both use `sort<int>`.
  C++ typically delegates to the linker the job of coalescing multiple,
  identical versions of the same generic function.

* reasonable build time:

  it is expected to be achieved / achievable even with reified generics, see the previous item

* type inference:

  Go extensively uses (and encourages to use) type inference instead
  of explicitly declaring the type of a variable.\
  Example: `a := foo()` rather than `var a int = foo()`.

  When an expression returns multiple values,
  Go actively pushes the programmer to use type inference. Example:
  ```
  n, err := fmt.Println("foo")
  ```
  is more verbose without type inference, because each `var`
  declaration can only reference one type:
  ```
  var n int
  var err error
  n, err = fmt.Println("foo")
  ```

  The goal for generics is to preserve and extend support for type inference,
  for example by allowing the syntax
  ```
  slice := make([]int, n)
  sort(slice)
  ```
  and automatically inferring that it means
  ```
  slice := make([]int, n)
  sort#[int](slice)
  ```

**TO BE CONTINUED**

# Design space

There are many possible ways to implement generics - one could say **too** many -
and they can be extremely different in usability, expressiveness, implementation complexity,
compile-time performance and run-time performance.

**TO BE CONTINUED**
