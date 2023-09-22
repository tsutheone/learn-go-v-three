package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("outer %v \t inner %v\n", i, j)
		}
		fmt.Printf("\n")
	}
}
