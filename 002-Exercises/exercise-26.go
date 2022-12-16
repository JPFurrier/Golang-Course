package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I'm", p.age, "years old")

}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   33,
	}

	p2 := person{
		first: "Mary",
		last:  "Huston",
		age:   20,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.speak()
	p2.speak()
}
