package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	x := 2
	return x
}

func bar() (int, string) {
	y := 1984
	z := "Babylon"

	return y, z
}
