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

	ma := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range ma {
		fmt.Println(v)
		for _, d := range v.favIC {
			fmt.Println(v.first, v.last, d)
		}
	}

}
