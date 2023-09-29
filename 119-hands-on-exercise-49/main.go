package main

import "fmt"

func main() {
	mssls := map[string][]string{
		`bond_james`:  {`shaken, not stirred`, `martinis`, `fast cars`},
		`money_penny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:       {`cats`, `ice cream`, `sunsets`},
	}
	for k, d := range mssls {
		fmt.Println(k)
		for i, v := range d {
			fmt.Println(i, v)
		}
		fmt.Println("")
	}
}
