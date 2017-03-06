package main

import (
	"fmt"
	"time"
)

// fibonacci(10):
// compiled: 832040, elapsed time: 0.00285 s
// interpreted: 832040 <uint> // eval time 3.371358 s
// the interpreter is 1180 times slower than compiled code...

func benchmarks() {
	// fibonacciBenchmark(30)
}

func fibonacciBenchmark(n uint) {
	t1 := time.Now()
	result := fibonacci(n)
	delta := time.Now().Sub(t1)
	fmt.Printf("fibonacci(%d) = %d, elapsed time: %g s\n", n, result, float64(delta)/float64(time.Second))
}

// func fibonacci(n uint) uint { if n <= 2 { return 1 }; return fibonacci(n-1) + fibonacci(n-2) }
func fibonacci(n uint) uint {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
