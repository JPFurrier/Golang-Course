package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	person1 := person{
		firstName:   "Robson",
		lastName:    "Pereira",
		favIceCream: []string{"Morango", "banana", "pistache"},
	}

	person2 := person{
		firstName:   "Lailson",
		lastName:    "da Rosa",
		favIceCream: []string{"Chocolate", "creme", "napolitano"},
	}

	fmt.Println(person1, person2)
	for i, v := range person1.favIceCream {
		fmt.Println(i, v)
	}
	for i, v := range person2.favIceCream {
		fmt.Println(i, v)
	}

}
