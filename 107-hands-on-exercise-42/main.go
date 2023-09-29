package main

import "fmt"

func main() {
	ai := [5]int{}
	for i := 0; i < len(ai); i++ {
		ai[i] = i
	}
	for i, v := range ai {
		fmt.Printf("%v - %v %T\n", i, v, v)
	}

}
