package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	hilux := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheel: true,
	}
	fmt.Println(hilux)

	jetta := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println(jetta)
}
