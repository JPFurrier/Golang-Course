package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should not print")
	case (2 == 2):
		fmt.Println("Should Print")
	case (3 == 1):
		fmt.Println("should not print")
	}

}
