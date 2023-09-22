package main

import (
	"fmt"
)

func main() {
	c := 42
	for {
		if c < 20 {
			break
		}
		fmt.Printf("%v \n", c)
		c--
	}
}
