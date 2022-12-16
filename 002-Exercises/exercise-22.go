package main

import "fmt"

// anonymous struct

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jamil":     3245,
			"Carlinhos": 1123,
			"Jhonson":   6677,
		},
		favDrinks: []string{
			"martini",
			"water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
