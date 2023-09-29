package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range xi {
		fmt.Printf("%v - %v %T\n", i, v, v)
	}
	// optional for comparison only
	fmt.Printf("%T \t %#v\n", xi, xi)
	fmt.Printf("%T \t %v\n", xi, xi)

}
