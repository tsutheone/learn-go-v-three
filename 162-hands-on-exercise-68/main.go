// Package main: Anonymous functions.
package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm anonymous function.")
	}()
}
