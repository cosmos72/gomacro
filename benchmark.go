package main

import (
	"fmt"
	"time"
)

// gcd(659634968758758759, 659634968758758759):
// output: 7, elapsed time: 5e-07 s
// the interpreter is 1600 times slower than compiled code...

func benchmarks() {
	gcdBenchmark(659634968758758759, 875875179346596349)
}

func gcdBenchmark(a, b uint) {
	t1 := time.Now()
	result := gcd(a, b)
	delta := time.Now().Sub(t1)
	fmt.Printf("gcd(%d, %d) = %d, elapsed time: %g s\n", a, b, result, float64(delta)/float64(time.Second))
}

func gcd(a, b uint) uint {
	for a != b && a > 1 && b > 1 {
		for a > b {
			a -= b
		}
		for b > a {
			b -= a
		}
	}
	if a > b {
		return a
	}
	return b
}
