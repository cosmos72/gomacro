// empty file. stops "go build" from complaining that
// no buildable files are in the directory "examples"

package main

func pair(a, b int) Pair { var p Pair; p.A = a; p.B = b; return p }

func main() {
	// pair(1, 2).A = 3
}
