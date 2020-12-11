// interrupt_interpreter.go
package main

import (
	"fmt"
	"time"

	"github.com/cosmos72/gomacro/fast"
)

func main() {
	interp := fast.New()

	go stop(interp)

	// you should use interp in the same goroutine where it was created,
	// otherwise interp.Interrupt() may not work
	run(interp)
}

func stop(interp *fast.Interp) {
	fmt.Println("sleeping 3 seconds")
	time.Sleep(3 * time.Second)

	fmt.Println("slept. stopping interpreter..")
	// tell interpreter to exit the infinite loop
	interp.Interrupt(nil)
}

func run(interp *fast.Interp) {
	defer func() {
		p := recover()
		fmt.Printf("interpreter: infinite loop exited with panic = %T(%#v) %v\n", p, p, p)
	}()
	fmt.Println("interpreter: entering infinite loop")

	// this is an infinite loop
	interp.Eval(`
			func main() { for {} }
			main()
	`)
}
