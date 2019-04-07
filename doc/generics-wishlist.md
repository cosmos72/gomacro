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

  a single `sort#[T]` function that can sort any slice of any ordered type.\
  a single `cache#[K,V]` type that can cache key/value pairs of any type,
  provided that keys can be compared.\
  a single `sortedmap#[K,V]` type, similar to the existing `map[K]V`
  but keeps its entries sorted, like C++ `map<K,V>` or Java `TreeMap<K,V>`

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
  only once, provided that Go implements a cache of instantiated functions and types,
  similarly to how it implements a cache of compiled packages.\
  Instead C++ `#include` mechanism typically needs to compile again
  the same generic function - say `sort<int>` even if it's instantiated with the same types
  from two different source files - for example `a.cpp` and `b.cpp` both use `sort<int>`.
  C++ compilers typically delegates to the linker the job of coalescing multiple,
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
  becomes more verbose without type inference, because each `var`
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

* constraints:

  when writing a generic function or type, it should be possible to specify constraints
  on their type arguments. This is an extensively discussed topic, for many reasons:

  1) constraints are expected to simplify compiler error messages, and make them
     more understandable. For example, a `sort#[T]` function would specify that values
	 of `T` must be ordered - the following syntax is just for illustration purposes:
	 ```Go
	 func sort#[T: Ordered](slice []T) {
		// ...
	 }
	 ```
	 Then, attempting to sort a non-ordered type as for example `func ()` could produce
	 an error message like `sort: type func() is not Ordered` instead
	 of some deeply nested error due to `a < b` used on `func()` values.

  2) constraints allow programmers writing generic code to specify explicitly
     the requirements of their code, i.e. on which types it can be used and why.

	 Without them, it is not always simple to understand if a complicated generic function
	 or type written by someone else can be used with a certain concrete type `T`,
	 and what are the requirements on such `T`:\
	 the author of generic code could document the requirements, for example in a comment,
	 but he/she may forget it, or the comment could become stale/erroneous if the generic
	 code gets updated.

	 A machine-readable, compiled information is less likely to become stale/erroneous,
	 especially if the compiler actually validates it.

  3) if the compiler assumes that constraints specify the **only** operations
     supported by the constrained types, it could detect immediately if a constrained
	 type is used improperly in generic code, without having to wait until it gets
	 instantiated (possibly by someone else) on concrete types - for example if methods
	 or arithmetic operations are used on a type that is only constrained as `T: Ordered`

	 For reference, Haskell does exactly that: a constraint specifies the only operations
	 allowed on a type.\
	 Actually, Haskell does even more: if a constraint for a type `T` is not specified,
	 the compiler infers it from the operations actually performed on `T` values
	 (it's not obvious whether such constraint inference is appropriate for Go).

  It should also be possible to specify multiple constraints on a type.
  For example, if a type `T` must be both `Ordered` and `Printable`,
  one could imagine a syntax like:
  ```Go
  func foo#[T: Ordered, Printable](arg T) {
	// ...
  }
  ```

* constraints implementation:

  An important question is: what should a constraint tell about a type?

  1) The signature of one or more methods?

  2) The signature of one or more functions and/or operators?

  3) The name and type of one or more fields?

  4) A combination of the above?

  It is surely tempting to answer 1. and reuse interfaces as constraints:
  this would spare us from inventing yet another language construct, but is it enough?

  ### Option 1. constraints declare type's methods

  Let's check with a relatively simple case: the `Ordered` constraint.\
  It describes types that can be ordered, and there's immediately a difficulty:
  Go operator `<` only works on basic types (integers and floats), and cannot be overloaded
  i.e. cannot be extended to support further types.
  It will work on types whose underlying type is integer or float, as for example
  ```Go
    package time

    type Duration int64
  ```
  but even in such case you cannot define a custom implementation:\
  operator `<` compares `time.Duration` values as it would compare `int64`.

  So let's say that `Ordered` will instead use a **function** `Less()` to compare values.\
  Here we hit another Go (intentional) limitation: function overloading is not supported,\
  i.e. it's not possible to define multiple functions with the same name and different signatures.

  Ok, then let's say that `Ordered` will use a **method** `Less()` to compare values.\
  How do we express that a type must have a method `Less()` to be `Ordered`?\
  With an interface, of course:
  ```Go
    type Ordered interface {
	  Less(/*what here?*/) bool
    }
  ```
  We are getting close: we need to express that the single argument of `Less()`
  is the same as the receiver. Go does not support this either, but we are trying to
  extend it with generics, and the addition "we can give a name to the receiver type"
  feels quite minimal.\
  What about the following?
  ```Go
    type Ordered#[T] interface {
	  func (T) Less(T) bool
    }
  ```
  It's supposed to mean that `Ordered` is a generic interface, i.e. it's polymorphic,
  and has a single type argument `T`. To satisfy `Ordered`, a type must have a method
  `Less(T)` where `T` is also the receiver type (the `func (T)` part).
  I chose the syntax `func (T) Less ...` because that's exactly how we already declare
  methods, and the shorter `(T) Less ...` did not sound familiar enough.

  There are still a couple of issues.

  First issue: builtin integers and floats do not have any method, so they cannot implement `Ordered`.
  This can only be solved with a Go language specs change which adds methods to basic types.
  On the other hand user-defined types, including standard library ones as `time.Duration`,
  could add a method `Less()`.

  Second issue: methods must be declared in the same package as their receiver.
  In other words, it's not possible to import a type `foo.Bar` and add a method `Less()` to it:
  either the method is already there because the author forecasted the need, or it's not there
  and there's no way to add it (unless you fork the package `foo` and modify it -
  something that should be a last resort, not the normal case).
  This cannot be solved reasonably - at most it can become an intentional limitation.




**TO BE CONTINUED**

# Anti-goals

Things the author does not want from Go generics

* a compile-time sub-language:

  Go generics should be an extension of idiomatic Go, not a whole sub-language
  to be used for compile time operations.

  For example, we should avoid compile-time Turing completeness and "expression templates",
  two accidental features of C++ templates that together created a sub-language of C++
  made of template specializations and recursive templates.\
  Such sub-language also provides arbitrary computation at compile-time (possibly a good thing)
  with a terrible syntax and no alternative with cleaner syntax.\
  The much more recent C++ `constexpr` actually provides the desired alternative, clean syntax
  for compile-time Turing completeness, but it is more limited: it can only manipulate values,
  not types.

# Design space

There are many possible ways to implement generics - one could say **too** many -
and they can be extremely different in usability, expressiveness, implementation complexity,
compile-time performance and run-time performance.

**TO BE CONTINUED**
