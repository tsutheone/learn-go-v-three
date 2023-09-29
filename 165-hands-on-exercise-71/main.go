// Package main: Callbacks.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(cb(square, 42))
	fmt.Println(profSolution(square, 42))
}

func square(n int) int {
	return n * n
}

func cb(f func(int) int, i int) string {
	return strconv.Itoa(f(i))
}

// Using Sprintf which returns and string, we don't need strconv.Itoa function.
func profSolution(f func(int) int, i int) string {
	x := f(i)
	return fmt.Sprintf("The square of %v is %v", i, x)
}

/*
● pass a func into a func as an argument
	○ func square(n int) int
	○ printSquare(f func(int)int, int) string
*/
