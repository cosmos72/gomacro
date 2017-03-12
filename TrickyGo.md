A collection of tricky go code

```
true := false
println(true)
```

```
func uint(x int) int { return x + 7 }
println(uint(1))
```

```
import "os"
func getGoPath() string {
	dir := os.Getenv("GOPATH")
	if len(dir) == 0 {
		dir := os.Getenv("HOME") // shadows outer "dir", does NOT modify it
		if len(dir) == 0 {
			panic("cannot determine go source directory: both $GOPATH and $HOME are unset or empty")
		}
		dir += "/go"
	}
	return dir // inner "dir" is not seen -> always returns os.Getenv("GOPATH")
}
