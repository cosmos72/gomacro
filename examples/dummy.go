// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	"go/importer"
	"go/token"
	"go/types"
	r "reflect"

	"golang.org/x/tools/go/types/typeutil"
)

func main() {
	imp := importer.Default()

	for _, path := range []string{"archive/tar", "archive/zip", "bufio", "bytes", "compress/bzip2", "compress/flate", "compress/gzip", "compress/lzw", "compress/zlib", "container/heap", "container/list", "container/ring", "context", "crypto", "crypto/aes", "crypto/cipher", "crypto/des", "crypto/dsa", "crypto/ecdsa", "crypto/elliptic", "crypto/hmac", "crypto/md5", "crypto/rand", "crypto/rc4", "crypto/rsa", "crypto/sha1", "crypto/sha256", "crypto/sha512", "crypto/subtle", "crypto/tls", "crypto/x509", "crypto/x509/pkix", "database/sql", "database/sql/driver", "debug/dwarf", "debug/elf", "debug/gosym", "debug/macho", "debug/pe", "debug/plan9obj", "encoding", "encoding/ascii85", "encoding/asn1", "encoding/base32", "encoding/base64", "encoding/binary", "encoding/csv", "encoding/gob", "encoding/hex", "encoding/json", "encoding/pem", "encoding/xml", "errors", "expvar", "flag", "fmt", "go/ast", "go/build", "go/constant", "go/doc", "go/format", "go/importer", "go/parser", "go/printer", "go/scanner", "go/token", "go/types", "hash", "hash/adler32", "hash/crc32", "hash/crc64", "hash/fnv", "html", "html/template", "image", "image/color", "image/color/palette", "image/draw", "image/gif", "image/jpeg", "image/png", "index/suffixarray", "io", "io/ioutil", "log", "log/syslog", "math", "math/big", "math/cmplx", "math/rand", "mime", "mime/multipart", "mime/quotedprintable", "net", "net/http", "net/http/cgi", "net/http/cookiejar", "net/http/fcgi", "net/http/httptest", "net/http/httptrace", "net/http/httputil", "net/http/pprof", "net/mail", "net/rpc", "net/rpc/jsonrpc", "net/smtp", "net/textproto", "net/url", "os", "os/exec", "os/signal", "os/user", "path", "path/filepath", "plugin", "reflect", "regexp", "regexp/syntax", "runtime", "runtime/debug", "runtime/pprof", "runtime/trace", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "testing", "testing/iotest", "testing/quick", "text/scanner", "text/tabwriter", "text/template", "text/template/parse", "time", "unicode", "unicode/utf16", "unicode/utf8", "unsafe"} {
		pkg, err := imp.Import(path)
		if err != nil {
			fmt.Printf("error importing package %q: %v\n", path, err)
			continue
		}
		inspectScope(pkg.Scope())
	}
}

type method struct {
	mtd   *types.Func
	index []int
}

func (m method) String() string {
	return fmt.Sprintf("%v\t%v", m.index, m.mtd)
}

type analyzer map[string][]method

func concat(list []int, tail int) []int {
	list = append([]int{}, list...) // make a copy
	return append(list, tail)
}

func inspectScope(scope *types.Scope) {
	for _, name := range scope.Names() {
		if obj, ok := scope.Lookup(name).(*types.TypeName); ok && obj.Exported() {
			if t, ok := obj.Type().(*types.Named); ok {
				if _, ok := t.Underlying().(*types.Struct); ok {
					a := new(analyzer)
					a.inspect(t, nil)
					a.show(t)
				}
			}
		}
	}
}

func (a *analyzer) add(mtd *types.Func, index []int) {
	if *a == nil {
		*a = make(map[string][]method)
	}
	name := mtd.Name()
	(*a)[name] = append((*a)[name], method{mtd, index})
}

func (a *analyzer) inspect(t *types.Named, index []int) {

	for i, n := 0, t.NumMethods(); i < n; i++ {
		mtd := t.Method(i)
		if mtd.Exported() {
			a.add(mtd, index)
		}
	}

	if u, ok := t.Underlying().(*types.Struct); ok {
		for i, n := 0, u.NumFields(); i < n; i++ {
			if f := u.Field(i); f.Anonymous() {
				switch ft := f.Type().(type) {
				case *types.Named:
					a.inspect(ft, concat(index, i))
				case *types.Pointer:
					if ft, ok := ft.Elem().(*types.Named); ok {
						a.inspect(ft, concat(index, i))
					}
				}
			}
		}
	}
}

func (a *analyzer) show(t *types.Named) {
	for _, list := range *a {
		if len(list[0].index) == 0 {
			// one method declared in the outermost type. no ambiguity
			continue
		}
		fmt.Printf("type %s\n", t)
		for _, m := range list {
			fmt.Printf("%s\n", m)
		}
	}
}

func main4() {
	type Pair struct{ A, B int }
	type Triple struct {
		Pair
		C int
	}
	var t Triple
	fmt.Printf("%v\n", t)
	fmt.Printf("%#v\n", t)
}

func add(a, b int) int {
	return a + b
}

func curry(f func(a, b int) int, a int) func(b int) int {
	return func(b int) int {
		return f(a, b)
	}
}

func reflectCurry(f r.Value, a r.Value) func([]r.Value) []r.Value {
	return func(args []r.Value) []r.Value {
		args = append([]r.Value{a}, args...)
		return f.Call(args)
	}
}

func main3() {
	type EmbedBasicType struct {
		int
	}
	var e EmbedBasicType
	fmt.Println(e.int)
	v := r.ValueOf(e)
	vi := v.Field(0)
	vi.Interface()
}

func main2() {
	interface1 := types.NewInterface(nil, nil).Complete()
	fmt.Printf("%v\n", interface1)

	pkg := types.NewPackage("main", "main")
	ret := types.NewVar(token.NoPos, pkg, "", types.Typ[types.Int])
	sig := types.NewSignature(nil, types.NewTuple(), types.NewTuple(ret), false)
	cap := types.NewFunc(token.NoPos, pkg, "Cap", sig)
	len := types.NewFunc(token.NoPos, pkg, "Len", sig)
	interface2 := types.NewInterface([]*types.Func{cap, len}, nil).Complete()
	fmt.Printf("%v\n", interface2)

	var m typeutil.Map
	m.Set(interface1, interface1)
	m.Set(interface2, interface2)
	fmt.Printf("%v\n", m.At(interface1))
	fmt.Printf("%v\n", m.At(interface2))
	fmt.Printf("%v\n", m.Len())
}

func main1() {
	type Cons struct {
		A, B int
	}
	p := Cons{1, 2}
	var i interface{} = p
	q := i.(Cons)
	q.A = 8
	fmt.Printf("%#v\n", p)
}

func main0() {
	add1 := curry(add, 1)
	fmt.Printf("add1 = %T\n", add1)
	fmt.Printf("add1(5) = %v\n", add1(5))

	add1refl := reflectCurry(r.ValueOf(add), r.ValueOf(1))
	fmt.Printf("add1refl = %T\n", add1refl)
	fmt.Printf("add1refl(5) = %v\n", add1refl([]r.Value{r.ValueOf(5)})[0])
}
