package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	maker string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{electric: true},
		maker:  "Tesla",
		model:  "Cybertruck",
		doors:  4,
		color:  "Grey",
	}
	v2 := vehicle{
		engine: engine{electric: false},
		maker:  "Ferrari",
		model:  "Murcielago",
		doors:  4,
		color:  "Red",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.maker, v1.model, v1.doors, v1.color)
	fmt.Println(v2.electric, v2.maker, v2.model, v2.doors, v2.color)

}
