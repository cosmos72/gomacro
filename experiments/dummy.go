// empty file. stops "go build" from complaining that
// no buildable files are in the directory "experiments"

package experiments

func foo() {
	m := make(map[string]string)
	var v int = m[9]
}
