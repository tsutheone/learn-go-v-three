package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "Bjork",
		last:  "Magnanum",
		favIC: []string{`chocolate`, `banana`, `strawberry`},
	}
	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{`lemon`, `vanilla`, `mango`},
	}
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	for _, v := range p1.favIC {
		fmt.Println(p1.first, "favorite is", v)
	}
	for _, v := range p2.favIC {
		fmt.Println(p2.first, "favorite is", v)
	}
}
