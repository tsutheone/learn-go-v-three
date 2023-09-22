package main

import (
	"fmt"
)

func main() {
	c := 20
	for c >= 0 {
		fmt.Printf("Iteration %v.\n", c)
		c--
	}
}
