package cmd_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/cosmos72/gomacro/base"
	. "github.com/cosmos72/gomacro/fast"
	"github.com/kylelemons/godebug/diff"
)

func TestWriteDeclsAndStmts(t *testing.T) {
	testdata, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range testdata {
		if !strings.HasSuffix(file.Name(), ".gomacro") {
			continue
		}
		test := strings.TrimSuffix(file.Name(), ".gomacro")
		t.Run(test, func(t *testing.T) {
			gomacroFile := filepath.Join("testdata", test+".gomacro")
			goFile := filepath.Join("testdata", test+".go")
			errFile := filepath.Join("testdata", test+".err")
			ir := New()
			g := &ir.Comp.Globals
			g.Options |= OptCollectDeclarations | OptCollectStatements

			comments, got := ir.EvalFile(gomacroFile)

			if os.Getenv("UPDATE_SNAPSHOTS") != "" {
				if got != nil {
					if err := ioutil.WriteFile(errFile, []byte(got.Error()), 0644); err != nil {
						t.Fatal(err)
					}
				}

				g.WriteDeclsToFile(goFile, comments)

				t.Fail() // Test not performed
			} else {
				buf, err := ioutil.ReadFile(errFile)
				if os.IsNotExist(err) {
					if got != nil {
						t.Errorf("EvalFile returned an error: %q", got)
					}
				} else if err == nil {
					if got == nil {
						t.Errorf("EvalFile didn't return an expected error:\nwant %q", string(buf))
					} else if string(buf) != got.Error() {
						t.Errorf("EvalFile returned an unexpected error:\nwant %q;\n got %q", string(buf), got)
					}
				} else {
					t.Fatal(err)
				}

				if buf, err = ioutil.ReadFile(goFile); err != nil {
					if os.IsNotExist(err) {
						t.Fatalf("%s, try setting UPDATE_SNAPSHOTS=on", err)
					}
					t.Fatal(err)
				}
				out := &strings.Builder{}
				out.WriteString(comments)
				g.WriteDeclsToStream(out)
				if diff := diff.Diff(string(buf), out.String()); diff != "" {
					t.Errorf("EvalFile produced unexpected AST (lines starting with '-' are missing, '+' lines are superfluous):\n%s", diff)
				}
			}
		})
	}
}
