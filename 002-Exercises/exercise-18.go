package main

import "fmt"

func main() {
	person := map[string][]string{
		"Bond_James":  []string{"Shaken, not Stirred", "Martinis", "Women"},
		"Money_penny": []string{"James Bond", "Literature", "Computer Science"},
		"Dr_Who":      []string{"Being Evil", "Ice Cream", "Sunset"},
	}
	// fmt.Println(person)

	// adicionando um item ao mapa
	person["Flenningas_Mrs."] = []string{"Steaks", "Cigars", "Espionage"}

	// deletando um item do mapa
	delete(person, "Money_penny")

	for k, v := range person {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
