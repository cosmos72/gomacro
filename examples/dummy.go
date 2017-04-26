// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

import (
	"fmt"
)

func main() {
	mi := make(map[rune]byte)
	mi['@'] += 2
	fmt.Println(mi)
}
