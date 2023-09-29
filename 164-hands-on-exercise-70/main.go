// Package main: Function returns.
package main

import "fmt"

func main() {
	f := outer()
	fmt.Println(f())
}

func outer() func() int {
	return func() int {
		return 42
	}
}

/*
● Create a func
	○ which returns a func
		■ which returns 42
● assign the returned func to a variable
● call the returned func
● print
*/
