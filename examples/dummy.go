// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
	r "reflect"
	"time"
)

func main() {
	var d int64 = 1
	v := r.ValueOf(&d)
	fmt.Printf("%v <%v>\n", v, v.Type())
	v = v.Convert(r.TypeOf((*time.Duration)(nil)))
	fmt.Printf("%v <%v>\n", v, v.Type())
}
