package main

import "fmt"

var favSport string

func main() {
	favSport := "basquete"

	switch favSport {
	case "futebol":
		fmt.Printf("Se você gosta de %v, hoje tem jogo!", favSport)
	case "natação":
		fmt.Printf("Se você gosta de %v, se prepare para nadar!", favSport)
	case "basquete":
		fmt.Printf("Se você gosta de %v, Golden State entra em quadra hoje!", favSport)

	}
}
