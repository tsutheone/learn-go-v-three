// Package main: Function closure.
package main

import (
	"fmt"
	"math"
)

func main() {

	// Incrementor block
	f1 := encloseInc()
	fmt.Println(f1()) // 1
	fmt.Println(f1()) // 2
	fmt.Println(f1()) // 3
	fmt.Println(f1()) // 4
	fmt.Println(f1()) // 5

	fmt.Println("")

	// Pow block
	f2 := enclosePow(2)
	fmt.Println(f2()) // 2
	fmt.Println(f2()) // 4
	fmt.Println(f2()) // 8
	fmt.Println(f2()) // 16
	fmt.Println(f2()) // 32
	fmt.Println(f2()) // 64

	fmt.Println("")

	// Pow block
	f3 := enclosePowProfSol(2)
	fmt.Println(f3()) // 2
	fmt.Println(f3()) // 4
	fmt.Println(f3()) // 8
	fmt.Println(f3()) // 16
	fmt.Println(f3()) // 32
	fmt.Println(f3()) // 64
	fmt.Println(f3()) // 128
	fmt.Println(f3()) // 256

}

// Incrementor function using closure.
func encloseInc() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Pow function using closure (pow should be played using always float64).
func enclosePow(n float64) func() float64 {
	x := 1
	return func() float64 {
		r := 1.0
		for i := 0; i < x; i++ {
			r *= n
		}
		x++
		return r
	}
}

// Profesor solution uses math.Pow ¬¬.
func enclosePowProfSol(n float64) func() float64 {
	c := 0.0
	return func() float64 {
		c++
		return math.Pow(n, c)
	}
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable.
*/
