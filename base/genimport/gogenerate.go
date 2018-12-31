package genimport

import (
	"fmt"
	"os"
	"strings"
)

var sep = string(os.PathSeparator)

// GoGenerateMain allows gomacro to be run under
// go generate. Use a go generate comment like
//
// `//go:generate gomacro -g` or
// `//go:generate gomacro -g github.com/cosmos72/gomacro/classic`
//
// at the top of one of your package's .go files.
// If omitted, the optional import path that
// follows -g defaults to the current directory.
//
// When GoGenerateMain() is called,
// arg may be length 0, in which case we guess at
// the import path using the current working
// directory. If arg[0] does hold a string, it
// must give an import path. The new x_package.go
// bindings are written back to the same package.
func GoGenerateMain(arg []string, imp *Importer) error {
	var pkgpath string
	switch len(arg) {
	case 0:
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("gomacro -g error trying to get current dir: '%v'", err)
		}
		gopath := os.Getenv("GOPATH")
		prefix := gopath + sep + "src" + sep
		if strings.HasPrefix(cwd, prefix) {
			pkgpath = cwd[len(prefix):]
		} else {
			// guess it is after the first `src` in cwd,
			// since traditionally all packages are
			// after $GOPATH/src/
			splt := strings.SplitN(cwd, sep+"src"+sep, 2)
			if len(splt) > 1 {
				pkgpath = splt[1]
			} else {
				return fmt.Errorf(
					"error: under gomacro -g <import path>" +
						" must specify import path" +
						" as we could not guess it.")
			}
		}
	case 1:
		pkgpath = arg[0]
	default:
		return fmt.Errorf("error: extraneous argument(s) after gomacro -g: '%s'",
			strings.Join(arg[1:], " "))
	}
	_, err := imp.ImportPackageOrError("_i", pkgpath)
	return err
}
