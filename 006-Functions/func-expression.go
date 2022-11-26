package main

import "fmt"

func main() {

	// assign a function on a variable
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The discovery of Brazil happened in", x)
	}
	g(1500)
}
