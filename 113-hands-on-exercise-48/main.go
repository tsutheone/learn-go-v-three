package main

import "fmt"

func main() {

	xs := [][]string{{"James", "Bond", "Shaken, not stirred."}, {"Miss", "Moneypenny", "I'm 008."}}
	for i, r := range xs {
		fmt.Println(i, r)
		for i, d := range r {
			fmt.Println(i, d)
		}
	}
}
