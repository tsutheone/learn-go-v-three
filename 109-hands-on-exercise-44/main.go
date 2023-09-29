package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	/* Easy showing */
	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])

	/*
		// True new slice with new underlying array
		xia := make([]int, 5)
		xib := make([]int, 5)
		xic := make([]int, 5)
		xid := make([]int, 5)

		copy(xia, xi[0:5])
		copy(xib, xi[5:])
		copy(xic, xi[2:7])
		copy(xid, xi[1:6])

		fmt.Printf("%T\t%#v\n", xia, xia)
		fmt.Printf("%T\t%#v\n", xib, xib)
		fmt.Printf("%T\t%#v\n", xic, xic)
		fmt.Printf("%T\t%#v\n", xid, xid)
	*/

}
