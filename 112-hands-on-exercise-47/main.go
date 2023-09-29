package main

import "fmt"

func main() {
	xs := make([]string, 0, 50)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	xs = append(xs,
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`,
		` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`,
		` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`,
		` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
		` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`,
		` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	for i := 0; i < len(xs); i++ {
		fmt.Printf("%v -%v\n", i, xs[i])
	}

	fmt.Println("")
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
}
