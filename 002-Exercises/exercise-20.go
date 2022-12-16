package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	person_1 := person{
		firstName:   "Robson",
		lastName:    "Pereira",
		favIceCream: []string{"Morango", "banana", "pistache"},
	}

	person_2 := person{
		firstName:   "Lailson",
		lastName:    "da Rosa",
		favIceCream: []string{"Chocolate", "creme", "napolitano"},
	}

	m := map[string]person{
		person_1.lastName: person_1,
		person_2.lastName: person_2,
	}
	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)

		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("-----------")

	}
}
