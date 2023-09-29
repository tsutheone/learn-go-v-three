package main

import "fmt"

func main() {

	mxi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(foo(mxi...))
	fmt.Println(bar(mxi))
}

func foo(ii ...int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func bar(xi []int) int {
	t := 0
	for _, v := range xi {
		t += v
	}
	return t
}

/*
● create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/
