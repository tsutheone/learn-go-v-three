// Package main: Function expression.
package main

import "fmt"

func main() {
	a1 := func() {
		fmt.Println("I am an anonymous function too!")
	}
	a1()
}
