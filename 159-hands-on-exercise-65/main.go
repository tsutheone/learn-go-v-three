/* Unit Test (3) */
package main

import (
	"errors"
	"fmt"
)

func main() {
	cxr := []string{"Blue", "Yellow", "Pink", "Black"}
	cxr, err := rangerAdd("White", cxr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cxr)
}

func rangerAdd(rname string, xs []string) ([]string, error) {
	cl := len(xs)
	xs = append(xs, rname)
	l := len(xs)
	if l != cl+1 {
		return xs, errors.New("rangerAdd: length has not incremented 1, append failed")
	} else {
		return xs, nil
	}
}
