package main

import "fmt"

func main() {
	x := 77

	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)
}
