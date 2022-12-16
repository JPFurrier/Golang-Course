package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	person_1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	person_2 := person{
		first: "Lord",
		last:  "Palpatine",
		age:   166,
	}

	agent_1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   33,
		},
		ltk: true,
	}
	agent_2 := secretAgent{
		person: person_2,
		ltk:    true,
	}
	fmt.Println(person_1, person_2)
	fmt.Println(agent_1)
	fmt.Println(agent_1.ltk)
	fmt.Println(agent_1.first)
	fmt.Println(agent_1.last)
	fmt.Println(agent_1.age)

	fmt.Println(agent_2)
	fmt.Println(agent_2.ltk)
	fmt.Println(agent_2.first)
	fmt.Println(agent_2.last)
	fmt.Println(agent_2.age)

	// anonymous structs

	person_3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Jaime",
		last:  "da Silva",
		age:   14,
	}
	fmt.Println(person_3)
}
