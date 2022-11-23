package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	//checar se um valor existe no mapa

	fmt.Println(m["João Pedro"])
	v, ok := m["João Pedro"]

	fmt.Println(v)
	fmt.Println(ok)

	// outra forma de chegar com if

	if v, ok := m["João Pedro"]; !ok {
		fmt.Println("Este nome não tem valor no map", v)
	}
}
