package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	//address where a value is stored (&)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)

	// dereference the address and get the value
	// gives you the value stored at an address when you have the address
	fmt.Println(*b)
}
