package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print 2")
	case (3 == 3):
		fmt.Println("This should print")
		fallthrough
	case (3 == 15):
		fmt.Println("NOT TRUE")
		fallthrough
	case (3 == 11):
		fmt.Println("ALSO NOT TRUE")
		fallthrough
	case (4 == 4):
		fmt.Println("This also true, does it print?")
	default:
		fmt.Println("This is default")
	}

}
