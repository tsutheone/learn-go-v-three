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

	fmt.Println("----------------------------------")
	fmt.Println("")

	mssls[`fleming_iam`] = []string{`steaks`, `cigars`, `espionage`}

	for k, d := range mssls {
		fmt.Println(k)
		for i, v := range d {
			fmt.Println(i, v)
		}
		fmt.Println("")
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")

	delete(mssls, `no_dr`)

	for k, d := range mssls {
		fmt.Println(k)
		for i, v := range d {
			fmt.Println(i, v)
		}
		fmt.Println("")
	}

}
